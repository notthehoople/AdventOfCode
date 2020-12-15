package main

import (
	"fmt"
)

func processFloatingStrings(binaryNumber string) (string, string) {
	var zeroString, oneString []byte
	var changeMade bool = false

	zeroString = make([]byte, len(binaryNumber))
	oneString = make([]byte, len(binaryNumber))

	for i := len(binaryNumber) - 1; i >= 0; i-- {
		switch binaryNumber[i] {
		case '0':
			fallthrough
		case '1':
			zeroString[i] = binaryNumber[i]
			oneString[i] = binaryNumber[i]
			break
		case 'X':
			if !changeMade {
				zeroString[i] = '0'
				oneString[i] = '1'
				changeMade = true
			} else {
				zeroString[i] = 'X'
				oneString[i] = 'X'
			}
			break
		default:
			panic("Bad number passed into function")
		}
	}
	return string(zeroString), string(oneString)
}

func checkForFloating(stringToCheck string) bool {
	for _, value := range stringToCheck {
		if value == 'X' {
			return true
		}
	}
	return false
}

func applyMaskToMemAddr(mask string, number int, debug bool) []int {
	//func applyMaskToMemAddr(mask string, number int) {

	binaryString := convertNumberToBinaryString(number)

	for i := len(mask) - 1; i >= 0; i-- {
		switch mask[i] {
		case 'X':
			binaryString[i] = 'X'
			break
		case '1':
			binaryString[i] = '1'
			break
		case '0':
			break
		default:
			panic("Corrupt mask in input")
		}
	}
	if debug {
		fmt.Printf("-----------------------------------------------\n")
		fmt.Printf("Mask          : %s\n", mask)
		fmt.Printf("Memory Address: %s\n", binaryString)
	}

	// Loop until no Xs left
	//   give memory address with floating Xs in it to function
	//   get back a pair of strings with the first X changed to 0 and 1
	// EndLoop
	var stringOfMemAdds [10000]string
	var keepLooping = true
	stringOfMemAdds[0], stringOfMemAdds[1] = processFloatingStrings(string(binaryString))

	var readingFrom int = 0
	var writingTo int = 2

	for keepLooping {
		stringOfMemAdds[writingTo], stringOfMemAdds[writingTo+1] = processFloatingStrings(stringOfMemAdds[readingFrom])
		writingTo += 2
		readingFrom++
		keepLooping = checkForFloating(stringOfMemAdds[readingFrom])
	}

	if debug {
		fmt.Println("String of Mem Adds:", stringOfMemAdds[readingFrom:])
	}
	// We have our list. Now build a list of ints to return

	var returnList []int
	for _, answer := range stringOfMemAdds[readingFrom:] {
		if answer != "" {
			returnList = append(returnList, convertBinaryStringToNumber(answer))
		} else {
			break
		}
	}

	if debug {
		fmt.Println(returnList)
	}
	return returnList
}

func calcMemoryAddressesB(filename string, part byte, debug bool) int {
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
		} else {
			// Must be a memory address
			matchedMemAddress, _ = fmt.Sscanf(line, "mem[%d] = %d", &memoryAddress, &memoryValue)
			if matchedMemAddress < 1 {
				panic("Duff content in input file")
			} else {
				for _, i := range applyMaskToMemAddr(rawmask, memoryAddress, debug) {
					memory[i] = memoryValue
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
