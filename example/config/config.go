package config

type ExampleData struct {
	Mid    []float64              `json:"mid"`
	Spawns map[string][][]float64 `json:"spawns"`
}

func (m ExampleData) IsMapData() {}

type ExampleTeamData struct {
}

func (d ExampleTeamData) IsTeamData() {}

type ExampleParticipantData struct {
}

func (d ExampleParticipantData) IsParticipantData() {}
