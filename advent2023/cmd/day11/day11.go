package main

import (
	"AdventOfCode-go/advent2023/utils"
	"fmt"
	"slices"
	"strings"
)

//
//

func day11(filename string, part byte, debug bool) int {
	var results int

	puzzleInput, _ := utils.ReadFile(filename)

	skyGrid := make([][]byte, len(puzzleInput))
	for line, puzzleLine := range puzzleInput {
		skyGrid[line] = make([]byte, len(puzzleLine))
		for key, value := range puzzleLine {
			skyGrid[line][key] = byte(value)
		}
	}

	// Check for row being empty of Galaxies

	utils.Print2DArrayByte(skyGrid)

	for i := 0; i < len(skyGrid); i++ {
		if strings.Contains(string(skyGrid[i]), "#") {
			fmt.Println("skyGrid[i]:", skyGrid[i], "contains a galaxy")
		} else {
			skyGrid = slices.Insert(skyGrid, i, skyGrid[i])
			fmt.Println("skyGrid[i]:", skyGrid[i], "is empty")
		}
	}

	return results
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Result is: %d\n", day11(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("Result is: %d\n", day11(filenamePtr, execPart, debug))
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
