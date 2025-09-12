package modes

type Normal struct {
}

func (n Normal) String() string {
	return "normal"
}

func (n Normal) MinimumTotalPlayers() int {
	return 0
}

func (n Normal) MaximumTotalPlayers() int {
	return 100
}

func (n Normal) NumberOfPlayersPerTeam() int {
	return 100
}
