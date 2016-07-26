package edsm

import (
	"regexp"
)

const (
	distance_command_regex = `^(.*)\s?\/\s?(.*)`
	sphere_command_regex   = `^(.*)\s+?(\d*\.?\d*)`
)

// Regex that matches commands like:
// "sothis / bava".
// Two capture groups are returned  (.*), (sothis), (bava)
func RegexMatchDistanceCommand(msg string) []string {
	placesPattern := regexp.MustCompile(distance_command_regex)
	return placesPattern.FindStringSubmatch(msg)
}

// Regex that matches commands like:
// "sothis 14.33ly".
// Two capture groups are returned  (.*), (sothis), (14.33)
func RegexMatchSphereCommand(msg string) []string {
	placesPattern := regexp.MustCompile(sphere_command_regex)
	return placesPattern.FindStringSubmatch(msg)
}
