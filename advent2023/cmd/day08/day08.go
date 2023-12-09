package main

import (
	"AdventOfCode-go/advent2023/utils"
	"fmt"
	"strings"
)

// File: First line is the movement instructions. Collection of L or R letters to signify movement
//    Rest of the file contains direction nodes which consist of a left element and a right element (BBB, CCC)
//    When running the movement phase the L or R command tells you which element to pick from where you are located

type locNode struct {
	left  string
	right string
}

func day08(filename string, part byte, debug bool) int {

	locIndex := make(map[string]locNode)

	puzzleInput, _ := utils.ReadFile(filename)

	var movements string
	for line, puzzleLine := range puzzleInput {
		if line == 0 {
			movements = puzzleLine
			continue
		}
		if len(puzzleLine) == 0 {
			continue
		}

		var location, left, right string
		fields := strings.Fields(puzzleLine)
		location = fields[0]
		left = fields[2][1:4]
		right = fields[3][0:3]

		locIndex[location] = locNode{left, right}
		if debug {
			fmt.Printf("location: %s left: %s right: %s map: %v\n", location, left, right, locIndex[location])
		}
	}

	if debug {
		fmt.Println(movements)
		fmt.Println(locIndex)
	}

	var currStep int
	var currLoc string = "AAA"
	var totalSteps int

	for while := false; !while; {
		currMove := movements[currStep%len(movements)]
		if debug {
			fmt.Println("currLoc:", currLoc)
		}

		switch currMove {
		case 'L':
			currLoc = locIndex[currLoc].left
		case 'R':
			currLoc = locIndex[currLoc].right
		default:
			panic("Bad move in the movements list")
		}

		totalSteps++
		if currLoc == "ZZZ" {
			while = true
		}

		currStep++
	}

	return totalSteps
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Result is: %d\n", day08(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("Result is: %d\n", day08(filenamePtr, execPart, debug))
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
