package dfutils

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	_ "unsafe"
)

//go:linkname Session github.com/df-mc/dragonfly/server/player.(*Player).session
func Session(_ *player.Player) *session.Session

//go:linkname WritePacket github.com/df-mc/dragonfly/server/session.(*Session).writePacket
func WritePacket(_ *session.Session, _ packet.Packet)

//go:linkname ParseEntityMetadata github.com/df-mc/dragonfly/server/session.(*Session).parseEntityMetadata
func ParseEntityMetadata(_ *session.Session, _ world.Entity) protocol.EntityMetadata

//go:linkname EntityRuntimeID github.com/df-mc/dragonfly/server/session.(*Session).entityRuntimeID
func EntityRuntimeID(_ *session.Session, _ world.Entity) uint64
