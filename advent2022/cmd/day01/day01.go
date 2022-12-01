package main

import (
	"AdventOfCode-go/advent2022/utils"
	"fmt"
	"strconv"
)

func rankCalories(firstElfCalories int, secondElfCalories int, thirdElfCalories int, calories int) (int, int, int) {
	switch {
	case calories > firstElfCalories:
		thirdElfCalories = secondElfCalories
		secondElfCalories = firstElfCalories
		firstElfCalories = calories
	case calories > secondElfCalories && calories < firstElfCalories:
		thirdElfCalories = secondElfCalories
		secondElfCalories = calories
	case calories > thirdElfCalories && calories < secondElfCalories:
		thirdElfCalories = calories
	}
	return firstElfCalories, secondElfCalories, thirdElfCalories
}

func elfMunchies(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)
	var elf int = 1
	//var topElf int
	//var topElfCalories int
	var firstElfCalories, secondElfCalories, thirdElfCalories int
	var calories int
	for i, j := range puzzleInput {

		if j == "" {
			fmt.Println("Blank Line found")

			firstElfCalories, secondElfCalories, thirdElfCalories = rankCalories(firstElfCalories, secondElfCalories, thirdElfCalories, calories)

			calories = 0
			elf++
		} else {
			foodCalories, _ := strconv.Atoi(j)
			fmt.Printf("Elf: %d Carrying food: %d\n", elf, foodCalories)
			calories += foodCalories

			if i == len(puzzleInput)-1 {
				firstElfCalories, secondElfCalories, thirdElfCalories = rankCalories(firstElfCalories, secondElfCalories, thirdElfCalories, calories)
			}
		}
	}

	if part == 'a' {
		return firstElfCalories
	} else {
		fmt.Printf("First place elf: %d\n", firstElfCalories)
		fmt.Printf("Second place elf: %d\n", secondElfCalories)
		fmt.Printf("Third place elf: %d\n", thirdElfCalories)
		return firstElfCalories + secondElfCalories + thirdElfCalories
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
