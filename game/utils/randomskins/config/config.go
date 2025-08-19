package config

const (
	CurrentConfigVersion = "0.0.1"
	ConfigurationFile    = "config.json"
)

type File struct {
	ConfigVersion    string `json:"version"`
	EditableSkin     string `json:"edit_skin"`
	RandomizerFolder string `json:"randomizer_folder"`
	LastGeneration   string `json:"last_generation"`
	//TODO: implement that:
	GenerationConfig Generation `json:"gen_config"`
}

type Generation struct {
	BaseGenConfig     Part `json:"base"`
	HeadGenConfig     Part `json:"head"`
	BodyGenConfig     Part `json:"body"`
	LeftArmGenConfig  Part `json:"left_arm"`
	RightArmGenConfig Part `json:"right_arm"`
	LeftLegGenConfig  Part `json:"left_leg"`
	RightLegGenConfig Part `json:"right_leg"`
}

type Part struct {
	Layer1Enabled bool `json:"layer_0"`
	Layer2Enabled bool `json:"layer_1"`
}
