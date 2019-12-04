package main

import (
	"flag"
	"fmt"
	"strconv"
)

// Takes an integer and returns true if two adjacent digits are the same
func containsDoubleLetters(numberToCheck int) bool {
	tempString := strconv.Itoa(numberToCheck)

	for i := 1; i < len(tempString); i++ {
		if tempString[i] == tempString[i-1] {
			return true
		}
	}
	return false
}

// Takes an integer and returns true if, Going from left to right, the digits never decrease; they only ever increase or stay the same
func isAscendingDigits(numberToCheck int) bool {
	tempString := strconv.Itoa(numberToCheck)

	for i := 1; i < len(tempString); i++ {
		if tempString[i] < tempString[i-1] {
			return false
		}
	}

	return true
}

// Need to count the number of possible passwords in the range passed through
// A password is a potential IF:
//   - Rule1: It is a six-digit number
//   - Rule2: Two adjacent digits are the same (like 22 in 122345).
//   - Rule3: Going from left to right, the digits never decrease; they only ever increase or stay the same
//
// func countUniquePasswords
func countUniquePasswords(startRange string, endRange string, part byte) int {
	var countPotential int

	// Data validation. Rule 1 - startPtr and endPtr must be exactly 6 digits
	if len(startRange) != 6 || len(endRange) != 6 {
		fmt.Println("Invalid start or end range. Must be 6 digits")
		return 0
	}

	// Make sure that startRange < endRange so we don't count forever
	if startRange > endRange {
		fmt.Println("Invalid input. Start must be < end")
		return 0
	}

	// Convert from string to integer for easy looping
	currentCheckInt, _ := strconv.Atoi(startRange)
	endRangeInt, _ := strconv.Atoi(endRange)

	// Loop through our potential passwords from start to end
	for ; currentCheckInt <= endRangeInt; currentCheckInt++ {

		// Check if Rule3 (ascending or equal digits) is followed
		if isAscendingDigits(currentCheckInt) {
			// Check if Rule2 (double digit) is followed
			if containsDoubleLetters(currentCheckInt) {
				countPotential++
			}
		}
	}

	return countPotential
}

// Main routine
func main() {
	startPtr := flag.String("start", "111111", "Start range of the passwords")
	endPtr := flag.String("end", "222222", "End range of the passwords")
	execPartPtr := flag.String("part", "a", "Which part of day18 do you want to calc (a or b)")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Number of different passwords:", countUniquePasswords(*startPtr, *endPtr, 'a'))
	case "b":
		fmt.Println("Part b - Not implemented yet")
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
