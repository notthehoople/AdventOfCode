package main

import (
	"AdventOfCode-go/advent2023/utils"
	"fmt"
	"strconv"
	"strings"
)

// In part b you win copies of subsequent cards
// The result is the number of cards you have at the end, including the originals

func day04b(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)

	// First let's build a map of the cards we have
	scratchCards := make(map[int]int)
	for _, puzzleLine := range puzzleInput {
		var cardNumber int
		fmt.Sscanf(puzzleLine, "Card %d:", &cardNumber)
		scratchCards[cardNumber] = 1
		if debug {
			fmt.Println("Card Number:", cardNumber)
		}
	}

	// Now let's look for winning numbers and increase card count appropriately
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

		// Loop through this based on number in scratchCards
		ourcard := strings.Fields(strings.TrimSpace(card[1]))
		var cardNumber int
		fmt.Sscanf(puzzleLine, "Card %d:", &cardNumber)

		var cardScore int = 0
		for _, numString := range ourcard {
			number, _ := strconv.Atoi(numString)
			if winningNumbers[number] {
				cardScore++
			}
		}
		if debug {
			fmt.Printf("Card: %s Num: %d Wins: %d\n", card[1], cardNumber, cardScore)
		}

		for i := 1; i <= scratchCards[cardNumber]; i++ {
			for j := 1; j <= cardScore; j++ {
				scratchCards[cardNumber+j]++
			}
		}
	}

	// Now count our scratch cards and see what we've won
	var result int
	for _, i := range scratchCards {
		result += i
	}
	if debug {
		fmt.Println(scratchCards)
	}
	return result
}

// Any number adjacent to a symbol, even diagonally, is a "part number" and should be
// included in the sum. Periods '.' do not count as a symbol.
// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53

func day04a(filename string, part byte, debug bool) int {
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

	switch execPart {
	case 'a':
		fmt.Printf("Result is: %d\n", day04a(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("Result is: %d\n", day04b(filenamePtr, execPart, debug))
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
