package main

import (
	"fmt"
	"math"
)

func getAngle(xPos int, yPos int, xxPos int, yyPos int) float64 {
	var dx float64
	// Minus to correct for coord re-mapping
	var dy float64
	var radToDeg float64

	dx = float64(-(xPos)) - float64(xxPos)
	dy = float64(-(yPos)) - float64(yyPos)
	radToDeg = 180 / math.Pi

	inRads := math.Atan2(dy, dx)

	// We need to map to coord system when 0 degree is at 3 O'clock, 270 at 12 O'clock
	if inRads < 0 {
		inRads = math.Abs(inRads)
	} else {
		inRads = 2*math.Pi - inRads
	}

	return inRads * radToDeg
}

func main() {
	fmt.Println("Degrees:", getAngle(0, 0, 0, 5))
}
