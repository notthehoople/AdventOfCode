package main

import (
	"AdventOfCode-go/advent2022/utils"
	"fmt"
	"strconv"
)

func elfMunchies(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)
	if part == 'a' {
		var elf int = 1
		var topElf int
		var topElfCalories int
		var calories int
		for _, j := range puzzleInput {

			if j == "" {
				fmt.Println("Blank Line found")
				if calories > topElfCalories {
					topElf = elf
					topElfCalories = calories
				}
				calories = 0
				elf++
			} else {
				foodCalories, _ := strconv.Atoi(j)
				fmt.Printf("Elf: %d Carrying food: %d\n", elf, foodCalories)
				calories += foodCalories
			}
		}

		fmt.Printf("Top Elf: %d Carrying calories: %d\n", topElf, topElfCalories)

		return topElfCalories
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
		fmt.Printf("Result is: %d\n", elfMunchies(filenamePtr, execPart, debug))
	}
}
