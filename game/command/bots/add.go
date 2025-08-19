package bots

import (
	"fmt"
	"github.com/ThronesMC/game/game"
	"github.com/ThronesMC/game/game/handler_custom"
	"github.com/ThronesMC/game/game/mechanic/bot"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/npc"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"log"
	"path"
	"time"
)

type AddSubCommand struct {
	Add    cmd.SubCommand    `cmd:"add"`
	Number cmd.Optional[int] `cmd:"number"`
}

func (asc AddSubCommand) Run(source cmd.Source, output *cmd.Output, tx *world.Tx) {
	p, ok := source.(*player.Player)
	if !ok {
		output.Error("Must be a player")
		return
	}

	number, ok := asc.Number.Load()
	if number <= 0 || !ok {
		number = 1
	}

	g := game.GetGame()
	needed := g.Settings.MaxPlayers - g.ParticipantLen()

	if needed <= 0 {
		output.Error("No bots needed, game is already full.")
		return
	}

	if number > needed {
		output.Print(text.Colourf("<yellow>You can only add %d bot(s) right now.</yellow>", needed))
		number = needed
	}

	go func() {
		for i := 0; i < number; i++ {
			if err := game.SkinManager.GenerateSkin(i); err != nil {
				log.Fatalf("Could not generate a new skin: %v", err)
			}
			p.H().ExecWorld(func(tx *world.Tx, e world.Entity) {
				skin := npc.MustSkin(npc.MustParseTexture(path.Join(".", "skins", fmt.Sprintf("skin_%v.png", i))), npc.DefaultModel)
				bot.AddBot(tx, p.Position(), p.Rotation(), skin, func(b *player.Player) {
					b.Handle(g.PlayerHandler)
					b.Handler().(handler_custom.JoinHandler).HandleJoin(b)
				})
			})
			time.Sleep(500 * time.Millisecond)
		}
	}()

	output.Print(text.Colourf("<green>%d bot(s) added successfully</green>", number))
}
