package bot

import (
	"github.com/bwmarrin/discordgo"
)

// Qilbot representation of the instance of qilbot.
type Qilbot struct {
	botID           string
	stopChannel     chan struct{}
	config          *QilbotConfig
	session         *discordgo.Session
	commands        map[string]*QilbotCommand
	commandSettings map[string]*commandSettings
}

// QilbotConfig representation of the configuration for qilbot.
type QilbotConfig struct {
	Email    string
	Password string
	Token    string
	Debug    bool
}
