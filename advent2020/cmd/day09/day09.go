package main

import (
	"fmt"
	"strconv"
)

func findNumber(filename string, preamble int, part byte, debug bool) int {
	var numberList []int

	puzzleInput, _ := readFile(filename)
	numberList = make([]int, len(puzzleInput))

	// Process the code into a more usable form
	for item, number := range puzzleInput {
		numberList[item], _ = strconv.Atoi(number)
	}

	var keepLooking bool = true
	var checkDigit int = 0
	var failedNumber int = 0
	var foundIndex int = 0

	for i := preamble; i < len(numberList); i++ {
		checkDigit = numberList[i]
		keepLooking = true
		for j := i - preamble; j < i && keepLooking; j++ {
			for k := j + 1; k < i && keepLooking; k++ {
				if numberList[k] != numberList[j] && numberList[k]+numberList[j] == checkDigit {
					keepLooking = false
				}
			}
		}
		if keepLooking {
			// didn't find a number
			failedNumber = checkDigit
			foundIndex = i
			break
		}
	}

	if part == 'b' {
		//The final step in breaking the XMAS encryption relies on the invalid number you just found: you must find a
		//contiguous set of at least two numbers in your list which sum to the invalid number from step 1
		//var firstDigit int
		var sumOfAll int
		var smallest int
		var largest int

		for i := foundIndex - 1; i >= 0; i-- {
			sumOfAll = numberList[i]
			keepLooking = true
			smallest = 9999999
			largest = 0

			for j := i - 1; j >= 0 && keepLooking; j-- {
				sumOfAll += numberList[j]
				if numberList[j] < smallest {
					smallest = numberList[j]
				}
				if numberList[j] > largest {
					largest = numberList[j]
				}
				if sumOfAll == failedNumber {
					//fmt.Printf("Found it! Looking for: %d Smallest number: %d Largest number: %d\n", numberList[i], smallest, largest)
					return smallest + largest
				} else if sumOfAll > failedNumber {
					// This time isn't going to work. Break out and try again
					keepLooking = false
				}
			}
		}
	}

	return failedNumber
}

// Main routine
func main() {
	filenamePtr, execPart, debug, test, preamble := catchUserInput()

	if test {
		return
	}

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else if execPart == 'a' {
		fmt.Println("first failed number:", findNumber(filenamePtr, preamble, execPart, debug))
	} else {
		fmt.Println("failed number items:", findNumber(filenamePtr, preamble, execPart, debug))
	}
}
