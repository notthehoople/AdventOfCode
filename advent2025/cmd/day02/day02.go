package main

import (
	"AdventOfCode-go/advent2025/utils"
	"fmt"
	"strconv"
	"strings"
)

func day02(filename string, part byte, debug bool) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)

	puzzleItems := strings.Split(puzzleInput[0], ",")
	for _, puzzleItem := range puzzleItems {

		if debug {
			fmt.Printf("%s\n", puzzleItem)
		}

		extremes := strings.Split(puzzleItem, "-")
		firstID, _ := strconv.Atoi(extremes[0])
		lastID, _ := strconv.Atoi(extremes[1])

		if part == 'a' {
			// Look for the first half of a testID being the same as the second
			for testID := firstID; testID <= lastID; testID++ {
				if debug {
					fmt.Println("testID:", testID)
				}

				testIDstr := strconv.Itoa(testID)
				// We only care about even number of digits
				if len(testIDstr)%2 == 0 {
					compare := testIDstr[0 : len(testIDstr)/2]

					if strings.Count(testIDstr, compare) > 1 {
						if debug {
							fmt.Println("Invalid ID:", testIDstr)
						}

						result += testID
					}
				}
			}
		} else { // part b
			// Repeats can be any length, including 1 character. We can't filter out strings that have an odd number of characters
			// Try brute force approach to start
			for testID := firstID; testID <= lastID; testID++ {
				// Starting with 1, increment through patterns of more an more characters looking for a match
				// We can filter out whenever the pattern length doesn't divide into the overall length of the testID

				testIDstr := strconv.Itoa(testID)

				for patternSize := 1; patternSize <= len(testIDstr)/2; patternSize++ {
					compare := testIDstr[0:patternSize]
					if debug {
						fmt.Println("testID:", testID, "patternSize:", patternSize, "compare:", compare)
					}

					// If the length of the patternSize exactly divides into the testIDstr
					// AND the compare string is present in a pattern in the testIDstr
					// THEN we have an invalid ID
					// - edge case - need to stop comparing when we've already matched an invalid ID so we don't count it again
					if (len(testIDstr)%patternSize == 0) && (strings.Count(testIDstr, compare) == len(testIDstr)/patternSize) {
						if debug {
							fmt.Println("Invalid ID:", testIDstr)
						}

						result += testID

						// Break out of the inner loop as we've found the invalid ID and don't need to try other options
						break
					}
				}
			}
		}
	}

	return result
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", day02(filenamePtr, execPart, debug))
	}
}
