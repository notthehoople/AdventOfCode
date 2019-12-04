package main

import (
	"flag"
	"fmt"
)

// func countUniquePasswords
func countUniquePasswords(startRange string, endRange string, part byte) int {

	// Need to count the number of possible passwords in the range passed through
	// A password is a potential IF:
	//   - It is a six-digit number
	//   - The value is within the range bounded by startRange and endRange
	//   - Two adjacent digits are the same (like 22 in 122345).
	//   - Going from left to right, the digits never decrease; they only ever increase or stay the same

	fmt.Printf("Start Range: %s End Range: %s Part: %c\n", startRange, endRange, part)

	// Data validation. Rule 1 - startPtr and endPtr must be exactly 6 digits
	if len(startRange) != 6 || len(endRange) != 6 {
		fmt.Println("Invalid start or end range. Must be 6 digits")
		return 0
	}

	// Start counting from startRange
	//   If reached endRange then done
	//   For the increasing digit, count from same as previous digit upwards (speed up)
	//     If 2 digits the same then count it

	return 0
}

// Main routine
func main() {
	var runTests bool

	flag.BoolVar(&runTests, "runtests", false, "runtests to Run initial tests")
	startPtr := flag.String("start", "111111", "Start range of the passwords")
	endPtr := flag.String("end", "222222", "End range of the passwords")
	execPartPtr := flag.String("part", "a", "Which part of day18 do you want to calc (a or b)")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		if runTests {
			fmt.Println("Part a - test 1:", countUniquePasswords("111111", "222222", 'a'))
		} else {
			fmt.Println("Part a - Number of different passwords:", countUniquePasswords(*startPtr, *endPtr, 'a'))
		}
	case "b":
		fmt.Println("Part b - Not implemented yet")
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
