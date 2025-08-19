package spawn

import (
	"github.com/ThronesMC/game/game/utils/maputils"
	"math/rand"

	"github.com/google/uuid"
)

var (
	occupiedSpawns = maputils.NewMap[string, map[int]uuid.UUID]()
	spawnConfig    = make(map[string][][]float64)
)

// InitSpawns initializes spawn points for each team.
func InitSpawns(config map[string][][]float64) {
	spawnConfig = config
	for teamID := range config {
		if _, ok := occupiedSpawns.Load(teamID); !ok {
			occupiedSpawns.Store(teamID, make(map[int]uuid.UUID))
		}
	}
}

// GetFreeSpawnIndex returns a free spawn index in the given team for a player.
func GetFreeSpawnIndex(teamID string, playerUUID uuid.UUID) int {
	if teamSpawns, ok := occupiedSpawns.Load(teamID); ok {
		// If the player already has a spawn assigned in this team
		for idx, u := range teamSpawns {
			if u == playerUUID {
				return idx
			}
		}

		// Collect free spawn indices
		var free []int
		spawnCount := len(spawnConfig[teamID])
		for i := 0; i < spawnCount; i++ {
			if _, ok := teamSpawns[i]; !ok {
				free = append(free, i)
			}
		}

		if len(free) == 0 {
			return -1
		}

		idx := free[rand.Intn(len(free))]
		teamSpawns[idx] = playerUUID
		return idx
	}
	return -1
}

// SetSpawnOccupied marks a spawn as occupied by a player.
func SetSpawnOccupied(teamID string, idx int, playerUUID uuid.UUID) {
	if teamSpawns, ok := occupiedSpawns.Load(teamID); ok {
		teamSpawns[idx] = playerUUID
	}
}

// FreePlayerSpawn frees the spawn occupied by a player in a team.
func FreePlayerSpawn(teamID string, playerUUID uuid.UUID) {
	if teamSpawns, ok := occupiedSpawns.Load(teamID); ok {
		for idx, v := range teamSpawns {
			if v == playerUUID {
				delete(teamSpawns, idx)
				break
			}
		}
	}
}
