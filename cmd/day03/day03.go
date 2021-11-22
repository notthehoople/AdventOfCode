package main

import (
	"aoc/advent2017/utils"
	"fmt"
)

type coords struct {
	x int
	y int
}

func sumAdjacentSquares(spiralPattern map[coords]int, currentPosition coords) int {
	var posValue, posSumTotal int
	var posExists bool

	posValue, posExists = spiralPattern[coords{currentPosition.x - 1, currentPosition.y - 1}]
	if posExists {
		posSumTotal += posValue
	}
	posValue, posExists = spiralPattern[coords{currentPosition.x, currentPosition.y - 1}]
	if posExists {
		posSumTotal += posValue
	}
	posValue, posExists = spiralPattern[coords{currentPosition.x + 1, currentPosition.y - 1}]
	if posExists {
		posSumTotal += posValue
	}
	posValue, posExists = spiralPattern[coords{currentPosition.x - 1, currentPosition.y}]
	if posExists {
		posSumTotal += posValue
	}
	posValue, posExists = spiralPattern[coords{currentPosition.x + 1, currentPosition.y}]
	if posExists {
		posSumTotal += posValue
	}
	posValue, posExists = spiralPattern[coords{currentPosition.x - 1, currentPosition.y + 1}]
	if posExists {
		posSumTotal += posValue
	}
	posValue, posExists = spiralPattern[coords{currentPosition.x, currentPosition.y + 1}]
	if posExists {
		posSumTotal += posValue
	}
	posValue, posExists = spiralPattern[coords{currentPosition.x + 1, currentPosition.y + 1}]
	if posExists {
		posSumTotal += posValue
	}
	return posSumTotal
}

func calcAdjacentValues(valueToCheck int, debug bool) int {
	/*
		As a stress test on the system, the programs here clear the grid and then
		store the value 1 in square 1. Then, in the same allocation order as shown
		above, they store the sum of the values in all adjacent squares, including
		diagonals.

		147  142  133  122   59
		304    5    4    2   57
		330   10    1    1   54
		351   11   23   25   26
		362  747  806--->   ...
	*/

	const LEFT = 1
	const UP = 2
	const RIGHT = 3
	const DOWN = 4

	// The value to create in each position. Starts at 1
	var currentPosValue int = 1
	var movingDirection int = RIGHT
	var currentPosition coords

	spiralPattern := make(map[coords]int)
	// Set where 1 starts from
	currentPosition.x = 0
	currentPosition.y = 0
	spiralPattern[currentPosition] = currentPosValue

	//for currentPosValue < valueToCheck {
	for {

		for movingDirection == RIGHT {
			// Move one space to the right. Add the currentValue to that position
			currentPosition.x++
			currentPosValue = sumAdjacentSquares(spiralPattern, currentPosition)
			spiralPattern[currentPosition] = currentPosValue
			if currentPosValue > valueToCheck {
				return currentPosValue
			}

			// If space ABOVE is empty then we start going upwards
			_, posExists := spiralPattern[coords{currentPosition.x, currentPosition.y - 1}]
			if !posExists {
				movingDirection = UP
			}
		}

		for movingDirection == UP {
			// Move one space to UP. Add the currentValue to that position
			currentPosition.y--
			currentPosValue = sumAdjacentSquares(spiralPattern, currentPosition)
			spiralPattern[currentPosition] = currentPosValue
			if currentPosValue > valueToCheck {
				return currentPosValue
			}

			// If space LEFT is empty then we start going left
			_, posExists := spiralPattern[coords{currentPosition.x - 1, currentPosition.y}]
			if !posExists {
				movingDirection = LEFT
			}
		}

		for movingDirection == LEFT {
			// Move one space to LEFT. Add the currentValue to that position
			currentPosition.x--
			currentPosValue = sumAdjacentSquares(spiralPattern, currentPosition)
			spiralPattern[currentPosition] = currentPosValue
			if currentPosValue > valueToCheck {
				return currentPosValue
			}

			// If space DOWN is empty then we start going down
			_, posExists := spiralPattern[coords{currentPosition.x, currentPosition.y + 1}]
			if !posExists {
				movingDirection = DOWN
			}
		}

		for movingDirection == DOWN {
			// Move one space to DOWN. Add the currentValue to that position
			currentPosition.y++
			currentPosValue = sumAdjacentSquares(spiralPattern, currentPosition)
			spiralPattern[currentPosition] = currentPosValue
			if currentPosValue > valueToCheck {
				return currentPosValue
			}

			// If space RIGHT is empty then we start going right
			_, posExists := spiralPattern[coords{currentPosition.x + 1, currentPosition.y}]
			if !posExists {
				movingDirection = RIGHT
			}
		}
	}
}

