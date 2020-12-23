package main

import (
	"container/ring"
	"fmt"
)

func playCrabCups(cupsRing *ring.Ring, moves int, debug bool) *ring.Ring {
	var maxCupLabel int
	var ringLength int

	ringLength = cupsRing.Len()

	// loop through the whole ring and grab the maximum cup value from the list
	for i := 1; i <= ringLength; i++ {
		if cupsRing.Value.(int) > maxCupLabel {
			maxCupLabel = cupsRing.Value.(int)
		}
		cupsRing = cupsRing.Next()
	}

	// ring is back at the starting position
	var destinationCup int = 0
	for turn := 1; turn <= moves; turn++ {
		if debug {
			fmt.Printf("\n-- move %d --\n", turn)
			fmt.Printf("cups: ")
			for i := 1; i <= ringLength; i++ {
				fmt.Printf("%d ", cupsRing.Value)
				cupsRing = cupsRing.Next()
			}
			fmt.Printf("\n")
			fmt.Println("current cup:", cupsRing.Value)
		}

		// remove the 3 entries clockwise from the current cup. Unlink removes from cupsRing.Next onwards
		pickedCups := cupsRing.Unlink(3)
		if debug {
			fmt.Printf("pick up: %d, %d, %d\n", pickedCups.Value, pickedCups.Next().Value, pickedCups.Next().Next().Value)
		}

		destinationCup = cupsRing.Value.(int) - 1

		var keepLooping bool = true
		for keepLooping {
			if destinationCup < 1 {
				destinationCup = maxCupLabel
			}
			if pickedCups.Value != destinationCup && pickedCups.Next().Value != destinationCup && pickedCups.Next().Next().Value != destinationCup {
				// found a value destination so done
				keepLooping = false
			} else {
				destinationCup--
			}
		}
		if debug {
			fmt.Printf("destination: %d\n", destinationCup)
		}

		// Place the removed items at the destination
		var currentCup int
		currentCup = cupsRing.Value.(int)

		for i := 1; i < ringLength; i++ {
			if cupsRing.Value.(int) == destinationCup {
				cupsRing = cupsRing.Link(pickedCups)
				break
			}
			cupsRing = cupsRing.Next()
		}

		// Must be a better way of doing this.
		// After adding to the ring, search through the ring looking for where we were
		for i := 1; i < ringLength; i++ {
			if cupsRing.Value.(int) == currentCup {
				break
			}
			cupsRing = cupsRing.Next()
		}
		//fmt.Printf("After adding current cup is %d and should be %d\n", cupsRing.Value, currentCup)

		// move clockwise in the ring
		cupsRing = cupsRing.Next()
	}

	if debug {
		fmt.Printf("\n-- final --\n")
		fmt.Printf("cups: ")
		for i := 1; i <= ringLength; i++ {
			fmt.Printf("%d ", cupsRing.Value)
			cupsRing = cupsRing.Next()
		}
		fmt.Printf("\n")
	}

	return cupsRing
}

func crabGamePartA(puzzleInput string, moves int, debug bool) string {
	var resultString []byte
	resultString = make([]byte, len(puzzleInput))

	cupsRing := ring.New(len(puzzleInput))
	for i := 0; i < len(puzzleInput); i++ {
		cupsRing.Value = int(puzzleInput[i] - '0')
		cupsRing = cupsRing.Next()
	}

	cupsRing = playCrabCups(cupsRing, moves, debug)

	// Position ourselves in the ring after the label '1'
	for {
		if cupsRing.Value.(int) == 1 {
			cupsRing = cupsRing.Next()
			break
		}
		cupsRing = cupsRing.Next()
	}

	// Build our answer string
	for i := 1; i <= len(puzzleInput); i++ {
		resultString[i-1] = byte(cupsRing.Value.(int) + '0')
		cupsRing = cupsRing.Next()
		if cupsRing.Value.(int) == 1 {
			break
		}
	}

	return string(resultString)
}

// Main routine
func main() {
	_, execPart, debug, test := catchUserInput()

	if test {
		return
	}

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else if execPart == 'a' {
		fmt.Println("Part a test:", crabGamePartA("389125467", 10, debug))
		fmt.Println("Part a test:", crabGamePartA("389125467", 100, debug))

		fmt.Println("Part a real:", crabGamePartA("716892543", 100, debug))
	} else {
	}
}
