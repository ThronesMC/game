package handler

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/thronesmc/game/game"
)

type GlobalWorldHandler struct{}

func (GlobalWorldHandler) HandleLiquidFlow(ctx *world.Context, from, into cube.Pos, liquid world.Liquid, replaced world.Block) {
	game.GetGame().WorldHandler.HandleLiquidFlow(ctx, from, into, liquid, replaced)
}

func (GlobalWorldHandler) HandleLiquidDecay(ctx *world.Context, pos cube.Pos, before, after world.Liquid) {
	game.GetGame().WorldHandler.HandleLiquidDecay(ctx, pos, before, after)
}

func (GlobalWorldHandler) HandleLiquidHarden(ctx *world.Context, hardenedPos cube.Pos, liquidHardened, otherLiquid, newBlock world.Block) {
	game.GetGame().WorldHandler.HandleLiquidHarden(ctx, hardenedPos, liquidHardened, otherLiquid, newBlock)
}

func (GlobalWorldHandler) HandleSound(ctx *world.Context, s world.Sound, pos mgl64.Vec3) {
	game.GetGame().WorldHandler.HandleSound(ctx, s, pos)
}

func (GlobalWorldHandler) HandleFireSpread(ctx *world.Context, from, to cube.Pos) {
	game.GetGame().WorldHandler.HandleFireSpread(ctx, from, to)
}

func (GlobalWorldHandler) HandleBlockBurn(ctx *world.Context, pos cube.Pos) {
	game.GetGame().WorldHandler.HandleBlockBurn(ctx, pos)
}

func (GlobalWorldHandler) HandleCropTrample(ctx *world.Context, pos cube.Pos) {
	game.GetGame().WorldHandler.HandleCropTrample(ctx, pos)
}

func (GlobalWorldHandler) HandleLeavesDecay(ctx *world.Context, pos cube.Pos) {
	game.GetGame().WorldHandler.HandleLeavesDecay(ctx, pos)
}

func (GlobalWorldHandler) HandleEntitySpawn(tx *world.Tx, e world.Entity) {
	game.GetGame().WorldHandler.HandleEntitySpawn(tx, e)
}

func (GlobalWorldHandler) HandleEntityDespawn(tx *world.Tx, e world.Entity) {
	game.GetGame().WorldHandler.HandleEntityDespawn(tx, e)
}

func (GlobalWorldHandler) HandleExplosion(ctx *world.Context, position mgl64.Vec3, entities *[]world.Entity, blocks *[]cube.Pos, itemDropChance *float64, spawnFire *bool) {
	game.GetGame().WorldHandler.HandleExplosion(ctx, position, entities, blocks, itemDropChance, spawnFire)
}

func (GlobalWorldHandler) HandleClose(tx *world.Tx) {
	game.GetGame().WorldHandler.HandleClose(tx)
}
