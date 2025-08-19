package handler

import (
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/thronesmc/game/example/config"
	"github.com/thronesmc/game/game"
	"github.com/thronesmc/game/game/handler_custom"
	"github.com/thronesmc/game/game/mechanic/cage"
	"github.com/thronesmc/game/game/mechanic/spawn"
	"github.com/thronesmc/game/game/utils/playerutils"
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

	pt, ok := g.GetParticipant(p)
	if !ok {
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

	g.BroadcastMessagef(
		"<yellow>%s</yellow> <green>has joined (<yellow>%d</yellow>/<yellow>%d</yellow>)!</green>",
		p.Name(), g.ParticipantLen(), g.Settings.MaxPlayers,
	)
}

func (PlayerHandler) HandleQuit(p *player.Player) {
	g := game.GetGame()

	pt, ok := g.GetParticipant(p)
	if !ok {
		return
	}

	team, ok := g.TeamOf(pt)
	if !ok {
		return
	}

	spawn.FreePlayerSpawn(team.GetID(), p.UUID())
	g.RemoveFromTeam(pt)
	cage.RemoveCage(p.Tx(), p.UUID())

	g.BroadcastMessagef(
		"<yellow>%s</yellow> <red>has left (<yellow>%d</yellow>/<yellow>%d</yellow>)!</red>",
		p.Name(), g.ParticipantLen()-1, g.Settings.MaxPlayers,
	)
}
