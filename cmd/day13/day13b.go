package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calcBusSequences(busRequirements string, debug bool) int {
	var earliestTimeStamp int
	var busToTakeID, minutesWaiting int = 0, 999999
	var runningBus []int

	busSchedule := strings.Split(busRequirements, ",")

	runningBus = make([]int, len(busSchedule))

	var busCounter int = 0
	for _, bus := range busSchedule {
		if bus != "x" {
			runningBus[busCounter], _ = strconv.Atoi(bus)
		}
		// If bus is "x" then leave entry as 0
		busCounter++
	}
	/*
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
	*/
	return earliestTimeStamp
}

func calcBusCompetition(filename string, part byte, debug bool) int {

	// Test cases
	//fmt.Println("7,13,x,x,59,x,31,19:", calcBusSequences("7,13,x,x,59,x,31,19", debug))
	fmt.Println("17,x,13,19:", calcBusSequences("17,x,13,19", debug))
	/*
		calcBusSequences("67,7,59,61", debug)
		calcBusSequences("67,x,7,59,61", debug)
		calcBusSequences("67,7,x,59,61", debug)
		calcBusSequences("1789,37,47,1889", debug)
	*/
	// Do it for real

	/*
		puzzleInput, _ := readFile(filename)
		return calcBusSequences(puzzleInput[1], debug)
	*/
	return 0
}
