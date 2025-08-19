package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/thronesmc/game/game"
)

type ResumeCommand struct {
}

func (rc ResumeCommand) Run(source cmd.Source, output *cmd.Output, _ *world.Tx) {
	_, ok := source.(*player.Player)
	if !ok {
		output.Error("Must be a player")

		return
	}

	series := game.GetGame().StateSeries

	if !series.IsPaused() {
		output.Error("Game is not paused")

		return
	}

	series.SetPaused(false)
	output.Print(text.Colourf("<green>Game resumed successfully</green>"))
}
