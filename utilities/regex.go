package utilities

import (
	"regexp"
)

const (
	commandRegex = `^<@([0-9]+)>\s([a-z]+)\s?(.*)`
)

// RegexMatchBotCommand find command matches in message strings.
// regex that matches commands like:
// "@<7897978789899> distance sothis / bava".
// Three capture groups are returned  (.*), (7897978789899), (distance), (sothis / bava)
func RegexMatchBotCommand(msg string) []string {
	actionPattern := regexp.MustCompile(commandRegex)
	return actionPattern.FindStringSubmatch(msg)
}
