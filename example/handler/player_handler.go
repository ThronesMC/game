package handler

import (
	"github.com/ThronesMC/game/example/config"
	"github.com/ThronesMC/game/game"
	"github.com/ThronesMC/game/game/handler_custom"
	"github.com/ThronesMC/game/game/mechanic/cage"
	"github.com/ThronesMC/game/game/mechanic/spawn"
	"github.com/ThronesMC/game/game/utils/playerutils"
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type PlayerHandler struct {
	handler_custom.NopJoinHandler
}

func (PlayerHandler) HandleJoin(p *player.Player) {
	g := game.GetGame()

	if g.IsFull() {
		p.Disconnect("game is full.")
		return
	}

	pt := g.GetParticipant(p)
	if pt == nil {
		p.Disconnect("participant not found.")
		return
	}

	team, ok := g.BalancedAvailableTeam()
	if !ok {
		p.Disconnect("no available teams.")
		return
	}

	g.AssignTeamToParticipant(pt, team)

	spawnIndex := spawn.GetFreeSpawnIndex(team.GetID(), p.UUID())
	if spawnIndex == -1 {
		p.Disconnect("no free spawns available.")
		return
	}

	cfg := game.GetMapData[*config.ExampleData]()
	spawnPos := cfg.Spawns[team.GetID()][spawnIndex]
	blockPos := cube.Pos{int(spawnPos[0]), int(spawnPos[1]), int(spawnPos[2])}

	playerutils.ResetPlayer(p, nil)
	p.Teleport(blockPos.Vec3Centre())

	spawn.SetSpawnOccupied(team.GetID(), spawnIndex, p.UUID())
	cage.BuildCage(p.Tx(), p.UUID(), blockPos, block.Glass{})

	p.Messagef(text.Colourf("<orange>You joined team %s", team.GetColour().AsTextColour(team.GetName())))

	g.World.Exec(func(tx *world.Tx) {
		g.BroadcastMessagef(
			tx,
			"<yellow>%s</yellow> <green>has joined (<yellow>%d</yellow>/<yellow>%d</yellow>)!</green>",
			p.Name(), g.ParticipantLen(), g.Settings.Mode.MaximumTotalPlayers(),
		)
	})
}

func (PlayerHandler) HandleQuit(p *player.Player) {
	g := game.GetGame()

	pt := g.GetParticipant(p)
	if pt == nil {
		return
	}

	team := g.TeamOf(pt)
	if team == nil {
		return
	}

	spawn.FreePlayerSpawn(team.GetID(), p.UUID())
	g.RemoveFromTeam(pt)
	cage.RemoveCage(p.Tx(), p.UUID())

	g.World.Exec(func(tx *world.Tx) {
		g.BroadcastMessagef(
			tx,
			"<yellow>%s</yellow> <red>has left (<yellow>%d</yellow>/<yellow>%d</yellow>)!</red>",
			p.Name(), g.ParticipantLen()-1, g.Settings.Mode.MaximumTotalPlayers(),
		)
	})
}
