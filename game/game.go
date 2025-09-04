package game

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ThronesMC/game/game/config"
	"github.com/ThronesMC/game/game/handler_custom"
	"github.com/ThronesMC/game/game/participant"
	"github.com/ThronesMC/game/game/settings"
	"github.com/ThronesMC/game/game/team"
	"github.com/ThronesMC/game/game/utils/maputils"
	"github.com/ThronesMC/game/game/utils/ziputils"
	"github.com/df-mc/dragonfly/server/item/inventory"
	"iter"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/sandertv/gophertunnel/minecraft/text"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/title"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/google/uuid"
	"github.com/josscoder/fsmgo/state"
)

var gameInstance *Game

type Game struct {
	Settings *settings.Settings
	Teams    []*team.Team

	PlayerHandler    handler_custom.JoinHandler
	InventoryHandler inventory.Handler

	StateSeries  *state.ScheduledStateSeries
	Participants *maputils.Map[uuid.UUID, *participant.Participant]

	MapLoaded bool
	mapConfig config.MapData

	World       *world.World
	WorldFolder string
}

func NewGame(settings *settings.Settings, teams []*team.Team, states []state.State, playerHandler handler_custom.JoinHandler, invHandler inventory.Handler) *Game {
	if playerHandler == nil {
		panic("player handler cannot be nil")
	}

	series := state.NewScheduledStateSeries(states, 100*time.Millisecond)

	game := &Game{
		Settings:         settings,
		Teams:            teams,
		PlayerHandler:    playerHandler,
		InventoryHandler: invHandler,
		StateSeries:      series,
		Participants:     maputils.NewMap[uuid.UUID, *participant.Participant](),
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

func (g *Game) GetParticipant(p *player.Player) *participant.Participant {
	pt, _ := g.Participants.Load(p.UUID())
	return pt
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
		fn(pt)
	}
}

func (g *Game) BroadcastMessage(tx *world.Tx, msg string) {
	g.ParticipantsCallback(func(pt *participant.Participant) {
		pt.TXPlayer(tx).Message(text.Colourf("%s", msg))
	})
}

func (g *Game) BroadcastMessagef(tx *world.Tx, format string, args ...any) {
	g.ParticipantsCallback(func(pt *participant.Participant) {
		pt.TXPlayer(tx).Message(text.Colourf(format, args...))
	})
}

func (g *Game) BroadcastTitle(tx *world.Tx, t title.Title) {
	g.ParticipantsCallback(func(pt *participant.Participant) {
		pt.TXPlayer(tx).SendTitle(t)
	})
}

func (g *Game) RandomAvailableTeam() (*team.Team, bool) {
	var available []*team.Team
	for _, t := range g.Teams {
		if t.Teammates.Len() < g.Settings.TeamSize {
			available = append(available, t)
		}
	}

	if len(available) == 0 {
		return nil, false
	}

	return available[rand.Intn(len(available))], true
}

func (g *Game) BalancedAvailableTeam() (*team.Team, bool) {
	var bestTeam *team.Team
	minCount := g.Settings.TeamSize + 1

	for _, t := range g.Teams {
		count := t.Teammates.Len()
		if count < g.Settings.TeamSize && count < minCount {
			minCount = count
			bestTeam = t
		}
	}

	if bestTeam == nil {
		return nil, false
	}
	return bestTeam, true
}

func (g *Game) AssignTeamToParticipant(pt *participant.Participant, team *team.Team) {
	team.Teammates.Store(pt.Player().UUID(), pt)
}

func (g *Game) RemoveFromTeam(pt *participant.Participant) {
	playerUUID := pt.Player().UUID()
	if t := g.TeamOf(pt); t != nil {
		t.Teammates.Delete(playerUUID)
	}
}

func (g *Game) GetTeamByID(id string) *team.Team {
	for _, t := range g.Teams {
		if t.GetID() == id {
			return t
		}
	}
	return nil
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

func (g *Game) TeamOf(pt *participant.Participant) *team.Team {
	playerUUID := pt.Player().UUID()
	for _, t := range g.Teams {
		if _, ok := t.Teammates.Load(playerUUID); ok {
			return t
		}
	}
	return nil
}

func (g *Game) InSameTeam(a, b *participant.Participant) bool {
	teamA := g.TeamOf(a)
	teamB := g.TeamOf(b)
	return teamA != nil && teamA == teamB
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

func (g *Game) Stop(tx *world.Tx) {
	g.StateSeries.End()

	g.ParticipantsCallback(func(pt *participant.Participant) {
		pt.TXPlayer(tx).Disconnect("game server shutdown")
	})

	if err := os.RemoveAll(g.WorldFolder); err != nil {
		fmt.Println("warning: failed to remove world folder:", err)
	}
}
