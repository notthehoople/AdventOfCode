package main

import (
	"aoc/advent2021/utils"
	"fmt"
	"sort"
)

var chunkPartners = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var closePartners = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

// Part b scoring
var autocorrectScore = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func calcAutocorrectCompletion(navigationLine string, debug bool) int {
	var chunkList []rune

	if debug {
		fmt.Println("====================================")
		fmt.Println("Line:", navigationLine)
	}

	// First, let's deal with the easy closing chunks from part a
	for _, bracket := range navigationLine {

		if openBracket, ok := chunkPartners[bracket]; ok {
			if chunkList[len(chunkList)-1] == openBracket {
				chunkList = chunkList[:len(chunkList)-1]
			}
		} else {
			// Build a list of the brackets we find
			chunkList = append(chunkList, bracket)
		}
	}
	if debug {
		fmt.Println("Remaining list to complete:", chunkList)
	}

	// Need to get to the end of the string and start working backwards, closing as we go
	var scoreResult int
	var keepGoing bool = true

	for keepGoing {
		if len(chunkList)-1 < 0 {
			keepGoing = false
		} else {
			currentBracket := chunkList[len(chunkList)-1]
			if closeBracket, ok := closePartners[currentBracket]; ok {
				chunkList = chunkList[:len(chunkList)-1]
				scoreResult *= 5
				scoreResult += autocorrectScore[closeBracket]
			}
		}
	}
	return scoreResult
}

// Part a scoring
var closeBracketScore = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func calcScoreOnNavigationLine(navigationLine string, debug bool) int {
	var chunkList []rune

	if debug {
		fmt.Println("====================================")
		fmt.Println("Line:", navigationLine)
	}

	for _, bracket := range navigationLine {
		if debug {
			fmt.Println("ChunkList:", chunkList)
		}

		if openBracket, ok := chunkPartners[bracket]; ok {
			if debug {
				fmt.Printf("We have opener:%c and closer:%c\n", openBracket, bracket)
			}
			if chunkList[len(chunkList)-1] == openBracket {
				if debug {
					fmt.Println("Found a match:", chunkList)
				}

				chunkList = chunkList[:len(chunkList)-1]
			} else {
				return closeBracketScore[bracket]
			}
		} else {
			// Build a list of the brackets we find
			chunkList = append(chunkList, bracket)
		}
	}
	return 0
}

func solveDay(filename string, part byte, debug bool) int {
	puzzleInput, _ := utils.ReadFile(filename)

	if part == 'a' {
		// Calc a score on the corrupted lines
		var syntaxErrorScore int

		for _, navigationLine := range puzzleInput {
			syntaxErrorScore += calcScoreOnNavigationLine(navigationLine, debug)
		}
		return syntaxErrorScore
	} else {
		var autocorrectScore []int

		for _, navigationLine := range puzzleInput {
			// Incomplete lines score '0' for part a so let's use that to filter out corrupt lines
			if calcScoreOnNavigationLine(navigationLine, debug) == 0 {
				autocorrectScore = append(autocorrectScore, calcAutocorrectCompletion(navigationLine, debug))
			}
		}

		sort.Ints(autocorrectScore)
		if debug {
			fmt.Println(autocorrectScore)
		}
		return autocorrectScore[len(autocorrectScore)/2]
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
