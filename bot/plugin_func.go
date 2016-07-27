package bot

import "fmt"

// GetHelpText generic implemetation for help text.
func (plugin *Plugin) GetHelpText() (msg string) {
	return fmt.Sprintf("Plugin **%s**, Description: *%s*", plugin.Name, plugin.Description)
}

// GetCommands generic implementation for getting commands.
func (plugin *Plugin) GetCommands() []CommandInformation {
	return plugin.Commands
}
