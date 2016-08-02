package bot

import (
	"github.com/bwmarrin/discordgo"
)

// Plugin describes a plugin in qilbot.
type Plugin struct {
	IPlugin
	ID          string
	Name        string
	Description string
	Qilbot      *Qilbot
	Commands    []CommandInformation
}

// CommandInformation describes a command in available for a plugin.
type CommandInformation struct {
	Command     string
	Template    string
	Description string
	Execute     func(s *discordgo.Session, m *discordgo.MessageCreate, commandText string)
}

// IPlugin interface used by qilbot.
type IPlugin interface {
	GetCommands() []CommandInformation
	GetHelpText() string
}
