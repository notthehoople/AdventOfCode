package main

import (
	"AdventOfCode-go/advent2025/utils"
	"fmt"
)

type Coords struct {
	x int
	y int
}

func adjacentRolls(currentPos Coords, areaMap map[Coords]byte, debug bool) int {
	var count int

	// Top Row
	if areaMap[Coords{currentPos.x - 1, currentPos.y - 1}] == '@' {
		count++
	}
	if areaMap[Coords{currentPos.x, currentPos.y - 1}] == '@' {
		count++
	}
	if areaMap[Coords{currentPos.x + 1, currentPos.y - 1}] == '@' {
		count++
	}
	// Middle Row
	if areaMap[Coords{currentPos.x - 1, currentPos.y}] == '@' {
		count++
	}
	if areaMap[Coords{currentPos.x + 1, currentPos.y}] == '@' {
		count++
	}
	// Bottom Row
	if areaMap[Coords{currentPos.x - 1, currentPos.y + 1}] == '@' {
		count++
	}
	if areaMap[Coords{currentPos.x, currentPos.y + 1}] == '@' {
		count++
	}
	if areaMap[Coords{currentPos.x + 1, currentPos.y + 1}] == '@' {
		count++
	}

	return count
}

func day04(filename string, part byte, debug bool) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)

	// Build the map of the printing department
	areaMap := make(map[Coords]byte)

	for y := 0; y < len(puzzleInput); y++ {
		for x := 0; x < len(puzzleInput[y]); x++ {
			areaMap[Coords{x, y}] = puzzleInput[y][x]
		}

		if debug {
			fmt.Println(areaMap)
		}
	}

	if part == 'a' {
		// Now look for all rolls of paper '@' that have fewer than 4 rolls of paper in the 8 adjacent spaces.
		for currentPos := range areaMap {
			if areaMap[currentPos] == '@' {
				count := adjacentRolls(currentPos, areaMap, debug)
				if count < 4 {
					result++
				}
			}
		}
	} else {
		// Look for all rolls of paper '@' that have fewer than 4 rolls of paper in the 8 adjacent spaces
		// Remove each then repeat the search until we can't remove any more

		looping := true
		var removedRolls int
		for looping {
			removedRolls = 0
			for currentPos := range areaMap {
				if areaMap[currentPos] == '@' {
					count := adjacentRolls(currentPos, areaMap, debug)
					if count < 4 {
						areaMap[currentPos] = 'x'
						removedRolls++
						result++
					}
				}
			}
			if removedRolls == 0 {
				looping = false
			}
		}
	}

	return result
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", day04(filenamePtr, execPart, debug))
	}
}
