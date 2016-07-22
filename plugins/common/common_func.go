package common

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dmportella/qilbot/bot"
	"github.com/dmportella/qilbot/utilities"
)

const (
	NAME        = "Qilbot Common plugin"
	DESCRIPTION = "Common plugin for qibot a a place for generic commands."
)

func New() CommonPlugin {
	return CommonPlugin{bot.Plugin{Name: NAME, Description: DESCRIPTION}}
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
			for _, plugin := range self.Plugin.Qilbot.Plugins {
				_, _ = s.ChannelMessageSend(m.ChannelID, plugin.GetHelpText())
			}

			break
		case "help":
			_, _ = s.ChannelMessageSend(m.ChannelID, "TODO: list commands in all plugins")
			break
		default:
			return
		}
	}
}
