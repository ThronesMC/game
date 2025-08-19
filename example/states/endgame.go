package states

import (
	"github.com/df-mc/dragonfly/server/player/bossbar"
	"github.com/josscoder/fsmgo/state"
	"github.com/thronesmc/game/game"
	"github.com/thronesmc/game/game/participant"
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

	duration := s.GetDuration().Seconds()
	remaining := s.GetRemainingTime().Seconds()
	progress := math.Max(0, math.Min(1, remaining/duration))

	game.GetGame().ParticipantsCallback(func(pt *participant.Participant) {
		pt.Player().SendBossBar(bossbar.New("EndGameState").WithColour(bossbar.Red()).WithHealthPercentage(progress))
	})
}

func (s *EndGameState) OnEnd() {
	log.Println("EndGameState onEnd")
}

func (s *EndGameState) GetDuration() time.Duration {
	return 10 * time.Second
}
