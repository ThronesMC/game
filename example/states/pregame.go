package states

import (
	"fmt"
	"github.com/ThronesMC/game/example/config"
	"github.com/ThronesMC/game/game"
	"github.com/ThronesMC/game/game/mechanic/bot"
	"github.com/ThronesMC/game/game/mechanic/nametag"
	"github.com/ThronesMC/game/game/mechanic/spawn"
	"github.com/ThronesMC/game/game/participant"
	"github.com/df-mc/dragonfly/server/player/bossbar"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/josscoder/fsmgo/state"
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
				if !bot.IsBot(p1.TXPlayer(tx)) {
					nametag.RefreshNameTag(tx, p1.TXPlayer(tx), p2)
				}
			}
		}
	})

	duration := float64(s.GetDuration().Milliseconds())
	remaining := float64(s.GetRemainingTime().Milliseconds())
	progress := math.Max(0, math.Min(1, remaining/duration))

	remainingSecs := float64(s.GetRemainingTime().Milliseconds()) / 1000.0
	remainingStr := fmt.Sprintf("%.1f s", remainingSecs)

	g.World.Exec(func(tx *world.Tx) {
		g.ParticipantsCallback(func(pt *participant.Participant) {
			pt.TXPlayer(tx).SendBossBar(
				bossbar.New("PreGameState - " + remainingStr).
					WithColour(bossbar.Grey()).
					WithHealthPercentage(progress),
			)
		})
	})
}

func (s *PreGameState) OnEnd() {
	log.Println("PreGameState OnEnd")
}

func (s *PreGameState) GetDuration() time.Duration {
	return 30 * time.Second
}
