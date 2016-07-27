package common

import (
	"bytes"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/dmportella/qilbot/bot"
	"github.com/dmportella/qilbot/utilities"
)

// New creates a new instance of Common Plugin.
func New() CommonPlugin {
	const (
		Name        = "Qilbot Common plugin"
		Description = "Common plugin for qibot a a place for generic commands."
	)
	return CommonPlugin{
		bot.Plugin{
			Name:        Name,
			Description: Description,
			Commands: []bot.CommandInformation{
				{
					Command:     "plugins",
					Template:    "plugins",
					Description: "Display a list of plugins enabled on qilbot.",
				},
				{
					Command:     "help",
					Template:    "help",
					Description: "Display a list of commands available to qilbot.",
				},
			},
		},
	}
}

// Initialize the init method for the common plugin
func (plugin *CommonPlugin) Initialize(qilbot *bot.Qilbot) {
	plugin.Qilbot = qilbot
	qilbot.AddHandler(plugin.messageCreate)
}

func (plugin *CommonPlugin) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if plugin.Plugin.Qilbot.IsBot(m.Author.ID) {
		return
	}

	matches := utilities.RegexMatchBotCommand(m.Content)

	if len(matches) >= 3 && plugin.Plugin.Qilbot.IsBot(matches[1]) {
		switch matches[2] {
		case "plugins":
			plugin.displayPluginList(s, m)
			break
		case "help":
			plugin.displayHelp(s, m)
			break
		default:
			return
		}
	}
}

func (plugin *CommonPlugin) displayPluginList(s *discordgo.Session, m *discordgo.MessageCreate) {
	var buffer bytes.Buffer
	for _, item := range plugin.Plugin.Qilbot.Plugins {
		buffer.WriteString(item.GetHelpText() + "\n")
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, buffer.String())
}

func (plugin *CommonPlugin) displayHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	var buffer bytes.Buffer
	buffer.WriteString("List of Commands available to Qilbot.\n")
	for _, item := range plugin.Plugin.Qilbot.Plugins {
		for _, command := range item.GetCommands() {
			buffer.WriteString(fmt.Sprintf("**%s** (%s): %s\n", command.Command, command.Template, command.Description))

		}
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, buffer.String())
}
