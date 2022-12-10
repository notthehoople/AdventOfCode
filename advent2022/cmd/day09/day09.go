package main

import (
	"AdventOfCode-go/advent2022/utils"
	"fmt"
	"math"
)

type Coords struct {
	x int
	y int
}

func (coord Coords) add(change Coords) Coords {
	coord.x += change.x
	coord.y += change.y
	return coord
}

func (coord Coords) follow(head Coords) Coords {
	if coord.y == head.y {
		// On same row so check for distance
		if coord.x > head.x {
			if (coord.x - head.x) > 1 {
				coord.x--
			}
		} else {
			if (head.x - coord.x) > 1 {
				coord.x++
			}
		}
		return coord
	} else if coord.x == head.x {
		// On same column so check for distance
		if coord.y > head.y {
			if (coord.y - head.y) > 1 {
				coord.y--
			}
		} else {
			if (head.y - coord.y) > 1 {
				coord.y++
			}
		}
		return coord
	}

	// More complex. Need to work out the distance then move diagonally towards the head
	if math.Abs(float64((coord.x-head.x))) > 1 || math.Abs(float64((coord.y-head.y))) > 1 {
		if coord.x > head.x {
			coord.x--
		} else {
			coord.x++
		}
		if coord.y > head.y {
			coord.y--
		} else {
			coord.y++
		}
	}

	return coord
}

func ropeFun(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)

	areaCovered := make(map[Coords]bool)

	var currentHead, currentTail Coords
	areaCovered[currentTail] = true

	for _, instruction := range puzzleInput {
		var direction string
		var amount int
		var transpose Coords
		fmt.Sscanf(instruction, "%s %d\n", &direction, &amount)

		if debug {
			fmt.Printf("Direction: %s Amount: %d\n", direction, amount)
		}

		switch direction {
		case "R":
			transpose = Coords{x: 1, y: 0}
		case "L":
			transpose = Coords{x: -1, y: 0}
		case "U":
			transpose = Coords{x: 0, y: -1}
		case "D":
			transpose = Coords{x: 0, y: 1}
		default:
			fmt.Println("Corrupt input:", direction)
			return 0
		}

		for i := 1; i <= amount; i++ {
			currentHead = currentHead.add(transpose)
			currentTail = currentTail.follow(currentHead)
			if debug {
				fmt.Println("CurrentHead", currentHead)
				fmt.Println("CurrentTail", currentTail)
			}
			areaCovered[currentTail] = true
		}
	}

	if debug {
		fmt.Println("Map", areaCovered)
	}

	return len(areaCovered)
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Result is: %d\n", ropeFun(filenamePtr, execPart, debug))
	case 'b':
		fmt.Println("Not implemented yet")
	case 'z':
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
