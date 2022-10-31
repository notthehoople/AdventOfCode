package main

import (
	"strconv"
	"strings"
)

func checkPotential(checkNumber int, firstBus int, secondBus int, position int) bool {

	if (checkNumber+position)%secondBus == 0 {
		return true
	}

	return false
}

func checkBusTimesSimple(busScheduleRaw string, useHint bool) int {

	var busSchedule []int
	var n int = 1
	var keepLooping bool = true
	var testNumber int
	var lookingGood bool

	// Build the bus schedule
	puzzleInput := strings.Split(busScheduleRaw, ",")
	busSchedule = make([]int, len(puzzleInput))
	var j int = 0
	for _, busID := range puzzleInput {
		if busID == "x" {
			busSchedule[j] = 0
		} else {
			busSchedule[j], _ = strconv.Atoi(busID)
		}
		j++
	}

	firstBus := busSchedule[0]
	lastPosition := len(busSchedule) - 1
	lastBus := busSchedule[lastPosition]

	n = 1

	var skipAhead int = 0

	for keepLooping {
		if skipAhead > 0 {
			testNumber += skipAhead
			n = testNumber / firstBus
			skipAhead = 0
		} else {
			testNumber = firstBus * n
		}

		if (testNumber+lastPosition)%lastBus == 0 {
			// found one matching bus. let's check the others
			// since we know these match and the buses are prime numbers, we can skip ahead by busID1 * busID2 to
			//   find the next potential match

			skipAhead = firstBus * lastBus

			lookingGood = true
			for checkPos, checkBus := range busSchedule {
				if checkPos != 0 && checkBus != 0 {
					lookingGood = checkPotential(testNumber, firstBus, checkBus, checkPos)
					if !lookingGood {
						// not a match so we give up checking here
						break
					} else {
						// another match! If nothing else we've got more things to skip on next try
						skipAhead = skipAhead * checkBus
					}
				}
			}
			if lookingGood {
				return testNumber
			}

		}
		if skipAhead == 0 {
			n++
		}
	}
	return 0
}
