package bot

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/skin"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/npc"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/samber/lo"
	"github.com/thronesmc/game/game/utils/dfutils"
)

func GetBotNames(tx *world.Tx) []string {
	names := make([]string, 0)
	for e := range tx.Entities() {
		if p, ok := e.(*player.Player); ok && IsBot(p) {
			names = append(names, p.Name())
		}
	}
	return names
}

func generateBotName() string {
	return "Bot-" + lo.RandomString(5, lo.AlphanumericCharset)
}

func AddBot(tx *world.Tx, pos mgl64.Vec3, rot cube.Rotation, s skin.Skin, onCreate func(p *player.Player)) {
	settings := npc.Settings{
		Name:       generateBotName(),
		Scale:      1,
		Position:   pos,
		Skin:       s,
		Yaw:        rot.Yaw(),
		Pitch:      rot.Pitch(),
		Immobile:   false,
		Vulnerable: true,
	}

	newBot := npc.Create(settings, tx, nil)

	onCreate(newBot)
}

func RemoveBot(tx *world.Tx, name string) bool {
	for e := range tx.Entities() {
		if p, ok := e.(*player.Player); ok && IsBot(p) && p.Name() == name {
			_ = p.Close()

			return true
		}
	}
	return false
}

func RemoveAllBots(tx *world.Tx) {
	for e := range tx.Entities() {
		if p, ok := e.(*player.Player); ok && IsBot(p) {
			_ = p.Close()
		}
	}
}

func IsBot(p *player.Player) bool {
	return dfutils.Session(p) == session.Nop
}
