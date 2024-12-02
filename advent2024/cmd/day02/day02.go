package main

import (
	"AdventOfCode-go/advent2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func isReportSafeWithTolerate(reportToCheck []int) bool {
	// Check using isReportSafe
	// Check if removing a digit makes it safe

	if isReportSafe(reportToCheck) {
		return true
	}

	// Ok, so the report isn't safe. Let's remove a digit at a time and re-test
	for i:=0; i < len(reportToCheck); i++ {

		slicePos := 0
		tolerateSlice := make([]int, len(reportToCheck)-1)
		for j:=0; j < len(reportToCheck); j++ {
			if j != i {
				tolerateSlice[slicePos] = reportToCheck[j]
				slicePos++
			}
		}

		if isReportSafe(tolerateSlice) {
			return true
		}

	}
	return false
}

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

	for _, puzzleLine := range puzzleInput {
		puzzleItems := strings.Split(puzzleLine, " ")

		var reportVals []int
		for _, value := range puzzleItems {
			singleVal, _ := strconv.Atoi(value)
			reportVals = append(reportVals, singleVal)
		}

		if part == 'a' {
			if isReportSafe(reportVals) {
				result++
			}
		} else {
			// part b
			if isReportSafeWithTolerate(reportVals) {
				result++
			}
		}
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
