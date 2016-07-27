package common

import (
	"bytes"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/dmportella/qilbot/bot"
	"github.com/dmportella/qilbot/utilities"
)

const (
	NAME        = "Qilbot Common plugin"
	DESCRIPTION = "Common plugin for qibot a a place for generic commands."
)

func New() CommonPlugin {
	return CommonPlugin{
		bot.Plugin{
			Name:        NAME,
			Description: DESCRIPTION,
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

func (self *CommonPlugin) Initialize(qilbot *bot.Qilbot) {
	self.Qilbot = qilbot
	qilbot.AddHandler(self.messageCreate)
}

func (self *CommonPlugin) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if self.Plugin.Qilbot.IsBot(m.Author.ID) {
		return
	}

	matches := utilities.RegexMatchBotCommand(m.Content)

	if len(matches) >= 3 && self.Plugin.Qilbot.IsBot(matches[1]) {
		switch matches[2] {
		case "plugins":
			self.displayPluginList(s, m)
			break
		case "help":
			self.displayHelp(s, m)
			break
		default:
			return
		}
	}
}

func (self *CommonPlugin) displayPluginList(s *discordgo.Session, m *discordgo.MessageCreate) {
	var buffer bytes.Buffer
	for _, plugin := range self.Plugin.Qilbot.Plugins {
		buffer.WriteString(plugin.GetHelpText() + "\n")
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, buffer.String())
}

func (self *CommonPlugin) displayHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	var buffer bytes.Buffer
	buffer.WriteString("List of Commands available to Qilbot.\n")
	for _, plugin := range self.Plugin.Qilbot.Plugins {
		for _, command := range plugin.GetCommands() {
			buffer.WriteString(fmt.Sprintf("**%s** (%s): %s\n", command.Command, command.Template, command.Description))

		}
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, buffer.String())
}
