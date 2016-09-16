package wow

import (
	"bytes"
	"github.com/bwmarrin/discordgo"
	"github.com/dmportella/qilbot/bot"
	"github.com/dmportella/qilbot/logging"
)

// NewPlugin creates a new instance of WOW Plugin.
func NewPlugin(qilbot *bot.Qilbot) (plugin *Plugin) {
	const (
		Name        = "Qilbot World of Warcraft plugin"
		Description = "A collection of tools and commands for the World of Warcraft API."
	)

	plugin = &Plugin{
		bot.Plugin{
			Qilbot:      qilbot,
			Name:        Name,
			Description: Description,
			Commands: []bot.CommandInformation{
				{
					Command:     "armory",
					Template:    "armory **item name**",
					Description: "Search the armory for items.",
					Execute: func(s *bot.DiscordSession, m *discordgo.MessageCreate, commandText string) {
						plugin.armoryCommand(s, m, commandText)
					},
				},
			},
		},
	}

	qilbot.AddPlugin(plugin)

	qilbot.AddCommand(&plugin.Commands[0])

	return
}

func (plugin *Plugin) armoryCommand(s *bot.DiscordSession, m *discordgo.MessageCreate, commandText string) {
	var buffer bytes.Buffer

	logging.Info.Println("comand text", commandText)

	buffer.WriteString("This will give you armory link and all stuff you like")

	s.RespondToUser(m, buffer.String())
}
