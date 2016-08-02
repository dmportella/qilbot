package utilities

import (
	"regexp"
)

const (
	//commandRegex = `^<@([0-9]+)>\s([a-z]+)\s?(.*)`
	commandRegex = `^\!(\w+)\s?(.*)`
)

// RegexMatchBotCommand find command matches in message strings.
// regex that matches commands like:
// "/distance sothis, bava".
// Three capture groups are returned  (.*), (distance), (sothis / bava)
func RegexMatchBotCommand(msg string) []string {
	actionPattern := regexp.MustCompile(commandRegex)
	return actionPattern.FindStringSubmatch(msg)
}
