package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calcBusSchedules(filename string, part byte, debug bool) int {
	var busToTakeID, minutesWaiting int = 0, 999999
	var runningBus []int

	puzzleInput, _ := readFile(filename)

	earliestTime, _ := strconv.Atoi(puzzleInput[0])
	busSchedule := strings.Split(puzzleInput[1], ",")

	runningBus = make([]int, len(busSchedule))

	fmt.Println("Earliest time:", earliestTime)
	var busCounter int = 0
	for _, bus := range busSchedule {
		if bus != "x" {
			runningBus[busCounter], _ = strconv.Atoi(bus)
			busCounter++
		}
	}

	var busArrivalTime int
	for _, bus := range runningBus {
		if bus != 0 {
			// Working with ints so division will remove remainder and give whole numbers
			busArrivalTime = (earliestTime/bus)*bus + bus
			fmt.Printf("Bus: %d Bus Time: %d\n", bus, busArrivalTime)
			if busArrivalTime-earliestTime < minutesWaiting {
				busToTakeID = bus
				minutesWaiting = busArrivalTime - earliestTime
			}
		}
	}

	return busToTakeID * minutesWaiting
}

// Main routine
func main() {
	filenamePtr, execPart, debug, test := catchUserInput()

	if test {
		return
	}

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else if execPart == 'a' {
		fmt.Println("Part a answer:", calcBusSchedules(filenamePtr, execPart, debug))
	} else {
		fmt.Println("Part b answer:", calcBusCompetition(filenamePtr, execPart, debug))
	}
}
