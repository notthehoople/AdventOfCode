package main

import (
	"AdventOfCode-go/advent2021/utils"
	"fmt"
	"strconv"
	"strings"
)

func calcLanternfish(lanternfish []int, maxDays int, part byte, debug bool) int {
	currentFish := make([]int, len(lanternfish))

	for day := 0; day < maxDays; day++ {
		if debug {
			fmt.Printf("Day:%d ", day)
			for i := 0; i < len(lanternfish); i++ {
				fmt.Printf("%d,", lanternfish[i])
			}
			fmt.Printf("\n")
		}

		for i := 0; i <= 8; i++ {
			currentFish[i] = lanternfish[i]
		}
		lanternfish[7] = currentFish[8]
		lanternfish[6] = currentFish[7]
		lanternfish[5] = currentFish[6]
		lanternfish[4] = currentFish[5]
		lanternfish[3] = currentFish[4]
		lanternfish[2] = currentFish[3]
		lanternfish[1] = currentFish[2]
		lanternfish[0] = currentFish[1]
		lanternfish[8] = currentFish[0]
		lanternfish[6] += currentFish[0]
	}

	var countFish int
	for i := 0; i < 9; i++ {
		countFish += lanternfish[i]
	}
	return countFish
}

func solveDay(filename string, days int, part byte, debug bool) int {
	var startFish int
	puzzleInput, _ := utils.ReadFile(filename)
	puzzleInputSplit := strings.Split(puzzleInput[0], ",")
	lanternfish := make([]int, 10)
	for i := 0; i < len(puzzleInputSplit); i++ {
		startFish, _ = strconv.Atoi(puzzleInputSplit[i])
		lanternfish[startFish]++
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
