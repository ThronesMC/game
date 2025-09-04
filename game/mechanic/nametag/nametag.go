package nametag

import (
	"github.com/ThronesMC/game/game"
	"github.com/ThronesMC/game/game/participant"
	"github.com/ThronesMC/game/game/utils/dfutils"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

// RefreshNameTag updates how `pt`’s name tag is displayed to `viewer`,
// applying teammate or enemy formatting based on their team relationship
// and sending the updated metadata to the viewer’s client.
func RefreshNameTag(tx *world.Tx, viewer *player.Player, pt *participant.Participant) {
	g := game.GetGame()

	viewerSession := dfutils.Session(viewer)
	md := dfutils.ParseEntityMetadata(viewerSession, pt.TXPlayer(tx))

	md[protocol.EntityDataKeyName] = text.Colourf(g.Settings.NameFormat(tx, viewer, pt))

	dfutils.WritePacket(viewerSession, &packet.SetActorData{
		EntityRuntimeID: dfutils.EntityRuntimeID(viewerSession, pt.TXPlayer(tx)),
		EntityMetadata:  md,
	})
}
