package bot

import "github.com/bwmarrin/discordgo"

// Qilbot representation of the instance of qilbot.
type Qilbot struct {
	BotID   string
	config  *QilbotConfig
	session *discordgo.Session
	Plugins []IPlugin
}

// QilbotConfig representation of the configuration for qilbot.
type QilbotConfig struct {
	Email    string
	Password string
	Token    string
}
