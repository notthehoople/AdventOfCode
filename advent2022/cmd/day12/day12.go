package main

import (
	"AdventOfCode-go/advent2022/utils"
	"fmt"
)

type Coords struct {
	x int
	y int
}

func (c Coords) printPosition() {
	fmt.Printf("x: %d y: %d\n", c.x, c.y)
}

func calcFewestSteps(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)

	areaArray := buildAreaArray(puzzleInput)

	return 0
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Monkey Business: %d\n", calcFewestSteps(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("Not implemented yet\n")
	case 'z':
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
