package modes

type FFA struct {
	MaxTotalPlayers int
}

func (f FFA) String() string {
	return "ffa"
}

func (f FFA) MinimumTotalPlayers() int {
	return 2
}

func (f FFA) MaximumTotalPlayers() int {
	return f.MaxTotalPlayers
}

func (f FFA) NumberOfPlayersPerTeam() int {
	return 1
}
