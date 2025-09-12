package modes

type Duo struct {
}

func (d Duo) String() string {
	return "duo"
}

func (d Duo) MinimumTotalPlayers() int {
	return 4
}

func (d Duo) MaximumTotalPlayers() int {
	return 16
}

func (d Duo) NumberOfPlayersPerTeam() int {
	return 2
}
