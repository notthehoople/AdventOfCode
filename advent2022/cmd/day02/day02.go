package main

import (
	"AdventOfCode-go/advent2022/utils"
	"fmt"
)

func getGameScore(elfPlay byte, myPlay byte) int {
	// A for Rock
	// B for Paper
	// C for Scissors
	// My play: X for Rock (score 1)
	// My play: Y for Paper (score 2)
	// My play: Z for Scissors (score 3)
	// Score for the round: 0 if lost; 3 is draw; 6 if won
	// Scissors beats paper
	// Paper beats rock
	// Rock beats Scissors

	switch elfPlay {
	case 'A':
		switch myPlay {
		case 'X':
			// It's a draw, and I played rock so 1 (rock) + 3 (draw)
			return 1 + 3
		case 'Y':
			// Paper wins and I played it so 2 (paper) + 6 (win)
			return 2 + 6
		case 'Z':
			// Scissors loses and I played it so 3 (scissors) + 0 (loss)
			return 3 + 0
		}
	case 'B':
		switch myPlay {
		case 'X':
			// Rock loses and I played it so 1 (rock) + 0 (loss)
			return 1 + 0
		case 'Y':
			// Paper draws and I played it so 2 (paper) + 3 (draw)
			return 2 + 3
		case 'Z':
			// Scissors wins and I played it so 3 (scissors) + 6 (win)
			return 3 + 6
		}
	case 'C':
		switch myPlay {
		case 'X':
			// Rock beats scissors and I played it so 1 (rock) + 6 (win)
			return 1 + 6
		case 'Y':
			// Paper loses and I played it so 2 (paper) + 0 (loss)
			return 2 + 0
		case 'Z':
			// Scissors draws and I played it so 3 (scissors) + 3 (draw)
			return 3 + 3
		}
	}
	return 0
}

func elfRockPaperScissors(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)
	fmt.Println("Processed file")

	var elfPlay, myPlay byte
	var currentScore int
	for _, j := range puzzleInput {
		fmt.Sscanf(j, "%c %c\n", &elfPlay, &myPlay)

		fmt.Printf("%s Elf played %c and I played %c\n", j, elfPlay, myPlay)

		currentScore += getGameScore(elfPlay, myPlay)
	}

	return currentScore
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", elfRockPaperScissors(filenamePtr, execPart, debug))
	}
}
