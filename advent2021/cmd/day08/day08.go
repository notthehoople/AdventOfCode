package main

import (
	"aoc/advent2021/utils"
	"fmt"
	"sort"
	"strings"
)

func decodeOutputDigits(displayDigits []string, part byte, debug bool) int {
	var totalOutput int
	for _, i := range displayDigits {
		digitsSplit := strings.Split(i, " | ")

		var singleOutputAnswer int
		for outputPos, outputDigit := range strings.Split(digitsSplit[1], " ") {

			if outputPos > 0 {
				singleOutputAnswer *= 10 ^ outputPos
			}

			switch len(outputDigit) {
			case 2: // Digit 1
				singleOutputAnswer += 1
			case 3: // Digit 7
				singleOutputAnswer += 7
			case 4: // Digit 4
				singleOutputAnswer += 4
			case 7: // Digit 8
				singleOutputAnswer += 8
			default:
				sorted := []rune(outputDigit)
				sort.Slice(sorted, func(i int, j int) bool { return sorted[i] < sorted[j] })
				outputDigit = string(sorted)

				fmt.Printf("Switching: %s\n", outputDigit)
				switch outputDigit {
				case "ab": // Digit 1
					singleOutputAnswer += 1
				case "acdfg": // Digit 2
					singleOutputAnswer += 2
				case "abcdf": // Digit 3
					singleOutputAnswer += 3
				case "abef": // Digit 4
					singleOutputAnswer += 4
				case "bcdef": // Digit 5
					singleOutputAnswer += 5
				case "bcdefg": // Digit 6
					singleOutputAnswer += 6
				case "abd": // Digit 7
					singleOutputAnswer += 7
				case "abcdefg": // Digit 8
					singleOutputAnswer += 8
				case "abcdef": // Digit 9
					singleOutputAnswer += 9
				case "abcdeg": // Digit 0
				default:
					fmt.Println("DISASTER DISASTER DISASTER:", outputDigit)
					return 0
				}
			}
		}
		totalOutput += singleOutputAnswer
	}
	return totalOutput
}

func countUniqueNumbers(displayDigits []string, part byte, debug bool) int {
	var countUnique int
	for _, i := range displayDigits {
		digitsSplit := strings.Split(i, " | ")

		for _, outputDigit := range strings.Split(digitsSplit[1], " ") {
			switch len(outputDigit) {
			case 2: // Digit 1
				countUnique++
			case 3: // Digit 7
				countUnique++
			case 4: // Digit 4
				countUnique++
			case 7: // Digit 8
				countUnique++
			default:
				// Not a unique digit
			}
		}
	}
	return countUnique
}

func solveDay(filename string, part byte, debug bool) int {
	puzzleInput, _ := utils.ReadFile(filename)

	if part == 'a' {
		return countUniqueNumbers(puzzleInput, part, debug)
	} else {
		return decodeOutputDigits(puzzleInput, part, debug)
	}
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", solveDay(filenamePtr, execPart, debug))
	}
}
