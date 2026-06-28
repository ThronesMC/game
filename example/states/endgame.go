package states

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/ThronesMC/game/game"
	"github.com/ThronesMC/game/game/participant"
	"github.com/df-mc/dragonfly/server/player/bossbar"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/josscoder/fsmgo/state"
)

// EndGameState manages the example end-game phase with a boss bar countdown.
type EndGameState struct {
	*state.BaseState
}

// Compile-time check that EndGameState implements state.Lifecycle.
var _ state.Lifecycle = (*EndGameState)(nil)

// NewEndGameState creates a new example end-game state.
func NewEndGameState() *EndGameState {
	s := &EndGameState{}
	s.BaseState = state.NewBaseState(s)
	return s
}

func (s *EndGameState) OnStart() {
	log.Println("EndGameState onStart")
}

func (s *EndGameState) OnUpdate(_ time.Duration) {
	log.Println("EndGameState onUpdate")

	duration := float64(s.GetDuration().Milliseconds())
	remaining := float64(s.GetRemainingTime().Milliseconds())
	progress := math.Max(0, math.Min(1, remaining/duration))

	remainingSecs := float64(s.GetRemainingTime().Milliseconds()) / 1000.0
	remainingStr := fmt.Sprintf("%.1f s", remainingSecs)

	game.GetGame().World.Exec(func(tx *world.Tx) {
		game.GetGame().ParticipantsCallback(func(pt *participant.Participant) {
			pt.TXPlayer(tx).SendBossBar(
				bossbar.New("EndGameState - " + remainingStr).
					WithColour(bossbar.Red()).
					WithHealthPercentage(progress),
			)
		})
	})
}

func (s *EndGameState) OnEnd() {
	log.Println("EndGameState onEnd")
}

func (s *EndGameState) GetDuration() time.Duration {
	return 10 * time.Second
}
