package handler

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type ChainedWorldHandler struct {
	Middle world.Handler
	Next   world.Handler
}

func (cwh ChainedWorldHandler) HandleLiquidFlow(ctx *world.Context, from, into cube.Pos, liquid world.Liquid, replaced world.Block) {
	if cwh.Middle != nil {
		cwh.Middle.HandleLiquidFlow(ctx, from, into, liquid, replaced)
		if ctx.Cancelled() {
			return
		}
	}
	cwh.Next.HandleLiquidFlow(ctx, from, into, liquid, replaced)
}

func (cwh ChainedWorldHandler) HandleLiquidDecay(ctx *world.Context, pos cube.Pos, before, after world.Liquid) {
	if cwh.Middle != nil {
		cwh.Middle.HandleLiquidDecay(ctx, pos, before, after)
		if ctx.Cancelled() {
			return
		}
	}
	cwh.Next.HandleLiquidDecay(ctx, pos, before, after)
}

func (cwh ChainedWorldHandler) HandleLiquidHarden(ctx *world.Context, hardenedPos cube.Pos, liquidHardened, otherLiquid, newBlock world.Block) {
	if cwh.Middle != nil {
		cwh.Middle.HandleLiquidHarden(ctx, hardenedPos, liquidHardened, otherLiquid, newBlock)
		if ctx.Cancelled() {
			return
		}
	}
	cwh.Next.HandleLiquidHarden(ctx, hardenedPos, liquidHardened, otherLiquid, newBlock)
}

func (cwh ChainedWorldHandler) HandleSound(ctx *world.Context, s world.Sound, pos mgl64.Vec3) {
	if cwh.Middle != nil {
		cwh.Middle.HandleSound(ctx, s, pos)
		if ctx.Cancelled() {
			return
		}
	}
	cwh.Next.HandleSound(ctx, s, pos)
}

func (cwh ChainedWorldHandler) HandleFireSpread(ctx *world.Context, from, to cube.Pos) {
	if cwh.Middle != nil {
		cwh.Middle.HandleFireSpread(ctx, from, to)
		if ctx.Cancelled() {
			return
		}
	}
	cwh.Next.HandleFireSpread(ctx, from, to)
}

func (cwh ChainedWorldHandler) HandleBlockBurn(ctx *world.Context, pos cube.Pos) {
	if cwh.Middle != nil {
		cwh.Middle.HandleBlockBurn(ctx, pos)
		if ctx.Cancelled() {
			return
		}
	}
	cwh.Next.HandleBlockBurn(ctx, pos)
}

func (cwh ChainedWorldHandler) HandleCropTrample(ctx *world.Context, pos cube.Pos) {
	if cwh.Middle != nil {
		cwh.Middle.HandleCropTrample(ctx, pos)
		if ctx.Cancelled() {
			return
		}
	}
	cwh.Next.HandleCropTrample(ctx, pos)
}

func (cwh ChainedWorldHandler) HandleLeavesDecay(ctx *world.Context, pos cube.Pos) {
	if cwh.Middle != nil {
		cwh.Middle.HandleLeavesDecay(ctx, pos)
		if ctx.Cancelled() {
			return
		}
	}
	cwh.Next.HandleLeavesDecay(ctx, pos)
}

func (cwh ChainedWorldHandler) HandleEntitySpawn(tx *world.Tx, e world.Entity) {
	if cwh.Middle != nil {
		cwh.Middle.HandleEntitySpawn(tx, e)
	}
	cwh.Next.HandleEntitySpawn(tx, e)
}

func (cwh ChainedWorldHandler) HandleEntityDespawn(tx *world.Tx, e world.Entity) {
	if cwh.Middle != nil {
		cwh.Middle.HandleEntityDespawn(tx, e)
	}
	cwh.Next.HandleEntityDespawn(tx, e)
}

func (cwh ChainedWorldHandler) HandleExplosion(ctx *world.Context, position mgl64.Vec3, entities *[]world.Entity, blocks *[]cube.Pos, itemDropChance *float64, spawnFire *bool) {
	if cwh.Middle != nil {
		cwh.Middle.HandleExplosion(ctx, position, entities, blocks, itemDropChance, spawnFire)
		if ctx.Cancelled() {
			return
		}
	}
	cwh.Next.HandleExplosion(ctx, position, entities, blocks, itemDropChance, spawnFire)
}

func (cwh ChainedWorldHandler) HandleClose(tx *world.Tx) {
	if cwh.Middle != nil {
		cwh.Middle.HandleClose(tx)
	}
	cwh.Next.HandleClose(tx)
}

func WorldChainHandlers(final world.Handler, middles ...world.Handler) world.Handler {
	h := final
	for i := len(middles) - 1; i >= 0; i-- {
		h = ChainedWorldHandler{
			Middle: middles[i],
			Next:   h,
		}
	}
	return h
}
