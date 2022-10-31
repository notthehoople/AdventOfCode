package main

import (
	"fmt"
)

// empty seats that see no occupied seats become occupied
// five or more visible occupied seats for an occupied seat to become vacant
// seat matching no rule don't change

type directionStruct struct {
	xChange int
	yChange int
}

func checkAreaB(tempArea [][]byte, yCoord int, xCoord int, debug bool) int {
	var directions = []directionStruct{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1}}
	var occupiedSeats int = 0
	var checkX, checkY int
	var keepLooping bool

	for i := 0; i < len(directions); i++ {

		checkX = xCoord
		checkY = yCoord

		keepLooping = true
		for keepLooping {

			checkX += directions[i].xChange
			checkY += directions[i].yChange

			if checkX < 0 || checkX > len(tempArea[yCoord])-1 {
				keepLooping = false
			}
			if checkY < 0 || checkY > len(tempArea)-1 {
				keepLooping = false
			}

			if keepLooping {
				if tempArea[checkY][checkX] == '#' {
					occupiedSeats++
					keepLooping = false
				} else if tempArea[checkY][checkX] == 'L' {
					keepLooping = false
				}
			}
		}
	}

	return occupiedSeats
}

func applySeatingRulesB(oldSeatPlan [][]byte, newSeatPlan [][]byte, debug bool) ([][]byte, bool) {
	var occupiedSeats int
	var somethingChanged bool = false

	// Loop through old array and apply rules when copying to new array
	for y := 0; y < len(oldSeatPlan); y++ {
		for x := 0; x < len(oldSeatPlan[0]); x++ {
			newSeatPlan[y][x] = oldSeatPlan[y][x]
			switch oldSeatPlan[y][x] {
			case 'L':
				occupiedSeats = checkAreaB(oldSeatPlan, y, x, debug)
				if occupiedSeats == 0 {
					newSeatPlan[y][x] = '#'
					somethingChanged = true
				}
				break
			case '#':
				occupiedSeats = checkAreaB(oldSeatPlan, y, x, debug)
				if occupiedSeats >= 5 {
					newSeatPlan[y][x] = 'L'
					somethingChanged = true
				}
				break
			case '.':
				break
			default:
				fmt.Printf("Found a dodgy area at x:%d y:%d\n", x, y)
				break
			}
		}
	}

	if debug {
		fmt.Println("====== SEAT PLAN =====")
		print2DArray(oldSeatPlan)
	}

	for y := 0; y < len(newSeatPlan); y++ {
		for x := 0; x < len(newSeatPlan[0]); x++ {
			oldSeatPlan[y][x] = newSeatPlan[y][x]
		}
	}

	return oldSeatPlan, somethingChanged
}

func howManyFilledSeatsB(filename string, part byte, debug bool) int {
	var seatPlan [][]byte
	var scratchSeatPlan [][]byte

	puzzleInput, _ := readFile(filename)
	seatPlan = make([][]byte, len(puzzleInput))
	scratchSeatPlan = make([][]byte, len(puzzleInput))

	for i := 0; i < len(puzzleInput); i++ {
		seatPlan[i] = make([]byte, len(puzzleInput[0]))
		scratchSeatPlan[i] = make([]byte, len(puzzleInput[0]))
	}

	setSeatPlan(puzzleInput, seatPlan)

	var somethingChanged bool = true
	for somethingChanged {
		seatPlan, somethingChanged = applySeatingRulesB(seatPlan, scratchSeatPlan, debug)
	}

	return countOccupiedSeats(seatPlan)
}
