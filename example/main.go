package main

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/thronesmc/game/game/command"
	"github.com/thronesmc/game/game/handler"
	"github.com/thronesmc/game/game/handler_custom"
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

	w.Handle(handler.GlobalWorldHandler{})

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
