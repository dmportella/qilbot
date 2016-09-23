package wow

import (
	"bytes"
	"github.com/dmportella/qilbot/bot"
	"github.com/dmportella/qilbot/logging"
)

const (
	name        = "Qilbot World of Warcraft plugin"
	description = "A collection of tools and commands for the World of Warcraft API."
)

// NewPlugin creates a new instance of WOW Plugin.
func NewPlugin(qilbot *bot.Qilbot) (plugin *Plugin) {
	plugin = &Plugin{}

	armory := bot.QilbotCommand{
		Command:     "armory",
		Template:    "armory **item name**",
		Description: "Search the armory for items.",
		Execute: func(ctx *bot.QilbotCommandContext) {
			plugin.armoryCommand(ctx)
		},
	}

	qilbot.AddCommand(&armory)

	return
}

func (plugin *Plugin) Name() string {
	return name
}

func (plugin *Plugin) Description() string {
	return description
}

func (plugin *Plugin) armoryCommand(ctx *bot.QilbotCommandContext) {
	var buffer bytes.Buffer

	logging.Info.Println("comand text", ctx.CommandText)

	buffer.WriteString("This will give you armory link and all stuff you like")

	ctx.RespondToUser(buffer.String())
}
