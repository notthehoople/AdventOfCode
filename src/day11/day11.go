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
func checkPower(powerMap [][]int, xCoord int, yCoord int) int {
	var powerLevel int = 0

	for x := xCoord; x < xCoord + 3; x++ {
		for y := yCoord; y < yCoord + 3; y++ {
			powerLevel += powerMap[y][x]
		}
	}

	return powerLevel
}

// Finds the 3x3 grid that has the highest power output in the overall large grid
// returns the x, y coords of the top left of the best 3x3 grid, then the best power level seen
func findBestPowerInMap(powerMap [][]int) (int, int) {
	var testPower, bestXCoord, bestYCoord, bestPower int = 0, 0, 0, 0

	for x := 1; x < len(powerMap) - 2; x++ {
		for y := 1; y < len(powerMap[x]) - 2; y++ {
			testPower = checkPower(powerMap, x, y)
			if testPower > bestPower {
				bestXCoord = y
				bestYCoord = x
				bestPower = testPower
			}
		}
	}

	return bestXCoord, bestYCoord
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
func powerCalc(puzzleInput int, gridSize int, part byte, xPrint int, yPrint int) (int, int) {
	var bestXCoord, bestYCoord int = 0, 0
	
	powerMap := make([][]int, gridSize)
	for i := 1; i < len(powerMap); i++ {
		powerMap[i] = make([]int, gridSize)
	}

	calcPowerOverMap(powerMap, puzzleInput)

	bestXCoord, bestYCoord = findBestPowerInMap(powerMap)

	return bestXCoord, bestYCoord
}

// Main routine
func main() {
	var puzzleInput, gridSize int = 0, 0
	var xcoord, ycoord int = 0, 0
	var xPrint, yPrint int = 0, 0

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
		xcoord, ycoord = powerCalc(puzzleInput, gridSize, 'a', xPrint, yPrint)
		fmt.Println("Part a - Coords of the highest power 3x3:", xcoord, ycoord)
	case "b":
		fmt.Println("Part b - Not here yet")
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}