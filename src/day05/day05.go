package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

// Day05 extensions: New Opcodes
//    3: takes a single integer as input and saves it to the position given by its only parameter. For example, the instruction 3,50 would take an input value and store it at address 50.
//    4: outputs the value of its only parameter. For example, the instruction 4,50 would output the value at address 50.
//
// Day05 extensions: Parameter Modes
//    0: parameter is interpreteted as a position. if the parameter is 50, its value is the value stored at address 50 in memory.
//		   Until now, ALL parameters have worked in POSITION mode
//    1: immediate mode. Parmeter is taken to be a VALUE. If the parameter is 50, its value is simply 50.
//
// Day05 opcode extension into 5 digits
//    ABCDE: DE - 2 digit opcode (e.g. 01, 02, 03, 04)
//			  C - mode of 1st parameter (if not present then 0)
//			  B - mode of 2nd parameter
//			  A - mode of 3rd parameter
//
// Read opcode:
//    1: adds together 2 numbers, stores in 3rd. Next 3 numbers are: number A, number B and where to store
//    2: multiplys together 2 numbers, stores in 3rd. Next 3 numbers are: number A, number B and where to store
//   99: exit
//  any: anything else means things have gone wrong
//
// When you've done your op code, step forward 4 positions to work on the next

func intcodeInitiation(filename string, inputInstruction int, debug bool, part byte) int {
	var currPos int
	var opcode, firstParamMode, secondParamMode, thirdParamMode int
	var firstValue, secondValue, thirdValue int
	var diagnosticCode int

	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// Only reading 1 line from the file and ignoring errors like a bad person
	lineRead, _ := reader.Read()
	// added explicit close in case we need to call the routine many times
	csvFile.Close()

	// Create an array the same size as the records we've read from the file, then assign corresponding entries to the array
	programArray := make([]int, len(lineRead))
	for i := 0; i < len(lineRead); i++ {
		programArray[i], _ = strconv.Atoi(lineRead[i])
	}

	if debug {
		fmt.Println(programArray)
	}

	// While something to do
	// Read op code at current position
	// If 99, exit and output our result
	// If 1, use next 3 numbers as positions and work on them
	// If 2, use next 3 numbers as positions and work on them
	// If anything else, output an error and quit

	for {
		// Extra digits from the opcode. ABCDE: DE - 2 digit opcode (e.g. 01), C, D and E are the mode of the 1st, 2nd, 3rd parameters
		opcode = (programArray[currPos]/10%10)*10 + programArray[currPos]%10
		firstParamMode = programArray[currPos] / 100 % 10
		secondParamMode = programArray[currPos] / 1000 % 10
		thirdParamMode = programArray[currPos] / 10000 % 10

		if debug {
			fmt.Printf("opcode: %2d first: %d second: %d third: %d\n", opcode, firstParamMode, secondParamMode, thirdParamMode)
		}

		switch opcode {
		case 99: // Exit
			if debug {
				fmt.Println("Time to exit")
			}
			return diagnosticCode

		case 1: // Addition
			if firstParamMode == 0 {
				firstValue = programArray[currPos+1]
			} else {
				firstValue = currPos + 1
			}

			if secondParamMode == 0 {
				secondValue = programArray[currPos+2]
			} else {
				secondValue = currPos + 2
			}

			if thirdParamMode == 0 {
				thirdValue = programArray[currPos+3]
			} else {
				fmt.Println("why are we here? Immediate mode for the result should never happen")
				thirdValue = currPos + 3
			}

			if debug {
				fmt.Printf("opcode %d: adding %d to %d and storing in position %d\n",
					opcode,
					programArray[firstValue],
					programArray[secondValue],
					programArray[thirdValue])
			}

			programArray[thirdValue] = programArray[firstValue] + programArray[secondValue]
			currPos += 4

			if debug {
				fmt.Println("After addition:", programArray)
			}

		case 2: // Multiply
			if firstParamMode == 0 {
				firstValue = programArray[currPos+1]
			} else {
				firstValue = currPos + 1
			}

			if secondParamMode == 0 {
				secondValue = programArray[currPos+2]
			} else {
				secondValue = currPos + 2
			}

			if thirdParamMode == 0 {
				thirdValue = programArray[currPos+3]
			} else {
				fmt.Println("why are we here? Immediate mode for the result should never happen")
				thirdValue = currPos + 3
			}

			if debug {
				fmt.Printf("opcode %d: multiplying %d to %d and storing in position %d\n",
					opcode,
					programArray[firstValue],
					programArray[secondValue],
					programArray[thirdValue])
			}

			programArray[thirdValue] = programArray[firstValue] * programArray[secondValue]
			currPos += 4

			if debug {
				fmt.Println("After multiply:", programArray)
			}

		case 3:
			// Opcode 3 takes a single integer as input and saves it to the position given by its only parameter. For example,
			// the instruction 3,50 would take an input value and store it at address 50.
			if firstParamMode == 0 {
				programArray[programArray[currPos+1]] = inputInstruction
			} else {
				programArray[currPos+1] = inputInstruction
			}
			currPos += 2

		case 4:
			// 4: outputs the value of its only parameter. For example, the instruction 4,50 would output the value at address 50.
			if firstParamMode == 0 {
				fmt.Println("Output:", programArray[programArray[currPos+1]])
				diagnosticCode = programArray[programArray[currPos+1]]
			} else {
				fmt.Println("Output:", programArray[currPos+1])
				diagnosticCode = programArray[currPos+1]
			}
			currPos += 2

		default: // This shouldn't happen
			fmt.Printf("Code not implemented yet for instruction %d\n", programArray[currPos])
			return 0
		}
	}
}

// Main routine
func main() {
	var debug bool

	filenamePtr := flag.String("file", "input.txt", "Filename containing the program to run")
	execPartPtr := flag.String("part", "a", "Which part of day18 do you want to calc (a or b)")
	inputPtr := flag.Int("input", 1, "Input instruction for the intcode computer")
	flag.BoolVar(&debug, "debug", false, "Turn debug on")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - intcode diagnostic code:", intcodeInitiation(*filenamePtr, *inputPtr, debug, 'a'))
	case "b":
		fmt.Println("Part b - Not implemented yet")
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
