package main

import (
	"aoc/advent2021/utils"
	"fmt"
	"strconv"
	"strings"
)

func calcDepthAndPositionUsingAim(movementCommands []string) int {
	/*
	   	down X increases your aim by X units.
	   	up X decreases your aim by X units.
	   	forward X does two things:
	       	It increases your horizontal position by X units.
	       	It increases your depth by your aim multiplied by X.
	*/
	var currentPosition, currentDepth, currentAim, moveAmount int
	var moveCommand []string

	for i := 0; i < len(movementCommands); i++ {
		moveCommand = strings.Split(movementCommands[i], " ")
		moveAmount, _ = strconv.Atoi(moveCommand[1])

		switch moveCommand[0] {
		case "forward":
			currentPosition += moveAmount
			currentDepth += currentAim * moveAmount
		case "down":
			currentAim += moveAmount
		case "up":
			currentAim -= moveAmount
		}
	}

	return currentPosition * currentDepth
}

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
		return calcDepthAndPositionUsingAim(puzzleInput)
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
