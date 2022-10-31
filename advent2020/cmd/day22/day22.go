package main

import (
	"fmt"
	"strconv"
)

func buildQueue(puzzleInput []string, sectionBreak string) (resultQueue []int) {
	var sectionProcessed bool = false
	var tmpCard int

	resultQueue = make([]int, 0)

	for _, line := range puzzleInput {
		if len(line) == 0 {
			// We've reached the break. Have we seen section break? If so quit
			if sectionProcessed {
				return resultQueue
			}
		}

		if line == sectionBreak {
			sectionProcessed = true
		} else if sectionProcessed {
			tmpCard, _ = strconv.Atoi(line)
			resultQueue = append(resultQueue, tmpCard)
		}
	}

	return resultQueue
}

/* HOW TO USE A SLICE AS A QUEUE
queue := make([]int, 0)
// Push to the queue
queue = append(queue, 1)
// Top (just get next element, don't remove it)
x = queue[0]
// Discard top element
queue = queue[1:]
// Is empty ?
if len(queue) == 0 {
    fmt.Println("Queue is empty !")
}
*/

func playCombat(player1Hand []int, player2Hand []int) []int {
	for len(player1Hand) > 0 && len(player2Hand) > 0 {
		if player1Hand[0] > player2Hand[0] {
			player1Hand = append(player1Hand, player1Hand[0])
			player1Hand = append(player1Hand, player2Hand[0])
			player1Hand = player1Hand[1:]
			player2Hand = player2Hand[1:]
		} else if player2Hand[0] > player1Hand[0] {
			player2Hand = append(player2Hand, player2Hand[0])
			player2Hand = append(player2Hand, player1Hand[0])
			player1Hand = player1Hand[1:]
			player2Hand = player2Hand[1:]
		} else {
			fmt.Println("ERROR: Draw!")
		}
	}

	if len(player1Hand) > 1 {
		return player1Hand
	}
	return player2Hand
}

// part a
func calcWinningScore(filename string, part byte, debug bool) int {
	var result int

	puzzleInput, _ := readFile(filename)

	player1Hand := buildQueue(puzzleInput, "Player 1:")
	player2Hand := buildQueue(puzzleInput, "Player 2:")

	if debug {
		fmt.Println("player1Hand:", player1Hand)
		fmt.Println("player2Hand:", player2Hand)
	}

	winningDeck := playCombat(player1Hand, player2Hand)

	posScore := len(winningDeck)
	for _, card := range winningDeck {
		result += card * posScore
		posScore--
	}

	return result
}

// Main routine
func main() {
	filenamePtr, execPart, debug, test := catchUserInput()

	if test {
		return
	}

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else if execPart == 'a' {
		fmt.Println("Winning Score:", calcWinningScore(filenamePtr, execPart, debug))
	} else {
	}
}
