package main

import (
	"AdventOfCode-go/advent2024/utils"
	"fmt"
	"strconv"
	"strings"
)

// countNumber used in part b to count the number of times an element appears in a list
func isReportSafe(reportToCheck []int) bool {
	var ascending string = "unset"

	for i:=1; i<len(reportToCheck); i++ {
		if (utils.Abs(reportToCheck[i] - reportToCheck[i-1]) > 3) ||
			(reportToCheck[i] == reportToCheck[i-1]) {
			return false
		}

		// There's no edge case where a number is the same as the previous. That's a fail
		switch ascending {
		case "unset":
			if (reportToCheck[i] > reportToCheck[i-1]) {
					ascending = "TRUE"
			} else {
				ascending = "FALSE"
			}
		case "TRUE":
			if reportToCheck[i] < reportToCheck[i-1] {
				return false
			}
		case "FALSE":
			if reportToCheck[i] > reportToCheck[i-1] {
				return false
			}
		}
		
	}
	return true
}

func day02(filename string, part byte, debug bool) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)

	if part == 'a' {
		// Part 1: Find the distances between the 2 lists.

		for _, puzzleLine := range puzzleInput {
			puzzleItems := strings.Split(puzzleLine, " ")

			var reportVals []int
			for _, value := range puzzleItems {
				singleVal, _ := strconv.Atoi(value)
				reportVals = append(reportVals, singleVal)
			}

			if isReportSafe(reportVals) {
				result++
			}
		}

		//return calcDistance(firstList, secondList)
		return result
	}

	return result
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", day02(filenamePtr, execPart, debug))
	}
}
