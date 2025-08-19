package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type LocationCommand struct{}

func (lc LocationCommand) Run(source cmd.Source, output *cmd.Output, _ *world.Tx) {
	p, ok := source.(*player.Player)
	if !ok {
		output.Error("Must be a player")
		return
	}

	pos := p.Position()
	rot := p.Rotation()

	output.Print(text.Colourf(
		"<aqua>Position: </aqua><grey>%.2f, %.2f, %.2f</grey>\n<aqua>Rotation: </aqua><grey>%.2f, %.2f</grey>",
		pos.X(), pos.Y(), pos.Z(), rot.Yaw(), rot.Pitch(),
	))
}
