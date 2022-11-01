package main

import (
	"AdventOfCode-go/advent2017/utils"
	"fmt"
	"strconv"
)

func getOutOfMaze(filename string, part byte, debug bool) int {
	//var instructionList []int

	puzzleInput, _ := utils.ReadFile(filename)
	instructionList := make([]int, len(puzzleInput))

	// Loop through passphrases
	// if valid passphrase then count++
	for i := 0; i < len(puzzleInput); i++ {
		instructionList[i], _ = strconv.Atoi(puzzleInput[i])
	}

	if debug {
		fmt.Println(instructionList)
	}

	var movement, instructionPos, stepsTaken int
	instructionPos = 0
	stepsTaken = 0
	for {
		movement = instructionList[instructionPos]
		instructionList[instructionPos]++

		instructionPos += movement

		if debug {
			fmt.Println(instructionList)
			fmt.Println("====================")
		}

		stepsTaken++
		if instructionPos > len(instructionList)-1 {
			return stepsTaken
		}
	}

	// Run through the list and carry out jumps.
	// If we go out of the list then we're finished
	// Count number of steps completed to exit the list

	return 0
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", getOutOfMaze(filenamePtr, execPart, debug))
	}
}
