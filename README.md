# Game Framework

A comprehensive game framework built on top of Dragonfly for creating multiplayer Minecraft games. This library provides a complete architecture for team-based games, player management, world handling, and game state management.

## Features

- 🎮 **Multiple Game Modes**: Support for Solo, Duo, Squad, FFA (Free for All), and Normal team modes
- 👥 **Team Management**: Complete team system with configurable teams and colors
- 🤖 **Bot Support**: Built-in bot system for testing and gameplay enhancement
- 🗺️ **Map System**: Flexible map loading and configuration system
- 🎯 **State Management**: Finite state machine for game phases (pregame, gameplay, endgame)
- ⚙️ **Configurable Settings**: Extensive configuration options for games and maps
- 🏷️ **Participant System**: Advanced player management with custom name formatting
- 🛠️ **Handler System**: Modular event handling for players, inventory, and world events
- 📦 **Utility Libraries**: Rich set of utilities for common game operations

## Installation

```bash
go get github.com/ThronesMC/game
```

### Prerequisites

- Go 1.25 or higher
- Dragonfly server framework

## Quick Start

Here's a simple example of creating a game:

```go
package main

import (
    "github.com/ThronesMC/game/game"
    "github.com/ThronesMC/game/game/modes"
    "github.com/ThronesMC/game/game/settings"
    "github.com/ThronesMC/game/game/team"
)

func main() {
    // Create game settings
    gameSettings := settings.NewGameSettings(
        "My Game",           // Game name
        "maps/",            // Maps folder
        "MyMap",            // Map name
        modes.Solo{},       // Game mode
        nameFormatFunc,     // Name formatting function
    )

    // Define teams
    teams := []*team.Team{
        team.NewTeam("red", "Red Team", team.Red, &teamConfig),
        team.NewTeam("blue", "Blue Team", team.Blue, &teamConfig),
    }

    // Create game states
    states := []state.State{
        states.NewPreGameState(),
        states.NewGameState(),
        states.NewEndGameState(),
    }

    // Initialize game
    g := game.NewGame(
        gameSettings,
        teams,
        states,
        playerHandler,
        inventoryHandler,
    )

    // Load map
    if err := g.LoadGameMapWithConfig(&mapConfig); err != nil {
        panic(err)
    }

    // Start game
    g.Start()
}
```

## Architecture Overview

The framework is built around several core components:

### Core Components

- **Game**: The main game instance that orchestrates all components
- **Settings**: Configuration management for games and maps
- **Teams**: Team management system with configurable properties
- **Participants**: Player management with extended functionality
- **Modes**: Different game mode implementations (Solo, Duo, Squad, etc.)
- **State Machine**: Game phase management (pregame, gameplay, endgame)

### Utility Systems

- **Handlers**: Event handling for players, inventory, and world interactions
- **Mechanics**: Game mechanics like bots, cages, damage systems, spawning, voting
- **Utils**: Helper libraries for common operations (maps, players, random skins, etc.)

## Game Modes

The framework supports various game modes out of the box:

- **Solo**: Individual player competition
- **Duo**: Two-player teams
- **Squad**: Larger team-based gameplay
- **FFA (Free for All)**: Every player for themselves
- **Normal**: Traditional team-based gameplay

Each mode defines minimum and maximum player counts and team configurations.

## Map System

The framework includes a flexible map system:

- **Map Loading**: Automatic world loading from ZIP files
- **Configuration**: JSON-based map configuration
- **Multiple Maps**: Support for multiple map variants
- **Custom Settings**: Per-map custom configuration options

## Documentation

- [Getting Started Guide](docs/GETTING_STARTED.md) - Step-by-step setup instructions
- [API Documentation](docs/API.md) - Detailed API reference
- [Configuration Guide](docs/CONFIG.md) - Settings and configuration options
- [Examples](docs/EXAMPLES.md) - Code examples and use cases

## Project Structure

```
game/
├── game/                    # Core game engine
│   ├── command/            # Game commands
│   ├── config/             # Configuration types
│   ├── handler_custom/     # Custom event handlers
│   ├── mechanic/           # Game mechanics
│   ├── modes/              # Game mode implementations
│   ├── participant/        # Player management
│   ├── settings/           # Settings management
│   ├── team/               # Team system
│   └── utils/              # Utility libraries
├── example/                # Example implementation
├── maps/                   # Map files
└── skins/                  # Skin configurations
```

## Contributing

Contributions are welcome! Please feel free to submit pull requests, create issues, or suggest improvements.

### Development Setup

1. Clone the repository
2. Install dependencies: `go mod download`
3. Run the example: `go run example/main.go`
4. Make your changes and test them

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

For support, questions, or discussions:

- Create an issue on GitHub
- Join our Discord community
- Check the documentation in the `docs/` folder

## Credits

Built on top of the excellent [Dragonfly](https://github.com/df-mc/dragonfly) Minecraft server framework.