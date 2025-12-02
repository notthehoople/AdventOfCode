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

		if part == 'a' {
			firstID, _ := strconv.Atoi(extremes[0])
			lastID, _ := strconv.Atoi(extremes[1])

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
							fmt.Println("Valid ID:", testIDstr)
						}

						result += testID
					}
				}
			}
		} else { // part b

		}

	}

	if part == 'a' {

		return result
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
