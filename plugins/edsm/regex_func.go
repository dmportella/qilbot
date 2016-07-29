package edsm

import (
	"regexp"
)

const (
	distanceCommandRegex	= `^(.*)\s?\/\s?(.*)`
	sphereCommandRegex		= `^(.*)\s+?(\d*\.?\d*)`
)

// regexMatchDistanceCommand matches a distance command.
// Regex that matches commands like:
// "sothis / bava".
// Two capture groups are returned  (.*), (sothis), (bava)
func regexMatchDistanceCommand(msg string) []string {
	placesPattern := regexp.MustCompile(distanceCommandRegex)
	return placesPattern.FindStringSubmatch(msg)
}

// regexMatchSphereCommand matches a sphere command.
// Regex that matches commands like:
// "sothis 14.33ly".
// Two capture groups are returned  (.*), (sothis), (14.33)
func regexMatchSphereCommand(msg string) []string {
	placesPattern := regexp.MustCompile(sphereCommandRegex)
	return placesPattern.FindStringSubmatch(msg)
}