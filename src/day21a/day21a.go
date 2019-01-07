package main

import (
	"fmt"
//	"flag"
)

// Main routine
func runProgram(reg0Val int, maxLoop int) int {
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
		return reg5
		if reg5 == reg0 {
			return reg0
		} else {
			goto goto6
		}
	}

// Main routine
// In Day 21 part a we're looking at the FIRST value that makes it through to the comparison.
func main() {
	var maxLoop = 50000

	// Having worked out what's going on with my first "brute force" day21 part a answer, here's one that returns the same result faster
	
	fmt.Println("Answer is:", runProgram(0, maxLoop))
}