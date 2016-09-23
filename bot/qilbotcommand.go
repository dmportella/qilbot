package bot

// QilbotCommand describes a command in available for a plugin.
type QilbotCommand struct {
	// Command the name of the command used in the chat to activate it.
	Command string

	// Template the markdown displayed when the command help text is rendered.
	Template string

	// Description the markdown displayed when the command help text is rendered.
	Description string

	// Execute called when the command is called from chat.
	Execute func(ctx *QilbotCommandContext)

	settings *commandSettings
}
