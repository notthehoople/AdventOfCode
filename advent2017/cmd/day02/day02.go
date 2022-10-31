package main

import (
	"AdventOfCode-go/advent2017/utils"
	"fmt"
	"strconv"
	"strings"
)

func calcRowDifferenceDivisors(row string, debug bool) int {
	/*
		For each row, determine the difference between the largest value and the smallest value;
		the checksum is the sum of all of these differences.
	*/

	var currValue1, currValue2 int

	for _, i := range strings.Fields(row) {
		for _, j := range strings.Fields(row) {
			if i != j {
				currValue1, _ = strconv.Atoi(i)
				currValue2, _ = strconv.Atoi(j)
				if currValue1 > currValue2 {
					if currValue1%currValue2 == 0 {
						return currValue1 / currValue2
					}
				} else {
					if currValue2%currValue1 == 0 {
						return currValue2 / currValue1
					}
				}
			}
		}
	}
	return 0
}

func calcRowDifference(row string, debug bool) int {
	/*
		For each row, determine the difference between the largest value and the smallest value;
		the checksum is the sum of all of these differences.
	*/
	var minValue int = 9999999
	var maxValue int = 0
	var currValue int

	for _, j := range strings.Fields(row) {
		currValue, _ = strconv.Atoi(j)
		if currValue < minValue {
			minValue = currValue
		}
		if currValue > maxValue {
			maxValue = currValue
		}
	}

	return maxValue - minValue
}

func solveChecksum(filename string, part byte, debug bool) int {
	var checksum int

	puzzleInput, _ := utils.ReadFile(filename)
	if part == 'a' {
		for i := range puzzleInput {
			checksum += calcRowDifference(puzzleInput[i], debug)
		}

		return checksum
	} else {
		for i := range puzzleInput {
			checksum += calcRowDifferenceDivisors(puzzleInput[i], debug)
		}

		return checksum
	}
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", solveChecksum(filenamePtr, execPart, debug))
	}
}
