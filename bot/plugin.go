package bot

// Plugin describes a plugin in qilbot.
type Plugin struct {
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
}

// IPlugin interface used by qilbot.
type IPlugin interface {
	GetCommands() []CommandInformation
	GetHelpText() string
	Initialize(qilbot *Qilbot)
}
