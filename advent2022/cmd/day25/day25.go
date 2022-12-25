package main

import (
	"AdventOfCode-go/advent2022/utils"
	"fmt"
)

type SNAFUstruct struct {
	SNAFU  string
	number int
}

func convertDecimaToSNAFU(number int) string {
	var SNAFUNumber string

	for number > 0 {
		switch (number + 2) % 5 {
		case 0:
			SNAFUNumber = "=" + SNAFUNumber
		case 1:
			SNAFUNumber = "-" + SNAFUNumber
		case 2:
			SNAFUNumber = "0" + SNAFUNumber
		case 3:
			SNAFUNumber = "1" + SNAFUNumber
		case 4:
			SNAFUNumber = "2" + SNAFUNumber
		}
		number = (number + 2) / 5
	}

	return SNAFUNumber
}

func convertSNAFUToDecimal(line string) int {
	var number int
	var position int = 1

	for loop := len(line) - 1; loop >= 0; loop-- {
		switch line[loop] {
		case '2':
			number += 2 * position
		case '1':
			number += 1 * position
		case '0':
			// Don't really need to do anything here
			number += 0 * position
		case '-':
			number += -1 * position
		case '=':
			number += -2 * position
		}
		position *= 5
	}

	return number
}

func buildSNAFUArray(puzzleInput []string, debug bool) []SNAFUstruct {

	SNAFUArray := make([]SNAFUstruct, len(puzzleInput))

	for index, line := range puzzleInput {
		SNAFUArray[index].SNAFU = line
		SNAFUArray[index].number = convertSNAFUToDecimal(line)
	}

	return SNAFUArray
}

func calcSNAFUNumber(filename string, part byte, debug bool) string {

	puzzleInput, _ := utils.ReadFile(filename)
	SNAFUArray := buildSNAFUArray(puzzleInput, debug)

	var fuelSum int
	for _, element := range SNAFUArray {
		fuelSum += element.number
	}

	return convertDecimaToSNAFU(fuelSum)
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("SNAFU number: %s\n", calcSNAFUNumber(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("Day25 has no puzzle for part B\n")
	case 'z':
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
