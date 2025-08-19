package playerutils

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
)

type ResetOpts struct {
	ResetInventory    bool
	ResetArmour       bool
	ClearEffects      bool
	ResetXP           bool
	ResetHealth       bool
	ResetFood         bool
	ResetScale        bool
	Extinguish        bool
	ResetFallDistance bool
	HealAmount        float64
	MaxHealth         float64
}

func DefaultResetOpts() ResetOpts {
	return ResetOpts{
		ResetInventory:    true,
		ResetArmour:       true,
		ClearEffects:      true,
		ResetXP:           true,
		ResetHealth:       true,
		ResetFood:         true,
		ResetScale:        true,
		Extinguish:        true,
		ResetFallDistance: true,
		HealAmount:        20,
		MaxHealth:         20,
	}
}

type ResetPlayerHealSource struct{}

func (ResetPlayerHealSource) HealingSource() {}

func ResetPlayer(p *player.Player, opts *ResetOpts) {
	if opts == nil {
		def := DefaultResetOpts()
		opts = &def
	}

	if opts.ResetInventory {
		p.Inventory().Clear()
	}
	if opts.ResetArmour {
		p.Armour().Clear()
	}
	p.MoveItemsToInventory() //
	p.SetHeldItems(item.Stack{}, item.Stack{})

	if opts.Extinguish {
		p.Extinguish()
	}
	if opts.ClearEffects {
		for _, effect := range p.Effects() {
			p.RemoveEffect(effect.Type())
		}
	}
	if opts.ResetFallDistance {
		p.ResetFallDistance()
	}
	if opts.ResetXP {
		p.SetExperienceProgress(0)
		p.SetExperienceLevel(0)
	}
	if opts.ResetHealth {
		p.SetMaxHealth(opts.MaxHealth)
		p.Heal(opts.HealAmount, ResetPlayerHealSource{})
	}
	if opts.ResetScale {
		p.SetScale(1)
	}
	if opts.ResetFood {
		p.SetFood(20)
	}
}
