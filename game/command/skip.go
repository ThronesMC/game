package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/thronesmc/game/game"
)

type SkipCommand struct {
}

func (sc SkipCommand) Run(source cmd.Source, output *cmd.Output, _ *world.Tx) {
	_, ok := source.(*player.Player)
	if !ok {
		output.Error("Must be a player")

		return
	}

	game.GetGame().StateSeries.Skip()
	output.Print(text.Colourf("<green>Game skipped to the next state successfully</green>"))
}
