package main

import (
	"fmt"
)

// Abs returns the absolute value of x.
func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Returns: newDirection, newX, newY
func moveWaypoint(command byte, moveAmount int, startX int, startY int) (int, int) {
	var newX, newY int

	newX = startX
	newY = startY

	switch command {
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
	default:
		panic("Duff direction command given")
	}

	return newX, newY
}

// Return: shipX, shipY position
func moveShip(amount int, shipX int, shipY int, wayptX int, wayptY int) (int, int) {
	shipX += wayptX * amount
	shipY += wayptY * amount

	return shipX, shipY
}

// Rotate the way point around the ship
// Return: wayptX, wayptY Waypoint position
func rotateWayPoint(command byte, amount int, shipX int, shipY int, wayptX int, wayptY int) (int, int) {
	var newX, newY int

	if command == 'R' {
		switch amount {
		case 90:
			newX = -wayptY
			newY = wayptX
			break
		case 180:
			newX = -wayptX
			newY = -wayptY
			break
		case 270:
			newX = wayptY
			newY = -wayptX
			break
		default:
			panic("Bad rotation in rotateWayPoint")
		}

		return newX, newY
	}

	// Rotate Left

	switch amount {
	case 90:
		newX = wayptY
		newY = -wayptX
		break
	case 180:
		newX = -wayptX
		newY = -wayptY
		break
	case 270:
		newX = -wayptY
		newY = wayptX
		break
	default:
		panic("Bad rotation in rotateWayPoint")
	}

	return newX, newY
}

func calcShipMovementB(filename string, part byte, debug bool) int {
	var wayptX, wayptY int = 10, -1
	var shipX, shipY int = 0, 0
	var startX, startY int = 0, 0
	var command byte
	var amount int

	puzzleInput, _ := readFile(filename)

	for _, line := range puzzleInput {
		fmt.Sscanf(line, "%c%d", &command, &amount)

		switch command {
		case 'N':
			fallthrough
		case 'S':
			fallthrough
		case 'E':
			fallthrough
		case 'W':
			wayptX, wayptY = moveWaypoint(command, amount, wayptX, wayptY)
			break
		case 'F':
			shipX, shipY = moveShip(amount, shipX, shipY, wayptX, wayptY)
			break
		case 'L':
			fallthrough
		case 'R':
			wayptX, wayptY = rotateWayPoint(command, amount, shipX, shipY, wayptX, wayptY)
			break
		default:
			panic("Duff commands given")
		}
	}

	return manhattanDistance2D(startX, startY, shipX, shipY)
}
