package team

import (
	"fmt"
	"github.com/ThronesMC/game/game/config"
	"github.com/ThronesMC/game/game/participant"
	"github.com/ThronesMC/game/game/utils/maputils"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/google/uuid"
)

type Team struct {
	id    string
	name  string
	color TeamColour

	Teammates *maputils.Map[uuid.UUID, *participant.Participant]

	teamData config.TeamData
}

func NewTeam(id string, name string, color TeamColour, data config.TeamData) *Team {
	return &Team{
		id:        id,
		name:      name,
		color:     color,
		Teammates: maputils.NewMap[uuid.UUID, *participant.Participant](),
		teamData:  data,
	}
}

func (t *Team) GetID() string {
	return t.id
}

func (t *Team) GetName() string {
	return t.name
}

func (t *Team) GetColour() TeamColour {
	return t.color
}

type TeamColour int

func (c TeamColour) AsTextColour(text string) string {
	switch c {
	case Black:
		return fmt.Sprintf("<black>%s</black>", text)
	case DarkBlue:
		return fmt.Sprintf("<dark_blue>%s</dark_blue>", text)
	case DarkGreen:
		return fmt.Sprintf("<dark_green>%s</dark_green>", text)
	case DarkAqua:
		return fmt.Sprintf("<dark_aqua>%s</dark_aqua>", text)
	case DarkRed:
		return fmt.Sprintf("<dark_red>%s</dark_red>", text)
	case DarkPurple:
		return fmt.Sprintf("<dark_purple>%s</dark_purple>", text)
	case Gold:
		return fmt.Sprintf("<gold>%s</gold>", text)
	case Grey:
		return fmt.Sprintf("<gray>%s</gray>", text)
	case DarkGrey:
		return fmt.Sprintf("<dark_gray>%s</dark_gray>", text)
	case Blue:
		return fmt.Sprintf("<blue>%s</blue>", text)
	case Green:
		return fmt.Sprintf("<green>%s</green>", text)
	case Aqua:
		return fmt.Sprintf("<aqua>%s</aqua>", text)
	case Red:
		return fmt.Sprintf("<red>%s</red>", text)
	case Purple:
		return fmt.Sprintf("<light_purple>%s</light_purple>", text)
	case Yellow:
		return fmt.Sprintf("<yellow>%s</yellow>", text)
	case White:
		return fmt.Sprintf("<white>%s</white>", text)
	case DarkYellow:
		return fmt.Sprintf("<dark_yellow>%s</dark_yellow>", text)
	case Quartz:
		return fmt.Sprintf("<quartz>%s</quartz>", text)
	case Iron:
		return fmt.Sprintf("<iron>%s</iron>", text)
	case Netherite:
		return fmt.Sprintf("<netherite>%s</netherite>", text)
	case Redstone:
		return fmt.Sprintf("<redstone>%s</redstone>", text)
	case Copper:
		return fmt.Sprintf("<copper>%s</copper>", text)
	case Emerald:
		return fmt.Sprintf("<emerald>%s</emerald>", text)
	case Diamond:
		return fmt.Sprintf("<diamond>%s</diamond>", text)
	case Lapis:
		return fmt.Sprintf("<lapis>%s</lapis>", text)
	case Amethyst:
		return fmt.Sprintf("<amethyst>%s</amethyst>", text)
	default:
		return text
	}
}

func (c TeamColour) AsItemColour() item.Colour {
	switch c {
	case Black:
		return item.ColourBlack()
	case DarkBlue:
		return item.ColourBlue()
	case DarkGreen:
		return item.ColourGreen()
	case DarkAqua:
		return item.ColourCyan()
	case DarkRed:
		return item.ColourRed()
	case DarkPurple:
		return item.ColourPurple()
	case Gold:
		return item.ColourOrange()
	case Grey:
		return item.ColourLightGrey()
	case DarkGrey:
		return item.ColourGrey()
	case Blue:
		return item.ColourBlue()
	case Green:
		return item.ColourLime()
	case Aqua:
		return item.ColourCyan()
	case Red:
		return item.ColourRed()
	case Purple:
		return item.ColourMagenta()
	case Yellow:
		return item.ColourYellow()
	case White:
		return item.ColourWhite()
	case DarkYellow:
		return item.ColourYellow()
	case Quartz:
		return item.ColourLightGrey()
	case Iron:
		return item.ColourGrey()
	case Netherite:
		return item.ColourBrown()
	case Redstone:
		return item.ColourRed()
	case Copper:
		return item.ColourOrange()
	case Emerald:
		return item.ColourLime()
	case Diamond:
		return item.ColourLightBlue()
	case Lapis:
		return item.ColourBlue()
	case Amethyst:
		return item.ColourPurple()
	default:
		return item.ColourWhite()
	}
}

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
	Purple
	Yellow
	White
	DarkYellow
	Quartz
	Iron
	Netherite
	Redstone
	Copper
	Emerald
	Diamond
	Lapis
	Amethyst
)

func GetTeamData[T config.TeamData](t *Team) T {
	return t.teamData.(T)
}
