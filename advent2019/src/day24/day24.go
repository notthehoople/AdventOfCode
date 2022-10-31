package main

import (
	"flag"
	"fmt"
)

func calcBiodiversity(tempBugArea [][]byte) int {
	var tileValue int = 1
	var totalBiodiversity int

	for i := 0; i < len(tempBugArea); i++ {
		for j := 0; j < len(tempBugArea[i]); j++ {
			if tempBugArea[i][j] == '#' {
				totalBiodiversity += tileValue
			}
			tileValue *= 2
			//fmt.Printf("Char: %c at x:%d y:%d bugs:%d biodiversity:%d\n", tempBugArea[i][j], j, i, checkArea(tempBugArea, i, j), totalBiodiversity)
		}
	}

	return totalBiodiversity
}

// func checkArea
// returns: bugs around us
func checkArea(tempArea [][]byte, yCoord int, xCoord int) int {
	var numBugs int = 0
	var startX, startY, maxX, maxY int

	// To consider: the edges. So if xCoord / yCoord == 0 or xCoord / yCoord = len(tempArea)

	// x x   x o   o x   x x   x o x   x x x   x x   x x
	// o x   x x   x x   x o   x x x   x o x   o x   x o
	//                                         x x   x x

	startX = xCoord - 1
	startY = yCoord - 1
	maxX = xCoord + 1
	maxY = yCoord + 1

	if xCoord == 0 {
		startX = 0
	} else {
		if xCoord == len(tempArea[yCoord])-1 {
			maxX = xCoord
		}
	}
	if yCoord == 0 {
		startY = 0
	} else {
		if yCoord == len(tempArea)-1 {
			maxY = yCoord
		}
	}

	for i := startY; i <= maxY; i++ {
		for j := startX; j <= maxX; j++ {
			if i == yCoord && j == xCoord {
				// Do nothing
			} else {
				// Ignore the diagonals
				if i == yCoord || j == xCoord {
					if tempArea[i][j] == '#' {
						numBugs++
					}
				}
			}
		}
	}

	return numBugs
}

// func
// Returns:
func findBiodiversityRating(filename string, debug bool, part byte) int {
	var biodiversityMap map[int]int
	var newBiodiversity int
	var minuteCount int
	var ok, keepLooping bool

	// Read the map from the file given
	baseBugArea, _ := readLines(filename)
	if debug {
		printMap(baseBugArea)
	}

	// Create a 2d Map of bytes to hold the current state of the bug area
	currentBugArea := make([][]byte, len(baseBugArea))
	for i := 0; i < len(baseBugArea); i++ {
		currentBugArea[i] = make([]byte, len(baseBugArea[0]))
	}
	// Create a 2d Map of bytes to hold the future state of the bug area so we can modify as we loop
	futureBugArea := make([][]byte, len(baseBugArea))
	for i := 0; i < len(baseBugArea); i++ {
		futureBugArea[i] = make([]byte, len(baseBugArea[0]))
	}

	readInitialState(baseBugArea, currentBugArea)
	copy2DSlice(currentBugArea, futureBugArea)

	biodiversityMap = make(map[int]int)

	// Entire area is a 5x5 grid. There are bugs (#) and space (.)
	// Bugs live and die based on the number of bugs in the FOUR adjacent tiles:
	//   - a bug dies unless there is exactly one bug adjacent to it
	//   - an empty space becomes infested with a bug IF exactly one or two bugs are adjacent to it
	// Otherwise the bug or empty space REMAINS THE SAME
	//
	// Tiles at the edge of the grid have less than four adjacent tiles. The missing tiles count as empty space
	// The process happens simultaneously in ever square
	//
	// We're interested in the first time an area matches a previous iteration
	// Then work out the biodiversity of the area:
	//   - Start from the top left, and consider each tile left-to-right then onto the next row.
	//   - Each tile is worth points equal to increasing powers of two: 1, 2, 4, 8, 16, 32 etc
	//   - Add up the biodiversity points for tiles with bugs

	keepLooping = true
	for keepLooping {
		for i := 0; i < len(currentBugArea); i++ {
			for j := 0; j < len(currentBugArea[i]); j++ {
				//fmt.Printf("Char:%c\n", currentBugArea[i][j])
				switch checkArea(currentBugArea, i, j) {
				case 1:
					// bug doesn't die
					// empty space becomes infested
					if currentBugArea[i][j] == '.' {
						futureBugArea[i][j] = '#'
					}
				case 2:
					// bug dies
					if currentBugArea[i][j] == '#' {
						futureBugArea[i][j] = '.'
					}
					// empty space becomes infested
					if currentBugArea[i][j] == '.' {
						futureBugArea[i][j] = '#'
					}
				default:
					// bug dies
					if currentBugArea[i][j] == '#' {
						futureBugArea[i][j] = '.'
					}
					// everything else stays the same
				}
			}
		}
		minuteCount++
		newBiodiversity = calcBiodiversity(futureBugArea)
		_, ok = biodiversityMap[newBiodiversity]
		if ok {
			// We have found a repeat
			fmt.Println("Found a repeat!")
			return newBiodiversity
		} else {
			biodiversityMap[newBiodiversity] = minuteCount
		}

		fmt.Println("START ==============", minuteCount)
		print2DSlice(futureBugArea)

		copy2DSlice(futureBugArea, currentBugArea)
	}
	return 0
}

// Main routine
func main() {
	var debug bool

	filenamePtr := flag.String("file", "input.txt", "Filename containing the program to run")
	execPartPtr := flag.String("part", "a", "Which part of day06 do you want to calc (a or b)")
	flag.BoolVar(&debug, "debug", false, "Turn debug on")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Biodiversity rating:", findBiodiversityRating(*filenamePtr, debug, 'a'))
	case "b":
		fmt.Println("Part b - Not implemented yet")

	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
