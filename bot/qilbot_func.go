package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dmportella/qilbot/logging"
)

// New creates a new instance of Qilbot
func New(config *QilbotConfig) (bot *Qilbot, err error) {
	bot = &Qilbot{config: config}

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

	bot.Plugins = []IPlugin{}

	return
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

// Stop not implemented
func (qilbot *Qilbot) Stop() {
	// discordgo package doesn't seem to have any close or stop functionality.
}

// AddPlugin Add a plugin to qilbot that will be initialised with a instance for the discord session.
func (qilbot *Qilbot) AddPlugin(plugin IPlugin) {
	qilbot.Plugins = append(qilbot.Plugins, plugin)
}

// AddHandler Adds an event handler for discord events
func (qilbot *Qilbot) AddHandler(handler interface{}) {
	qilbot.session.AddHandler(handler)
}

// IsBot Simple check to see if the an ID matches the bot id.
func (qilbot *Qilbot) IsBot(id string) bool {
	return id == qilbot.BotID
}
