package main

import (
	"AdventOfCode-go/advent2025/utils"
	"fmt"
	"strconv"
	"strings"
)

// Extracts a number from a position in an array of strings. Number is vertically presented in the string
func extractNumber(stringArray []string, pos int) int {
	var number int

	for i := 0; i < len(stringArray)-1; i++ {
		if stringArray[i][pos] != ' ' {
			number = number*10 + int(stringArray[i][pos]-'0')
		}
	}

	return number
}

func day06(filename string, part byte, debug bool) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)

	// Get the size of the fresh ingredients array that we need
	countArraySize := len(puzzleInput)

	holdArray := make([][]string, countArraySize)
	for i, puzzleLine := range puzzleInput {
		// Remove leading and trailing white space, and split based on whitespace
		holdArray[i] = strings.Fields(puzzleLine)
	}

	if part == 'a' {
		for sumNumber := 0; sumNumber < len(holdArray[0]); sumNumber++ {
			// Work out our numbers
			sumResult, _ := strconv.Atoi(holdArray[0][sumNumber])

			for j := 1; j < len(holdArray)-1; j++ {
				sumTemp, _ := strconv.Atoi(holdArray[j][sumNumber])

				if holdArray[len(holdArray)-1][sumNumber] == "+" {
					sumResult += sumTemp
				} else {
					sumResult *= sumTemp
				}
			}
			result += sumResult
		}
	} else {
		// Position is everything, so can't use the previous method
		// working backwards from the end of the string, pull out numbers for each of the character positions for each line
		// once build a number, add it to the current sum
		// once all numbers have been collected, run through the sums and produce the answer

		var maxLength int
		for _, puzzleLine := range puzzleInput {
			if maxLength < len(puzzleLine) {
				maxLength = len(puzzleLine)
				// Note: all the lines are the same length in the input file, so we don't need to cater for this
			}
		}

		positionsFound := len(holdArray[len(holdArray)-1]) - 1
		operatorPos := len(holdArray) - 1

		var tempSum int
		var firstDigit bool = true
		for stringPos := len(puzzleInput[0]) - 1; stringPos >= 0; stringPos-- {
			numberFound := extractNumber(puzzleInput, stringPos)

			if numberFound == 0 {
				// No number found so we are done with the sum - there's only a single blank between each sum
				result += tempSum
				positionsFound--
				tempSum = 0
				firstDigit = true
			} else {
				if holdArray[operatorPos][positionsFound] == "+" {
					tempSum += numberFound
				} else {
					if firstDigit {
						tempSum = numberFound
						firstDigit = false
					} else {
						tempSum *= numberFound
					}
				}
			}
		}
		result += tempSum // Catch the last sum that we missed as the for loop exited
	}

	return result
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", day06(filenamePtr, execPart, debug))
	}
}
