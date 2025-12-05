package main

import (
	"AdventOfCode-go/advent2025/utils"
	"fmt"
)

func findMaxNumber(batteryBank string, startPos int, maxDigits int) (int, int) {
	var maxNext, maxNextPos int = 0, 0

	// Subtract 1 from maxDigits as we start at Pos 0
	maxDigits--
	for i := startPos; i < len(batteryBank)-maxDigits; i++ {
		if int(batteryBank[i]-'0') > maxNext {
			maxNext = int(batteryBank[i] - '0')
			maxNextPos = i + 1
		}
	}

	return maxNext, maxNextPos
}

func day04(filename string, part byte, debug bool) int {
	var result, characters int

	puzzleInput, _ := utils.ReadFile(filename)

	if part == 'a' {
		characters = 2
	} else {
		characters = 12
	}

	for _, batteryBank := range puzzleInput {

		var joltage int = 0
		var nextPos int = 0
		var maxPower int = 0
		for count := characters; count > 0; count-- {

			maxPower, nextPos = findMaxNumber(batteryBank, nextPos, count)
			joltage = joltage*10 + maxPower

		}
		result += joltage
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
