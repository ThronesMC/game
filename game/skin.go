package game

import (
	"github.com/ThronesMC/game/game/utils/randomskins/skin"
	"log"
)

var SkinManager skin.Manager

func init() {
	skinManager, err := skin.SetupSkinManager()
	if err != nil {
		log.Fatalf("Could not set up a new skin manager: %v", err)
	}

	SkinManager = skinManager
}
