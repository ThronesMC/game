package command

import (
	"github.com/ThronesMC/game/game"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type PauseCommand struct {
}

func (pc PauseCommand) Run(source cmd.Source, output *cmd.Output, _ *world.Tx) {
	_, ok := source.(*player.Player)
	if !ok {
		output.Error("Must be a player")

		return
	}

	series := game.GetGame().StateSeries

	if series.IsPaused() {
		output.Error("Game is already paused")

		return
	}

	series.SetPaused(true)
	output.Print(text.Colourf("<green>Game pause successfully</green>"))
}
