package modes

type Mode interface {
	String() string
	MinimumTotalPlayers() int
	MaximumTotalPlayers() int
	NumberOfPlayersPerTeam() int
}
