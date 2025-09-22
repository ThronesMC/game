# Getting Started Guide

Simple guide to set up and run your first game.

## Prerequisites

- Go 1.25 or higher
- Basic Go knowledge

## Installation

```bash
mkdir my-game
cd my-game
go mod init my-game
go get github.com/ThronesMC/game
```

## Basic Setup

### 1. Main File

Create `main.go`:

```go
package main

import (
    "github.com/ThronesMC/game/game"
    "github.com/ThronesMC/game/game/modes"
    "github.com/ThronesMC/game/game/settings"
    "github.com/ThronesMC/game/game/team"
    "github.com/df-mc/dragonfly/server"
    "github.com/df-mc/dragonfly/server/player/chat"
    "github.com/df-mc/dragonfly/server/world"
    "log/slog"
)

func main() {
    chat.Global.Subscribe(chat.StdoutSubscriber{})
    
    // Create name format function
    nameFormat := func(tx *world.Tx, viewer *player.Player, pt *participant.Participant) string {
        return pt.TXPlayer(tx).Name()
    }
    
    // Create game settings
    gameSettings := settings.NewStaticSettings(
        "My Game",
        "maps/",
        "TestMap", 
        nameFormat,
    )
    
    // Create teams
    teams := []*team.Team{
        team.NewTeam("red", "Red", team.Red, &ExampleTeamConfig{}),
        team.NewTeam("blue", "Blue", team.Blue, &ExampleTeamConfig{}),
    }
    
    // Create game
    g := game.NewGame(gameSettings, teams, []state.State{}, 
                     &PlayerHandler{}, &InventoryHandler{})
    
    // Load map
    if err := g.LoadGameMapWithConfig(&MapConfig{}); err != nil {
        panic(err)
    }
    
    // Setup server
    config := server.DefaultConfig()
    config.World.Folder = g.WorldFolder
    
    conf, err := config.Config(slog.Default())
    if err != nil {
        panic(err)
    }
    
    conf.Generator = func(dim world.Dimension) world.Generator {
        return world.NopGenerator{}
    }
    
    srv := conf.New()
    srv.CloseOnProgramEnd()
    g.World = srv.World()
    
    if err := g.Start(); err != nil {
        panic(err)
    }
    
    srv.Listen()
    for p := range srv.Accept() {
        p.Handle(g.PlayerHandler)
        g.PlayerHandler.HandleJoin(p)
    }
}

type ExampleTeamConfig struct{}
func (ExampleTeamConfig) IsTeamData() {}

type MapConfig struct{}
func (MapConfig) IsMapData() {}

type PlayerHandler struct{}
func (p *PlayerHandler) HandleJoin(player *player.Player) {}
func (p *PlayerHandler) HandleLeave(player *player.Player) {}

type InventoryHandler struct{}
func (i *InventoryHandler) HandleTake(ctx context.Context, slot int, it item.Stack) bool { return true }
func (i *InventoryHandler) HandlePlace(ctx context.Context, slot int, it item.Stack) bool { return true }
```

### 2. Map Setup

Create directory structure:
```
maps/
└── TestMap/
    ├── config.json
    └── world.zip
```

Create `maps/TestMap/config.json`:
```json
{
  "mid": [0, 64, 0, 0, 0],
  "spawns": {
    "red": [[10, 64, 0, 0, 0]],
    "blue": [[-10, 64, 0, 180, 0]]
  }
}
```

You need to provide a `world.zip` file containing your Minecraft world.

### 3. Running

```bash
go run main.go
```

The server will start on port 19132. Connect with a Minecraft Bedrock client.

## Next Steps

- Study the example in the `example/` folder
- Check out the API documentation
- Explore different game modes
- Add custom handlers and mechanics