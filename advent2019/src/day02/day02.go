package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

// Read opcode:
//    1: adds together 2 numbers, stores in 3rd. Next 3 numbers are: number A, number B and where to store
//    2: multiplys together 2 numbers, stores in 3rd. Next 3 numbers are: number A, number B and where to store
//   99: exit
//  any: anything else means things have gone wrong
//
// When you've done your op code, step forward 4 positions to work on the next

func gravityAssistProgram(filename string, noun int, verb int, debug bool, part byte) int {
	var currPos int

	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// Only reading 1 line from the file and ignoring errors like a bad person
	lineRead, _ := reader.Read()
	// added explicit close so we can call gravityAssistProgram many times
	csvFile.Close()

	// Create an array the same size as the records we've read from the file, then assign corresponding entries to the array
	programArray := make([]int, len(lineRead))
	for i := 0; i < len(lineRead); i++ {
		programArray[i], _ = strconv.Atoi(lineRead[i])
	}

	programArray[1] = noun
	programArray[2] = verb

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
		switch programArray[currPos] {
		case 99: // Exit
			if debug {
				fmt.Println("Time to exit")
			}
			return programArray[0]
		case 1:
			if debug {
				fmt.Printf("opcode 1: adding %d to %d and storing in position %d\n",
					programArray[programArray[currPos+1]],
					programArray[programArray[currPos+2]],
					programArray[currPos+3])
			}

			programArray[programArray[currPos+3]] = programArray[programArray[currPos+1]] + programArray[programArray[currPos+2]]
			currPos += 4

			if debug {
				fmt.Println("After addition:", programArray)
			}

		case 2:
			if debug {
				fmt.Printf("opcode 2: multiplying %d to %d and storing in position %d\n",
					programArray[programArray[currPos+1]],
					programArray[programArray[currPos+2]],
					programArray[currPos+3])
			}

			programArray[programArray[currPos+3]] = programArray[programArray[currPos+1]] * programArray[programArray[currPos+2]]
			currPos += 4

			if debug {
				fmt.Println("After multiply:", programArray)
			}

		default: // This shouldn't happen
			fmt.Println("Things have gone horribly wrong. Exiting")
			return 0
		}
	}

	return 0
}

// Main routine
func main() {
	var debug bool

	filenamePtr := flag.String("file", "input.txt", "Filename containing the program to run")
	execPartPtr := flag.String("part", "a", "Which part of day18 do you want to calc (a or b)")
	nounPtr := flag.Int("noun", 12, "Noun to change for part a")
	verbPtr := flag.Int("verb", 2, "verb to change for part a")
	flag.BoolVar(&debug, "debug", false, "Turn debug on")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Gravity Assit Program output:", gravityAssistProgram(*filenamePtr, *nounPtr, *verbPtr, debug, 'a'))
	case "b":
		for nounCount := 0; nounCount < 99; nounCount++ {
			for verbCount := 0; verbCount < 99; verbCount++ {
				if debug {
					fmt.Printf("nounCount: %d, verbCount: %d\n", nounCount, verbCount)
				}

				if gravityAssistProgram(*filenamePtr, nounCount, verbCount, debug, 'b') == 19690720 {
					if debug {
						fmt.Printf("Noun: %d Verb: %d\n", nounCount, verbCount)
					}
					fmt.Println("Part b - 100 * Noun + Verb:", 100*nounCount+verbCount)
					return
				}
			}
		}
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
