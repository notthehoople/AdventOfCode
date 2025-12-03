package main

import (
	"AdventOfCode-go/advent2025/utils"
	"fmt"
)

func day03(filename string, part byte, debug bool) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)

	for _, batteryBank := range puzzleInput {

		var maxPos, maxPower int = 0, 0
		for i, battery := range batteryBank {
			// Look for the highest FIRST battery power. Can't be at the last position on the line

			currBattery := int(battery - '0')

			if (currBattery > maxPower) && (i != len(batteryBank)-1) {
				maxPower = currBattery
				maxPos = i
			}

		}

		// After we have the highest FIRST battery power, look for the next highest in the rest of the string

		var maxNext int = 0
		for i := maxPos + 1; i < len(batteryBank); i++ {
			if int(batteryBank[i]-'0') > maxNext {
				maxNext = int(batteryBank[i] - '0')
			}
		}

		result += maxPower*10 + maxNext
		if debug {
			fmt.Printf("MaxPower: %d MaxNext: %d Actual: %d\n", maxPower, maxNext, maxPower*10+maxNext)
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
		fmt.Printf("Result is: %d\n", day03(filenamePtr, execPart, debug))
	}
}
