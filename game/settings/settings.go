package settings

import (
	"github.com/ThronesMC/game/game/modes"
	"github.com/ThronesMC/game/game/participant"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type Settings struct {
	GameName   string
	MapsFolder string
	MapName    string
	Mode       modes.Mode

	NameFormat nameFormat
}

func NewStaticSettings(gameName, mapsFolder, mapName string, nameFormat nameFormat) *Settings {
	if gameName == "" {
		panic("game name cannot be empty")
	}
	if mapsFolder == "" {
		panic("maps folder cannot be empty")
	}
	if mapName == "" {
		panic("map name cannot be empty")
	}
	return &Settings{
		GameName:   gameName,
		MapsFolder: mapsFolder,
		MapName:    mapName,
		Mode:       modes.Normal{},
		NameFormat: nameFormat,
	}
}

func NewGameSettings(gameName, mapsFolder, mapName string, mode modes.Mode, nameFormat nameFormat) *Settings {
	if gameName == "" {
		panic("game name cannot be empty")
	}
	if mapsFolder == "" {
		panic("maps folder cannot be empty")
	}
	if mapName == "" {
		panic("map name cannot be empty")
	}
	if mode.MinimumTotalPlayers() <= 0 || mode.MaximumTotalPlayers() <= 0 {
		panic("minPlayers and maxPlayers must be greater than 0")
	}
	if mode.MinimumTotalPlayers() > mode.MaximumTotalPlayers() {
		panic("minPlayers cannot be greater than maxPlayers")
	}
	if mode.NumberOfPlayersPerTeam() <= 0 {
		panic("teamSize must be greater than 0")
	}
	if mode.MaximumTotalPlayers()%mode.NumberOfPlayersPerTeam() != 0 {
		panic("maxPlayers must be divisible by teamSize")
	}
	return &Settings{
		GameName:   gameName,
		MapsFolder: mapsFolder,
		MapName:    mapName,
		Mode:       mode,
		NameFormat: nameFormat,
	}
}

type nameFormat func(tx *world.Tx, viewer *player.Player, pt *participant.Participant) string
