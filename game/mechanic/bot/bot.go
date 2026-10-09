package bot

import (
	"github.com/ThronesMC/game/game"
	"github.com/ThronesMC/game/game/utils/dfutils"
	"github.com/ThronesMC/game/game/utils/handlerutils"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/skin"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/npc"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/samber/lo"
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

	// npc.Create installs a handler of its own, which keeps a chunk loader
	// following the bot so it does not unload. Put the game's handlers in front
	// of that one rather than over it.
	//
	// Without the game's handlers a bot takes no part in the game's damage or
	// death handling at all: whatever a game does when a player dies - score it,
	// put them in spectator, respawn them - never happens, so a bot can be
	// killed over and over without ever dying.
	//
	// Its join handler is deliberately not run. A bot is placed by whatever
	// spawned it, which has already decided its skin, team and inventory, and
	// the join chain would undo those.
	if g := game.GetGame(); g != nil && g.PlayerHandler != nil {
		newBot.Handle(handlerutils.PlayerChainHandlers(joinable{Handler: newBot.Handler()}, g.PlayerHandler))
	}

	onCreate(newBot)
}

// joinable adapts a plain player.Handler to the JoinHandler the chain helpers
// take, for handlers that have nothing to do when a player joins.
type joinable struct {
	player.Handler
}

func (joinable) HandleJoin(*player.Player) {}

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
