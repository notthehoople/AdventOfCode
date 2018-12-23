package main

import (
	"fmt"
	"flag"
	"strconv"
	"math"
)

// Calc the power values over the entire map
func calcPowerOverMap(powerMap [][]int, puzzleInput int) {
	var rackID, powerLevel int = 0, 0

	for x := 1; x < len(powerMap); x++ {
		for y := 1; y < len(powerMap[x]); y++ {
			rackID = x + 10
			powerLevel = ((rackID * y) + puzzleInput) * rackID
			powerLevel = int(math.Mod(float64(powerLevel / 100),10))
			powerLevel -= 5

			powerMap[x][y] = powerLevel
		}
	}
}

// Checks the 3x3 grid starting at xCoord, yCoord and calcs the power of the grid
// Returns: the power level seen
func checkPower(powerMap [][]int, xCoord int, yCoord int, part byte) (int, int) {
	var powerLevel int = 0
	var bestPower, bestGridSize int = 0, 0

	if part == 'a' {
 		for x := xCoord; x < xCoord + 3; x++ {
			for y := yCoord; y < yCoord + 3; y++ {
				powerLevel += powerMap[y][x]
			}
		}
		fmt.Printf("adding x: %d y: %d powerlevel: %d\n", xCoord, yCoord, powerLevel)

		return powerLevel, 3
	} else {
		for t := 3; (xCoord + t) < 300 && (yCoord + t) < 300; t++ {
			powerLevel = 0
			//fmt.Printf("t: %d xCoord: %d yCoord: %d\n", t, xCoord, yCoord)
			for x := xCoord; x < xCoord + t; x++ {
				for y := yCoord; y < yCoord + t; y++ {
					powerLevel += powerMap[y][x]
				}
			}
			//fmt.Printf("adding x: %d y: %d grid size: %d powerlevel: %d\n", xCoord, yCoord, t, powerLevel)
			if powerLevel > bestPower {
				bestPower = powerLevel
				bestGridSize = t
			}
		}
		return bestPower, bestGridSize

	}
}

// Finds the 3x3 grid that has the highest power output in the overall large grid
// returns the x, y coords of the top left of the best 3x3 grid, then the best power level seen
func findBestPowerInMap(powerMap [][]int, part byte) (int, int, int) {
	var testPower, bestXCoord, bestYCoord, bestPower, bestSizeOfgrid, sizeOfGrid int = 0, 0, 0, 0, 0, 0

	if part == 'a' {
		for x := 1; x < len(powerMap) - 2; x++ {
			for y := 1; y < len(powerMap[x]) - 2; y++ {
				testPower, sizeOfGrid = checkPower(powerMap, x, y, part)
				if testPower > bestPower {
					bestXCoord = y
					bestYCoord = x
					bestPower = testPower
					bestSizeOfgrid = sizeOfGrid
				}
			}
		}

		return bestXCoord, bestYCoord, 3
	} else {
		for x := 1; x < len(powerMap) - 2; x++ {
			for y := 1; y < len(powerMap[x]) - 2; y++ {
				testPower, sizeOfGrid = checkPower(powerMap, x, y, part)
				if testPower > bestPower {
					bestXCoord = y
					bestYCoord = x
					bestPower = testPower
					bestSizeOfgrid = sizeOfGrid
				}
			}
		}

		return bestXCoord, bestYCoord, bestSizeOfgrid
	}
}

func printPowerMap(powerMap [][]int, grid int, xPrint int, yPrint int) {
	fmt.Printf("%d \n", powerMap[xPrint][yPrint])
}

// Find the fuel cell's rack ID, which is its X coordinate plus 10.
// Begin with a power level of the rack ID times the Y coordinate.
// Increase the power level by the value of the grid serial number (your puzzle input).
// Set the power level to itself multiplied by the rack ID.
// Keep only the hundreds digit of the power level (so 12345 becomes 3; numbers with no hundreds digit become 0).
// Subtract 5 from the power level.
func powerCalc(puzzleInput int, gridSize int, part byte, xPrint int, yPrint int) (int, int, int) {
	var bestXCoord, bestYCoord, sizeOfGrid int = 0, 0, 0
	
	powerMap := make([][]int, gridSize)
	for i := 1; i < len(powerMap); i++ {
		powerMap[i] = make([]int, gridSize)
	}

	calcPowerOverMap(powerMap, puzzleInput)

	bestXCoord, bestYCoord, sizeOfGrid = findBestPowerInMap(powerMap, part)

	return bestXCoord, bestYCoord, sizeOfGrid
}

// Main routine
func main() {
	var puzzleInput, gridSize int = 0, 0
	var xcoord, ycoord int = 0, 0
	var xPrint, yPrint int = 0, 0
	var sizeOfGrid int = 0

	puzzleInputPtr := flag.String("puzzle", "7165", "Puzzle input value")
	gridSizePtr := flag.String("grid", "300", "Size of grid to calc power values for")
	xPrintPtr := flag.String("x", "10", "xcoord to print")
	yPrintPtr := flag.String("y", "10", "y coord to print")
	execPartPtr := flag.String("part", "a", "Which part of day10 do you want to calc (a or b)")

	flag.Parse()

	puzzleInput, _ = strconv.Atoi(*puzzleInputPtr)
	gridSize, _ = strconv.Atoi(*gridSizePtr)
	xPrint, _ = strconv.Atoi(*xPrintPtr)
	yPrint, _ = strconv.Atoi(*yPrintPtr)

	switch *execPartPtr {
	case "a":
		xcoord, ycoord, sizeOfGrid = powerCalc(puzzleInput, gridSize, 'a', xPrint, yPrint)
		fmt.Println("Part a - Coords of the highest power 3x3:", xcoord, ycoord)
	case "b":
		xcoord, ycoord, sizeOfGrid = powerCalc(puzzleInput, gridSize, 'b', xPrint, yPrint)
		fmt.Println("Part b - Coords of the highest power 3x3:", xcoord, ycoord, sizeOfGrid)
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}