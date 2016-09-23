package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/dmportella/qilbot/logging"
	"github.com/dmportella/qilbot/utilities"
	"path"
)

// New creates a new instance of Qilbot
func New(config *QilbotConfig) (*Qilbot, error) {
	bot := Qilbot{
		config:          config,
		stopChannel:     make(chan struct{}),
		commands:        make(map[string]*QilbotCommand),
		commandSettings: make(map[string]*commandSettings),
	}

	ok1 := bot.initialiseData()

	if ok1 != nil {
		return nil, ok1
	}

	ok2 := bot.initialiseDiscord()

	if ok2 != nil {
		return nil, ok2
	}

	bot.initialiseCommands()

	return &bot, nil
}

func (qilbot *Qilbot) saveData() {
	if currentFolder, ok1 := utilities.GetCurrentFolder(); ok1 == nil {
		commandSettingsPath := path.Join(currentFolder, "command-settings.json")

		if data, ok2 := utilities.ToJSON(&qilbot.commandSettings); ok2 == nil {
			ok3 := utilities.SaveToFile(data, commandSettingsPath)

			logging.Info.Println("Command settings saved.", len(data))

			if ok3 != nil {
				logging.Error.Println("Could not save command settings.", ok3)
			}
		} else {
			logging.Error.Println("Could not encode command settings.", ok2)
		}
	}
}

func (qilbot *Qilbot) initialiseData() error {
	if currentFolder, ok1 := utilities.GetCurrentFolder(); ok1 == nil {
		commandSettingsPath := path.Join(currentFolder, "command-settings.json")

		if data, ok2 := utilities.ReadFromFile(commandSettingsPath); ok2 == nil && data != nil {
			utilities.FromJSON(data, &qilbot.commandSettings)

			logging.Info.Printf("Loaded command settings at '%s'.", commandSettingsPath)

			return ok2
		}

		return ok1
	}

	return nil
}

func (qilbot *Qilbot) initialiseDiscord() (err error) {
	// Create a new Discord session using the provided login information.
	if dg, ok := discordgo.New(qilbot.config.Email, qilbot.config.Password, qilbot.config.Token); ok != nil {
		logging.Error.Println("Could not create discord session, ", ok)
		err = ok
	} else {
		qilbot.session = dg
	}

	// Get the account information.
	if u, ok := qilbot.session.User("@me"); ok != nil {
		logging.Error.Println("Could not fetch bot account details, ", ok)
		err = ok
	} else {
		// store bot user id for later use.
		qilbot.botID = u.ID
	}

	qilbot.AddHandler(qilbot.discordCreateMessage)

	return
}

func (qilbot *Qilbot) discordCreateMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if qilbot.IsBot(m.Author.ID) {
		return
	}

	matches := utilities.RegexMatchBotCommand(m.Content)

	if len(matches) > 0 {
		logging.Info.Println(matches)

		commandCalled := matches[1]

		if command, ok := qilbot.commands[commandCalled]; ok {
			ctx := QilbotCommandContext{
				Message:        m,
				CommandText:    matches[2],
				command:        command,
				discordSession: qilbot.session,
			}
			logging.Trace.Printf("%p", qilbot)

			go command.Execute(&ctx)
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

	qilbot.saveData()
}

// AddCommand Add a command to the list of commands available to qilbot.
func (qilbot *Qilbot) AddCommand(command *QilbotCommand) (err error) {
	if _, ok := qilbot.commands[command.Command]; ok {
		err = fmt.Errorf("Another command is registered with this '%s' name", command.Command)
		return
	}

	if settings, ok := qilbot.commandSettings[command.Command]; ok {
		command.settings = settings
	} else {
		emptySettings := commandSettings{"", true, false}

		qilbot.commandSettings[command.Command] = &emptySettings

		command.settings = &emptySettings
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
	return id == qilbot.botID
}
