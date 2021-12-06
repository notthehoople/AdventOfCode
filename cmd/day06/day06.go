package main

import (
	"aoc/advent2021/utils"
	"fmt"
	"strconv"
	"strings"
)

func calcLanternfish(lanternfish []int, maxDays int, part byte, debug bool) int {
	var currentFish int

	for day := 0; day < maxDays; day++ {
		fmt.Println("Day:", day)
		if debug {
			fmt.Printf("Day:%d ", day)
			for i := 0; i < len(lanternfish); i++ {
				fmt.Printf("%d,", lanternfish[i])
			}
			fmt.Printf("\n")
		}

		currentFish = len(lanternfish)
		for i := 0; i < currentFish; i++ {

			if lanternfish[i] == 0 {
				// Birth of a new fishy
				lanternfish[i] = 6
				lanternfish = append(lanternfish, 8)
			} else {
				lanternfish[i]--
			}
		}
	}

	return len(lanternfish)
}

func solveDay(filename string, days int, part byte, debug bool) int {
	puzzleInput, _ := utils.ReadFile(filename)
	puzzleInputSplit := strings.Split(puzzleInput[0], ",")
	lanternfish := make([]int, len(puzzleInputSplit), len(puzzleInputSplit)*100000)
	for i := 0; i < len(puzzleInputSplit); i++ {
		lanternfish[i], _ = strconv.Atoi(puzzleInputSplit[i])
	}
	return calcLanternfish(lanternfish, days, part, debug)
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf(" 18 day result is: %d\n", solveDay(filenamePtr, 18, execPart, debug))
		fmt.Printf(" 80 day result is: %d\n", solveDay(filenamePtr, 80, execPart, debug))
		fmt.Printf("256 day result is: %d\n", solveDay(filenamePtr, 256, execPart, debug))
	}
}
