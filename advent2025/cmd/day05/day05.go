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

func smallest(first int, second int) int {
	if first < second {
		return first
	}
	return second
}

func largest(first int, second int) int {
	if first > second {
		return first
	}
	return second
}

func overlappingRange(first freshRange, second freshRange) bool {
	if ((first.end >= second.start) && (first.end <= second.end)) ||
		((first.start >= second.start) && (first.start <= second.end)) {
		return true
	}
	return false
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

			if part == 'a' {
				// Check each ingredient in the fresh ingredient DB we previously built
				ingredient, _ := strconv.Atoi(puzzleLine)
				for _, freshStuff := range ingredientDB {
					if ingredient >= freshStuff.start && ingredient <= freshStuff.end {
						result++
						break
					}
				}
			} else {
				break
			}
		}
	}

	if part == 'a' {
		return result
	}

	// part b. Don't process the ingredient list. Focus on the DB

	// Need to work out how many unique fresh IDs there are
	// Use the ranges in the DB
	// look for overlapping ranges and COMBINE them
	// once everything has been combined, add up the remaining IDs

	fmt.Println(ingredientDB)
	expandedDB := make([]freshRange, countArraySize)

	for i := 0; i < len(ingredientDB); i++ {
		for j := i + 1; j < len(ingredientDB); j++ {
			if overlappingRange(ingredientDB[i], ingredientDB[j]) {
				// Combine the ranges
				fmt.Println("Overlapping ranges:", ingredientDB[i], ingredientDB[j])

				expandedDB[i].start = smallest(ingredientDB[i].start, ingredientDB[j].start)
				expandedDB[i].end = largest(ingredientDB[i].end, ingredientDB[j].end)
				// To combine the ranges, take the smallest start and the largest end and add them together
			}
		}
		// If nothing has been put into the expandedDB array at this point, just copy from ingredientsDB
		// Need to deal with multiple entries. How?
		// Once an overlap has been found should we remove that entry from the ingredientDB?
	}
	fmt.Println(expandedDB)

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
