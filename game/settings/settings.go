package settings

import "github.com/ThronesMC/game/game/participant"

type Mode string

const (
	NormalMode Mode = "Normal"
	SoloMode   Mode = "Solo"
	DuosMode   Mode = "Duos"
	DuelsMode  Mode = "Duels"
	SquadsMode Mode = "Squads"
)

type Settings struct {
	GameName   string
	MapsFolder string
	MapName    string
	Mode       Mode
	MinPlayers int
	MaxPlayers int
	TeamSize   int

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
		Mode:       NormalMode,
		MinPlayers: 0,
		MaxPlayers: -1,
		TeamSize:   0,
		NameFormat: nameFormat,
	}
}

func NewGameSettings(gameName, mapsFolder, mapName string, mode Mode, minPlayers, maxPlayers, teamSize int, nameFormat nameFormat) *Settings {
	if gameName == "" {
		panic("game name cannot be empty")
	}
	if mapsFolder == "" {
		panic("maps folder cannot be empty")
	}
	if mapName == "" {
		panic("map name cannot be empty")
	}
	if minPlayers <= 0 || maxPlayers <= 0 {
		panic("minPlayers and maxPlayers must be greater than 0")
	}
	if minPlayers > maxPlayers {
		panic("minPlayers cannot be greater than maxPlayers")
	}
	if teamSize <= 0 {
		panic("teamSize must be greater than 0")
	}
	if maxPlayers%teamSize != 0 {
		panic("maxPlayers must be divisible by teamSize")
	}
	return &Settings{
		GameName:   gameName,
		MapsFolder: mapsFolder,
		MapName:    mapName,
		Mode:       mode,
		MinPlayers: minPlayers,
		MaxPlayers: maxPlayers,
		TeamSize:   teamSize,
		NameFormat: nameFormat,
	}
}

type nameFormat func(viewer, pt *participant.Participant) string
