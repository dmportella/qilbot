package bot

import (
	"github.com/bwmarrin/discordgo"
)

// QilbotCommandContext the context of the command call encapsulates all the plugin access to discord.
type QilbotCommandContext struct {
	Message        *discordgo.MessageCreate
	CommandText    string
	command        *QilbotCommand
	discordSession *discordgo.Session
}
