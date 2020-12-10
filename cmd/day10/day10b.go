package main

import (
	"fmt"
	"sort"
	"strconv"
)

var resultsCache map[int]int

func calcPathsRecurse(workNumber int, joltList []int, debug bool) int {
	var pathCount int = 0
	var loopCount int = 0
	var result int
	// workNumber: the item to look at
	// joltList: the remaining list of adapters to examine

	// Work out how many of the list *could* work from here
	//   Recurse with each potential and a subset of the list
	//     Return numberOfArrangements

	if debug {
		fmt.Println("===================")
		fmt.Printf("Checking: %d\n", workNumber)
		fmt.Println("Remaining List:", joltList)
	}

	for i := 0; i < len(joltList); i++ {
		loopCount++
		if joltList[i]-workNumber > 3 {
			// List is sorted so we quit the loop when we're found all we can
			break
		}

		if resultsCache[joltList[i]] > 0 {
			pathCount += resultsCache[joltList[i]]
		} else {
			result = calcPathsRecurse(joltList[i], joltList[i+1:], debug)
			resultsCache[joltList[i]] = result
			pathCount += result
		}

		if loopCount > 1 {
			pathCount++
		}
	}

	return pathCount
}

func calcNumberJoltArrangements(filename string, part byte, debug bool) int {
	var distinctArrangements int = 0
	var joltList []int

	puzzleInput, _ := readFile(filename)
	joltList = make([]int, len(puzzleInput))

	// Process the jolt list into a more usable form
	for item, number := range puzzleInput {
		joltList[item], _ = strconv.Atoi(number)
	}
	sort.Ints(joltList)

	resultsCache = make(map[int]int)
	// Don't have to use every adapter now
	// Recurse with number list
	//   Work out how many of the list *could* work from here
	//     Recurse with each potential and a subset of the list
	//       Return numberOfArrangements

	fmt.Println("Recursing....")
	distinctArrangements = calcPathsRecurse(0, joltList, debug)
	distinctArrangements++
	fmt.Println(".....back from recurse")

	return distinctArrangements
}
