package modes

type Solo struct {
}

func (s Solo) String() string {
	return "solo"
}

func (s Solo) MinimumTotalPlayers() int {
	return 2
}

func (s Solo) MaximumTotalPlayers() int {
	return 16
}

func (s Solo) NumberOfPlayersPerTeam() int {
	return 1
}
