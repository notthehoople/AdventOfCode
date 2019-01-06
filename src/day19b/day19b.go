package main

import (
	"fmt"
)

// Main routine
func main() {
	var reg0, reg2, reg3, reg4 int = 0, 0, 0, 0

	reg4 = 1
	reg3 = 10551408

	two:
		reg2 = 1

	three:
		if (reg4 * reg2) == reg3 {
			fmt.Printf("Looping. Reg0: %d Reg2: %d Reg3: %d Reg4: %d\n", reg0, reg2, reg3, reg4)
			reg0 += reg4
			reg4++
			goto two
		} else {
			reg2++
			if (reg4 * reg2) > reg3 {
				reg4++
				if reg4 > reg3 {
					fmt.Println("Reg0 is:", reg0)
					return
				} else {
					goto two
				}
			} else {
				goto three
			}
		}
}