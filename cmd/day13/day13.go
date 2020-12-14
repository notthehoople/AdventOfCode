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
		fmt.Println("7,13,x,x,59,x,31,19 Result is:", checkBusTimesSimple("7,13,x,x,59,x,31,19", false))
		fmt.Println("17,x,13,19 Result is:", checkBusTimesSimple("17,x,13,19", false))
		fmt.Println("67,7,59,61 Result is:", checkBusTimesSimple("67,7,59,61", false))
		fmt.Println("67,x,7,59,61 Result is:", checkBusTimesSimple("67,x,7,59,61", false))
		fmt.Println("67,7,x,59,61 Result is:", checkBusTimesSimple("67,7,x,59,61", false))
		fmt.Println("1789,37,47,1889 Result is:", checkBusTimesSimple("1789,37,47,1889", false))
		fmt.Println("Real Result is:", checkBusTimesSimple("29,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,433,x,x,x,x,x,x,x,x,x,x,x,x,13,17,x,x,x,x,19,x,x,x,23,x,x,x,x,x,x,x,977,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,41", false))
	}
}
