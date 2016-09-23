package bot

import (
	"regexp"
)

const (
	keyValueCommandRegex = `^(\w+)\s?(.*)`
)

// regexMatchKeyValueCommand matches a key value pair.
func regexMatchKeyValueCommand(msg string) []string {
	placesPattern := regexp.MustCompile(keyValueCommandRegex)
	return placesPattern.FindStringSubmatch(msg)
}
