package main

import (
	"AdventOfCode-go/advent2023/utils"
	"fmt"
	"strconv"
	"strings"
)

// File: First line is the movement instructions. Collection of L or R letters to signify movement

func day09(filename string, part byte, debug bool) int {
	var extrapolatedValueSum int

	puzzleInput, _ := utils.ReadFile(filename)
	for _, puzzleLine := range puzzleInput {
		sequenceStr := strings.Fields(puzzleLine)
		sequence := make([]int, len(sequenceStr))

		for step, value := range sequenceStr {
			sequence[step], _ = strconv.Atoi(value)
		}

		// Loop through the sequence array. Calulate the difference between the numbers
		// If the resulting array of numbers are not all equal, loop again using the result array of numbers as the input
		// Once you get to a resulting array of numbers being all equal, use that array (and the others) to calc next number

		// This produces too many array items. HOW TO FIX THIS?
		resultsArray := make([][]int, len(sequence))
		var notFinished bool = true

		stepCount := 1
		resultsArray[0] = sequence

		for i := 0; i < len(sequence)-1 && notFinished; i++ {
			resultsArray[stepCount] = make([]int, len(sequence)-stepCount)

			var prevResult int
			for j := 0; j < len(resultsArray[stepCount-1])-1; j++ {
				var calcResult = resultsArray[i][j+1] - resultsArray[i][j]
				resultsArray[stepCount][j] = calcResult
				if prevResult == calcResult {
					notFinished = false
				} else {
					notFinished = true
				}

				prevResult = calcResult
			}

			stepCount++
		}

		// We now have the built array of sequences. Now to predict the next in the sequence

		var resultAdd, predictAdd int
		for i := len(resultsArray) - 1; i >= 0; i-- {
			if resultsArray[i] != nil {
				if debug {
					fmt.Println("Last digit:", resultsArray[i][len(resultsArray[i])-1])
				}

				predictAdd = predictAdd + resultsArray[i][len(resultsArray[i])-1]
				if i > 0 {
					resultAdd = resultsArray[i-1][len(resultsArray[i-1])-1] + predictAdd
				}
			}
		}
		extrapolatedValueSum += resultAdd
	}

	return extrapolatedValueSum
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Result is: %d\n", day09(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("Result is: %d\n", day09(filenamePtr, execPart, debug))
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
