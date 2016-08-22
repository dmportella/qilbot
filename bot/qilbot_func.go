package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/dmportella/qilbot/logging"
	"github.com/dmportella/qilbot/utilities"
)

// New creates a new instance of Qilbot
func New(config *QilbotConfig) (bot *Qilbot, err error) {
	bot = &Qilbot{
		config:      config,
		Plugins:     []IPlugin{},
		stopChannel: make(chan struct{}),
		commands:    make(map[string]*CommandInformation),
	}

	// Create a new Discord session using the provided login information.
	if dg, ok := discordgo.New(bot.config.Email, bot.config.Password, bot.config.Token); ok != nil {
		logging.Error.Println("Could not create discord session, ", err)
		err = ok
	} else {
		bot.session = dg
	}

	// Get the account information.
	if u, ok := bot.session.User("@me"); ok != nil {
		logging.Error.Println("Could not fetch bot account details, ", err)
		err = ok
	} else {
		// store bot user id for later use.
		bot.BotID = u.ID
	}

	bot.AddHandler(bot.discordCreateMessage)

	return
}

func (qilbot *Qilbot) discordCreateMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if qilbot.IsBot(m.Author.ID) {
		return
	}

	matches := utilities.RegexMatchBotCommand(m.Content)

	logging.Info.Println(matches)

	if len(matches) > 0 {
		commandCalled := matches[1]

		if command, ok := qilbot.commands[commandCalled]; ok {
			go command.Execute(s, m, matches[2])
		}
	}
}

// InDebugMode returns true if qilbot is running in debug mode.
func (qilbot *Qilbot) InDebugMode() bool {
	return qilbot.config.Debug
}

// Start Opens a WebSocket connection with discord.
func (qilbot *Qilbot) Start() (err error) {
	// Open the websocket and begin listening.
	if ok := qilbot.session.Open(); ok != nil {
		logging.Error.Println("error opening connection,", err)
		err = ok
	}
	return
}

// Stop signal all go routines to stop.
func (qilbot *Qilbot) Stop() {
	close(qilbot.stopChannel)
}

// AddPlugin Add a plugin to qilbot.
func (qilbot *Qilbot) AddPlugin(plugin IPlugin) {
	qilbot.Plugins = append(qilbot.Plugins, plugin)
}

// AddCommand Add a command to the list of commands available to qilbot.
func (qilbot *Qilbot) AddCommand(command *CommandInformation) (err error) {
	if _, ok := qilbot.commands[command.Command]; ok {
		err = fmt.Errorf("Another command is registered with this '%s' name", command.Command)
		return
	}

	qilbot.commands[command.Command] = command
	return
}

// AddHandler Adds an event handler for discord events
func (qilbot *Qilbot) AddHandler(handler interface{}) {
	qilbot.session.AddHandler(handler)
}

// IsBot Simple check to see if the an ID matches the bot id.
func (qilbot *Qilbot) IsBot(id string) bool {
	return id == qilbot.BotID
}
