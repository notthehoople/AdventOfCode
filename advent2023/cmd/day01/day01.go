package main

import (
	"AdventOfCode-go/advent2023/utils"
	"fmt"
)

func day01(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)

	fmt.Printf(puzzleInput[0], debug)

	if part == 'a' {
		return 0
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
		fmt.Printf("Result is: %d\n", day01(filenamePtr, execPart, debug))
	}
}
