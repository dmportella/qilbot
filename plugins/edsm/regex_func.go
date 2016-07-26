package edsm

import (
	"regexp"
)

const (
	DISTANCE_COMMAND_REGEX = `^(.*)\s?\/\s?(.*)`
	SPHERE_COMMAND_REGEX   = `^(.*)\s+?(\d*\.?\d*)`
)

// Regex that matches commands like:
// "sothis / bava".
// Two capture groups are returned  (.*), (sothis), (bava)
func RegexMatchDistanceCommand(msg string) []string {
	placesPattern := regexp.MustCompile(DISTANCE_COMMAND_REGEX)
	return placesPattern.FindStringSubmatch(msg)
}

// Regex that matches commands like:
// "sothis 14.33ly".
// Two capture groups are returned  (.*), (sothis), (14.33)
func RegexMatchSphereCommand(msg string) []string {
	placesPattern := regexp.MustCompile(SPHERE_COMMAND_REGEX)
	return placesPattern.FindStringSubmatch(msg)
}
