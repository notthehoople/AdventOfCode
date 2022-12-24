package main

import (
	"AdventOfCode-go/advent2022/utils"
	"container/ring"
	"fmt"
	"strconv"
)

//type Ring struct {
//	Value any // for use by client; untouched by this library
// contains filtered or unexported fields
//}

type position struct {
	index  int
	number int
}

func buildGroveRing(numbersList []int, debug bool) map[position]*ring.Ring {

	positionMap := make(map[position]*ring.Ring, len(numbersList))
	r := ring.New(len(numbersList))

	for i, j := range numbersList {
		positionMap[position{index: i, number: j}] = r
		r.Value = j
		r = r.Next()
	}

	return positionMap
}

func calcGroveCoords(filename string, part byte, debug bool) int {
	const decryptKey int = 811589153

	zeroPos := position{}

	puzzleInput, _ := utils.ReadFile(filename)
	numbersList := make([]int, len(puzzleInput))
	for i := 0; i < len(puzzleInput); i++ {
		numbersList[i], _ = strconv.Atoi(puzzleInput[i])
		if part == 'b' {
			numbersList[i] *= decryptKey
		}
		// We need the zero pos so we can calc the coords starting from 0. Grab it when we see it go passed
		if numbersList[i] == 0 {
			zeroPos.index = i
			zeroPos.number = 0
		}
	}

	// We need a map to point into the ring so we don't lose track of where things are as we're moving things around
	groveCoords := buildGroveRing(numbersList, debug)

	// Loop through the list of numbers, using them to look into the map to get the item to be moved
	var loopCount int = 1
	if part == 'b' {
		loopCount = 10
	}

	// Testing speedup code
	listLength := len(numbersList) - 1
	halfLength := listLength >> 1

	for loop := 0; loop < loopCount; loop++ {
		fmt.Println("Loop:", loop)

		for index, number := range numbersList {
			// Grab the loop pointer for the position we're at in the list of numbers
			ringPos := groveCoords[position{index: index, number: number}].Prev()
			movingItem := ringPos.Unlink(1)

			// With thanks to dhruvmanila, this optimisation from Python's deque.rotate method took my solution from several hours to under a second
			if (number > halfLength) || (number < -halfLength) {
				number %= listLength
				switch {
				case number > halfLength:
					number -= listLength
				case number < -halfLength:
					number += listLength
				}
			}

			ringPos.Move(number).Link(movingItem)
		}
	}
	// Now use the zeroPos we grabbed earlier as the starting point. Do 3 x move(1000) and grab the value from each

	var groveCoordsResult int
	ringPos := groveCoords[zeroPos]
	ringPos = ringPos.Move(1000)
	groveCoordsResult += ringPos.Value.(int)
	ringPos = ringPos.Move(1000)
	groveCoordsResult += ringPos.Value.(int)
	ringPos = ringPos.Move(1000)
	groveCoordsResult += ringPos.Value.(int)

	return groveCoordsResult
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Grove Coords: %d\n", calcGroveCoords(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("Grove Coords: %d\n", calcGroveCoords(filenamePtr, execPart, debug))
	case 'z':
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
