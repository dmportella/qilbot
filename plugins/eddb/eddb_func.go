package eddb

import (
	"bytes"
	"github.com/dmportella/qilbot/bot"
)

const (
	name        = "Qilbot EDDB plugin"
	description = "A collection of tools and commands for EDDB."
)

// NewPlugin creates a new instance of WOW Plugin.
func NewPlugin(qilbot *bot.Qilbot) (plugin *Plugin) {
	plugin = &Plugin{}

	commodities := bot.QilbotCommand{
		Command:     "commodities",
		Template:    "commodities **item name**",
		Description: "Search the database for commodity information.",
		Execute: func(ctx *bot.QilbotCommandContext) {
			plugin.underConstructionCommand(ctx)
		},
	}

	stations := bot.QilbotCommand{
		Command:     "stations",
		Template:    "stations **system name**",
		Description: "Returns the stations available in specified system.",
		Execute: func(ctx *bot.QilbotCommandContext) {
			plugin.underConstructionCommand(ctx)
		},
	}

	modules := bot.QilbotCommand{
		Command:     "modules",
		Template:    "modules **module name**",
		Description: "Returns the information about specified module.",
		Execute: func(ctx *bot.QilbotCommandContext) {
			plugin.underConstructionCommand(ctx)
		},
	}

	bodies := bot.QilbotCommand{
		Command:     "bodies",
		Template:    "bodies **body name**",
		Description: "Returns the information about specified stellar body.",
		Execute: func(ctx *bot.QilbotCommandContext) {
			plugin.underConstructionCommand(ctx)
		},
	}

	qilbot.AddCommand(&commodities)
	qilbot.AddCommand(&stations)
	qilbot.AddCommand(&modules)
	qilbot.AddCommand(&bodies)

	return
}

// Name returns the name of the plugin
func (plugin *Plugin) Name() string {
	return name
}

// Description returns the description of the plugin
func (plugin *Plugin) Description() string {
	return description
}

func (plugin *Plugin) underConstructionCommand(ctx *bot.QilbotCommandContext) {
	var buffer bytes.Buffer

	buffer.WriteString("this command is currently under construction")

	ctx.RespondToUser(buffer.String())
}
