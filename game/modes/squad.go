package modes

type Squad struct {
}

func (s Squad) String() string {
	return "squad"
}

func (s Squad) MinimumTotalPlayers() int {
	return 8
}

func (s Squad) MaximumTotalPlayers() int {
	return 16
}

func (s Squad) NumberOfPlayersPerTeam() int {
	return 4
}
