package main

import (
	"container/ring"
	"fmt"
)

func playCrabCups(cupsRing *ring.Ring, moves int, debug bool) (*ring.Ring, int, int) {
	var maxCupLabel int
	var ringLength int

	ringLength = cupsRing.Len()
	ringIndex := make(map[int]*ring.Ring, ringLength)

	// loop through the whole ring and grab the maximum cup value from the list
	for i := 1; i <= ringLength; i++ {
		if cupsRing.Value.(int) > maxCupLabel {
			maxCupLabel = cupsRing.Value.(int)
		}
		ringIndex[cupsRing.Value.(int)] = cupsRing
		cupsRing = cupsRing.Next()
	}

	// ring is back at the starting position
	var destinationCup int = 0
	for turn := 1; turn <= moves; turn++ {
		if debug {
			fmt.Printf("\n-- move %d --\n", turn)
			// Only print out the ring if we're not doing the huge ring size
			if ringLength < 100 {
				fmt.Printf("cups: ")
				for i := 1; i <= ringLength; i++ {
					fmt.Printf("%d ", cupsRing.Value)
					cupsRing = cupsRing.Next()
				}
				fmt.Printf("\n")
			}
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

		ringIndex[destinationCup].Link(pickedCups)

		// move clockwise in the ring
		cupsRing = cupsRing.Next()
	}

	if debug && ringLength < 100 {
		fmt.Printf("\n-- final --\n")
		fmt.Printf("cups: ")
		for i := 1; i <= ringLength; i++ {
			fmt.Printf("%d ", cupsRing.Value)
			cupsRing = cupsRing.Next()
		}
		fmt.Printf("\n")
	}

	return cupsRing, ringIndex[1].Next().Value.(int), ringIndex[1].Next().Next().Value.(int)
}

func crabGamePartB(puzzleInput string, moves int, debug bool) int {
	var maxPuzzleInput int
	var oneRight, twoRight int

	// Crab has many cups. 1 million. On a raft
	cupsRing := ring.New(1000000)
	for i := 0; i < len(puzzleInput); i++ {
		cupsRing.Value = int(puzzleInput[i] - '0')
		if cupsRing.Value.(int) > maxPuzzleInput {
			maxPuzzleInput = cupsRing.Value.(int)
		}
		cupsRing = cupsRing.Next()
	}

	for i := maxPuzzleInput + 1; i <= 1000000; i++ {
		cupsRing.Value = i
		cupsRing = cupsRing.Next()
	}

	_, oneRight, twoRight = playCrabCups(cupsRing, moves, debug)

	return oneRight * twoRight
}

func crabGamePartA(puzzleInput string, moves int, debug bool) string {
	var resultString []byte
	resultString = make([]byte, len(puzzleInput))

	cupsRing := ring.New(len(puzzleInput))
	for i := 0; i < len(puzzleInput); i++ {
		cupsRing.Value = int(puzzleInput[i] - '0')
		cupsRing = cupsRing.Next()
	}

	cupsRing, _, _ = playCrabCups(cupsRing, moves, debug)

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
		fmt.Println("Part b test:", crabGamePartB("389125467", 10, debug))
		fmt.Println("Part b test:", crabGamePartB("389125467", 10000000, debug))

		fmt.Println("Part b real:", crabGamePartB("716892543", 10000000, debug))

	}
}
