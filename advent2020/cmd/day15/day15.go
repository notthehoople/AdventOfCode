package main

import (
	"fmt"
	"strconv"
	"strings"
)

func playNumberGame(puzzleInput string, lastRound int, part byte, debug bool) int {
	var numberSpoken int
	var keepLooping bool
	var turnNumber int = 1
	var previousSpoken int

	memory := make(map[int]int, 2022)

	// this is all one too early. Need to record the previousNumber AFTER we're considering it, not before
	// here I'm saving the previousNumber one step too soon so it *always* exists

	// First we go through the puzzle input to set the scene
	result := strings.Split(puzzleInput, ",")
	for item, strNumber := range result {
		if item > 0 {
			numberSpoken, _ = strconv.Atoi(strNumber)
			memory[previousSpoken] = turnNumber - 1
			if debug {
				fmt.Printf("Turn: %d Number said: %d\n", turnNumber, numberSpoken)
			}
			turnNumber++
			previousSpoken = numberSpoken
		} else {
			previousSpoken, _ = strconv.Atoi(strNumber)
		}
	}

	if debug {
		fmt.Println("Previous:", previousSpoken)
		fmt.Println(memory)
	}

	var tmpval int
	var tmpok bool

	keepLooping = true
	for keepLooping {
		// consider the previous number spoken

		tmpval, tmpok = memory[previousSpoken]
		if debug {
			fmt.Printf("Test for %d is tmpval: %d tmpok %t\n", previousSpoken, tmpval, tmpok)
		}

		if tmpok {
			// It's been spoken before. We need to say "previous turn number" - "last time said turn number"
			numberSpoken = (turnNumber - 1) - memory[previousSpoken]
			memory[previousSpoken] = turnNumber - 1
			if debug {
				fmt.Printf("Old Turn: %d Number said: %d\n", turnNumber, numberSpoken)
			}
			turnNumber++
			previousSpoken = numberSpoken
		} else {
			// If we have no memory of the number, it must be new so say "0"
			numberSpoken = 0
			memory[previousSpoken] = turnNumber - 1
			if debug {
				fmt.Printf("New Turn: %d Number said: %d\n", turnNumber, numberSpoken)
			}
			turnNumber++
			previousSpoken = numberSpoken
		}

		if debug {
			fmt.Println(memory)
		}

		if turnNumber == lastRound {
			break
		}
	}

	return numberSpoken
}

// Main routine
func main() {
	_, execPart, debug, test := catchUserInput()

	if test {
		return
	}

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else if execPart == 'a' {
		fmt.Println("2020th number spoken for [0,3,6]:", playNumberGame("0,3,6", 2020, execPart, debug))
		fmt.Println("2020th number spoken for [1,3,2]:", playNumberGame("1,3,2", 2020, execPart, debug))
		fmt.Println("2020th number spoken for [2,1,3]:", playNumberGame("2,1,3", 2020, execPart, debug))
		fmt.Println("2020th number spoken for [1,2,3]:", playNumberGame("1,2,3", 2020, execPart, debug))
		fmt.Println("2020th number spoken for [2,3,1]:", playNumberGame("2,3,1", 2020, execPart, debug))
		fmt.Println("2020th number spoken for [3,2,1]:", playNumberGame("3,2,1", 2020, execPart, debug))
		fmt.Println("2020th number spoken for [3,1,2]:", playNumberGame("3,1,2", 2020, execPart, debug))
		fmt.Println("2020th number spoken for [1,0,15,2,10,13]:", playNumberGame("1,0,15,2,10,13", 2020, execPart, debug))
	} else {
		fmt.Println("30000000th number spoken for [0,3,6]:", playNumberGame("0,3,6", 30000000, execPart, debug))
		fmt.Println("30000000th number spoken for [1,3,2]:", playNumberGame("1,3,2", 30000000, execPart, debug))
		fmt.Println("30000000th number spoken for [2,1,3]:", playNumberGame("2,1,3", 30000000, execPart, debug))
		fmt.Println("30000000th number spoken for [1,2,3]:", playNumberGame("1,2,3", 30000000, execPart, debug))
		fmt.Println("30000000th number spoken for [2,3,1]:", playNumberGame("2,3,1", 30000000, execPart, debug))
		fmt.Println("30000000th number spoken for [3,2,1]:", playNumberGame("3,2,1", 30000000, execPart, debug))
		fmt.Println("30000000th number spoken for [3,1,2]:", playNumberGame("3,1,2", 30000000, execPart, debug))
		fmt.Println("30000000th number spoken for [1,0,15,2,10,13]:", playNumberGame("1,0,15,2,10,13", 30000000, execPart, debug))
	}
}
