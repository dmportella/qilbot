package bot

type Plugin struct {
	Name        string
	Description string
	Qilbot      *Qilbot
	Commands    []CommandInformation
}

type CommandInformation struct {
	Command     string
	Template    string
	Description string
}

type IPlugin interface {
	GetHelpText() string
	Initialize(qilbot *Qilbot)
}
