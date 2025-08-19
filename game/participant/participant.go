package participant

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/thronesmc/game/game/config"
)

type Participant struct {
	player *player.Player
	handle *world.EntityHandle
	state  State

	participantData config.ParticipantData
}

func NewParticipant(player *player.Player, handle *world.EntityHandle) *Participant {
	return &Participant{player, handle, nil, nil}
}

func (pt *Participant) Player() *player.Player {
	return pt.player
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
