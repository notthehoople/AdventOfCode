package main

import (
	"AdventOfCode-go/advent2022/utils"
	"fmt"
)

func isResultCycle(cycleCount int) bool {
	return (cycleCount-20)%40 == 0
}

func calcSignalStrengths(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)

	var registerX int = 1
	var cycleCount int = 1
	var signalStrength int

	for _, instruction := range puzzleInput {
		var command string
		var valueV int
		fmt.Sscanf(instruction, "%s %d\n", &command, &valueV)
		//fmt.Printf("Command: %s Value: %d\n", command, valueV)

		switch command {
		case "noop":
			cycleCount++

		case "addx":
			// Note: the instructions specifically say that the addition happens at the END of the 2 cycles that addx costs
			//       Need to check that we haven't tripped over one of the "answer delivery" cycles (20, then every 40 after that) in the middle of this command
			cycleCount++
			if isResultCycle(cycleCount) {
				fmt.Printf("CycleCount: %d registerX: %d signalStrength: %d\n", cycleCount, registerX, (registerX * cycleCount))
				signalStrength += (registerX * cycleCount)
			}
			cycleCount++
			registerX += valueV

		default:
			fmt.Println("Corrupt error file at", command, valueV)
			return 0
		}

		if isResultCycle(cycleCount) {
			fmt.Printf("CycleCount: %d registerX: %d signalStrength: %d\n", cycleCount, registerX, (registerX * cycleCount))

			signalStrength += (registerX * cycleCount)
		}
	}

	return signalStrength
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Result is: %d\n", calcSignalStrengths(filenamePtr, execPart, debug))
	case 'b':
		fmt.Println("Not implemented yet")
	case 'z':
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