// valueToCheck is the value that we want to build the spiral to, then
// check the manhattan distance between it and 1
func calcSteps(valueToCheck int, debug bool) int {
	/*
		Each square on the grid is allocated in a spiral pattern starting at a location
		marked 1 and then counting up while spiraling outward. For example, the first
		few squares are allocated like this:

		17  16  15  14  13
		18   5   4   3  12
		19   6   1   2  11
		20   7   8   9  10
		21  22  23---> ...

		While this is very space-efficient (no squares are skipped), requested data must
		be carried back to square 1 (the location of the only access port for this memory
		system) by programs that can only move up, down, left, or right. They always
		take the shortest path: the Manhattan Distance between the location of the data
		and square 1.
	*/

	//var spiralPattern map[int,int]int

	const LEFT = 1
	const UP = 2
	const RIGHT = 3
	const DOWN = 4

	// The value to create in each position. Starts at 1
	var currentPosValue int = 1
	var movingDirection int = RIGHT
	var currentPosition coords

	spiralPattern := make(map[coords]int)
	// Set where 1 starts from
	currentPosition.x = 0
	currentPosition.y = 0
	spiralPattern[currentPosition] = currentPosValue

	for currentPosValue <= valueToCheck {

		for movingDirection == RIGHT {
			// Move one space to the right. Add the currentValue to that position
			currentPosition.x++
			currentPosValue++
			spiralPattern[currentPosition] = currentPosValue

			// If space ABOVE is empty then we start going upwards
			_, posExists := spiralPattern[coords{currentPosition.x, currentPosition.y - 1}]
			if !posExists {
				movingDirection = UP
			}
		}

		for movingDirection == UP {
			// Move one space to UP. Add the currentValue to that position
			currentPosition.y--
			currentPosValue++
			spiralPattern[currentPosition] = currentPosValue

			// If space LEFT is empty then we start going left
			_, posExists := spiralPattern[coords{currentPosition.x - 1, currentPosition.y}]
			if !posExists {
				movingDirection = LEFT
			}
		}

		for movingDirection == LEFT {
			// Move one space to LEFT. Add the currentValue to that position
			currentPosition.x--
			currentPosValue++
			spiralPattern[currentPosition] = currentPosValue

			// If space DOWN is empty then we start going down
			_, posExists := spiralPattern[coords{currentPosition.x, currentPosition.y + 1}]
			if !posExists {
				movingDirection = DOWN
			}
		}

		for movingDirection == DOWN {
			// Move one space to DOWN. Add the currentValue to that position
			currentPosition.y++
			currentPosValue++
			spiralPattern[currentPosition] = currentPosValue

			// If space RIGHT is empty then we start going right
			_, posExists := spiralPattern[coords{currentPosition.x + 1, currentPosition.y}]
			if !posExists {
				movingDirection = RIGHT
			}
		}
	}
	if debug {
		fmt.Println(spiralPattern)
	}

	var resultPosition coords

	for i, j := range spiralPattern {
		if j == valueToCheck {
			resultPosition = i
		}
	}

	if debug {
		fmt.Println(resultPosition)
	}
	return utils.ManhattanDistance2D(0, 0, resultPosition.x, resultPosition.y)
}

func solveSteps(inputData int, part byte, debug bool) int {
	if part == 'a' {
		return calcSteps(inputData, debug)
	} else {
		return calcAdjacentValues(inputData, debug)
	}
}

// Main routine
func main() {
	_, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", solveSteps(368078, execPart, debug))
	}
}
