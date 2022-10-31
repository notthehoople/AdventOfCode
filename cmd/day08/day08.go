package main

import (
	"fmt"
)

type bootCodeStruct struct {
	operator string
	opAmount int
}

func runSingleOp(singleOp bootCodeStruct, currentLine int, accumulator int, part byte, debug bool) (int, int) {
	var nextLine int

	switch singleOp.operator {
	case "nop": // Does nothing. Ignore the opAmount.
		nextLine = currentLine + 1
		break
	case "acc":
		accumulator += singleOp.opAmount
		nextLine = currentLine + 1
	case "jmp":
		nextLine = currentLine + singleOp.opAmount
	default:
		fmt.Println("Code is invalid")
		break
	}

	return nextLine, accumulator
}

func runAllBootCode(filename string, part byte, debug bool) int {
	var codeAlreadyRun [1000]bool
	var bootCode [1000]bootCodeStruct
	var accumulator int

	puzzleInput, _ := readFile(filename)

	// Process the boot code into a more usable form
	for item, operatorLine := range puzzleInput {
		fmt.Sscanf(operatorLine, "%s %d", &bootCode[item].operator, &bootCode[item].opAmount)
	}

	if debug {
		// Print the boot code program
		for i := 0; i < len(bootCode); i++ {
			if bootCode[i].operator == "" {
				break
			}
			fmt.Printf("%d Op: %s OpAmount: %d\n", i, bootCode[i].operator, bootCode[i].opAmount)
		}
	}

	// Run the boot code program
	var currentLine = 0
	for true {
		if codeAlreadyRun[currentLine] {
			break
		}
		codeAlreadyRun[currentLine] = true
		currentLine, accumulator = runSingleOp(bootCode[currentLine], currentLine, accumulator, part, debug)
	}

	return accumulator
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
		fmt.Println("Accumulator:", runAllBootCode(filenamePtr, execPart, debug))
	} else {
		fmt.Println("Accumulator on working code:", runAllBootCodePartB(filenamePtr, execPart, debug))
	}
}
