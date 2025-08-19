package bots

import (
	"github.com/ThronesMC/game/game/mechanic/bot"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type RemoveSubCommand struct {
	Remove  cmd.SubCommand `cmd:"remove"`
	BotName bots           `cmd:"botName"`
}

func (rsc RemoveSubCommand) Run(_ cmd.Source, output *cmd.Output, tx *world.Tx) {
	botName := string(rsc.BotName)

	if botName == "all" {
		bot.RemoveAllBots(tx)
		output.Print(text.Colourf("<green>All bots removed successfully.</green>"))
		return
	}

	if ok := bot.RemoveBot(tx, botName); ok {
		output.Print(text.Colourf("<green>Bot %s removed successfully.</green>", botName))
		return
	}

	output.Print(text.Colourf("<yellow>No bot found with the name '%s'.</yellow>", botName))
}

type bots string

func (bots) Type() string {
	return "botName"
}

func (bots) Options(source cmd.Source) []string {
	p, ok := source.(*player.Player)
	if !ok {
		return []string{}
	}

	names := bot.GetBotNames(p.Tx())
	return append([]string{"all"}, names...)
}
