package main

import (
	"AdventOfCode-go/advent2025/utils"
	"fmt"
	"strconv"
	"strings"
)

func day05(filename string, part byte, debug bool) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)

	ingredientDB := make(map[int]bool)

	var checkIngredients bool = false
	for _, puzzleLine := range puzzleInput {
		if !checkIngredients {
			// Process fresh ingredient ID ranges until we see the blank line

			if len(puzzleLine) == 0 {
				fmt.Println("found the blank line")
				checkIngredients = true
			} else {
				ingredientRange := strings.Split(puzzleLine, "-")
				startRange, _ := strconv.Atoi(ingredientRange[0])
				endRange, _ := strconv.Atoi(ingredientRange[1])

				for id := startRange; id <= endRange; id++ {
					//fmt.Println("id", id)
					ingredientDB[id] = true
				}
			}
		} else {
			// Check each ingredient in the map we previously built
			ingredient, _ := strconv.Atoi(puzzleLine)
			if ingredientDB[ingredient] {
				//fmt.Printf("Ingredient %d is fresh\n", ingredient)
				result++
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
