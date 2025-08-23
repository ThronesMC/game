package main

import (
	"github.com/ThronesMC/game/example/config"
	gamehandler "github.com/ThronesMC/game/example/handler"
	"github.com/ThronesMC/game/example/states"
	"github.com/ThronesMC/game/game"
	"github.com/ThronesMC/game/game/participant"
	"github.com/ThronesMC/game/game/settings"
	"github.com/ThronesMC/game/game/team"
	"github.com/ThronesMC/game/game/utils/handlerutils"
	"github.com/josscoder/fsmgo/state"
	"github.com/sandertv/gophertunnel/minecraft/text"
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
			func(viewer, pt *participant.Participant) string {
				if game.GetGame().InSameTeam(viewer, pt) {
					return text.Colourf("<green>%v</green>", pt.Player().Name())
				} else {
					return text.Colourf("<red>%v</red>", pt.Player().Name())

				}
			},
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
		handlerutils.PlayerChainHandlers(
			gamehandler.PlayerHandler{},
		),
	)

	var cfg config.ExampleData
	if err := g.LoadGameMapWithConfig(&cfg); err != nil {
		panic(err)
	}

	g.World.Handle(handlerutils.WorldChainHandlers(gamehandler.WorldHandler{}))

	return g
}
