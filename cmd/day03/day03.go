package main

import (
	"aoc/advent2021/utils"
	"fmt"
	"strconv"
)

func calcLifeSupportRating(subDiagnostics []string, debug bool) int {
	/*
	   To find oxygen generator rating, determine the most common value (0 or 1) in the
	   current bit position, and keep only numbers with that bit in that position. If 0
	   and 1 are equally common, keep values with a 1 in the position being considered.

	   To find CO2 scrubber rating, determine the least common value (0 or 1) in the
	   current bit position, and keep only numbers with that bit in that position. If 0
	   and 1 are equally common, keep values with a 0 in the position being considered.
	*/

	// Filter maps: keeps a bool to say if a diagnostic record is still "in" or not
	oxygenFilter := make(map[int]bool)
	co2Filter := make(map[int]bool)

	// Initialise the maps to true, corresponding to the number of diagnostic records
	for i := 0; i < len(subDiagnostics); i++ {
		oxygenFilter[i] = true
		co2Filter[i] = true
	}

	/*
		To find oxygen generator rating, determine the most common value (0 or 1) in the
		current bit position, and keep only numbers with that bit in that position. If 0
		and 1 are equally common, keep values with a 1 in the position being considered.
	*/

	var oneCount, zeroCount, entryCount int

	// oxygen
	for checkDigit := 0; checkDigit < len(subDiagnostics[0]); checkDigit++ {
		oneCount = 0
		zeroCount = 0
		for i := 0; i < len(subDiagnostics); i++ {
			if oxygenFilter[i] {
				if subDiagnostics[i][checkDigit] == '1' {
					oneCount++
				} else {
					zeroCount++
				}
			}
		}
		if oneCount < zeroCount {
			// Filter out diagnostics records for use as Oxygen entries
			entryCount = 0
			for oxygenKey, oxygenValue := range oxygenFilter {
				if oxygenValue {
					if subDiagnostics[oxygenKey][checkDigit] == '1' {
						oxygenFilter[oxygenKey] = false
					} else {
						entryCount++
					}
				}
			}
			if entryCount == 1 {
				break
			}
		} else {
			entryCount = 0
			for oxygenKey, oxygenValue := range oxygenFilter {
				if oxygenValue {
					if subDiagnostics[oxygenKey][checkDigit] == '0' {
						oxygenFilter[oxygenKey] = false
					} else {
						entryCount++
					}
				}
			}
			if entryCount == 1 {
				break
			}
		}
	}

	// co2
	for checkDigit := 0; checkDigit < len(subDiagnostics[0]); checkDigit++ {
		oneCount = 0
		zeroCount = 0
		for i := 0; i < len(subDiagnostics); i++ {
			if co2Filter[i] {
				if subDiagnostics[i][checkDigit] == '1' {
					oneCount++
				} else {
					zeroCount++
				}
			}
		}
		if oneCount < zeroCount {
			// Filter out diagnostics records for use as CO2 entries
			entryCount = 0
			for co2Key, co2Value := range co2Filter {
				if co2Value {
					if subDiagnostics[co2Key][checkDigit] == '0' {
						co2Filter[co2Key] = false
					} else {
						entryCount++
					}
				}
			}
			if entryCount == 1 {
				break
			}
		} else {
			entryCount = 0
			for co2Key, co2Value := range co2Filter {
				if co2Value {
					if subDiagnostics[co2Key][checkDigit] == '1' {
						co2Filter[co2Key] = false
					} else {
						entryCount++
					}
				}
			}
			if entryCount == 1 {
				break
			}
		}
	}

	if debug {
		fmt.Println(subDiagnostics)
		fmt.Println(oxygenFilter)
		fmt.Println(co2Filter)
	}

	var oxygenRatingString, co2RatingString string
	var oxygenRating, co2Rating int64
	for oxygenKey, oxygenValue := range oxygenFilter {
		if oxygenValue {
			oxygenRatingString = subDiagnostics[oxygenKey]
			oxygenRating, _ = strconv.ParseInt(oxygenRatingString, 2, 64)
		}
	}
	for co2Key, co2Value := range co2Filter {
		if co2Value {
			co2RatingString = subDiagnostics[co2Key]
			co2Rating, _ = strconv.ParseInt(co2RatingString, 2, 64)
		}
	}

	if debug {
		fmt.Printf("Oxygen: %s CO2: %s\n", oxygenRatingString, co2RatingString)
	}

	return int(oxygenRating * co2Rating)
}

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
		return calcLifeSupportRating(puzzleInput, debug)
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
