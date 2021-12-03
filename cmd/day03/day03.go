package main

import (
	"aoc/advent2021/utils"
	"fmt"
	"strconv"
)

func calcPowerConsumption(subDiagnostics []string, debug bool) int {
	/*
		Each bit in the gamma rate can be determined by finding the most common bit in
		the corresponding position of all numbers in the diagnostic report.
		The epsilon rate is calculated in a similar way; rather than use the most
		common bit, the least common bit from each position is used.
	*/
	oneCount := make(map[int]int)
	zeroCount := make(map[int]int)

	for i := 0; i < len(subDiagnostics); i++ {
		for bitPos := len(subDiagnostics[i]) - 1; bitPos >= 0; bitPos-- {
			if subDiagnostics[i][bitPos] == '1' {
				oneCount[bitPos]++
			} else {
				zeroCount[bitPos]++
			}
			if debug {
				fmt.Printf("Diagnostic Record: %s\n", subDiagnostics[i])
				fmt.Printf("Bit: %c\n", subDiagnostics[i][bitPos])
			}
		}
	}

	gammaRateSlice := make([]byte, len(subDiagnostics[0]))
	epsilonRateSlice := make([]byte, len(subDiagnostics[0]))
	for i := 0; i < len(subDiagnostics[0]); i++ {
		if oneCount[i] > zeroCount[i] {
			gammaRateSlice[i] = '1'
			epsilonRateSlice[i] = '0'
		} else {
			gammaRateSlice[i] = '0'
			epsilonRateSlice[i] = '1'
		}
	}
	gammaRate, _ := strconv.ParseInt(string(gammaRateSlice), 2, 64)
	epsilonRate, _ := strconv.ParseInt(string(epsilonRateSlice), 2, 64)

	return int(gammaRate * epsilonRate)
}

func solveDay(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)
	if part == 'a' {
		return calcPowerConsumption(puzzleInput, debug)
	} else {
		return 0
	}
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", solveDay(filenamePtr, execPart, debug))
	}
}
