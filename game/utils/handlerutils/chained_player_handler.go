package handlerutils

import (
	"github.com/ThronesMC/game/game"
	"github.com/ThronesMC/game/game/handler_custom"
	"github.com/df-mc/dragonfly/server/entity"
	"net"
	"time"

	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/skin"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type ChainedPlayerHandler struct {
	Middle handler_custom.JoinHandler
	Next   handler_custom.JoinHandler
}

func (cph ChainedPlayerHandler) HandleJoin(p *player.Player) {
	if err := game.GetGame().Join(p); err != nil {
		p.Disconnect(err.Error())
		return
	}
	if cph.Middle != nil {
		cph.Middle.HandleJoin(p)
	}
	cph.Next.HandleJoin(p)
}

func (cph ChainedPlayerHandler) HandleChangeWorld(p *player.Player, before, after *world.World) {
	if cph.Middle != nil {
		cph.Middle.HandleChangeWorld(p, before, after)
	}
	cph.Next.HandleChangeWorld(p, before, after)
}

func (cph ChainedPlayerHandler) HandleChat(ctx *player.Context, message *string) {
	if cph.Middle != nil {
		cph.Middle.HandleChat(ctx, message)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleChat(ctx, message)
}

func (cph ChainedPlayerHandler) HandleMove(ctx *player.Context, newPos mgl64.Vec3, newRot cube.Rotation) {
	if cph.Middle != nil {
		cph.Middle.HandleMove(ctx, newPos, newRot)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleMove(ctx, newPos, newRot)
}

func (cph ChainedPlayerHandler) HandleJump(p *player.Player) {
	if cph.Middle != nil {
		cph.Middle.HandleJump(p)
	}
	cph.Next.HandleJump(p)
}

func (cph ChainedPlayerHandler) HandleTeleport(ctx *player.Context, pos mgl64.Vec3) {
	if cph.Middle != nil {
		cph.Middle.HandleTeleport(ctx, pos)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleTeleport(ctx, pos)
}

func (cph ChainedPlayerHandler) HandleToggleSprint(ctx *player.Context, after bool) {
	if cph.Middle != nil {
		cph.Middle.HandleToggleSprint(ctx, after)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleToggleSprint(ctx, after)
}

func (cph ChainedPlayerHandler) HandleToggleSneak(ctx *player.Context, after bool) {
	if cph.Middle != nil {
		cph.Middle.HandleToggleSneak(ctx, after)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleToggleSneak(ctx, after)
}

func (cph ChainedPlayerHandler) HandleFoodLoss(ctx *player.Context, from int, to *int) {
	if cph.Middle != nil {
		cph.Middle.HandleFoodLoss(ctx, from, to)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleFoodLoss(ctx, from, to)
}

func (cph ChainedPlayerHandler) HandleHeal(ctx *player.Context, health *float64, src world.HealingSource) {
	if cph.Middle != nil {
		cph.Middle.HandleHeal(ctx, health, src)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleHeal(ctx, health, src)
}

func (cph ChainedPlayerHandler) HandleHurt(ctx *player.Context, damage *float64, immune bool, attackImmunity *time.Duration, src world.DamageSource) {
	if cph.Middle != nil {
		cph.Middle.HandleHurt(ctx, damage, immune, attackImmunity, src)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleHurt(ctx, damage, immune, attackImmunity, src)
}

func (cph ChainedPlayerHandler) HandleDeath(p *player.Player, src world.DamageSource, keepInv *bool) {
	if cph.Middle != nil {
		cph.Middle.HandleDeath(p, src, keepInv)
	}
	cph.Next.HandleDeath(p, src, keepInv)
}

func (cph ChainedPlayerHandler) HandleRespawn(p *player.Player, pos *mgl64.Vec3, w **world.World) {
	if cph.Middle != nil {
		cph.Middle.HandleRespawn(p, pos, w)
	}
	cph.Next.HandleRespawn(p, pos, w)
}

func (cph ChainedPlayerHandler) HandleSkinChange(ctx *player.Context, skin *skin.Skin) {
	if cph.Middle != nil {
		cph.Middle.HandleSkinChange(ctx, skin)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleSkinChange(ctx, skin)
}

func (cph ChainedPlayerHandler) HandleFireExtinguish(ctx *player.Context, pos cube.Pos) {
	if cph.Middle != nil {
		cph.Middle.HandleFireExtinguish(ctx, pos)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleFireExtinguish(ctx, pos)
}

func (cph ChainedPlayerHandler) HandleStartBreak(ctx *player.Context, pos cube.Pos) {
	if cph.Middle != nil {
		cph.Middle.HandleStartBreak(ctx, pos)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleStartBreak(ctx, pos)
}

func (cph ChainedPlayerHandler) HandleBlockBreak(ctx *player.Context, pos cube.Pos, drops *[]item.Stack, xp *int) {
	if cph.Middle != nil {
		cph.Middle.HandleBlockBreak(ctx, pos, drops, xp)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleBlockBreak(ctx, pos, drops, xp)
}

func (cph ChainedPlayerHandler) HandleBlockPlace(ctx *player.Context, pos cube.Pos, b world.Block) {
	if cph.Middle != nil {
		cph.Middle.HandleBlockPlace(ctx, pos, b)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleBlockPlace(ctx, pos, b)
}

func (cph ChainedPlayerHandler) HandleBlockPick(ctx *player.Context, pos cube.Pos, b world.Block) {
	if cph.Middle != nil {
		cph.Middle.HandleBlockPick(ctx, pos, b)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleBlockPick(ctx, pos, b)
}

func (cph ChainedPlayerHandler) HandleItemUse(ctx *player.Context) {
	if cph.Middle != nil {
		cph.Middle.HandleItemUse(ctx)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleItemUse(ctx)
}

func (cph ChainedPlayerHandler) HandleItemUseOnBlock(ctx *player.Context, pos cube.Pos, face cube.Face, clickPos mgl64.Vec3) {
	if cph.Middle != nil {
		cph.Middle.HandleItemUseOnBlock(ctx, pos, face, clickPos)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleItemUseOnBlock(ctx, pos, face, clickPos)
}

func (cph ChainedPlayerHandler) HandleItemUseOnEntity(ctx *player.Context, e world.Entity) {
	if cph.Middle != nil {
		cph.Middle.HandleItemUseOnEntity(ctx, e)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleItemUseOnEntity(ctx, e)
}

func (cph ChainedPlayerHandler) HandleItemRelease(ctx *player.Context, item item.Stack, dur time.Duration) {
	if cph.Middle != nil {
		cph.Middle.HandleItemRelease(ctx, item, dur)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleItemRelease(ctx, item, dur)
}

func (cph ChainedPlayerHandler) HandleItemConsume(ctx *player.Context, item item.Stack) {
	if cph.Middle != nil {
		cph.Middle.HandleItemConsume(ctx, item)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleItemConsume(ctx, item)
}

func (cph ChainedPlayerHandler) HandleAttackEntity(ctx *player.Context, e world.Entity, force, height *float64, critical *bool) {
	if cph.Middle != nil {
		cph.Middle.HandleAttackEntity(ctx, e, force, height, critical)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleAttackEntity(ctx, e, force, height, critical)
}

func (cph ChainedPlayerHandler) HandleExperienceGain(ctx *player.Context, amount *int) {
	if cph.Middle != nil {
		cph.Middle.HandleExperienceGain(ctx, amount)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleExperienceGain(ctx, amount)
}

func (cph ChainedPlayerHandler) HandlePunchAir(ctx *player.Context) {
	if cph.Middle != nil {
		cph.Middle.HandlePunchAir(ctx)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandlePunchAir(ctx)
}

func (cph ChainedPlayerHandler) HandleSignEdit(ctx *player.Context, pos cube.Pos, frontSide bool, oldText, newText string) {
	if cph.Middle != nil {
		cph.Middle.HandleSignEdit(ctx, pos, frontSide, oldText, newText)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleSignEdit(ctx, pos, frontSide, oldText, newText)
}

func (cph ChainedPlayerHandler) HandleLecternPageTurn(ctx *player.Context, pos cube.Pos, oldPage int, newPage *int) {
	if cph.Middle != nil {
		cph.Middle.HandleLecternPageTurn(ctx, pos, oldPage, newPage)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleLecternPageTurn(ctx, pos, oldPage, newPage)
}

func (cph ChainedPlayerHandler) HandleItemDamage(ctx *player.Context, i item.Stack, damage int) {
	if cph.Middle != nil {
		cph.Middle.HandleItemDamage(ctx, i, damage)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleItemDamage(ctx, i, damage)
}

func (cph ChainedPlayerHandler) HandleItemPickup(ctx *player.Context, i *item.Stack) {
	if cph.Middle != nil {
		cph.Middle.HandleItemPickup(ctx, i)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleItemPickup(ctx, i)
}

func (cph ChainedPlayerHandler) HandleHeldSlotChange(ctx *player.Context, from, to int) {
	if cph.Middle != nil {
		cph.Middle.HandleHeldSlotChange(ctx, from, to)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleHeldSlotChange(ctx, from, to)
}

func (cph ChainedPlayerHandler) HandleItemDrop(ctx *player.Context, s item.Stack) {
	if cph.Middle != nil {
		cph.Middle.HandleItemDrop(ctx, s)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleItemDrop(ctx, s)
}

func (cph ChainedPlayerHandler) HandleMount(ctx *player.Context, r entity.Rideable, seatPos *mgl64.Vec3, driver *bool) {
	if cph.Middle != nil {
		cph.Middle.HandleMount(ctx, r, seatPos, driver)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleMount(ctx, r, seatPos, driver)
}

func (cph ChainedPlayerHandler) HandleDismount(ctx *player.Context, r entity.Rideable) {
	if cph.Middle != nil {
		cph.Middle.HandleDismount(ctx, r)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleDismount(ctx, r)
}

func (cph ChainedPlayerHandler) HandleTransfer(ctx *player.Context, addr *net.UDPAddr) {
	if cph.Middle != nil {
		cph.Middle.HandleTransfer(ctx, addr)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleTransfer(ctx, addr)
}

func (cph ChainedPlayerHandler) HandleCommandExecution(ctx *player.Context, command cmd.Command, args []string) {
	if cph.Middle != nil {
		cph.Middle.HandleCommandExecution(ctx, command, args)
		if ctx.Cancelled() {
			return
		}
	}
	cph.Next.HandleCommandExecution(ctx, command, args)
}

func (cph ChainedPlayerHandler) HandleQuit(p *player.Player) {
	if cph.Middle != nil {
		cph.Middle.HandleQuit(p)
	}
	cph.Next.HandleQuit(p)
	game.GetGame().Quit(p)
}

func (cph ChainedPlayerHandler) HandleDiagnostics(p *player.Player, d session.Diagnostics) {
	if cph.Middle != nil {
		cph.Middle.HandleDiagnostics(p, d)
	}
	cph.Next.HandleDiagnostics(p, d)
}

func PlayerChainHandlers(final handler_custom.JoinHandler, middles ...handler_custom.JoinHandler) handler_custom.JoinHandler {
	h := final
	for i := len(middles) - 1; i >= 0; i-- {
		h = ChainedPlayerHandler{
			Middle: middles[i],
			Next:   h,
		}
	}
	return h
}
