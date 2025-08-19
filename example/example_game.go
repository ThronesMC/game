package main

import (
	"github.com/josscoder/fsmgo/state"
	"github.com/thronesmc/game/example/config"
	gamehandler "github.com/thronesmc/game/example/handler"
	"github.com/thronesmc/game/example/states"
	"github.com/thronesmc/game/game"
	"github.com/thronesmc/game/game/handler"
	"github.com/thronesmc/game/game/settings"
	"github.com/thronesmc/game/game/team"
)

func NewExampleGame() *game.Game {
	var teamCfg config.ExampleTeamData

	g := game.NewGame(
		settings.NewGameSettings(
			"Example",
			"maps/",
			"Example1",
			settings.SoloMode,
			2,
			16,
			4,
			"<green>%v</green>",
			"<red>%v</red>",
		),
		[]*team.Team{
			team.NewTeam("red", "Red", team.Red, &teamCfg),
			team.NewTeam("green", "Green", team.Green, &teamCfg),
			team.NewTeam("blue", "Blue", team.Blue, &teamCfg),
			team.NewTeam("yellow", "Yellow", team.Yellow, &teamCfg),
		},
		[]state.State{
			states.NewPreGameStateState(),
			states.NewEndGameState(),
		},
		handler.WorldChainHandlers(gamehandler.WorldHandler{}),
		handler.PlayerChainHandlers(
			handler.GlobalPlayerHandler{},
			gamehandler.PlayerHandler{},
		),
	)

	var cfg config.ExampleData
	if err := g.LoadGameMapWithConfig(&cfg); err != nil {
		panic(err)
	}

	return g
}
