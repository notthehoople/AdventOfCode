package main

import (
	"fmt"
	"strconv"
	//"strings"
)

func convertNumberToBinaryString(number int) []byte {
	builtString := []byte("000000000000000000000000000000000000")

	binaryString := strconv.FormatInt(int64(number), 2)
	copyPlace := len(binaryString) - 1
	for i := 35; i >= 0 && copyPlace >= 0; i-- {
		builtString[i] = binaryString[copyPlace]
		copyPlace--
	}

	return builtString
}

func convertBinaryStringToNumber(binaryNumber string) int {
	number, _ := strconv.ParseInt(binaryNumber, 2, 64)
	return int(number)
}

func applyMaskToNumber(mask string, number int) int {

	binaryString := convertNumberToBinaryString(number)

	for i := len(mask) - 1; i >= 0; i-- {
		switch mask[i] {
		case 'X':
			break
		case '1':
			binaryString[i] = '1'
			break
		case '0':
			binaryString[i] = '0'
			break
		default:
			panic("Corrupt mask in input")
		}
	}

	return convertBinaryStringToNumber(string(binaryString))
}

func calcMemoryAddresses(filename string, part byte, debug bool) int {
	var memory map[int]int
	var rawmask string
	var matchedMask int
	var matchedMemAddress, memoryAddress, memoryValue int

	puzzleInput, _ := readFile(filename)

	memory = make(map[int]int, len(puzzleInput))

	for _, line := range puzzleInput {

		matchedMask, _ = fmt.Sscanf(line, "mask = %s", &rawmask)
		if matchedMask > 0 {
			// Found a mask
			if debug {
				fmt.Printf("Mask: %s\n", rawmask)
			}
		} else {
			// Must be a memory address
			matchedMemAddress, _ = fmt.Sscanf(line, "mem[%d] = %d", &memoryAddress, &memoryValue)
			if matchedMemAddress < 1 {
				panic("Duff content in input file")
			} else {
				memory[memoryAddress] = applyMaskToNumber(rawmask, memoryValue)
				if debug {
					fmt.Printf("Found mem address %d to be set to value %d\n", memoryAddress, memoryValue)
					fmt.Printf("Memory address: %d now %d\n", memoryAddress, memory[memoryAddress])
				}
			}
		}
	}

	var totalMemory int = 0
	for _, value := range memory {
		totalMemory += value
	}

	return totalMemory
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
		fmt.Println("Part b answer:", calcMemoryAddressesB(filenamePtr, execPart, debug))
	}
}
