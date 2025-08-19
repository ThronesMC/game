package handler_custom

import "github.com/df-mc/dragonfly/server/player"

type JoinHandler interface {
	player.Handler
	HandleJoin(p *player.Player)
}

type NopJoinHandler struct {
	player.NopHandler
}

func (NopJoinHandler) HandleJoin(*player.Player) {}
