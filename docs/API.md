# API Documentation

Simple API reference for the Game Framework.

## Table of Contents

- [Core Game API](#core-game-api)
- [Settings](#settings)
- [Teams](#teams)
- [Participants](#participants)
- [Game Modes](#game-modes)
- [Utilities](#utilities)

## Core Game API

### Game Structure

```go
type Game struct {
    id       uuid.UUID
    Settings *settings.Settings
    Teams    []*team.Team
    
    PlayerHandler    handler_custom.JoinHandler
    InventoryHandler inventory.Handler
    
    StateSeries  *state.ScheduledStateSeries
    Participants *maputils.Map[uuid.UUID, *participant.Participant]
    
    MapLoaded bool
    mapConfig config.MapData
    
    World       *world.World
    WorldFolder string
}
```

### Constructor

```go
func NewGame(settings *settings.Settings, teams []*team.Team, states []state.State, 
            playerHandler handler_custom.JoinHandler, invHandler inventory.Handler) *Game

func GetGame() *Game  // Get current game instance
```

### Key Methods

```go
// Game lifecycle
func (g *Game) Start() error
func (g *Game) Stop(tx *world.Tx)

// Map management
func (g *Game) LoadGameMapWithConfig(config config.MapData) error

// Player management
func (g *Game) Join(p *player.Player) error
func (g *Game) Quit(p *player.Player)
func (g *Game) GetParticipant(p *player.Player) *participant.Participant
func (g *Game) GetParticipants() iter.Seq[*participant.Participant]
func (g *Game) ParticipantLen() int
func (g *Game) HasEnoughPlayers() bool
func (g *Game) IsFull() bool

// Team operations
func (g *Game) InSameTeam(p1, p2 *participant.Participant) bool
func (g *Game) GetTeamByID(id string) *team.Team
func (g *Game) TeamOf(pt *participant.Participant) *team.Team
func (g *Game) AssignTeamToParticipant(pt *participant.Participant, team *team.Team)
func (g *Game) RemoveFromTeam(pt *participant.Participant)
func (g *Game) RandomAvailableTeam() (*team.Team, bool)
func (g *Game) BalancedAvailableTeam() (*team.Team, bool)
func (g *Game) EnemiesOf(pt *participant.Participant) []*participant.Participant

// Broadcasting
func (g *Game) BroadcastMessage(tx *world.Tx, msg string)
func (g *Game) BroadcastMessagef(tx *world.Tx, format string, args ...any)
func (g *Game) BroadcastTitle(tx *world.Tx, t title.Title)
func (g *Game) ParticipantsCallback(fn func(pt *participant.Participant))
```

## Settings

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

### Constructors

```go
func NewStaticSettings(gameName, mapsFolder, mapName string, nameFormat nameFormat) *Settings

func NewGameSettings(gameName, mapsFolder, mapName string, mode modes.Mode, nameFormat nameFormat) *Settings
```

### Name Format Function

```go
type nameFormat func(tx *world.Tx, viewer *player.Player, pt *participant.Participant) string
```

## Teams

### Team Structure

```go
type Team struct {
    id        string
    name      string
    color     TeamColour
    Teammates *maputils.Map[uuid.UUID, *participant.Participant]
    TeamData  config.TeamData
}
```

### Team Methods

```go
func NewTeam(id string, name string, color TeamColour, data config.TeamData) *Team

func (t *Team) GetID() string
func (t *Team) GetName() string
func (t *Team) GetColour() TeamColour
```

### Team Colors

```go
const (
    Black TeamColour = iota
    DarkBlue
    DarkGreen
    DarkAqua
    DarkRed
    DarkPurple
    Gold
    Grey
    DarkGrey
    Blue
    Green
    Aqua
    Red
    LightPurple
    Yellow
    White
)
```

## Participants

### Participant Structure

```go
type Participant struct {
    player          *player.Player
    handle          *world.EntityHandle
    state           State
    participantData config.ParticipantData
}
```

### Participant Methods

```go
func NewParticipant(player *player.Player, handle *world.EntityHandle) *Participant

func (pt *Participant) Player() *player.Player
func (pt *Participant) TXPlayer(tx *world.Tx) *player.Player
func (pt *Participant) Handle() *world.EntityHandle
func (pt *Participant) State() State
func (pt *Participant) SetState(s State)
func (pt *Participant) InState(s State) bool
func (pt *Participant) Data() config.ParticipantData
func (pt *Participant) SetData(data config.ParticipantData)
```

### Participant States

```go
var (
    StateAlive              = alive{}
    StateEliminated         = eliminated{}
    StateRespawning         = respawning{}
    StateTemporarySpectator = spectator{IsPermanent: false}
    StatePermanentSpectator = spectator{IsPermanent: true}
)
```

## Game Modes

### Mode Interface

```go
type Mode interface {
    String() string
    MinimumTotalPlayers() int
    MaximumTotalPlayers() int
    NumberOfPlayersPerTeam() int
}
```

### Available Modes

```go
// Solo: 1 player per team, 2-16 total players
type Solo struct{}

// Duo: 2 players per team, 4-16 total players  
type Duo struct{}

// Squad: 4 players per team, 8-16 total players
type Squad struct{}

// FFA: 1 player per team, configurable max players
type FFA struct {
    MaxTotalPlayers int
}

// Normal: Up to 100 players per team, 0-100 total players
type Normal struct{}
```

## Utilities

### Map Utils

Generic map implementation:

```go
type Map[K comparable, V any] struct {
    // Internal implementation
}

func NewMap[K comparable, V any]() *Map[K, V]
func (m *Map[K, V]) Store(key K, value V)
func (m *Map[K, V]) Load(key K) (V, bool)
func (m *Map[K, V]) Delete(key K)
func (m *Map[K, V]) Len() int
func (m *Map[K, V]) Map() map[K]V
```

### Bot Utils

```go
func GetBotNames(tx *world.Tx) []string
func AddBot(tx *world.Tx, pos mgl64.Vec3, rot cube.Rotation, s skin.Skin, onCreate func(p *player.Player))
func RemoveBot(tx *world.Tx, name string) bool
func RemoveAllBots(tx *world.Tx)
func IsBot(p *player.Player) bool
```

### Random Skins

```go
type Manager struct {
    // Skin management implementation
}

func (sm *Manager) GenerateSkin(id int) error
```

### Handler Utils

```go
func PlayerChainHandlers(handlers ...handler_custom.JoinHandler) handler_custom.JoinHandler
func WorldChainHandlers(handlers ...world.Handler) world.Handler
```