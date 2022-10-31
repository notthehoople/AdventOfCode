package main

import (
	"aoc/advent2021/utils"
	"fmt"
	"strconv"
)

func countSlidingWindowIncreases(depthReadings []string) int {
	/*
	   Your goal now is to count the number of times the sum of measurements in this
	   sliding window increases from the previous sum. So, compare A with B, then
	   compare B with C, then C with D, and so on. Stop when there aren't enough
	   measurements left to create a new three-measurement sum.
	*/
	var numberOfIncreases int
	var firstReading, secondReading, thirdReading int

	slidingWindows := make([]int, len(depthReadings))
	slidingWindowPos := 0
	for i := 0; i < len(depthReadings); i++ {
		if i+2 < len(depthReadings) {
			firstReading, _ = strconv.Atoi(depthReadings[i])
			secondReading, _ = strconv.Atoi(depthReadings[i+1])
			thirdReading, _ = strconv.Atoi(depthReadings[i+2])

			slidingWindows[slidingWindowPos] += firstReading
			slidingWindows[slidingWindowPos] += secondReading
			slidingWindows[slidingWindowPos] += thirdReading
		}
		slidingWindowPos++
	}

	currentPos := slidingWindows[0]
	for j := 1; j < len(slidingWindows); j++ {
		if slidingWindows[j] > currentPos {
			numberOfIncreases++
		}
		currentPos = slidingWindows[j]
	}
	return numberOfIncreases
}

func countDepthIncreases(depthReadings []string) int {
	/*
		To do this, count the number of times a depth measurement increases from the
		previous measurement. (There is no measurement before the first measurement.)
	*/
	var numberOfIncreases int
	var currentReading, previousReading int

	previousReading, _ = strconv.Atoi(depthReadings[0])

	for i := 1; i < len(depthReadings); i++ {
		currentReading, _ = strconv.Atoi(depthReadings[i])
		if currentReading > previousReading {
			numberOfIncreases++
		}
		previousReading = currentReading
	}

	return numberOfIncreases
}

func solveDay(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)
	if part == 'a' {
		return countDepthIncreases(puzzleInput)
	} else {
		return countSlidingWindowIncreases(puzzleInput)
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
