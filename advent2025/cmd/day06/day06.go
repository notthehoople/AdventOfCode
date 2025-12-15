package main

import (
	"AdventOfCode-go/advent2025/utils"
	"fmt"
	"strconv"
	"strings"
)

func day06(filename string, part byte, debug bool) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)

	// Get the size of the fresh ingredients array that we need
	countArraySize := len(puzzleInput)
	holdArray := make([][]string, countArraySize)

	for i, puzzleLine := range puzzleInput {
		// Remove leading and trailing white space, and split based on whitespace
		holdArray[i] = strings.Fields(puzzleLine)
	}

	for sumNumber := 0; sumNumber < len(holdArray[0]); sumNumber++ {
		// Work out our numbers
		sumResult, _ := strconv.Atoi(holdArray[0][sumNumber])

		for j := 1; j < len(holdArray)-1; j++ {
			sumTemp, _ := strconv.Atoi(holdArray[j][sumNumber])

			if holdArray[len(holdArray)-1][sumNumber] == "+" {
				sumResult += sumTemp
			} else {
				sumResult *= sumTemp
			}
		}
		result += sumResult

	}

	return result
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", day06(filenamePtr, execPart, debug))
	}
}
