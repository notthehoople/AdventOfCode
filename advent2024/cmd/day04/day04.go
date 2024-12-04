package main

import (
	"AdventOfCode-go/advent2024/utils"
	"fmt"
)

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

	if part == 'a' {
		xmas := "XMAS"
		samx := "SAMX"
		maxX := len(puzzleInput[0])
		maxY := len(puzzleInput)

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
	return 0
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
