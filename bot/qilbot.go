package bot

import "github.com/bwmarrin/discordgo"

type Qilbot struct {
	BotID   string
	config  *QilbotConfig
	session *discordgo.Session
	plugins []*IPlugin
}

type QilbotConfig struct {
	Email    string
	Password string
	Token    string
}
