package main

import (
	"fmt"
)

func printBootCode(bootCode [1000]bootCodeStruct) {
	// Print the boot code program
	fmt.Println("==========")

	for i := 0; i < len(bootCode); i++ {
		if bootCode[i].operator == "" {
			break
		}
		fmt.Printf("%d Op: %s OpAmount: %d\n", i, bootCode[i].operator, bootCode[i].opAmount)
	}
}

func whichOpToChange(bootCode [1000]bootCodeStruct, changedCode [1000]bool) (int, string) {
	for i := 0; i < len(bootCode); i++ {
		switch bootCode[i].operator {
		case "nop": // If not changed so far, change to jmp
			if !changedCode[i] {
				return i, "jmp"
			}
			break
		case "acc": // No need to change
			break
		case "jmp": // If not changed so far, change to nop
			if !changedCode[i] {
				return i, "nop"
			}
			break
		default:
			break
		}
	}
	return 0, ""
}

func runAllBootCodePartB(filename string, part byte, debug bool) int {
	var origBootCode [1000]bootCodeStruct
	var runBootCode [1000]bootCodeStruct // Boot code changed to test
	var changedCode [1000]bool
	var accumulator int
	var trying int
	var newOperator string

	puzzleInput, _ := readFile(filename)

	// Process the boot code into a more usable form
	for item, operatorLine := range puzzleInput {
		fmt.Sscanf(operatorLine, "%s %d", &origBootCode[item].operator, &origBootCode[item].opAmount)
	}

	if debug {
		printBootCode(origBootCode)
	}

	var keepLooping = true
	for keepLooping {
		// Change one operator each time through the loop

		// Reset the code for another try
		runBootCode = origBootCode
		trying, newOperator = whichOpToChange(runBootCode, changedCode)
		changedCode[trying] = true
		runBootCode[trying].operator = newOperator

		if debug {
			printBootCode(runBootCode)
		}

		// Run the boot code program
		var currentLine = 0
		var accumulator = 0
		var codeAlreadyRun [1000]bool
		for true {
			if codeAlreadyRun[currentLine] {
				break
			} else if runBootCode[currentLine].operator == "" {
				// We made it to the end!
				keepLooping = false
				return accumulator
			}
			codeAlreadyRun[currentLine] = true
			currentLine, accumulator = runSingleOp(runBootCode[currentLine], currentLine, accumulator, part, debug)
		}
	}

	return accumulator
}
