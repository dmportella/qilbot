package main

import (
	//	"github.com/bwmarrin/discordgo"

	"fmt"
	"math"
)

// Build version of the binary
var Build string

// Revision number of the binary
var Revision string

// Branch name of the binary
var Branch string

func main() {
	fmt.Printf("Golang tutorial version %s, branch %s at revision %s.\n\rDaniel Portella (c) 2016\n\r", Build, Branch, Revision)

	bava := [3]float64{83.40625, -134.3125, -80.75}
	sothis := [3]float64{-352.78125, 10.5, -346.34375}

	deltaX := bava[0] - sothis[0]
	deltaY := bava[1] - sothis[1]
	deltaZ := bava[2] - sothis[2]

	distance := math.Sqrt(deltaX*deltaX + deltaY*deltaY + deltaZ*deltaZ)

	fmt.Printf("distance between bava and sothis is: %f\n", distance)

}
