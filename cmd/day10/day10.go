package main

import (
	"aoc/advent2021/utils"
	"fmt"
)

var chunkPartners = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

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
		var syntaxErrorScore int

		for _, navigationLine := range puzzleInput {
			syntaxErrorScore += calcScoreOnNavigationLine(navigationLine, debug)
		}
		return syntaxErrorScore
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
		fmt.Printf("Result is: %d\n", solveDay(filenamePtr, execPart, debug))
	}
}
