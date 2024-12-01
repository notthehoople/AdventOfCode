package main

import (
	"AdventOfCode-go/advent2024/utils"
	"fmt"
	"slices"
)

// countNumber used in part b to count the number of times an element appears in a list
func countNumber(listToSearch []int, intToFind int) int {
	var count int
	for _, item := range listToSearch {
		if item == intToFind {
			count++
		}
	}
	return count
}

// calcDistance used in part a to calc the distance between slice elements
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

	for i, puzzleLine := range puzzleInput {
		fmt.Sscanf(puzzleLine, "%d   %d\n", &firstList[i], &secondList[i])
	}

	if part == 'a' {
		// Part 1: Find the distances between the 2 lists.

		slices.Sort(firstList)
		slices.Sort(secondList)
		
		return calcDistance(firstList, secondList)
	}

	// Part B - find the similarity score between the two lists
	for _, item := range firstList {
		count := countNumber(secondList, item)
		result += count * item
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
