package edsm

import (
	"regexp"
)

const (
	DISTANCE_COMMAND_REGEX = `^(.*)\s?\/\s?(.*)`
)

// Regex that matches commands like:
// "sothis / bava".
// Two capture groups are returned  (.*), (sothis), (bava)
func RegexMatchDistanceCommand(msg string) []string {
	placesPattern := regexp.MustCompile(DISTANCE_COMMAND_REGEX)
	return placesPattern.FindStringSubmatch(msg)
}
