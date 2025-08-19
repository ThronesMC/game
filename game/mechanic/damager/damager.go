package damager

import (
	"github.com/ThronesMC/game/game/utils/maputils"
	"time"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/google/uuid"
)

var damagers = maputils.NewMap[uuid.UUID, []Info]()

// Info stores information about received damage.
type Info struct {
	Player    *player.Player
	Timestamp time.Time
}

// RegisterDamage registers that the victim was damaged by the damager.
func RegisterDamage(victim, damager *player.Player) {
	if victim == damager {
		return
	}

	victimUUID := victim.UUID()
	info := Info{Player: damager, Timestamp: time.Now()}
	list, _ := damagers.Load(victimUUID)

	// Remove recent duplicates
	for i := range list {
		if list[i].Player == damager {
			list = append(list[:i], list[i+1:]...)
			break
		}
	}

	list = append(list, info)
	damagers.Store(victimUUID, list)
}

// GetRecentDamagers returns the players who damaged the victim within the last N seconds.
func GetRecentDamagers(victim *player.Player, within time.Duration) []*player.Player {
	victimUUID := victim.UUID()
	list, _ := damagers.Load(victimUUID)

	now := time.Now()
	var attackers []*player.Player
	for _, info := range list {
		if now.Sub(info.Timestamp) <= within {
			attackers = append(attackers, info.Player)
		}
	}
	return attackers
}

// GetLastDamager returns the last valid damager (for kill credit) within the given time frame.
func GetLastDamager(victim *player.Player, within time.Duration) (*player.Player, bool) {
	victimUUID := victim.UUID()
	list, _ := damagers.Load(victimUUID)

	now := time.Now()
	for i := len(list) - 1; i >= 0; i-- {
		if now.Sub(list[i].Timestamp) <= within {
			return list[i].Player, true
		}
	}
	return nil, false
}

// ClearDamagers clears the damagers record for a player (for example, on death).
func ClearDamagers(victim *player.Player) {
	damagers.Delete(victim.UUID())
}
