package main

import (
	"fmt"
)

type cubeCoord4D struct {
	xCoord int
	yCoord int
	zCoord int
	wCoord int
}

func countActiveCubes4D(cubeMap map[cubeCoord4D]bool) int {
	var activeCubes int = 0

	for _, active := range cubeMap {
		if active {
			activeCubes++
		}
	}
	return activeCubes
}

func setInitialState4D(puzzleInput []string) (cubeMap map[cubeCoord4D]bool) {
	var tmpCoords cubeCoord4D
	cubeMap = make(map[cubeCoord4D]bool)

	for y := 0; y < len(puzzleInput); y++ {
		for x := 0; x < len(puzzleInput[0]); x++ {
			if puzzleInput[y][x] == '#' {
				tmpCoords = cubeCoord4D{xCoord: x, yCoord: y, zCoord: 0, wCoord: 0}
				cubeMap[tmpCoords] = true
			}
		}
	}
	return cubeMap
}

func checkCubeNeighbours4D(checkCoords cubeCoord4D, cubeMap map[cubeCoord4D]bool) int {
	var activeNeighbours int = 0

	for z := checkCoords.zCoord - 1; z <= checkCoords.zCoord+1; z++ {
		for y := checkCoords.yCoord - 1; y <= checkCoords.yCoord+1; y++ {
			for x := checkCoords.xCoord - 1; x <= checkCoords.xCoord+1; x++ {
				for w := checkCoords.wCoord - 1; w <= checkCoords.wCoord+1; w++ {
					if x == checkCoords.xCoord && y == checkCoords.yCoord && z == checkCoords.zCoord && w == checkCoords.wCoord {
						// This is our coords to check so dont count this one!
					} else {
						if cubeMap[cubeCoord4D{xCoord: x, yCoord: y, zCoord: z, wCoord: w}] {
							activeNeighbours++
						}
					}
				}
			}
		}
	}

	return activeNeighbours
}

func applyCycleRules4D(cubeMap map[cubeCoord4D]bool, debug bool) (newCubeMap map[cubeCoord4D]bool) {
	var activeNeighbours int
	newCubeMap = make(map[cubeCoord4D]bool)

	// Loop through old array and apply rules when copying to new array
	for checkCoords := range cubeMap {

		// We need to check all the neighbours of the stored cube
		for z := checkCoords.zCoord - 1; z <= checkCoords.zCoord+1; z++ {
			for y := checkCoords.yCoord - 1; y <= checkCoords.yCoord+1; y++ {
				for x := checkCoords.xCoord - 1; x <= checkCoords.xCoord+1; x++ {
					for w := checkCoords.wCoord - 1; w <= checkCoords.wCoord+1; w++ {

						activeNeighbours = checkCubeNeighbours4D(cubeCoord4D{xCoord: x, yCoord: y, zCoord: z, wCoord: w}, cubeMap)
						if cubeMap[cubeCoord4D{xCoord: x, yCoord: y, zCoord: z, wCoord: w}] {
							// cube is active. Stays active if has 2 or 3 active neighbours, otherwise goes inactive
							if activeNeighbours == 2 || activeNeighbours == 3 {
								newCubeMap[cubeCoord4D{xCoord: x, yCoord: y, zCoord: z, wCoord: w}] = true
							} else {
								newCubeMap[cubeCoord4D{xCoord: x, yCoord: y, zCoord: z, wCoord: w}] = false
							}
						} else {
							// cube is inactive. Becomes active if exactly 3 of its neighbours are active
							if activeNeighbours == 3 {
								newCubeMap[cubeCoord4D{xCoord: x, yCoord: y, zCoord: z, wCoord: w}] = true
							}
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

func howManyActiveCubesPartB(filename string, part byte, debug bool) int {
	var cubeMap map[cubeCoord4D]bool

	puzzleInput, _ := readFile(filename)
	cubeMap = setInitialState4D(puzzleInput)

	if debug {
		fmt.Println("--- Initial State ---")
		fmt.Println(cubeMap)
	}

	for bootCycle := 0; bootCycle < 6; bootCycle++ {
		cubeMap = applyCycleRules4D(cubeMap, debug)
	}

	return countActiveCubes4D(cubeMap)
}
