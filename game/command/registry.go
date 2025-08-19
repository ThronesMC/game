package command

import (
	bots2 "github.com/ThronesMC/game/game/command/bots"
	"github.com/df-mc/dragonfly/server/cmd"
)

func RegisterDevCommands() {
	cmd.Register(cmd.New(
		"pause",
		"Pause your game",
		nil,
		PauseCommand{},
	))
	cmd.Register(cmd.New(
		"resume",
		"Resume your game",
		nil,
		ResumeCommand{},
	))
	cmd.Register(cmd.New(
		"skip",
		"Skip to next state",
		nil,
		SkipCommand{},
	))
	cmd.Register(cmd.New(
		"location",
		"Show your position",
		[]string{"loc"},
		LocationCommand{},
	))
	cmd.Register(cmd.New(
		"bots",
		"Manage game bots",
		nil,
		bots2.AddSubCommand{},
		bots2.RemoveSubCommand{},
		bots2.FillSubCommand{},
	))
}
