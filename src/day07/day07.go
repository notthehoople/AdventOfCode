package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func processSequence(phase string) []int {
	phaseSequence := make([]int, len(phase))

	for i := 0; i < len(phase); i++ {
		phaseSequence[i], _ = strconv.Atoi(string(phase[i]))
	}
	return phaseSequence
}

// Read file and format it ready for use by the intcomputer
func readFileInput(filename string) []int {
	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// Only reading 1 line from the file and ignoring errors like a bad person
	lineRead, _ := reader.Read()
	// added explicit close in case we need to call the routine many times
	csvFile.Close()

	// Create an array the same size as the records we've read from the file, then assign corresponding entries to the array
	baseProgram := make([]int, len(lineRead))
	for i := 0; i < len(lineRead); i++ {
		baseProgram[i], _ = strconv.Atoi(lineRead[i])
	}
	return baseProgram
}

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
//    1: adds together 2 numbers, stores in 3rd. Next 3 numbers are: number A, number B and where to store
//    2: multiplys together 2 numbers, stores in 3rd. Next 3 numbers are: number A, number B and where to store
//    3: takes a single integer as input and saves it to the position given by its only parameter. For example, the instruction 3,50 would take an input value and store it at address 50.
//    4: outputs the value of its only parameter. For example, the instruction 4,50 would output the value at address 50.
//    5: jump-if-true: if first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise do nothing
//    6: jump-if-false: if first parameter is zero, it sets instruction pointer to the value from second parameter. Otherwise, do nothing.
//    7: is less than: if first parameter is less than second parameter, it stores 1 in the position given by third parameter. Otherwise, stores 0.
//    8: is equals: if first parameter is equal to second parameter, it stores 1 in the position given by third parameter. Otherwise, stores 0.
//   99: exit
//  any: anything else means things have gone wrong

func intcodeComputer(programArray []int, inputInstruction []int, debug bool, part byte) int {
	var currPos int
	var opcode, firstParamMode, secondParamMode, thirdParamMode int
	var firstValue, secondValue, thirdValue int
	var diagnosticCode int

	if debug {
		fmt.Println(programArray)
	}

	for {
		// Extra digits from the opcode. ABCDE: DE - 2 digit opcode (e.g. 01), C, D and E are the mode of the 1st, 2nd, 3rd parameters
		opcode = (programArray[currPos]/10%10)*10 + programArray[currPos]%10
		firstParamMode = programArray[currPos] / 100 % 10
		secondParamMode = programArray[currPos] / 1000 % 10
		thirdParamMode = programArray[currPos] / 10000 % 10

		if debug {
			fmt.Printf("[START] opcode: %2d first: %d second: %d third: %d\n", opcode, firstParamMode, secondParamMode, thirdParamMode)
		}

		if opcode == 1 || opcode == 2 || opcode == 7 || opcode == 8 {
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

			// Write positions are never immediate mode
			thirdValue = programArray[currPos+3]
		} else if opcode == 5 || opcode == 6 {
			if firstParamMode == 0 {
				firstValue = programArray[currPos+1]
			} else {
				firstValue = currPos + 1
			}
		}

		switch opcode {
		case 99: // Exit
			return diagnosticCode

		case 1: // Addition
			if debug {
				fmt.Println("[OP:01] programArray is:", programArray)
				fmt.Printf("[OP:01] opcode %d: adding %d to %d and storing in position %d\n",
					opcode,
					programArray[firstValue],
					programArray[secondValue],
					programArray[thirdValue])
			}

			programArray[thirdValue] = programArray[firstValue] + programArray[secondValue]
			currPos += 4

			if debug {
				fmt.Println("[OP:01] After addition:", programArray)
			}

		case 2: // Multiply
			if debug {
				fmt.Println("[OP:02] programArray is:", programArray)
				fmt.Printf("[OP:02] opcode %d: multiplying %d to %d and storing in position %d\n",
					opcode,
					programArray[firstValue],
					programArray[secondValue],
					programArray[thirdValue])
			}

			programArray[thirdValue] = programArray[firstValue] * programArray[secondValue]
			currPos += 4

			if debug {
				fmt.Println("[OP:02] After multiply:", programArray)
			}

		case 3:
			// Opcode 3 takes a single integer as input and saves it to the position given by its only parameter. For example,
			// the instruction 3,50 would take an input value and store it at address 50.

			// Write to is never in immediate mode
			if debug {
				fmt.Println("[OP:03] programArray is:", programArray)
				fmt.Printf("[OP:03] opcode %d: Takes input %d and stores in position %d\n", opcode, inputInstruction, programArray[firstValue])
			}
			programArray[programArray[currPos+1]] = inputInstruction[0]
			fmt.Println("[OP:03] inputInstruction is:", inputInstruction[0])
			inputInstruction = inputInstruction[1:]
			fmt.Println("[OP:03] inputInstruction is:", inputInstruction[0])

			if debug {
				fmt.Println("[OP:03] After input stored:", programArray)
			}
			currPos += 2

		case 4:
			// 4: outputs the value of its only parameter. For example, the instruction 4,50 would output the value at address 50.
			if debug {
				fmt.Println("[OP:04] programArray is:", programArray)
			}

			if firstParamMode == 0 {
				fmt.Println("[OP:04] Output:", programArray[programArray[currPos+1]])
				diagnosticCode = programArray[programArray[currPos+1]]
			} else {
				fmt.Println("[OP:04] Output:", programArray[currPos+1])
				diagnosticCode = programArray[currPos+1]
			}
			if debug {
				fmt.Println("[OP:04] After output:", programArray)
			}
			currPos += 2

		case 5:
			// 5: jump-if-true: if first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise do nothing
			if programArray[firstValue] != 0 {
				// Set instruction pointer to the value from the second parameter
				if secondParamMode == 0 {
					currPos = programArray[programArray[currPos+2]]
					if debug {
						fmt.Println("[OP:05] currPos after jump:", currPos)
					}
				} else {
					currPos = programArray[currPos+2]
					if debug {
						fmt.Println("[OP:05] currPos after jump:", currPos)
					}
				}
			} else { // Do nothing
				currPos += 3
				if debug {
					fmt.Println("[OP:05] currPos after DO NOTHING:", currPos)
				}
			}

		case 6:
			// 6: jump-if-false: if first parameter is zero, it sets instruction pointer to the value from second parameter. Otherwise, do nothing.

			if programArray[firstValue] == 0 {
				// Set instruction pointer to the value from the second parameter
				if secondParamMode == 0 {
					currPos = programArray[programArray[currPos+2]]
				} else {
					currPos = programArray[currPos+2]
				}
			} else { // Do nothing
				currPos += 3
			}

		case 7:
			// 7: is less than: if first parameter is less than second parameter, it stores 1 in the position given by third parameter. Otherwise, stores 0.
			if debug {
				fmt.Println("[OP:07] programArray is:", programArray)
				fmt.Printf("[OP:07] opcode %d: if %d is less than %d then store 1 in position %d\n",
					opcode,
					programArray[firstValue],
					programArray[secondValue],
					programArray[thirdValue])
			}

			if programArray[firstValue] < programArray[secondValue] {
				programArray[thirdValue] = 1
			} else {
				programArray[thirdValue] = 0
			}
			currPos += 4
			if debug {
				fmt.Println("[OP:07] After less than:", programArray)
			}

		case 8:
			// 8: is equals: if first parameter is equal to second parameter, it stores 1 in the position given by third parameter. Otherwise, stores 0.
			if debug {
				fmt.Println("[OP:08] programArray is:", programArray)
				fmt.Printf("[OP:08] opcode %d: if %d is equal to %d then store 1 in position %d\n",
					opcode,
					programArray[firstValue],
					programArray[secondValue],
					programArray[thirdValue])
			}

			if programArray[firstValue] == programArray[secondValue] {
				programArray[thirdValue] = 1
			} else {
				programArray[thirdValue] = 0
			}
			currPos += 4
			if debug {
				fmt.Println("[OP:08] After equals:", programArray)
			}

		default: // This shouldn't happen
			fmt.Printf("Code not implemented yet for instruction %d\n", programArray[currPos])
			return 0
		}
	}
}

