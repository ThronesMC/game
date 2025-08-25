package handler

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/inventory"
)

type InventoryHandler struct {
}

func (i InventoryHandler) HandleTake(ctx *inventory.Context, slot int, it item.Stack) {
	ctx.Cancel()
}

func (i InventoryHandler) HandlePlace(ctx *inventory.Context, slot int, it item.Stack) {
	ctx.Cancel()
}

func (i InventoryHandler) HandleDrop(ctx *inventory.Context, slot int, it item.Stack) {
	ctx.Cancel()
}
