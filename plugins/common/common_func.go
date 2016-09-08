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
					Template:    "!set *variable* on|off",
					Description: "Changes the settings of qilbot at runtime.",
					Execute: func(s *bot.DiscordSession, m *discordgo.MessageCreate, commandText string) {
						plugin.setCommand(s, m, commandText)
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

	channel, _ := s.UserChannelCreate(m.Author.ID)

	logging.Trace.Println(channel)

	_, _ = s.ChannelMessageSend(channel.ID, buffer.String())
}

func (plugin *Plugin) pluginsCommand(s *bot.DiscordSession, m *discordgo.MessageCreate, commandText string) {
	var buffer bytes.Buffer
	for _, item := range plugin.Plugin.Qilbot.Plugins {
		buffer.WriteString(item.GetHelpText() + "\n")
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, buffer.String())
}

func (plugin *Plugin) setCommand(s *bot.DiscordSession, m *discordgo.MessageCreate, commandText string) {
	var buffer bytes.Buffer

	if s.IsOwnerOfGuild(m) {
		buffer.WriteString("This would have done something.")
		_, _ = s.ChannelMessageSend(m.ChannelID, buffer.String())
	} else {
		buffer.WriteString("Only the Server owner can change the bot settings...")
		_, _ = s.ChannelMessageSend(m.ChannelID, buffer.String())
	}
}
