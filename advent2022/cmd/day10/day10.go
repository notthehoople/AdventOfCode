package main

import (
	"AdventOfCode-go/advent2022/utils"
	"fmt"
)

func isResultCycle(cycleCount int) bool {
	return (cycleCount-20)%40 == 0
}

// Print the display, converting the 1 and 0 values into '#' and '.' respectively
func printDisplay(toPrint [7][40]int) {
	for y := 0; y < len(toPrint); y++ {
		for x := 0; x < len(toPrint[y]); x++ {
			if toPrint[y][x] == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Printf("\n")
	}
}

// Set the Display Position on the CRT. The screen is 40 pixels by 6. If we've reached the end of the scan line, start at the beginning of the line below
func setDisplayPos(displayX int, displayY int) (int, int) {
	displayX++
	if (displayX % 40) == 0 {
		displayX = 0
		displayY++
	}
	return displayX, displayY
}

// Can we see the sprite or not? The sprite is 3 pixels wide, so check if we're drawing where the sprite is
// return 1 for visible sprite, 0 if not visible
func setDisplayPixel(displayX int, registerX int) int {
	if (displayX >= registerX-1) && (displayX <= registerX+1) {
		// draw the sprite
		return 1
	}
	return 0
}

func calcSignalStrengths(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)

	var registerX int = 1
	var cycleCount int = 1
	var signalStrength int
	var displayX, displayY int
	var display [7][40]int

	for _, instruction := range puzzleInput {
		var command string
		var valueV int
		fmt.Sscanf(instruction, "%s %d\n", &command, &valueV)

		switch command {
		case "noop":
			cycleCount++
			display[displayY][displayX] = setDisplayPixel(displayX, registerX)
			displayX, displayY = setDisplayPos(displayX, displayY)

		case "addx":
			// Note: the instructions specifically say that the addition happens at the END of the 2 cycles that addx costs
			//       Need to check that we haven't tripped over one of the "answer delivery" cycles (20, then every 40 after that) in the middle of this command
			cycleCount++
			display[displayY][displayX] = setDisplayPixel(displayX, registerX)
			displayX, displayY = setDisplayPos(displayX, displayY)

			display[displayY][displayX] = setDisplayPixel(displayX, registerX)
			displayX, displayY = setDisplayPos(displayX, displayY)

			if isResultCycle(cycleCount) {
				signalStrength += (registerX * cycleCount)
			}
			cycleCount++
			registerX += valueV

		default:
			fmt.Println("Corrupt error file at", command, valueV)
			return 0
		}

		if isResultCycle(cycleCount) {
			signalStrength += (registerX * cycleCount)
		}
	}

	printDisplay(display)
	return signalStrength
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Signal Strength: %d\n", calcSignalStrengths(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("Signal Strength: %d\n", calcSignalStrengths(filenamePtr, execPart, debug))
	case 'z':
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
