package bot

import "fmt"

func (self *Plugin) GetHelpText() (msg string) {
	return fmt.Sprintf("Plugin **%s**, Description: *%s*", self.Name, self.Description)
}

func (self *Plugin) GetCommands() []CommandInformation {
	return self.Commands
}
