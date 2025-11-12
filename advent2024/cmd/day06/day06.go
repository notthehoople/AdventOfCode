package main

import (
	"AdventOfCode-go/advent2024/utils"
	"fmt"
)

type Coords struct {
	x int
	y int
}

func day06(filename string, part byte, debug bool) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)

	// Obstructions are '#'. The the guard reaches an obstruction she turns right
	// if no obstruction the guard takes a step forward
	// Guard is '^', '>', '<', 'V' depending on direction faced
	// If reaches the edge of the map then the guard leaves the area
	// How many positions in the map will the guard visit before leaving the area?

	areaMap := make(map[Coords]byte)

	var currentPos Coords
	var movement Coords
	var direction string

	for y := 0; y < len(puzzleInput); y++ {
		for x := 0; x < len(puzzleInput[y]); x++ {
			if puzzleInput[y][x] != '.' {
				areaMap[Coords{x, y}] = puzzleInput[y][x]
			}
			if puzzleInput[y][x] == '^' {
				currentPos = Coords{x, y}
				movement = Coords{0, -1}
				direction = "UP"
			}
		}

		if debug {
			fmt.Println(areaMap)
			fmt.Println(currentPos)
		}
	}

	// Now loop until we're finished. Need bounds for the map to test against
	maxX := len(puzzleInput[0]) - 1
	maxY := len(puzzleInput) - 1

	var guardInArea bool = true

	for guardInArea {
		if currentPos.x < 0 || currentPos.x >= maxX || currentPos.y < 0 || currentPos.y >= maxY {
			guardInArea = false
		}

		areaMap[currentPos] = 'X'
		switch areaMap[Coords{currentPos.x + movement.x, currentPos.y + movement.y}] {
		case '#':
			switch direction {
			case "UP":
				direction = "RIGHT"
				movement = Coords{1, 0}
			case "RIGHT":
				direction = "DOWN"
				movement = Coords{0, 1}
			case "DOWN":
				direction = "LEFT"
				movement = Coords{-1, 0}
			case "LEFT":
				direction = "UP"
				movement = Coords{0, -1}
			}
		default:
			currentPos.x += movement.x
			currentPos.y += movement.y
		}
		if debug {
			fmt.Println("-------------------")
			fmt.Println(areaMap)
			fmt.Println(currentPos)
			fmt.Println(direction, movement)
		}
	}

	// Count the results
	for _, i := range areaMap {
		if i == 'X' {
			result++
		}
	}

	// Part b: add an obstruction that makes the guards route into an infinite loop
	// Perhaps need to put a number for the guards route instead of an 'X'
	// Add an obstruction at each point on the map. Set the guard moving.
	// When the guard is traversing a route with all the same numbers on the ground then answer
	// Continue until have tested every position in the map
	//
	// Run through part a to build a map of where the guard goes already. No point in trying
	// to put an obsticle in places that the guard doesn't go!
	// Then replace each 'X' with an obstacle to see if it makes an infinite loop for the guard

	return result
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", day06(filenamePtr, execPart, debug))
	}
}
