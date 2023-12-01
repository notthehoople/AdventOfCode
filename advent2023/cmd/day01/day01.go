package main

import (
	"AdventOfCode-go/advent2023/utils"
	"fmt"
	"strings"
)

// Part 1: look for the first and last digit (in that order) and use them to make a 2 digit number
//   Sum al the 2 digit numbers together
//   Edge case: if there's only 1 number in the string, use it for both digits

// Part 2: digits can be spelt out OR numerical

func findNumberAndText(puzzleLine string, forwards bool) int {
	var startPos, stopPos, step int
	var bestPos, bestValue int
	digits := [...]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	// First, look for text numbers and record the lowest (forwards=true)
	//   or highest (forwards=false) positions where they are found

	if forwards {
		bestPos = len(puzzleLine)
	} else {
		bestPos = -1
	}

	for textValue, textNum := range digits {
		var foundPos int
		if forwards {
			foundPos = strings.Index(puzzleLine, textNum)
		} else {
			foundPos = strings.LastIndex(puzzleLine, textNum)
		}
		if foundPos != -1 {
			if forwards && foundPos < bestPos {
				bestPos = foundPos
				bestValue = textValue
			}
			if !forwards && foundPos > bestPos {
				bestPos = foundPos
				bestValue = textValue
			}
		}
	}

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
			if forwards {
				if i < bestPos {
					return int(puzzleLine[i] - '0')
				} else {
					return bestValue
				}
			} else {
				if i > bestPos {
					return int(puzzleLine[i] - '0')
				} else {
					return bestValue
				}
			}
		}
	}

	return bestValue
}

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

	if part == 'a' {
		for _, puzzleLine := range puzzleInput {
			firstDigit := findNumber(puzzleLine, true)
			secondDigit := findNumber(puzzleLine, false)

			result += (firstDigit * 10) + secondDigit
		}
		return result
	}

	// Part b
	for _, puzzleLine := range puzzleInput {
		firstDigit := findNumberAndText(puzzleLine, true)
		secondDigit := findNumberAndText(puzzleLine, false)
		result += (firstDigit * 10) + secondDigit
		if debug {
			fmt.Println(puzzleLine, (firstDigit*10)+secondDigit)
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
		fmt.Printf("Result is: %d\n", day01(filenamePtr, execPart, debug))
	}
}
