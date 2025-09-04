package participant

import (
	"github.com/ThronesMC/game/game/config"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type Participant struct {
	player *player.Player
	handle *world.EntityHandle
	state  State

	participantData config.ParticipantData
}

func NewParticipant(player *player.Player, handle *world.EntityHandle) *Participant {
	return &Participant{
		player:          player,
		handle:          handle,
		state:           nil,
		participantData: nil,
	}
}

func (pt *Participant) Player() *player.Player {
	return pt.player
}

func (pt *Participant) TXPlayer(tx *world.Tx) *player.Player {
	e, ok := pt.handle.Entity(tx)
	if !ok {
		return nil
	}
	return e.(*player.Player)
}

func (pt *Participant) Handle() *world.EntityHandle {
	return pt.handle
}

func (pt *Participant) State() State {
	return pt.state
}

func (pt *Participant) InState(s State) bool {
	return pt.state == s
}

func (pt *Participant) SetState(s State) {
	pt.state = s
}
