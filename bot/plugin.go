package bot

type Plugin struct {
	Name        string
	Description string
	Qilbot      *Qilbot
}

type IPlugin interface {
	GetHelpText() string
	Initialize(qilbot *Qilbot)
}
