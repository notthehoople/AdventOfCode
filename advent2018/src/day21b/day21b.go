package main

import (
	"fmt"
//	"flag"
)

// Main routine
func runProgram(reg0Val int, maxLoop int) (int, int) {
		var reg0, reg2, reg3, reg4, reg5 int = 0, 0, 0, 0, 0
		var loopCounter int = 0

		reg0 = reg0Val
	
		reg5 = 0
		
	goto6:
		reg2 = (reg5 | 65536)
		reg5 = 7571367

	goto8:
		loopCounter++
		reg4 = (reg2 & 255)
		reg5 += reg4
		reg5 = (reg5 & 16777215)
		reg5 = reg5 * 65899
		reg5 = (reg5 & 16777215)
		if 256 > reg2 {
			goto goto28
		}

		reg4 = 0

	goto18:
		reg3 = reg4 + 1
		reg3 = reg3 * 256
		if reg3 > reg2 {
			goto goto26
		}
			
		reg4++
		goto goto18

	goto26:
		reg2 = reg4
		goto goto8

	goto28:
		if loopCounter > maxLoop {
			return 0, 0
		}
		fmt.Printf("reg0: %d reg5: %d\n", reg0, reg5)
		if reg5 == reg0 {
			return reg0, loopCounter
		} else {
			goto goto6
		}
	}

// Main routine
// In Day 21 part b we're looking at the LAST value before the loop repeats.
// No idea how to code that yet. Will hopefully come back to it. To find it run the code and dump to a text file
// Search for the LAST reg0: 0 code that doesn't repeat in the reg0: 0 loop. That's your answer.
func main() {
	var maxLoop = 50000

	// Unlike part a we don't need to loop lots. In fact, part a doesn't need to loop lots either. Doh
	for tryIt := 0; tryIt < 3; tryIt++ {
		result, loopResult := runProgram(tryIt, maxLoop)
		if result != 0 {
			fmt.Println("Result worked:", result, loopResult, maxLoop)
			if maxLoop > loopResult {
				maxLoop = loopResult + 1
			}
		} else {
			//fmt.Println("Failed:", tryIt)
		}
	}
}