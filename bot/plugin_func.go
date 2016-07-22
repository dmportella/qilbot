package bot

import "fmt"

func (self *Plugin) GetHelpText() (msg string) {
	return fmt.Sprintf("Plugin **%s**\nDescription: *%s*", self.Name, self.Description)
}
