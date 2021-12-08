package main

import (
	"aoc/advent2021/utils"
	"fmt"
	"strings"
)

func countUniqueNumbers(displayDigits []string, part byte, debug bool) int {
	var countUnique int
	for _, i := range displayDigits {
		digitsSplit := strings.Split(i, " | ")

		for _, outputDigit := range strings.Split(digitsSplit[1], " ") {
			switch len(outputDigit) {
			case 2: // Digit 1
				countUnique++
			case 3: // Digit 7
				countUnique++
			case 4: // Digit 4
				countUnique++
			case 7: // Digit 8
				countUnique++
			default:
				// Not a unique digit
			}
		}
	}
	return countUnique
}

func solveDay(filename string, part byte, debug bool) int {
	puzzleInput, _ := utils.ReadFile(filename)

	if part == 'a' {
		return countUniqueNumbers(puzzleInput, part, debug)
	} else {
		return 0
	}
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", solveDay(filenamePtr, execPart, debug))
	}
}
