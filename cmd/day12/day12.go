package main

import (
	"fmt"
)

// Returns: newDirection, newX, newY
func doSingleMove(command byte, moveAmount int, startDirection int, startX int, startY int) (int, int, int) {
	var newDirection, newX, newY int

	newX = startX
	newY = startY
	newDirection = startDirection

	switch command {
	case 'F': // Forward. Continue moving in the direction pointing
		switch startDirection {
		case 0:
			newY -= moveAmount
			break
		case 90:
			newX += moveAmount
			break
		case 180:
			newY += moveAmount
			break
		case 270:
			newX -= moveAmount
			break
		default:
			panic("Duff move forward command")
		}
		break
	case 'N': // Move North. Direction stays the same. North is negative Y
		newY = startY - moveAmount
		break
	case 'S': // Move South. Direction stays the same. South is positive Y
		newY = startY + moveAmount
		break
	case 'E': // Move East. Direction stays the same. East is positive X
		newX = startX + moveAmount
		break
	case 'W': // Move West. Direction stays the same. West is negative X
		newX = startX - moveAmount
		break
	case 'R': // Turn Right. Add degrees to startDirection. Mod on 360 degrees
		newDirection = (startDirection + moveAmount) % 360
		break
	case 'L': // Move Left.
		newDirection = startDirection - moveAmount
		if newDirection < 0 {
			newDirection += 360
		}
		break
	default:
		panic("Duff direction command given")
	}

	return newDirection, newX, newY
}

func calcShipMovement(filename string, part byte, debug bool) int {
	var currDirection, currX, currY int = 90, 0, 0
	var startX, startY int = 0, 0
	var command byte
	var amount int

	puzzleInput, _ := readFile(filename)

	for _, line := range puzzleInput {
		fmt.Sscanf(line, "%c%d", &command, &amount)
		fmt.Printf("%c%d\n", command, amount)

		currDirection, currX, currY = doSingleMove(command, amount, currDirection, currX, currY)
	}

	return manhattanDistance2D(startX, startY, currX, currY)

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
		fmt.Println("Ship changed distance:", calcShipMovement(filenamePtr, execPart, debug))
	} else {
	}
}
