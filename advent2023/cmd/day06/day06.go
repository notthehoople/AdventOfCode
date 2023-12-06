package main

import (
	"AdventOfCode-go/advent2023/utils"
	"fmt"
	"strconv"
	"strings"
)

// File format: 2 lines
// Line1: "Time:" followed by 1 or more time in ms of race lengths
// Line2: "Distance:" followed by 1 or more distances in mm for best distance during the corresponding race
//
// Toy boat starts at 0 mm per ms
// For each ms of button held down, boat's speed increases by 1 mm per ms
// Time that button is held down is deducted from time available for boat to move

func day06b(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)

	timesLine := strings.Split(puzzleInput[0], ":")
	raceTimesStr := strings.Join(strings.Fields(timesLine[1]), "")
	raceTimes, _ := strconv.Atoi(raceTimesStr)

	if debug {
		fmt.Println("Race times:", raceTimes)
	}

	distanceLine := strings.Split(puzzleInput[1], ":")
	raceDistanceStr := strings.Join(strings.Fields(distanceLine[1]), "")
	raceDistance, _ := strconv.Atoi(raceDistanceStr)

	if debug {
		fmt.Println("Race distances:", raceDistance)
	}

	var result int = 1

	// For each of the races
	var raceOptions int

	//   Loop through the potential button pressess then work out the distance travelled
	//     Start at 1ms press

	for buttonPress := 0; buttonPress < raceTimes; buttonPress++ {
		boatSpeed := buttonPress
		raceTimeLeft := raceTimes - buttonPress
		if (boatSpeed * raceTimeLeft) > raceDistance {
			if debug {
				fmt.Printf("Button Press: %d Boat Speed: %d Distance: %d RaceBest: %d\n", buttonPress, boatSpeed, boatSpeed*raceTimeLeft, raceDistance)
			}
			raceOptions++
		}
	}
	result *= raceOptions

	return result
}

func day06(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)

	timesLine := strings.Split(puzzleInput[0], ":")
	raceTimesStr := strings.Fields(strings.TrimSpace(timesLine[1]))
	raceTimes := make([]int, len(raceTimesStr))
	for i := 0; i < len(raceTimesStr); i++ {
		raceTimes[i], _ = strconv.Atoi(raceTimesStr[i])
	}

	if debug {
		fmt.Println("Race times:", raceTimes)
	}

	distanceLine := strings.Split(puzzleInput[1], ":")
	raceDistanceStr := strings.Fields(strings.TrimSpace(distanceLine[1]))
	raceDistance := make([]int, len(raceDistanceStr))
	for i := 0; i < len(raceDistanceStr); i++ {
		raceDistance[i], _ = strconv.Atoi(raceDistanceStr[i])
	}
	if debug {
		fmt.Println("Race distances:", raceDistance)
	}

	var result int = 1

	// For each of the races
	for race := 0; race < len(raceTimes); race++ {
		var raceOptions int
		if debug {
			fmt.Println("Race:", race)
		}

		//   Loop through the potential button pressess then work out the distance travelled
		//     Start at 1ms press

		for buttonPress := 0; buttonPress < raceTimes[race]; buttonPress++ {
			boatSpeed := buttonPress
			raceTimeLeft := raceTimes[race] - buttonPress
			if (boatSpeed * raceTimeLeft) > raceDistance[race] {
				if debug {
					fmt.Printf("Button Press: %d Boat Speed: %d Distance: %d RaceBest: %d\n", buttonPress, boatSpeed, boatSpeed*raceTimeLeft, raceDistance[race])
				}
				raceOptions++
			}
		}
		result *= raceOptions
	}

	return result
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Result is: %d\n", day06(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("Result is: %d\n", day06b(filenamePtr, execPart, debug))
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
