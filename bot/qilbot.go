package bot

// Qilbot representation of the instance of qilbot.
type Qilbot struct {
	// Publics

	Plugins []IPlugin

	// Privates

	botID           string
	stopChannel     chan struct{}
	config          *QilbotConfig
	session         *DiscordSession
	commands        map[string]*CommandInformation
	commandSettings map[string]*commandSettings
}

// QilbotConfig representation of the configuration for qilbot.
type QilbotConfig struct {
	Email    string
	Password string
	Token    string
	Debug    bool
}
