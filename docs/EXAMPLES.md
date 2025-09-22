# Examples

Practical examples showing how to use the ThronesMC Game Framework effectively.

## Basic Game Setup

### Simple Main Function

```go
package main

import (
    "log/slog"
    "github.com/df-mc/dragonfly/server"
    "github.com/df-mc/dragonfly/server/player/chat"
    "github.com/df-mc/dragonfly/server/world"
    "github.com/ThronesMC/game/game"
    "github.com/ThronesMC/game/game/modes"
    "github.com/ThronesMC/game/game/settings"
    "github.com/ThronesMC/game/game/team"
    "github.com/ThronesMC/game/game/participant"
    "github.com/df-mc/dragonfly/server/player"
    "github.com/sandertv/gophertunnel/minecraft/text"
)

func main() {
    chat.Global.Subscribe(chat.StdoutSubscriber{})

    g := createGame()

    config := server.DefaultConfig()
    config.Server.DisableJoinQuitMessages = true
    config.World.Folder = "world"

    conf, err := config.Config(slog.Default())
    if err != nil {
        panic(err)
    }

    conf.ReadOnlyWorld = true
    conf.MaxPlayers = 20
    conf.Generator = func(dim world.Dimension) world.Generator {
        return world.NopGenerator{}
    }

    server.New(conf, slog.Default()).Listen()
}

func createGame() *game.Game {
    // Name formatting function
    nameFormat := func(tx *world.Tx, viewer *player.Player, pt *participant.Participant) string {
        return text.Colourf("<white>%v</white>", pt.TXPlayer(tx).Name())
    }

    // Game settings
    gameSettings := settings.NewGameSettings(
        "Test Game",
        "maps/",
        "TestMap",
        modes.Solo{},
        nameFormat,
    )

    // Teams
    teams := []*team.Team{
        team.NewTeam("players", "Players", team.White, nil),
    }

    // Create game
    g := game.NewGame(gameSettings, teams, nil, nil, nil)

    return g
}
```

## Game Modes

### Solo Mode

```go
func createSoloGame() *game.Game {
    nameFormat := func(tx *world.Tx, viewer *player.Player, pt *participant.Participant) string {
        return text.Colourf("<white>%v</white>", pt.TXPlayer(tx).Name())
    }

    settings := settings.NewGameSettings(
        "Solo Battle",
        "maps/",
        "SoloMap",
        modes.Solo{},
        nameFormat,
    )

    teams := []*team.Team{
        team.NewTeam("solo", "Solo Players", team.White, nil),
    }

    return game.NewGame(settings, teams, nil, nil, nil)
}
```

### Duo Mode

```go
func createDuoGame() *game.Game {
    nameFormat := func(tx *world.Tx, viewer *player.Player, pt *participant.Participant) string {
        g := game.GetGame()
        if vpt := g.GetParticipant(viewer); vpt != nil {
            if g.InSameTeam(vpt, pt) {
                return text.Colourf("<green>%v</green>", pt.TXPlayer(tx).Name())
            }
            return text.Colourf("<red>%v</red>", pt.TXPlayer(tx).Name())
        }
        return text.Colourf("<grey>[SPEC] %v</grey>", pt.TXPlayer(tx).Name())
    }

    settings := settings.NewGameSettings(
        "Duo Battle",
        "maps/",
        "DuoMap", 
        modes.Duo{},
        nameFormat,
    )

    teams := []*team.Team{
        team.NewTeam("red", "Red Team", team.Red, nil),
        team.NewTeam("blue", "Blue Team", team.Blue, nil),
        team.NewTeam("green", "Green Team", team.Green, nil),
        team.NewTeam("yellow", "Yellow Team", team.Yellow, nil),
    }

    return game.NewGame(settings, teams, nil, nil, nil)
}
```

### Squad Mode

```go
func createSquadGame() *game.Game {
    nameFormat := func(tx *world.Tx, viewer *player.Player, pt *participant.Participant) string {
        g := game.GetGame()
        if vpt := g.GetParticipant(viewer); vpt != nil {
            if g.InSameTeam(vpt, pt) {
                return text.Colourf("<green>%v</green>", pt.TXPlayer(tx).Name())
            }
            return text.Colourf("<red>%v</red>", pt.TXPlayer(tx).Name())
        }
        return text.Colourf("<grey>[SPEC] %v</grey>", pt.TXPlayer(tx).Name())
    }

    settings := settings.NewGameSettings(
        "Squad Battle",
        "maps/",
        "SquadMap",
        modes.Squad{},
        nameFormat,
    )

    teams := []*team.Team{
        team.NewTeam("red", "Red Squad", team.Red, nil),
        team.NewTeam("blue", "Blue Squad", team.Blue, nil),
        team.NewTeam("green", "Green Squad", team.Green, nil),
        team.NewTeam("yellow", "Yellow Squad", team.Yellow, nil),
    }

    return game.NewGame(settings, teams, nil, nil, nil)
}
```

