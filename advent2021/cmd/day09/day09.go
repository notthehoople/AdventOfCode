package main

import (
	"AdventOfCode-go/advent2021/utils"
	"fmt"
)

type coords struct {
	x int
	y int
}

func lowPoint(location coords, height int, heightMap map[coords]int) bool {
	/*
		lowPoint: locations that are lower than any of its adjacent locations
		Most locations have four adjacent locations (up, down, left, and right)
		locations on the edge or corner of the map have three or two adjacent locations
		ignore diagonals
	*/

	// check y with location.x+1 & location.x-1
	// check x with location.y+1 & location.y-1
	// expect that some will return error so check for successful result

	var checkHeight int
	var ok bool
	checkHeight, ok = heightMap[coords{location.x + 1, location.y}]
	if ok {
		if checkHeight <= height {
			return false
		}
	}
	checkHeight, ok = heightMap[coords{location.x - 1, location.y}]
	if ok {
		if checkHeight <= height {
			return false
		}
	}

	checkHeight, ok = heightMap[coords{location.x, location.y + 1}]
	if ok {
		if checkHeight <= height {
			return false
		}
	}
	checkHeight, ok = heightMap[coords{location.x, location.y - 1}]
	if ok {
		if checkHeight <= height {
			return false
		}
	}

	return true
}

func calcBasinSize(heightMap map[coords]int, part byte, debug bool) int {

	basinCentre := make(map[coords]bool)

	// First let's find the low points of the map
	for location, height := range heightMap {
		if lowPoint(location, height, heightMap) {
			basinCentre[location] = true
		}
	}

	if debug {
		for location := range basinCentre {
			fmt.Printf("Low Point centre x:%d y:%d\n", location.x, location.y)
		}
	}

	// Now for each low point look for adjacent spaces which are greater than current,
	// but are not 9 and not the edge of the map
	basinReaches := make(map[coords]int)

	// For each low point in basinCentre
	//   Current basin = (increment)
	//   Loop till we're done
	//     Check each surrounding point. If greater, not 9, not edge add to basinReaches with current basin number
	// Loop through basinReaches
	//   Count sizes of each basin
	//   Multiply together the top 3 basin sizes. Return result

	return 0
}

func calcRiskLevel(heightMap map[coords]int, part byte, debug bool) int {
	/*
		risk level of each lowest point is 1 + height
	*/
	var riskLevelTotal int
	for location, height := range heightMap {
		if lowPoint(location, height, heightMap) {
			riskLevelTotal += 1 + height
		}
	}

	return riskLevelTotal
}

func solveDay(filename string, part byte, debug bool) int {
	puzzleInput, _ := utils.ReadFile(filename)
	heightMap := make(map[coords]int)

	for y := 0; y < len(puzzleInput); y++ {
		for x, height := range puzzleInput[y] {
			if debug {
				fmt.Printf("x:%d, y:%d, height:%c\n", x, y, height)
			}
			heightMap[coords{x: x, y: y}] = int(height - '0')
		}
	}

	if part == 'a' {
		return calcRiskLevel(heightMap, part, debug)
	} else {
		return calcBasinSize(heightMap, part, debug)
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
