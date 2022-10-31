package main

import (
	"AdventOfCode-go/advent2017/utils"
	"fmt"
)

func calcCaptcha(captcha string, jumpAheadBy int) int {
	/*
		The captcha requires you to review a sequence of digits (your puzzle input) and
		find the sum of all digits that match the next digit in the list. The list is
		circular, so the digit after the last digit is the first digit in the list.
	*/
	var result int
	var jumpAheadPos int
	var firstDigit, secondDigit int
	if len(captcha) == 0 {
		return 0
	}
	for i := 0; i < len(captcha); i++ {
		jumpAheadPos = (i + jumpAheadBy) % len(captcha)
		firstDigit = int(captcha[i]) - '0'
		secondDigit = int(captcha[jumpAheadPos]) - '0'
		if firstDigit == secondDigit {
			result += firstDigit
		}
	}

	return result
}

func solveCaptcha(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)
	if part == 'a' {
		return calcCaptcha(puzzleInput[0], 1)
	} else {
		return calcCaptcha((puzzleInput[0]), len(puzzleInput[0])/2)
	}
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", solveCaptcha(filenamePtr, execPart, debug))
	}
}
