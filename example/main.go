package main

import (
	"github.com/ThronesMC/game/game/command"
	"github.com/ThronesMC/game/game/handler_custom"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/world"
	"log/slog"
)

func main() {
	chat.Global.Subscribe(chat.StdoutSubscriber{})

	var g = NewExampleGame()

	srvConf := server.DefaultConfig()
	srvConf.Server.DisableJoinQuitMessages = true
	srvConf.Server.MuteEmoteChat = true
	srvConf.World.Folder = g.WorldFolder

	conf, err := srvConf.Config(slog.Default())
	if err != nil {
		panic(err)
	}

	conf.ReadOnlyWorld = true
	conf.DisableResourceBuilding = true
	conf.MaxPlayers = 100
	conf.Generator = func(dim world.Dimension) world.Generator {
		return world.NopGenerator{}
	}

	srv := conf.New()

	w := srv.World()
	w.StopWeatherCycle()
	w.StopRaining()
	w.StopThundering()
	w.SetTime(3000)
	w.StopTime()
	g.World = w

	command.RegisterDevCommands()

	if err := g.Start(); err != nil {
		panic(err)
	}

	srv.CloseOnProgramEnd()
	srv.Listen()

	for p := range srv.Accept() {
		p.Handle(g.PlayerHandler)
		p.Handler().(handler_custom.JoinHandler).HandleJoin(p)
	}
}
