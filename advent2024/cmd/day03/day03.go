package main

import (
	"AdventOfCode-go/advent2024/utils"
	"fmt"
	"regexp"
)

func day03(filename string, part byte, debug bool) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)

	if part == 'a' {
		for _, puzzleLine := range puzzleInput {
			reg := regexp.MustCompile(`mul\(\d+,\d+\)`)
			regMatched := reg.FindAllStringSubmatch(puzzleLine, -1)
		
			for _, item := range regMatched {
				var firstNum, secondNum int
				fmt.Sscanf(item[0], "mul(%d,%d)", &firstNum, &secondNum)
				result += firstNum * secondNum
			}
		}
		return result
	}

	// Part 2: we need to take notice of the do() and don't() commands in the instruction list
	var useStatement bool = true

	for _, puzzleLine := range puzzleInput {
		reg := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
		regMatched := reg.FindAllStringSubmatch(puzzleLine, -1)
	
		if debug {
			fmt.Println(regMatched)
		}
		for _, item := range regMatched {

			if debug {
				fmt.Println(item)
			}

			switch item[0] {
			case "do()":
				useStatement = true
			case "don't()":
				useStatement = false
			default:
				if useStatement {
					var firstNum, secondNum int
					fmt.Sscanf(item[0], "mul(%d,%d)", &firstNum, &secondNum)
					result += firstNum * secondNum
				}
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
		fmt.Printf("Result is: %d\n", day03(filenamePtr, execPart, debug))
	}
}
