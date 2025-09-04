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
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
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
			func(tx *world.Tx, viewer *player.Player, pt *participant.Participant) string {
				g := game.GetGame()
				if vpt := g.GetParticipant(viewer); vpt != nil {
					if game.GetGame().InSameTeam(vpt, pt) {
						return text.Colourf("<green>%v</green>", pt.TXPlayer(tx).Name())
					} else {
						return text.Colourf("<red>%v</red>", pt.TXPlayer(tx).Name())

					}
				} else {
					return text.Colourf("<grey>[SPECTATOR] %v</grey>", pt.TXPlayer(tx).Name())
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
		gamehandler.InventoryHandler{},
	)

	var cfg config.ExampleData
	if err := g.LoadGameMapWithConfig(&cfg); err != nil {
		panic(err)
	}

	g.World.Handle(handlerutils.WorldChainHandlers(gamehandler.WorldHandler{}))

	return g
}
