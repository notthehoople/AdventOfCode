package main

import (
	"AdventOfCode-go/advent2021/utils"
	"fmt"
)

type coords struct {
	x int
	y int
}

/*
func copyMap(sourceMap map[coords]int) (destinationMap map[coords]int) {
	destinationMap = make(map[coords]int)

	for k, v := range sourceMap {
		destinationMap[k] = v
	}

	return destinationMap
}
*/

func updateAdjacentOctopus(mapToUpdate map[coords]int, centre coords) {
	if centre.y > 0 {
		if centre.x > 0 {
			mapToUpdate[coords{centre.x - 1, centre.y - 1}]++
		}
		mapToUpdate[coords{centre.x, centre.y - 1}]++
		if centre.x < 9 {
			mapToUpdate[coords{centre.x + 1, centre.y - 1}]++
		}
	}

	if centre.x > 0 {
		mapToUpdate[coords{centre.x - 1, centre.y}]++
	}
	if centre.x < 9 {
		mapToUpdate[coords{centre.x + 1, centre.y}]++
	}

	if centre.y < 9 {
		if centre.x > 0 {
			mapToUpdate[coords{centre.x - 1, centre.y + 1}]++
		}
		mapToUpdate[coords{centre.x, centre.y + 1}]++
		if centre.x < 9 {
			mapToUpdate[coords{centre.x + 1, centre.y + 1}]++
		}
	}
}

func stepControl(octopusMap map[coords]int, numberOfSteps int, part byte, debug bool) int {
	var totalFlashCount int

	var stepLoopContinue bool = true
	var step int = 0

	for stepLoopContinue {

		step++
		if step > numberOfSteps && part == 'a' {
			stepLoopContinue = false
			continue
		}

		flashedOctopus := make(map[coords]bool)

		// First, the energy level of each octopus increases by 1.
		for k := range octopusMap {
			octopusMap[k]++
		}

		// Then, any octopus with an energy level greater than 9 flashes. This increases the energy level of all adjacent
		// octopuses by 1, including octopuses that are diagonally adjacent. If this causes an octopus to have an energy level
		// greater than 9, it also flashes. This process continues as long as new octopuses keep having their energy level
		// increased beyond 9. (An octopus can only flash at most once per step.)
		var keepGoing = true
		for keepGoing {
			keepGoing = false
			for k, v := range octopusMap {
				if v > 9 {
					// Time to flash for this octopus. Add it to the list of "flashes" then increase the energy of those around
					if _, ok := flashedOctopus[k]; ok {
						// This octopus has already flashed so ignore it
					} else {
						flashedOctopus[k] = true
						// Update all the octopus adjacent to this one
						updateAdjacentOctopus(octopusMap, k)
						keepGoing = true
					}
				}
			}
		}

		// Finally, any octopus that flashed during this step has its energy level set to 0
		var flashCount int
		for k := range flashedOctopus {
			octopusMap[k] = 0
			totalFlashCount++ // Total count of all flashes (part a)
			flashCount++      // Number of flashes in this step (part b)
		}

		if part == 'b' && flashCount == len(octopusMap) {
			return step
		}

		if debug {
			fmt.Println("Flashed Octopus:", flashedOctopus)
		}

	}

	if debug {
		fmt.Println("Map at Finish:", octopusMap)
		fmt.Printf("Step: %d Flashes: %d\n", numberOfSteps, totalFlashCount)
	}

	// Part a result
	return totalFlashCount
}

func solveDay(filename string, part byte, debug bool) int {
	puzzleInput, _ := utils.ReadFile(filename)
	octopusMap := make(map[coords]int)

	var numberOfFlashes int

	for y := 0; y < len(puzzleInput); y++ {
		for x, startPower := range puzzleInput[y] {
			if debug {
				fmt.Printf("x:%d, y:%d, power:%c\n", x, y, startPower)
			}
			octopusMap[coords{x: x, y: y}] = int(startPower - '0')
		}
	}

	numberOfFlashes = stepControl(octopusMap, 100, part, debug)

	return numberOfFlashes
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
