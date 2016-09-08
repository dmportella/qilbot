package bot

// Qilbot representation of the instance of qilbot.
type Qilbot struct {
	// Publics

	BotID   string
	Plugins []IPlugin

	// Privates

	stopChannel chan struct{}
	config      *QilbotConfig
	session     *DiscordSession
	commands    map[string]*CommandInformation
}

// QilbotConfig representation of the configuration for qilbot.
type QilbotConfig struct {
	Email    string
	Password string
	Token    string
	Debug    bool
}
