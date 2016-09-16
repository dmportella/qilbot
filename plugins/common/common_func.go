package common

import (
	"bytes"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/dmportella/qilbot/bot"
	"github.com/dmportella/qilbot/logging"
	"strings"
)

// NewPlugin creates a new instance of Common Plugin.
func NewPlugin(qilbot *bot.Qilbot) (plugin *Plugin) {
	const (
		Name        = "Qilbot Common plugin"
		Description = "Common plugin for qilbot a place for generic commands."
	)

	plugin = &Plugin{
		bot.Plugin{
			Qilbot:      qilbot,
			Name:        Name,
			Description: Description,
			Commands: []bot.CommandInformation{
				{
					Command:     "plugins",
					Template:    "plugins",
					Description: "Display a list of plugins enabled on qilbot.",
					Execute: func(s *bot.DiscordSession, m *discordgo.MessageCreate, commandText string) {
						plugin.pluginsCommand(s, m, commandText)
					},
				},
				{
					Command:     "help",
					Template:    "!help or !help *command*",
					Description: "Display a list of commands available to qilbot and more information about specific commands.",
					Execute: func(s *bot.DiscordSession, m *discordgo.MessageCreate, commandText string) {
						plugin.helpCommand(s, m, commandText)
					},
				},
				{
					Command:     "set",
					Template:    "!set *variable* *value*",
					Description: "Changes the settings of qilbot at runtime.",
					Execute: func(s *bot.DiscordSession, m *discordgo.MessageCreate, commandText string) {
						plugin.setCommand(s, m, commandText)
					},
				},
				{
					Command:     "get",
					Template:    "!get-variables",
					Description: "Returns a list of available settings on Qilbot.",
					Execute: func(s *bot.DiscordSession, m *discordgo.MessageCreate, commandText string) {
						plugin.getCommand(s, m, commandText)
					},
				},
			},
		},
	}

	qilbot.AddPlugin(plugin)

	qilbot.AddCommand(&plugin.Commands[0])
	qilbot.AddCommand(&plugin.Commands[1])
	qilbot.AddCommand(&plugin.Commands[2])

	return
}

func (plugin *Plugin) helpCommand(s *bot.DiscordSession, m *discordgo.MessageCreate, commandText string) {
	var buffer bytes.Buffer

	logging.Info.Println("comand text", commandText)

	specificCommand := commandText != ""

	for _, item := range plugin.Plugin.Qilbot.Plugins {
		for _, command := range item.GetCommands() {
			if strings.Compare(strings.ToLower(commandText), strings.ToLower(command.Command)) == 0 {
				buffer.WriteString(fmt.Sprintf("**%s** (%s): %s\n", command.Command, command.Template, command.Description))
			} else if !specificCommand {
				buffer.WriteString(fmt.Sprintf("**%s** (%s): %s\n", command.Command, command.Template, command.Description))
			}
		}
	}

	s.RespondToUser(m, buffer.String())
}

func (plugin *Plugin) pluginsCommand(s *bot.DiscordSession, m *discordgo.MessageCreate, commandText string) {
	var buffer bytes.Buffer

	for _, item := range plugin.Plugin.Qilbot.Plugins {
		buffer.WriteString(item.GetHelpText() + "\n")
	}

	s.RespondToUser(m, buffer.String())
}

func (plugin *Plugin) setCommand(s *bot.DiscordSession, m *discordgo.MessageCreate, commandText string) {
	var buffer bytes.Buffer

	if s.IsOwnerOfGuild(m) {
		buffer.WriteString("This would have done something.")
	} else {
		buffer.WriteString("Only the Server owner can change the bot settings...")
	}

	s.RespondToUser(m, buffer.String())
}

func (plugin *Plugin) getCommand(s *bot.DiscordSession, m *discordgo.MessageCreate, commandText string) {
	var buffer bytes.Buffer

	if s.IsOwnerOfGuild(m) {
		buffer.WriteString("This would have done something.")
	} else {
		buffer.WriteString("Only the Server owner can change the bot settings...")
	}

	s.RespondToUser(m, buffer.String())
}
