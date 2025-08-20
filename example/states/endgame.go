package states

import (
	"fmt"
	"github.com/ThronesMC/game/game"
	"github.com/ThronesMC/game/game/participant"
	"github.com/df-mc/dragonfly/server/player/bossbar"
	"github.com/josscoder/fsmgo/state"
	"log"
	"math"
	"time"
)

type EndGameState struct {
	*state.BaseState
}

func NewEndGameState() *EndGameState {
	s := &EndGameState{}
	s.BaseState = state.NewBaseState(s)
	return s
}

func (s *EndGameState) OnStart() {
	log.Println("EndGameState onStart")
}

func (s *EndGameState) OnUpdate() {
	log.Println("EndGameState onUpdate")

	duration := float64(s.GetDuration().Milliseconds())
	remaining := float64(s.GetRemainingTime().Milliseconds())
	progress := math.Max(0, math.Min(1, remaining/duration))

	remainingSecs := float64(s.GetRemainingTime().Milliseconds()) / 1000.0
	remainingStr := fmt.Sprintf("%.1f s", remainingSecs)

	game.GetGame().ParticipantsCallback(func(pt *participant.Participant) {
		pt.Player().SendBossBar(
			bossbar.New("EndGameState - " + remainingStr).
				WithColour(bossbar.Red()).
				WithHealthPercentage(progress),
		)
	})
}

func (s *EndGameState) OnEnd() {
	log.Println("EndGameState onEnd")
}

func (s *EndGameState) GetDuration() time.Duration {
	return 10 * time.Second
}
