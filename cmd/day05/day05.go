package main

import (
	"flag"
	"fmt"
)

func catchUserInput() (string, byte, bool) {
	var debug bool

	filenamePtr := flag.String("file", "testInput.txt", "Filename containing the program to run")
	execPartPtr := flag.String("part", "a", "Which part of the puzzle do you want to calc (a or b)")
	flag.BoolVar(&debug, "debug", false, "Turn debug on")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		return *filenamePtr, 'a', debug
	case "b":
		return *filenamePtr, 'b', debug

	default:
		return *filenamePtr, 'z', debug
	}
}

func calcRange(calcChar rune, lowerLimit int, upperLimit int, debug bool) (int, int) {
	var newSize int
	var newLowerLimit, newUpperLimit int

	size := upperLimit - lowerLimit + 1
	if debug {
		fmt.Printf("Before calc... size: %d lowerLimit: %d upperLimit: %d\n", size, lowerLimit, upperLimit)
	}

	newSize = size / 2
	if calcChar == 'F' || calcChar == 'L' {
		newLowerLimit = lowerLimit
		newUpperLimit = lowerLimit + newSize - 1
	} else {
		newLowerLimit = lowerLimit + newSize
		newUpperLimit = upperLimit
	}

	if debug {
		fmt.Printf("After calc... size: %d lowerLimit: %d upperLimit: %d\n", newSize, newLowerLimit, newUpperLimit)
	}
	return newLowerLimit, newUpperLimit
}

func decodeSinglePass(boardingPass string, part byte, debug bool) int {
	if debug {
		fmt.Println("Processing: ", boardingPass)
	}

	//var rows int = 128
	var lowerLimit int = 0
	var upperLimit int = 127
	var rowNumber int = 0
	for _, rowChar := range boardingPass[0:7] {
		if debug {
			fmt.Printf("Row Instruction: %c\n", rowChar)
		}
		lowerLimit, upperLimit = calcRange(rowChar, lowerLimit, upperLimit, debug)
	}
	rowNumber = lowerLimit

	lowerLimit = 0
	upperLimit = 7
	var columnNumber int = 0
	for _, columnChar := range boardingPass[7:] {
		if debug {
			fmt.Printf("Column Instruction: %c\n", columnChar)
		}
		lowerLimit, upperLimit = calcRange(columnChar, lowerLimit, upperLimit, debug)
	}
	columnNumber = lowerLimit

	return (rowNumber * 8) + columnNumber
}

func decodeBoardingPasses(filename string, part byte, debug bool) int {
	var highestSeatID int = 0
	var currentSeatID int

	puzzleInput, _ := readFile(filename)

	for _, boardingPass := range puzzleInput {
		currentSeatID = decodeSinglePass(boardingPass, part, debug)
		if currentSeatID > highestSeatID {
			highestSeatID = currentSeatID
		}
	}

	return highestSeatID
}

// Main routine
func main() {
	filenamePtr, execPart, debug := catchUserInput()
	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Println("Highest Seat Number: ", decodeBoardingPasses(filenamePtr, execPart, debug))
	}

	fmt.Printf("Boarding Card: %s Seat Number: %d\n", "FBFBBFFRLR", decodeSinglePass("FBFBBFFRLR", execPart, debug))
	fmt.Printf("Boarding Card: %s Seat Number: %d\n", "BFFFBBFRRR", decodeSinglePass("BFFFBBFRRR", execPart, debug))
	fmt.Printf("Boarding Card: %s Seat Number: %d\n", "FFFBBBFRRR", decodeSinglePass("FFFBBBFRRR", execPart, debug))
	fmt.Printf("Boarding Card: %s Seat Number: %d\n", "BBFFBBFRLL", decodeSinglePass("BBFFBBFRLL", execPart, debug))

}
