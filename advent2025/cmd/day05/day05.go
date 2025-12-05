package main

import (
	"AdventOfCode-go/advent2025/utils"
	"fmt"
	"strconv"
	"strings"
)

type freshRange struct {
	start int
	end   int
}

func day05(filename string, part byte, debug bool) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)

	// Get the size of the fresh ingredients array that we need
	var countArraySize int
	for _, puzzleLine := range puzzleInput {
		if len(puzzleLine) == 0 {
			break
		}
		countArraySize++
	}
	if debug {
		fmt.Println("Array Size:", countArraySize)
	}

	ingredientDB := make([]freshRange, countArraySize)

	var checkIngredients bool = false
	for i, puzzleLine := range puzzleInput {
		if !checkIngredients {
			// Process fresh ingredient ID ranges until we see the blank line

			if len(puzzleLine) == 0 {
				if debug {
					fmt.Println("found the blank line")
				}
				checkIngredients = true
			} else {
				ingredientRange := strings.Split(puzzleLine, "-")
				ingredientDB[i].start, _ = strconv.Atoi(ingredientRange[0])
				ingredientDB[i].end, _ = strconv.Atoi(ingredientRange[1])
			}
		} else {
			// Check each ingredient in the fresh ingredient DB we previously built
			ingredient, _ := strconv.Atoi(puzzleLine)
			for _, freshStuff := range ingredientDB {
				if ingredient >= freshStuff.start && ingredient <= freshStuff.end {
					result++
					break
				}
			}
		}
	}

	//fmt.Println(ingredientDB)

	return result
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", day05(filenamePtr, execPart, debug))
	}
}
