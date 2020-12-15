package main

import (
	"fmt"
	//"strconv"
	//"strings"
)

func convertStringToMask(rawmask string) int {
	var maskValue int = 0
	var placeValue int = 1

	for i := len(rawmask) - 1; i >= 0; i-- {
		fmt.Printf("Mask digit: %c\n", rawmask[i])
		switch rawmask[i] {
		case 'X':
			break
		case '1':
			maskValue += placeValue
			break
		case '0':
			break
		default:
			panic("Corrupt mask in input")
		}

		placeValue *= 2
	}

	return maskValue
}

func calcMemoryAddresses(filename string, part byte, debug bool) int {
	var memory map[int]int
	var rawmask string
	var matchedMask, usableMask int
	var matchedMemAddress, memoryAddress, memoryValue int

	puzzleInput, _ := readFile(filename)

	memory = make(map[int]int, len(puzzleInput))

	for _, line := range puzzleInput {
		matchedMask, _ = fmt.Sscanf(line, "mask = %s", &rawmask)
		if matchedMask > 0 {
			// Found a mask
			usableMask = convertStringToMask(rawmask)
			fmt.Printf("Found mask %s converted %d\n", rawmask, usableMask)

			// Delete when function built
			if debug {
				fmt.Println(memory[4], matchedMemAddress)
			}
		} else {
			// Must be a memory address
			matchedMemAddress, _ = fmt.Sscanf(line, "mem[%d] = %d", &memoryAddress, &memoryValue)
			fmt.Printf("Found mem address %d to be set to value %d\n", memoryAddress, memoryValue)
		}
	}

	return 0
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
		fmt.Println("Part a answer:", calcMemoryAddresses(filenamePtr, execPart, debug))
	} else {
		fmt.Println("Not implemented yet")
	}
}
