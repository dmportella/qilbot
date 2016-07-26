package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dmportella/qilbot/logging"
)

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

// Opens a WebSocket connection with discord.
func (self *Qilbot) Start() (err error) {
	// Open the websocket and begin listening.
	if ok := self.session.Open(); ok != nil {
		logging.Error.Println("error opening connection,", err)
		err = ok
	}
	return
}

func (self *Qilbot) Stop() {
	// discordgo package doesnt seem to have any close or stop functionality.
}

// Add a plugin to qilbot that will be initialised with a instance for the discord session.
func (self *Qilbot) AddPlugin(plugin IPlugin) {
	self.Plugins = append(self.Plugins, plugin)

	logging.Info.Println(plugin.GetHelpText())

	plugin.Initialize(self)
}

// Adds an event handler for discord events
func (self *Qilbot) AddHandler(handler interface{}) {
	self.session.AddHandler(handler)
}

// Simple check to see if the an ID matches the bot id.
func (self *Qilbot) IsBot(id string) bool {
	return id == self.BotID
}
