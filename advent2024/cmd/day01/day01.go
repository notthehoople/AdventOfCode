package main

import (
	"AdventOfCode-go/advent2024/utils"
	"fmt"
	"slices"
)

// Part 1: 

func calcDistance(firstList []int, secondList []int) int {
	var result int

	for i := 0; i < len(firstList); i ++ {
		result += utils.Abs(firstList[i] - secondList[i])
	}
	return result
}

func day01(filename string, part byte, debug bool) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)
	inputLength := len(puzzleInput)

	firstList := make([]int, inputLength)
	secondList := make([]int, inputLength)

	if part == 'a' {
		for i, puzzleLine := range puzzleInput {
			fmt.Sscanf(puzzleLine, "%d   %d\n", &firstList[i], &secondList[i])
		}

		slices.Sort(firstList)
		slices.Sort(secondList)
		
		return calcDistance(firstList, secondList)
	}

	return result
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
