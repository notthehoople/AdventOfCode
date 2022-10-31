package main

import (
	"fmt"
	"flag"
)

// Used to build a work list of water sources to be followed
type spaceCoordsStruct struct {	
	xCoord			int
	yCoord			int
	zCoord			int
	tCoord			int
	constellation	int
}

// func: constellationsControlA
// 
func constellationsControlA(fileName string, maxRange int, debug bool) (int) {
	var pointLocation []spaceCoordsStruct
	var constellationCounter int = 0
	var tempDistance int
	var tempConstellation1, tempConstellation2 int = 0, 0
	var countingStars[1000] bool

	fileContents, _ := readLines(fileName)
	pointLocation = processInputFile(fileContents)

	if debug {
		fmt.Println(pointLocation)
	}

	for i := 0; i < len(pointLocation) - 1; i++ {

		if pointLocation[i].constellation == 0 {
			constellationCounter++
			pointLocation[i].constellation = constellationCounter
		}

		for j := i+1; j < len(pointLocation); j++ {
	
			if debug {
				fmt.Println("================================")

				fmt.Println("First PointLocation:", i, pointLocation[i])
				fmt.Println("Second PointLocation:", j, pointLocation[j])
			}

			tempDistance = manhattanDistance4D(pointLocation[i], pointLocation[j])
			if debug {
				fmt.Printf("Distance from %d to %d is %d\n", i, j, tempDistance)
			}

			if tempDistance <= maxRange {
				tempConstellation1 = pointLocation[i].constellation
				tempConstellation2 = pointLocation[j].constellation

				if tempConstellation2 == 0 {
					pointLocation[j].constellation = tempConstellation1
				} else {
					if tempConstellation1 == tempConstellation2 {
						// Do nothing. Already in the same constellation
					} else {
						if debug {
							fmt.Println("Joining Constellations")
						}
						for z := 0; z < len(pointLocation); z++ {
							if pointLocation[z].constellation == tempConstellation2 {
								pointLocation[z].constellation = tempConstellation1
							}
						}
					}
				}
			}

		}
	}
	if debug {
		fmt.Println(pointLocation)
	}

	// Now need to count the number of unique consellations we have. We need to do this as joining constellations together
	// can leave us with a non-contiguous list of constellations

	for i := 0; i < len(pointLocation); i++ {
		//fmt.Println("Constellation:", pointLocation[i].constellation)
		countingStars[pointLocation[i].constellation] = true
	}

	constellationCounter = 0
	for i := 0; i < len(countingStars); i++ {
		if countingStars[i] {
			constellationCounter++
		}
	}

	return constellationCounter
}

// Main routine
func main() {
	var debug bool
	var maxRange int

	fileNamePtr := flag.String("file", "input1.txt", "A filename containing input strings")
	flag.BoolVar(&debug, "debug", false, "turns print debugging on")
	flag.IntVar(&maxRange, "range", 3, "Range to count as constellation")
	execPartPtr := flag.String("part", "a", "Which part of day18 do you want to calc (a or b)")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Number of constelations:", constellationsControlA(*fileNamePtr, maxRange, debug))
	case "b":
		fmt.Println("Not here yet")
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}