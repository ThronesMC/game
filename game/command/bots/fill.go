package bots

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/thronesmc/game/game"
)

type FillSubCommand struct {
	Fill cmd.SubCommand `cmd:"fill"`
}

func (asc FillSubCommand) Run(source cmd.Source, output *cmd.Output, tx *world.Tx) {
	p, ok := source.(*player.Player)
	if !ok {
		output.Error("Must be a player")
		return
	}

	g := game.GetGame()
	needed := g.Settings.MaxPlayers - g.ParticipantLen()

	p.ExecuteCommand(fmt.Sprintf("/bots add %d", needed))
}
