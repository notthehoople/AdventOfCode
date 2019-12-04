package main

import (
	"flag"
	"fmt"
)

// func countUniquePasswords
func countUniquePasswords(startRange string, endRange string, part byte) int {
	fmt.Printf("Start Range: %s End Range: %s Part: %c\n", startRange, endRange, part)

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
