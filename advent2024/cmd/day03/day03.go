package main

import (
	"AdventOfCode-go/advent2024/utils"
	"fmt"
	"regexp"
)

func day03(filename string, part byte, debug bool) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)

	for _, puzzleLine := range puzzleInput {

		reg := regexp.MustCompile(`mul\(\d+,\d+\)`)
		regMatched := reg.FindAllStringSubmatch(puzzleLine, -1)
		//fmt.Println(regMatched, len(regMatched))
		
		for _, item := range regMatched {
			var firstNum, secondNum int

			fmt.Sscanf(item[0], "mul(%d,%d)", &firstNum, &secondNum)
			//fmt.Println(item[0], firstNum, secondNum)

			result += firstNum * secondNum
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
		fmt.Printf("Result is: %d\n", day03(filenamePtr, execPart, debug))
	}
}
