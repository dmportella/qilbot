package utilities

import (
	"regexp"
)

const (
	COMMAND_REGEX = `^<@([0-9]+)>\s([a-z]+)\s?(.*)`
)

// Regex that matches commands like:
// "@<7897978789899> distance sothis / bava".
// Three capture groups are returned  (.*), (7897978789899), (distance), (sothis / bava)
func RegexMatchBotCommand(msg string) []string {
	actionPattern := regexp.MustCompile(COMMAND_REGEX)
	return actionPattern.FindStringSubmatch(msg)
}