### Free For All

```go  
func createFFAGame() *game.Game {
    nameFormat := func(tx *world.Tx, viewer *player.Player, pt *participant.Participant) string {
        if viewer.UUID() == pt.Player().UUID() {
            return text.Colourf("<green>%v</green>", pt.TXPlayer(tx).Name())
        }
        return text.Colourf("<red>%v</red>", pt.TXPlayer(tx).Name())
    }

    settings := settings.NewGameSettings(
        "Free For All",
        "maps/", 
        "FFAMap",
        modes.FFA{},
        nameFormat,
    )

    teams := []*team.Team{
        team.NewTeam("ffa", "All Players", team.White, nil),
    }

    return game.NewGame(settings, teams, nil, nil, nil)
}
```

## Map Configuration

### Basic Map Config

Create `maps/TestMap/config.json`:

```json
{
  "mid": [0, 64, 0, 0, 0],
  "spawns": {
    "players": [[0, 64, 0, 0, 0]]
  }
}
```

### Team-Based Map Config

Create `maps/TeamMap/config.json`:

```json
{
  "mid": [0, 64, 0, 0, 0],
  "spawns": {
    "red": [[20, 64, 0, 0, 0]],
    "blue": [[-20, 64, 0, 180, 0]],
    "green": [[0, 64, 20, 270, 0]],
    "yellow": [[0, 64, -20, 90, 0]]
  }
}
```

## Team Colors

Available team colors:

```go
teams := []*team.Team{
    team.NewTeam("red", "Red Team", team.Red, nil),
    team.NewTeam("blue", "Blue Team", team.Blue, nil),
    team.NewTeam("green", "Green Team", team.Green, nil),
    team.NewTeam("yellow", "Yellow Team", team.Yellow, nil),
    team.NewTeam("aqua", "Aqua Team", team.Aqua, nil),
    team.NewTeam("white", "White Team", team.White, nil),
    team.NewTeam("pink", "Pink Team", team.Pink, nil),
    team.NewTeam("grey", "Grey Team", team.Grey, nil),
}
```

## Project Structure

```
my-game/
├── main.go
├── go.mod
├── go.sum
├── maps/
│   └── TestMap/
│       ├── config.json
│       └── world.zip
└── world/ (generated)
```

## Running Your Game

1. Create the project structure
2. Add your map files  
3. Run: `go run main.go`
4. Connect with Minecraft Bedrock Edition to `localhost:19132`

## Complete Working Example

This example creates a simple working game:

```go
package main

import (
    "log/slog"
    "github.com/df-mc/dragonfly/server"
    "github.com/df-mc/dragonfly/server/player/chat"
    "github.com/df-mc/dragonfly/server/world"
    "github.com/ThronesMC/game/game"
    "github.com/ThronesMC/game/game/modes"
    "github.com/ThronesMC/game/game/settings"
    "github.com/ThronesMC/game/game/team"
    "github.com/ThronesMC/game/game/participant"
    "github.com/df-mc/dragonfly/server/player"
    "github.com/sandertv/gophertunnel/minecraft/text"
)

func main() {
    chat.Global.Subscribe(chat.StdoutSubscriber{})

    nameFormat := func(tx *world.Tx, viewer *player.Player, pt *participant.Participant) string {
        return text.Colourf("<white>%v</white>", pt.TXPlayer(tx).Name())
    }

    gameSettings := settings.NewGameSettings(
        "My Game",
        "maps/",
        "TestMap", 
        modes.Solo{},
        nameFormat,
    )

    teams := []*team.Team{
        team.NewTeam("players", "Players", team.White, nil),
    }

    g := game.NewGame(gameSettings, teams, nil, nil, nil)

    config := server.DefaultConfig()
    config.Server.DisableJoinQuitMessages = true
    config.World.Folder = "world"

    conf, err := config.Config(slog.Default())
    if err != nil {
        panic(err)
    }

    conf.ReadOnlyWorld = true
    conf.MaxPlayers = 20
    conf.Generator = func(dim world.Dimension) world.Generator {
        return world.NopGenerator{}
    }

    server.New(conf, slog.Default()).Listen()
}
```

This creates a basic server that players can join and play on.