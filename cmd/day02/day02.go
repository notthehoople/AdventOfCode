package main

import (
	"aoc/advent2021/utils"
	"fmt"
	"strconv"
	"strings"
)

func calcDepthAndPosition(movementCommands []string) int {
	/*
		forward X increases the horizontal position by X units.
		down X increases the depth by X units.
		up X decreases the depth by X units.
	*/
	var currentPosition, currentDepth, moveAmount int
	var moveCommand []string

	for i := 0; i < len(movementCommands); i++ {
		moveCommand = strings.Split(movementCommands[i], " ")
		moveAmount, _ = strconv.Atoi(moveCommand[1])

		switch moveCommand[0] {
		case "forward":
			currentPosition += moveAmount
		case "down":
			currentDepth += moveAmount
		case "up":
			currentDepth -= moveAmount
		}
	}

	return currentPosition * currentDepth
}

func solveDay(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)
	if part == 'a' {
		return calcDepthAndPosition(puzzleInput)
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
