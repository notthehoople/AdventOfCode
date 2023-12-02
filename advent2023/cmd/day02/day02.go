package main

import (
	"AdventOfCode-go/advent2023/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func validGame(puzzleLine string, maxRed int, maxGreen int, maxBlue int) int {
	// Pull out the game number from the puzzleLine
	// Split the line by ';'
	// Pull out the no. of red cubes, blue cubes, green cubes
	// test against the maximums we've been given. If good continue; if bad fail

	var gameNumber int
	var remainingLine string

	fmt.Sscanf(puzzleLine, "Game %d:", &gameNumber)
	remainingLine = strings.Split(puzzleLine, ": ")[1]

	for _, cubeSet := range strings.Split(remainingLine, "; ") {

		reg := regexp.MustCompile(`(\d+) (\w+)`)
		matched := reg.FindAllStringSubmatch(cubeSet, -1)

		for _, match := range matched {
			numCubes, _ := strconv.Atoi(match[1])
			switch match[2] {
			case "red":
				if numCubes > maxRed {
					return 0
				}
			case "green":
				if numCubes > maxGreen {
					return 0
				}
			case "blue":
				if numCubes > maxBlue {
					return 0
				}
			}

		}

	}

	return gameNumber
}

func day02(filename string, part byte, maxRed int, maxGreen int, maxBlue int, debug bool) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)

	if part == 'a' {
		for _, puzzleLine := range puzzleInput {

			result += validGame(puzzleLine, maxRed, maxGreen, maxBlue)
		}
		return result
	}

	// Part b

	return 0
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		var red, green, blue int = 12, 13, 14
		fmt.Printf("Result is: %d\n", day02(filenamePtr, execPart, red, green, blue, debug))
	}
}
