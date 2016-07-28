package edsm

import (
	"math"
)

func calculateDistance(sys1 *Coordinates, sys2 *Coordinates) float64 {
	deltaX := sys1.X - sys2.X
	deltaY := sys1.Y - sys2.Y
	deltaZ := sys1.Z - sys2.Z

	return math.Sqrt(deltaX*deltaX + deltaY*deltaY + deltaZ*deltaZ)
}
