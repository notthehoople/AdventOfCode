package main

import (
	"AdventOfCode-go/advent2023/utils"
	"fmt"
	"strconv"
	"strings"
)

// Any number adjacent to a symbol, even diagonally, is a "part number" and should be
// included in the sum. Periods '.' do not count as a symbol.
// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53

func day04(filename string, part byte, debug bool) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)

	for _, puzzleLine := range puzzleInput {
		// Split string on ':' and '|'
		macrocard := strings.Split(puzzleLine, ":")
		card := strings.Split(strings.TrimSpace(macrocard[1]), "|")

		winningNumbers := make(map[int]bool)

		wincard := strings.Fields(strings.TrimSpace(card[0]))
		for _, numString := range wincard {
			number, _ := strconv.Atoi(numString)
			winningNumbers[number] = true
		}

		ourcard := strings.Fields(strings.TrimSpace(card[1]))
		var cardScore int = 0
		for _, numString := range ourcard {
			number, _ := strconv.Atoi(numString)
			if winningNumbers[number] {
				if cardScore == 0 {
					cardScore = 1
				} else {
					cardScore *= 2
				}
			}
		}

		result += cardScore
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
