package main

import (
	"AdventOfCode-go/advent2024/utils"
	"fmt"
)

// checkForXmas - part b. Checking for a cross of MAS or SAM
func checkForXmas(puzzleInput []string, x int, y int, mas string, sam string, maxX int, maxY int) int {
	// Ditch the easy situations for speed
	// If the top left of the X doesn't match either an "M" or "S" then exit
	if puzzleInput[y][x] != mas[0] && puzzleInput[y][x] != sam[0] {
		return 0
	}

	// We're looking to the left and down ONLY. Check our boundaries and return if we're too close to the edges
	if x+2 >= maxX || y+2 >= maxY {
		return 0
	}

	// If the middle square isn't an 'A' then it can't be an X-MAS
	if puzzleInput[y+1][x+1] != 'A' {
		return 0
	}

	// Now check for MAS going down to the left and SAM going down to the left
	gotMAS := checkForString(puzzleInput, x, y, 1, 1, mas, maxX, maxY)
	gotSAM := checkForString(puzzleInput, x, y, 1, 1, sam, maxX, maxY)
	// Now check for MAS going down to the right and SAM going down to the right
	gotbackMAS := checkForString(puzzleInput, x+2, y, -1, 1, mas, maxX, maxY)
	gotbackSAM := checkForString(puzzleInput, x+2, y, -1, 1, sam, maxX, maxY)

	// If we got a MAS doing down to the left and EITHER a MAS or a SAM going down to the right, we have an X
	if gotMAS == 1 && (gotbackMAS == 1 || gotbackSAM == 1) {
		return 1
	}
	// If we got a SAM doing down to the left and EITHER a MAS or a SAM going down to the right, we have an X
	if gotSAM == 1 && (gotbackMAS == 1 || gotbackSAM == 1) {
		return 1
	}

	// We got nothing :-(
	return 0
}

// part a. Checking for the searchString in whatever direction we're asked to. Also referenced in part b
func checkForString(puzzleInput []string, x int, y int, chgX int, chgY int, searchString string, maxX int, maxY int) int {

	if puzzleInput[y][x] != searchString[0] {
		return 0
	}

	var xPos int = x
	var yPos int = y
	for check := 0; check < len(searchString); check++ {
		if xPos >= maxX || xPos < 0 || yPos >= maxY || yPos < 0 {
			return 0
		}

		if puzzleInput[yPos][xPos] != searchString[check] {
			return 0
		}
		xPos += chgX
		yPos += chgY

	}

	return 1
}

func day04(filename string, part byte, debug bool) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)

	maxX := len(puzzleInput[0])
	maxY := len(puzzleInput)

	if part == 'a' {
		xmas := "XMAS"
		samx := "SAMX"

		for y := 0; y < len(puzzleInput); y++ {
			for x := 0; x < len(puzzleInput[y]); x++ {

				//Check along from x,y for XMAS
				//Check along from x,y for SAMX
				result += checkForString(puzzleInput, x, y, 1, 0, xmas, maxX, maxY)
				result += checkForString(puzzleInput, x, y, 1, 0, samx, maxX, maxY)

				// Check Down from x,y for XMAS
				// Check Down from x,y for SAMX
				result += checkForString(puzzleInput, x, y, 0, 1, xmas, maxX, maxY)
				result += checkForString(puzzleInput, x, y, 0, 1, samx, maxX, maxY)

				// Check Down Left from x,y for XMAS
				// Check Down Left from x,y for SAMX
				result += checkForString(puzzleInput, x, y, -1, 1, xmas, maxX, maxY)
				result += checkForString(puzzleInput, x, y, -1, 1, samx, maxX, maxY)

				// Check Down Right from x,y foo XMAS
				// Check Down Right from x,y foo SAMX
				result += checkForString(puzzleInput, x, y, 1, 1, xmas, maxX, maxY)
				result += checkForString(puzzleInput, x, y, 1, 1, samx, maxX, maxY)
			}
		}

		return result
	}

	mas := "MAS"
	sam := "SAM"

	for y := 0; y < len(puzzleInput); y++ {
		for x := 0; x < len(puzzleInput[y]); x++ {

			// Starting from top left of the X we look 3 down to the right. If we don't see an 'A' in the middle we quit
			result += checkForXmas(puzzleInput, x, y, mas, sam, maxX, maxY)
		}
	}

	return result
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", day04(filenamePtr, execPart, debug))
	}
}
