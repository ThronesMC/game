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
func RefreshNameTag(tx *world.Tx, viewer, pt *participant.Participant) {
	g := game.GetGame()

	viewerEntity, ok := viewer.Handle().Entity(tx)
	if !ok {
		return
	}
	ptEntity, ok := pt.Handle().Entity(tx)
	if !ok {
		return
	}

	viewerPlayer, ok := viewerEntity.(*player.Player)
	if !ok {
		return
	}
	ptPlayer, ok := ptEntity.(*player.Player)
	if !ok {
		return
	}

	viewerSession := dfutils.Session(viewerPlayer)
	md := dfutils.ParseEntityMetadata(viewerSession, ptPlayer)

	md[protocol.EntityDataKeyName] = text.Colourf(g.Settings.NameFormat(viewer, pt))

	dfutils.WritePacket(viewerSession, &packet.SetActorData{
		EntityRuntimeID: dfutils.EntityRuntimeID(viewerSession, ptPlayer),
		EntityMetadata:  md,
	})
}
