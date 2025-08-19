package states

import (
	"github.com/df-mc/dragonfly/server/player/bossbar"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/josscoder/fsmgo/state"
	"github.com/thronesmc/game/example/config"
	"github.com/thronesmc/game/game"
	"github.com/thronesmc/game/game/mechanic/bot"
	"github.com/thronesmc/game/game/mechanic/nametag"
	"github.com/thronesmc/game/game/mechanic/spawn"
	"github.com/thronesmc/game/game/participant"
	"log"
	"math"
	"time"
)

type PreGameState struct {
	*state.BaseState
}

func NewPreGameStateState() *PreGameState {
	s := &PreGameState{}
	s.BaseState = state.NewBaseState(s)
	return s
}

func (s *PreGameState) OnStart() {
	cfg := game.GetMapData[*config.ExampleData]()
	spawn.InitSpawns(cfg.Spawns)

	log.Println("PreGameState onStart")
}

func (s *PreGameState) OnUpdate() {
	log.Println("PreGameState OnUpdate")

	g := game.GetGame()

	g.World.Exec(func(tx *world.Tx) {
		for p1 := range g.GetParticipants() {
			for p2 := range g.GetParticipants() {
				if !bot.IsBot(p1.Player()) {
					nametag.RefreshNameTag(tx, p1, p2)
				}
			}
		}
	})

	duration := s.GetDuration().Seconds()
	remaining := s.GetRemainingTime().Seconds()
	progress := math.Max(0, math.Min(1, remaining/duration))

	g.ParticipantsCallback(func(pt *participant.Participant) {
		pt.Player().SendBossBar(bossbar.New("PreGameState").WithColour(bossbar.Grey()).WithHealthPercentage(progress))
	})
}

func (s *PreGameState) OnEnd() {
	log.Println("PreGameState OnEnd")
}

func (s *PreGameState) GetDuration() time.Duration {
	return 30 * time.Second
}
