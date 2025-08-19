package game

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/thronesmc/game/game/config"
	"github.com/thronesmc/game/game/handler_custom"
	"github.com/thronesmc/game/game/mechanic/bot"
	"github.com/thronesmc/game/game/participant"
	"github.com/thronesmc/game/game/utils/maputils"
	"github.com/thronesmc/game/game/utils/ziputils"
	"iter"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/sandertv/gophertunnel/minecraft/text"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/title"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/google/uuid"
	"github.com/josscoder/fsmgo/state"
)

var gameInstance *Game

type Game struct {
	Settings *Settings
	Teams    []*Team

	WorldHandler  world.Handler
	PlayerHandler handler_custom.JoinHandler

	StateSeries  *state.ScheduledStateSeries
	Participants *maputils.Map[uuid.UUID, *participant.Participant]

	MapLoaded bool
	mapConfig config.MapData

	World       *world.World
	WorldFolder string
}

func NewGame(settings *Settings, teams []*Team, states []state.State, worldHandler world.Handler, playerHandler handler_custom.JoinHandler) *Game {
	if worldHandler == nil || playerHandler == nil {
		panic("world handler and player handler cannot be nil")
	}

	series := state.NewScheduledStateSeries(states)

	game := &Game{
		Settings:      settings,
		Teams:         teams,
		WorldHandler:  worldHandler,
		PlayerHandler: playerHandler,
		StateSeries:   series,
		Participants:  maputils.NewMap[uuid.UUID, *participant.Participant](),
	}
	gameInstance = game
	return game
}

func GetGame() *Game {
	return gameInstance
}

func (g *Game) Start() error {
	g.StateSeries.Start()
	return nil
}

func (g *Game) LoadGameMapWithConfig(config config.MapData) error {
	if g.MapLoaded {
		return errors.New("map already loaded")
	}

	mapDir := filepath.Join(g.Settings.MapsFolder, g.Settings.MapName)
	worldZip := filepath.Join(mapDir, "world.zip")

	stat, err := os.Stat(worldZip)
	if err != nil || stat.IsDir() {
		return fmt.Errorf("world.zip not found or is not a file: %s", worldZip)
	}

	rawConfig, err := os.ReadFile(filepath.Join(mapDir, "config.json"))
	if err != nil {
		return fmt.Errorf("failed reading config.json: %w", err)
	}

	if err := json.Unmarshal(rawConfig, config); err != nil {
		return fmt.Errorf("failed unmarshalling config.json: %w", err)
	}

	g.WorldFolder = "world"

	if err := os.RemoveAll(g.WorldFolder); err != nil {
		return fmt.Errorf("failed to remove existing world folder: %w", err)
	}

	if err := ziputils.UnZipFile(worldZip, g.WorldFolder); err != nil {
		return fmt.Errorf("failed to copy world: %w", err)
	}

	g.MapLoaded = true
	g.mapConfig = config

	return nil
}

func GetMapData[TM config.MapData]() TM {
	if !gameInstance.MapLoaded {
		panic("map is not loaded")
	}
	return gameInstance.mapConfig.(TM)
}

func (g *Game) GetParticipant(p *player.Player) (*participant.Participant, bool) {
	return g.Participants.Load(p.UUID())
}

func (g *Game) GetParticipants() iter.Seq[*participant.Participant] {
	return func(yield func(*participant.Participant) bool) {
		for _, par := range g.Participants.Map() {
			if !yield(par) {
				return
			}
		}
	}
}

func (g *Game) ParticipantLen() int {
	return g.Participants.Len()
}

func (g *Game) HasEnoughPlayers() bool {
	return g.ParticipantLen() >= g.Settings.MinPlayers
}

func (g *Game) IsFull() bool {
	return g.Settings.MaxPlayers != -1 && g.ParticipantLen() >= g.Settings.MaxPlayers
}

func (g *Game) ParticipantsCallback(fn func(pt *participant.Participant)) {
	for _, pt := range g.Participants.Map() {
		if !bot.IsBot(pt.Player()) {
			fn(pt)
		}
	}
}

func (g *Game) BroadcastMessage(msg string) {
	g.ParticipantsCallback(func(pt *participant.Participant) {
		pt.Player().Message(text.Colourf("%s", msg))
	})
}

func (g *Game) BroadcastMessagef(format string, args ...any) {
	g.ParticipantsCallback(func(pt *participant.Participant) {
		pt.Player().Message(text.Colourf(format, args...))
	})
}

func (g *Game) BroadcastTitle(t title.Title) {
	g.ParticipantsCallback(func(pt *participant.Participant) {
		pt.Player().SendTitle(t)
	})
}

func (g *Game) RandomAvailableTeam() (*Team, bool) {
	var available []*Team
	for _, team := range g.Teams {
		if team.Teammates.Len() < g.Settings.TeamSize {
			available = append(available, team)
		}
	}

	if len(available) == 0 {
		return nil, false
	}

	return available[rand.Intn(len(available))], true
}

func (g *Game) BalancedAvailableTeam() (*Team, bool) {
	var bestTeam *Team
	minCount := g.Settings.TeamSize + 1

	for _, team := range g.Teams {
		count := team.Teammates.Len()
		if count < g.Settings.TeamSize && count < minCount {
			minCount = count
			bestTeam = team
		}
	}

	if bestTeam == nil {
		return nil, false
	}
	return bestTeam, true
}

func (g *Game) AssignTeamToParticipant(pt *participant.Participant, team *Team) {
	team.Teammates.Store(pt.Player().UUID(), pt)
}

func (g *Game) RemoveFromTeam(pt *participant.Participant) {
	playerUUID := pt.Player().UUID()
	if team, ok := g.TeamOf(pt); ok {
		team.Teammates.Delete(playerUUID)
	}
}

func (g *Game) GetTeamByID(id string) (*Team, bool) {
	for _, t := range g.Teams {
		if t.GetID() == id {
			return t, true
		}
	}
	return nil, false
}

func (g *Game) EnemiesOf(pt *participant.Participant) []*participant.Participant {
	var enemies []*participant.Participant

	for other := range g.GetParticipants() {
		if other == pt || g.InSameTeam(pt, other) {
			continue
		}

		enemies = append(enemies, other)
	}

	return enemies
}

func (g *Game) TeamOf(pt *participant.Participant) (*Team, bool) {
	playerUUID := pt.Player().UUID()
	for _, team := range g.Teams {
		if _, ok := team.Teammates.Load(playerUUID); ok {
			return team, true
		}
	}
	return nil, false
}

func (g *Game) InSameTeam(a, b *participant.Participant) bool {
	teamA, okA := g.TeamOf(a)
	teamB, okB := g.TeamOf(b)
	return okA && okB && teamA.GetID() == teamB.GetID()
}

func (g *Game) Join(p *player.Player) error {
	if !g.MapLoaded {
		return errors.New("game map is not loaded yet")
	}

	g.Participants.Store(p.UUID(), participant.NewParticipant(p, p.H()))
	return nil
}

func (g *Game) Quit(p *player.Player) {
	g.Participants.Delete(p.UUID())
}

func (g *Game) Stop() {
	g.StateSeries.End()

	g.ParticipantsCallback(func(pt *participant.Participant) {
		pt.Player().Disconnect("game server shutdown")
	})

	if err := os.RemoveAll(g.WorldFolder); err != nil {
		fmt.Println("warning: failed to remove world folder:", err)
	}
}
