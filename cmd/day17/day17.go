package main

import (
	"fmt"
)

type cubeCoord struct {
	xCoord int
	yCoord int
	zCoord int
}

func countActiveCubes(cubeMap map[cubeCoord]bool) int {
	var activeCubes int = 0

	for _, active := range cubeMap {
		if active {
			activeCubes++
		}
	}
	return activeCubes
}

func setInitialState(puzzleInput []string) (cubeMap map[cubeCoord]bool) {
	var tmpCoords cubeCoord
	cubeMap = make(map[cubeCoord]bool)

	for y := 0; y < len(puzzleInput); y++ {
		for x := 0; x < len(puzzleInput[0]); x++ {
			if puzzleInput[y][x] == '#' {
				tmpCoords = cubeCoord{xCoord: x, yCoord: y, zCoord: 0}
				cubeMap[tmpCoords] = true
			}
		}
	}
	return cubeMap
}

func checkCubeNeighbours(checkCoords cubeCoord, cubeMap map[cubeCoord]bool) int {
	var activeNeighbours int = 0

	for z := checkCoords.zCoord - 1; z <= checkCoords.zCoord+1; z++ {
		for y := checkCoords.yCoord - 1; y <= checkCoords.yCoord+1; y++ {
			for x := checkCoords.xCoord - 1; x <= checkCoords.xCoord+1; x++ {
				if x == checkCoords.xCoord && y == checkCoords.yCoord && z == checkCoords.zCoord {
					// This is our coords to check so dont count this one!
				} else {
					if cubeMap[cubeCoord{xCoord: x, yCoord: y, zCoord: z}] {
						activeNeighbours++
					}
				}
			}
		}
	}

	return activeNeighbours
}

func applyCycleRules(cubeMap map[cubeCoord]bool, debug bool) (newCubeMap map[cubeCoord]bool) {
	var activeNeighbours int
	newCubeMap = make(map[cubeCoord]bool)

	// Loop through old array and apply rules when copying to new array
	for checkCoords := range cubeMap {

		// We need to check all the neighbours of the stored cube
		for z := checkCoords.zCoord - 1; z <= checkCoords.zCoord+1; z++ {
			for y := checkCoords.yCoord - 1; y <= checkCoords.yCoord+1; y++ {
				for x := checkCoords.xCoord - 1; x <= checkCoords.xCoord+1; x++ {

					activeNeighbours = checkCubeNeighbours(cubeCoord{xCoord: x, yCoord: y, zCoord: z}, cubeMap)
					if cubeMap[cubeCoord{xCoord: x, yCoord: y, zCoord: z}] {
						// cube is active. Stays active if has 2 or 3 active neighbours, otherwise goes inactive
						if activeNeighbours == 2 || activeNeighbours == 3 {
							newCubeMap[cubeCoord{xCoord: x, yCoord: y, zCoord: z}] = true
						} else {
							newCubeMap[cubeCoord{xCoord: x, yCoord: y, zCoord: z}] = false
						}
					} else {
						// cube is inactive. Becomes active if exactly 3 of its neighbours are active
						if activeNeighbours == 3 {
							newCubeMap[cubeCoord{xCoord: x, yCoord: y, zCoord: z}] = true
						}
					}
				}
			}
		}

	}

	if debug {
		fmt.Println("----------")
		fmt.Println(newCubeMap)
	}

	return newCubeMap
}

func howManyActiveCubes(filename string, part byte, debug bool) int {
	var cubeMap map[cubeCoord]bool

	puzzleInput, _ := readFile(filename)
	cubeMap = setInitialState(puzzleInput)

	if debug {
		fmt.Println("--- Initial State ---")
		fmt.Println(cubeMap)
	}

	for bootCycle := 0; bootCycle < 6; bootCycle++ {
		cubeMap = applyCycleRules(cubeMap, debug)
	}

	return countActiveCubes(cubeMap)
}

// ? Use a map to hold the known about cubes. struct of coords as the key, active or inactive as value
// Loop through the known cubes
// a neighbour is at most 1 coordinate different (e.g 1,1,1 has neighbours 0,1,1 and 2,1,1 continued)
// if a cube is active and 2 or 3 of its neighbours are also active, remain active OTHERWISE inactive
// if a cube is inactive but exactly 3 of its neighbours are active, cube becomes active OTHERWISE inactive

// Main routine
func main() {
	filenamePtr, execPart, debug, test := catchUserInput()

	if test {
		return
	}

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else if execPart == 'a' {
		fmt.Println("active cubes:", howManyActiveCubes(filenamePtr, execPart, debug))
	} else {
		//fmt.Println("occupied seats:", howManyFilledSeatsB(filenamePtr, execPart, debug))
	}
}
