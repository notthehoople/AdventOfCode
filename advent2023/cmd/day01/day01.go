package main

import (
	"AdventOfCode-go/advent2023/utils"
	"fmt"
)

// Part 1: look for the first and last digit (in that order) and use them to make a 2 digit number
//   Sum al the 2 digit numbers together
//   Edge case: if there's only 1 number in the string, use it for both digits

func findNumber(puzzleLine string, forwards bool) int {
	var startPos, stopPos, step int

	if forwards {
		startPos = 0
		stopPos = len(puzzleLine)
		step = 1
	} else {
		startPos = len(puzzleLine) - 1
		stopPos = -1
		step = -1
	}

	for i := startPos; i != stopPos; i += step {
		if puzzleLine[i] > '0' && puzzleLine[i] <= '9' {
			return int(puzzleLine[i] - '0')
		}
	}
	return 0
}

func day01(filename string, part byte, debug bool) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)

	for _, puzzleLine := range puzzleInput {
		firstDigit := findNumber(puzzleLine, true)
		secondDigit := findNumber(puzzleLine, false)

		result += (firstDigit * 10) + secondDigit
		//fmt.Println("Answer:", firstDigit, secondDigit)
	}

	if part == 'a' {
		return result
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
