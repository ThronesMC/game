# Configuration Guide

Simple configuration guide for the Game Framework.

## Table of Contents

- [Game Settings](#game-settings)
- [Map Configuration](#map-configuration)
- [Team Configuration](#team-configuration)
- [Skin Configuration](#skin-configuration)

## Game Settings

### Settings Structure

```go
type Settings struct {
    GameName   string
    MapsFolder string
    MapName    string
    Mode       modes.Mode
    NameFormat nameFormat
}
```

### Basic Settings

```go
// Static settings (no player limits)
settings := settings.NewStaticSettings(
    "My Game",        // Game name
    "maps/",          // Maps folder
    "MyMap",          // Map name
    nameFormatFunc,   // Name formatting function
)

// Game settings with modes and limits
settings := settings.NewGameSettings(
    "My Game",        // Game name
    "maps/",          // Maps folder
    "MyMap",          // Map name
    modes.Solo{},     // Game mode
    nameFormatFunc,   // Name formatting function
)
```

### Available Game Modes

```go
// Solo: 1 player per team, 2-16 total players
modes.Solo{}

// Duo: 2 players per team, 4-16 total players  
modes.Duo{}

// Squad: 4 players per team, 8-16 total players
modes.Squad{}

// FFA: 1 player per team, configurable max players
modes.FFA{MaxTotalPlayers: 20}

// Normal: Up to 100 players per team, 0-100 total players
modes.Normal{}
```

### Name Formatting

```go
func nameFormatFunc(tx *world.Tx, viewer *player.Player, pt *participant.Participant) string {
    g := game.GetGame()
    if vpt := g.GetParticipant(viewer); vpt != nil {
        if g.InSameTeam(vpt, pt) {
            return text.Colourf("<green>%v</green>", pt.TXPlayer(tx).Name())
        }
        return text.Colourf("<red>%v</red>", pt.TXPlayer(tx).Name())
    }
    return text.Colourf("<grey>[SPECTATOR] %v</grey>", pt.TXPlayer(tx).Name())
}
```

## Map Configuration

Maps use a simple folder structure with JSON configuration.

### Map Structure

```
maps/
└── YourMapName/
    ├── config.json    # Map configuration
    └── world.zip      # World data
```

### Configuration Format

Example `config.json`:

```json
{
  "mid": [0, 100, 0, 0, 0],
  "spawns": {
    "red": [
      [20, 100, 0, 90, 0],
      [22, 100, 2, 90, 0]
    ],
    "blue": [
      [-20, 100, 0, -90, 0],
      [-22, 100, -2, -90, 0]
    ],
    "green": [
      [0, 100, 20, 180, 0],
      [2, 100, 22, 180, 0]
    ],
    "yellow": [
      [0, 100, -20, 0, 0],
      [-2, 100, -22, 0, 0]
    ]
  }
}
```

### Spawn Format

Each spawn point: `[x, y, z, yaw, pitch]`

- **x, y, z**: World coordinates
- **yaw**: Horizontal rotation (degrees)
- **pitch**: Vertical rotation (degrees)

## Team Configuration

### Team Creation

```go
teams := []*team.Team{
    team.NewTeam("red", "Red Team", team.Red, &config.TeamData{}),
    team.NewTeam("blue", "Blue Team", team.Blue, &config.TeamData{}),
}
```

### Team Colors

Available colors:
- `team.Red`
- `team.Green` 
- `team.Blue`
- `team.Yellow`
- `team.Purple`
- `team.Orange`
- `team.Pink`
- `team.Cyan`
- `team.White`
- `team.Black`

## Skin Configuration

The framework supports random skin generation for players.

### Skin Config File

Create `skins/config.json`:

```json
{
    "version": "0.0.1",
    "edit_skin": "./skins",
    "randomizer_folder": "./skins/skin_template", 
    "last_generation": "2024-10-27T12:00:00Z",
    "gen_config": {
        "base": {
            "layer_0": true,
            "layer_1": true
        },
        "head": {
            "layer_0": true,
            "layer_1": true
        },
        "body": {
            "layer_0": true,
            "layer_1": true
        },
        "left_arm": {
            "layer_0": true,
            "layer_1": true
        },
        "right_arm": {
            "layer_0": true,
            "layer_1": true
        },
        "left_leg": {
            "layer_0": true,
            "layer_1": true
        },
        "right_leg": {
            "layer_0": true,
            "layer_1": true
        }
    }
}
```

### Skin Generation

The skin manager generates random skins by combining different parts:

- **Base**: Base skin layer
- **Head**: Head customizations
- **Body**: Body parts
- **Arms**: Left and right arm parts
- **Legs**: Left and right leg parts

Each part has two layers (layer_0 and layer_1) that can be enabled or disabled in the configuration.