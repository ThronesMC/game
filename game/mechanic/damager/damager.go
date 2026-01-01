package damager

import (
	"github.com/ThronesMC/game/game/utils/maputils"
	"time"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/google/uuid"
)

var damagers = maputils.NewMap[string, []Info]()

// Info stores information about received damage.
type Info struct {
	Damager   world.Entity
	DamagerID string // Stable identifier for the damager
	Damage    float64
	Timestamp time.Time
}

// RegisterDamage registers that the victim was damaged by the damager.
func RegisterDamage(victim, damager world.Entity) {
	RegisterDamageWithAmount(victim, damager, 1.0)
}

// RegisterDamageWithAmount registers damage with a specific amount for tracking damage contribution.
func RegisterDamageWithAmount(victim, damager world.Entity, amount float64) {
	if victim == damager {
		return
	}

	victimID := getEntityID(victim)
	damagerID := getEntityID(damager)

	info := Info{
		Damager:   damager,
		DamagerID: damagerID,
		Damage:    amount,
		Timestamp: time.Now(),
	}

	list, _ := damagers.Load(victimID)

	// Update existing entry if damager already in list
	found := false
	for i := range list {
		if list[i].DamagerID == damagerID {
			list[i].Damage += amount
			list[i].Timestamp = time.Now()
			list[i].Damager = damager // Update handle in case it changed
			found = true
			break
		}
	}

	if !found {
		list = append(list, info)
	}

	damagers.Store(victimID, list)
}

// GetRecentDamagers returns the entities who damaged the victim within the last N seconds.
func GetRecentDamagers(victim world.Entity, within time.Duration) []world.Entity {
	victimID := getEntityID(victim)
	list, _ := damagers.Load(victimID)

	now := time.Now()
	var attackers []world.Entity
	for _, info := range list {
		if now.Sub(info.Timestamp) <= within {
			attackers = append(attackers, info.Damager)
		}
	}
	return attackers
}

// GetRecentPlayerDamagers returns only the players who damaged the victim within the time frame.
func GetRecentPlayerDamagers(victim world.Entity, within time.Duration) []*player.Player {
	victimID := getEntityID(victim)
	list, _ := damagers.Load(victimID)

	now := time.Now()
	var attackers []*player.Player
	for _, info := range list {
		if now.Sub(info.Timestamp) <= within {
			if p, ok := info.Damager.(*player.Player); ok {
				attackers = append(attackers, p)
			}
		}
	}
	return attackers
}

// GetTopDamagers returns the top N damagers sorted by damage amount within the time frame.
func GetTopDamagers(victim world.Entity, within time.Duration, limit int) []world.Entity {
	victimID := getEntityID(victim)
	list, _ := damagers.Load(victimID)

	now := time.Now()
	var validInfos []Info
	for _, info := range list {
		if now.Sub(info.Timestamp) <= within {
			validInfos = append(validInfos, info)
		}
	}

	// Sort by damage (descending)
	for i := 0; i < len(validInfos)-1; i++ {
		for j := i + 1; j < len(validInfos); j++ {
			if validInfos[j].Damage > validInfos[i].Damage {
				validInfos[i], validInfos[j] = validInfos[j], validInfos[i]
			}
		}
	}

	var result []world.Entity
	for i := 0; i < len(validInfos) && i < limit; i++ {
		result = append(result, validInfos[i].Damager)
	}
	return result
}

// GetTopPlayerDamagers returns the top N player damagers sorted by damage amount.
func GetTopPlayerDamagers(victim world.Entity, within time.Duration, limit int) []*player.Player {
	allDamagers := GetTopDamagers(victim, within, limit*2) // Get more in case some aren't players

	var players []*player.Player
	for _, damager := range allDamagers {
		if p, ok := damager.(*player.Player); ok {
			players = append(players, p)
			if len(players) >= limit {
				break
			}
		}
	}
	return players
}

// GetDamageContribution returns the percentage of damage dealt by a specific damager.
func GetDamageContribution(victim, damager world.Entity, within time.Duration) float64 {
	victimID := getEntityID(victim)
	damagerID := getEntityID(damager)
	list, _ := damagers.Load(victimID)

	now := time.Now()
	var totalDamage, entityDamage float64

	for _, info := range list {
		if now.Sub(info.Timestamp) <= within {
			totalDamage += info.Damage
			if info.DamagerID == damagerID {
				entityDamage = info.Damage
			}
		}
	}

	if totalDamage == 0 {
		return 0
	}
	return entityDamage / totalDamage
}

// GetLastDamager returns the last valid damager (for kill credit) within the given time frame.
func GetLastDamager(victim world.Entity, within time.Duration) (world.Entity, bool) {
	victimID := getEntityID(victim)
	list, _ := damagers.Load(victimID)

	now := time.Now()
	for i := len(list) - 1; i >= 0; i-- {
		if now.Sub(list[i].Timestamp) <= within {
			return list[i].Damager, true
		}
	}
	return nil, false
}

// GetLastPlayerDamager returns the last player damager within the time frame.
func GetLastPlayerDamager(victim world.Entity, within time.Duration) (*player.Player, bool) {
	victimID := getEntityID(victim)
	list, _ := damagers.Load(victimID)

	now := time.Now()
	for i := len(list) - 1; i >= 0; i-- {
		if now.Sub(list[i].Timestamp) <= within {
			if p, ok := list[i].Damager.(*player.Player); ok {
				return p, true
			}
		}
	}
	return nil, false
}

// HasMultipleDamagers returns true if more than one entity damaged the victim within the time frame.
func HasMultipleDamagers(victim world.Entity, within time.Duration) bool {
	return len(GetRecentDamagers(victim, within)) > 1
}

// HasMultiplePlayerDamagers returns true if more than one player damaged the victim.
func HasMultiplePlayerDamagers(victim world.Entity, within time.Duration) bool {
	return len(GetRecentPlayerDamagers(victim, within)) > 1
}

// ClearDamagers clears the damagers record for an entity (for example, on death).
func ClearDamagers(victim world.Entity) {
	damagers.Delete(getEntityID(victim))
}

// getEntityID returns a stable identifier for an entity.
// For players, it uses UUID. For other entities, it uses name() or a combination.
func getEntityID(entity world.Entity) string {
	if p, ok := entity.(*player.Player); ok {
		return p.UUID().String()
	}

	// For other entities, use their name and world position as identifier
	// This works because entities maintain consistent names during their lifetime
	if named, ok := entity.(interface{ Name() string }); ok {
		return named.Name() + "_" + uuid.New().String() // Unique per entity instance
	}

	// Fallback: generate a unique ID
	return uuid.New().String()
}