// func intcodeMaxThrusterSignal
// Day07 part a. Run the given program through the intcode computer looking for the sequence that will give the highest output
func intcodeMaxThrusterSignal(filename string, phase string, debug bool, part byte) int {
	var outputSignal int
	var largestOutputSignal int

	baseProgram := readFileInput(filename)

	phaseSequence := processSequence(phase)
	outputSignal = 0
	inputInstruction := make([]int, 5)

	// Need to build a phaseSequence for each run through
	// Each setting is a number from 0 to 4 which is used ONCE per test run
	for a := 0; a < 5; a++ {
		for b := 0; b < 5; b++ {
			for c := 0; c < 5; c++ {
				for d := 0; d < 5; d++ {
					for e := 0; e < 5; e++ {
						if a == b || a == c || a == d || a == e || b == c || b == d || b == e || c == d || c == e || d == e {
							continue
						}
						phaseSequence[0] = a
						phaseSequence[1] = b
						phaseSequence[2] = c
						phaseSequence[3] = d
						phaseSequence[4] = e

						outputSignal = 0

						fmt.Println("Phase Sequence:", phaseSequence)

						for ampRun := 0; ampRun < 5; ampRun++ {
							// Reset the program for the next Amp run
							programArray := baseProgram

							inputInstruction[0] = phaseSequence[ampRun]
							inputInstruction[1] = outputSignal

							if debug {
								fmt.Println("BEFORE Amp:Program", ampRun, programArray)
								fmt.Println("BEFORE inputInstruction", inputInstruction)
							}
							outputSignal = intcodeComputer(programArray, inputInstruction, debug, part)
							if debug {
								fmt.Println("AFTER Amp:Program", ampRun, programArray)
								fmt.Println("AFTER inputInstruction", inputInstruction)
							}

							fmt.Printf("Output signal from Amp%d is %d\n", ampRun, outputSignal)
							if ampRun == 4 && outputSignal > largestOutputSignal {
								fmt.Println("Found a larger output signal:", outputSignal, largestOutputSignal)
								largestOutputSignal = outputSignal
							}
						}
					}
				}
			}
		}
	}

	return largestOutputSignal
}

// Main routine
func main() {
	var debug bool

	filenamePtr := flag.String("file", "input.txt", "Filename containing the program to run")
	execPartPtr := flag.String("part", "a", "Which part of day18 do you want to calc (a or b)")
	phasePtr := flag.String("phase", "43210", "Phase setting sequence")
	flag.BoolVar(&debug, "debug", false, "Turn debug on")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Max thruster signal:", intcodeMaxThrusterSignal(*filenamePtr, *phasePtr, debug, 'a'))
	case "b":
		fmt.Println("Part b - Not implemented yet")
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
