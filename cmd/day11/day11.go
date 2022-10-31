package main

import (
	"fmt"
)

func countOccupiedSeats(seatPlan [][]byte) int {
	var occupiedSeats int = 0

	for y := 0; y < len(seatPlan); y++ {
		for x := 0; x < len(seatPlan[0]); x++ {
			if seatPlan[y][x] == '#' {
				occupiedSeats++
			}
		}
	}
	return occupiedSeats
}

func setSeatPlan(puzzleInput []string, seatPlan [][]byte) {
	for y := 0; y < len(puzzleInput); y++ {
		for x := 0; x < len(puzzleInput[0]); x++ {
			seatPlan[y][x] = puzzleInput[y][x]
		}
	}
}

func checkArea(tempArea [][]byte, yCoord int, xCoord int) (int, int) {
	var occupiedSeats int = 0
	var emptySeats int = 0
	var startX, startY, maxX, maxY int

	// To consider: the edges. So if xCoord / yCoord == 0 or xCoord / yCoord = len(tempArea)

	// x x   x o   o x   x x   x o x   x x x   x x   x x
	// o x   x x   x x   x o   x x x   x o x   o x   x o
	//                                         x x   x x

	startX = xCoord - 1
	startY = yCoord - 1
	maxX = xCoord + 1
	maxY = yCoord + 1

	if xCoord == 0 {
		startX = 0
	} else {
		if xCoord == len(tempArea[yCoord])-1 {
			maxX = xCoord
		}
	}
	if yCoord == 0 {
		startY = 0
	} else {
		if yCoord == len(tempArea)-1 {
			maxY = yCoord
		}
	}

	for i := startY; i <= maxY; i++ {
		for j := startX; j <= maxX; j++ {
			if i == yCoord && j == xCoord {
				// Do nothing
			} else {
				if tempArea[i][j] == '#' {
					occupiedSeats++
				} else {
					if tempArea[i][j] == 'L' {
						emptySeats++
					}
				}
			}
		}
	}

	return occupiedSeats, emptySeats
}

func applySeatingRules(oldSeatPlan [][]byte, newSeatPlan [][]byte, debug bool) ([][]byte, bool) {
	var occupiedSeats int
	var somethingChanged bool = false

	// Loop through old array and apply rules when copying to new array
	for y := 0; y < len(oldSeatPlan); y++ {
		for x := 0; x < len(oldSeatPlan[0]); x++ {
			newSeatPlan[y][x] = oldSeatPlan[y][x]
			switch oldSeatPlan[y][x] {
			case 'L':
				occupiedSeats, _ = checkArea(oldSeatPlan, y, x)
				if occupiedSeats == 0 {
					newSeatPlan[y][x] = '#'
					somethingChanged = true
				}
				break
			case '#':
				occupiedSeats, _ = checkArea(oldSeatPlan, y, x)
				if occupiedSeats >= 4 {
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

	if debug {
		fmt.Println("====== NEW SEAT PLAN =====")
		print2DArray(oldSeatPlan)
	}

	return oldSeatPlan, somethingChanged
}

func howManyFilledSeats(filename string, part byte, debug bool) int {
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
	if debug {
		fmt.Println("====== START =====")
		print2DArray(seatPlan)
	}

	var somethingChanged bool = true
	for somethingChanged {
		seatPlan, somethingChanged = applySeatingRules(seatPlan, scratchSeatPlan, debug)
	}

	return countOccupiedSeats(seatPlan)
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
		fmt.Println("occupied seats:", howManyFilledSeats(filenamePtr, execPart, debug))
	} else {
		fmt.Println("occupied seats:", howManyFilledSeatsB(filenamePtr, execPart, debug))
	}
}
