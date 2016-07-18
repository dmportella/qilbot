package main

import (
	"github.com/bwmarrin/discordgo"

	"fmt"
)

// Build version of the binary
var Build string

// Revision number of the binary
var Revision string

// Branch name of the binary
var Branch string

func main() {
	fmt.Printf("Golang tutorial version %s, branch %s at revision %s.\n\rDaniel Portella (c) 2016\n\r", Build, Branch, Revision)

}
