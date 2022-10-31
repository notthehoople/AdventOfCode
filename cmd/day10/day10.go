package main

import (
	"fmt"
	"sort"
	"strconv"
)

func calcJoltDifferences(filename string, part byte, debug bool) int {
	var oneJoltCount, threeJoltCount int = 0, 0
	var joltList []int
	var difference int

	puzzleInput, _ := readFile(filename)
	joltList = make([]int, len(puzzleInput))

	// Process the jolt list into a more usable form
	for item, number := range puzzleInput {
		joltList[item], _ = strconv.Atoi(number)
	}

	sort.Ints(joltList)

	// First count the difference between the charging socket and the first item
	if joltList[0] == 1 {
		oneJoltCount++
	} else if joltList[0] == 3 {
		threeJoltCount++
	}

	for i := 1; i < len(joltList); i++ {
		difference = joltList[i] - joltList[i-1]
		if difference == 1 {
			oneJoltCount++
		} else if difference == 3 {
			threeJoltCount++
		}
	}

	threeJoltCount++ // device is 3 more than the largest

	if debug {
		fmt.Printf("one: %d three: %d\n", oneJoltCount, threeJoltCount)
	}

	return oneJoltCount * threeJoltCount
}

// Main routine
func main() {
	filenamePtr, execPart, debug, test := catchUserInput()

	if test {
		return
	}

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else if execPart == 'a' {
		fmt.Println("number of jolt differences:", calcJoltDifferences(filenamePtr, execPart, debug))
	} else {
		fmt.Println("number of arrangements:", calcNumberJoltArrangements(filenamePtr, execPart, debug))
	}
}
