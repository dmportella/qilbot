package common

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dmportella/qilbot/bot"
)

const (
	NAME        = "Qilbot Common plugin"
	DESCRIPTION = "Common plugin for qibot a a place for generic commands."
)

func New() CommonPlugin {
	return CommonPlugin{bot.Plugin{Name: NAME, Description: DESCRIPTION}}
}

func (self *CommonPlugin) Initialize(qilbot *bot.Qilbot) {
	//qilbot.AddEventHandler(self.messageCreate)
}

func (self *CommonPlugin) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

}
