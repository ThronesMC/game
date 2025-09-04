package cage

import (
	"github.com/ThronesMC/game/game/utils/maputils"
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/google/uuid"
)

var cages = maputils.NewMap[uuid.UUID, []cube.Pos]()

// BuildCage creates a cage at the given position for a player.
func BuildCage(tx *world.Tx, playerUUID uuid.UUID, pos cube.Pos, cageBlock world.Block) {
	var blocks []cube.Pos
	for i := 0; i < 3; i++ {
		blocks = append(blocks, pos.Add(cube.Pos{-1, i, 0}))
		blocks = append(blocks, pos.Add(cube.Pos{1, i, 0}))
		blocks = append(blocks, pos.Add(cube.Pos{0, i, -1}))
		blocks = append(blocks, pos.Add(cube.Pos{0, i, 1}))
	}
	blocks = append(blocks, pos.Add(cube.Pos{0, -1, 0})) // floor
	blocks = append(blocks, pos.Add(cube.Pos{0, 3, 0}))  // ceiling

	for _, b := range blocks {
		tx.SetBlock(b, cageBlock, nil)
	}

	cages.Store(playerUUID, blocks)
}

// RemoveCage removes the cage for a specific player.
func RemoveCage(tx *world.Tx, playerUUID uuid.UUID) {
	if blocks, ok := cages.Load(playerUUID); ok {
		for _, b := range blocks {
			tx.SetBlock(b, block.Air{}, nil)
		}
		cages.Delete(playerUUID)
	}
}

// RemoveAllCages removes all cages for all players.
func RemoveAllCages(tx *world.Tx) {
	for playerUUID, blocks := range cages.Map() {
		for _, b := range blocks {
			tx.SetBlock(b, block.Air{}, nil)
		}
		cages.Delete(playerUUID)
	}
}
