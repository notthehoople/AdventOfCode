package main

import (
	"fmt"
	"flag"
)

// Used to build a work list of water sources to be followed
type nanoCoordsStruct struct {	
	xCoord		int
	yCoord		int
	zCoord		int
	signalRange	int
}

// func: nanoBotControlB
// 
func nanoBotControlB(fileName string, debug bool) (int) {
	var nanoBotLocation []nanoCoordsStruct
	var minX, maxX, minY, maxY, minZ, maxZ int
	var botsInRange, tempDistance int = 0, 0
	var maxBotsInRange, maxArrayPos int = 0, 0

	fileContents, _ := readLines(fileName)
	nanoBotLocation = processInputFile(fileContents)

	minX, maxX, minY, maxY, minZ, maxZ = scanInputForMaxMins(nanoBotLocation)

	if debug {
		fmt.Printf("minX: %d maxX: %d minY: %d maxY: %d minZ: %d maxZ: %d\n", minX, maxX, minY, maxY, minZ, maxZ)
	}


	for i := 0; i < len(nanoBotLocation); i++ {
		botsInRange = 0
		for j := 0; j < len(nanoBotLocation); j++ {
			tempDistance = manhattanDistance(nanoBotLocation[j], nanoBotLocation[i])
			//if tempDistance <= nanoBotLocation[i].signalRange {
			if tempDistance <= 50000000 {
					botsInRange++
			}

		}
		if debug {
			if botsInRange > 300 {
				fmt.Printf("Bot at %d,%d,%d range %d has %d bots in range\n", nanoBotLocation[i].xCoord, nanoBotLocation[i].yCoord, nanoBotLocation[i].zCoord, nanoBotLocation[i].signalRange, botsInRange)
			}
		}
		if botsInRange > maxBotsInRange {
			maxArrayPos = i
			fmt.Printf("High Number found at pos: %d Bots: %d\n", maxArrayPos, botsInRange)
			maxBotsInRange = botsInRange
		}
	}

	fmt.Println("Max bots:", nanoBotLocation[maxArrayPos], maxBotsInRange)

	// Need to do some thinking here. How can we work out what is close?
	// Basically we're looking for the intersection of the most number of nanobots
	// If you take the coords of a nanobot and the signal range it desribes a "circular" area of space
	// Where the most intersections between nanobot circles happens is our answer
	//
	// So how do I build a model of all the circles and work out what intersects what?
	// Perhaps build a circle for point [i] then compare against all the other points' circles, counting intersections (and where they are)
	// at the end of it I should have an idea of roughly where the most intersections happened, then I can walk through each of the
	// points in the intersection area and test them all specifically against the nanobots

	return 0
}

// func: nanoBotControl-A
// 
func nanoBotControlA(fileName string, debug bool) (int) {
	var nanoBotLocation []nanoCoordsStruct
	var mostPowerfulPos, mostPowerfulRange int							// Array reference to the most powerful nanoBot
	var botsInRange int = 0
	var tempDistance int

	fileContents, _ := readLines(fileName)
	nanoBotLocation = processInputFile(fileContents)

	if debug {
		fmt.Println(nanoBotLocation)
	}

	mostPowerfulPos = findMostPowerful(nanoBotLocation)
	mostPowerfulRange = nanoBotLocation[mostPowerfulPos].signalRange

	if debug {
		fmt.Println("Most powerful nanobot is at:", mostPowerfulPos)
	}

	for i := 0; i < len(nanoBotLocation); i++ {
		tempDistance = manhattanDistance(nanoBotLocation[mostPowerfulPos], nanoBotLocation[i])
		if tempDistance <= mostPowerfulRange {
			botsInRange++
		}
		if debug {
			fmt.Printf("Nano %d,%d,%d is %d distance away\n", nanoBotLocation[i].xCoord, nanoBotLocation[i].yCoord, nanoBotLocation[i].zCoord, tempDistance)
		}
	}
	
	if debug {
		fmt.Println("Bots in range:", botsInRange)
	}

	return botsInRange
}

// Main routine
func main() {
	var debug bool

	fileNamePtr := flag.String("file", "input1.txt", "A filename containing input strings")
	flag.BoolVar(&debug, "debug", false, "turns print debugging on")
	execPartPtr := flag.String("part", "a", "Which part of day18 do you want to calc (a or b)")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Number of nanobots in range:", nanoBotControlA(*fileNamePtr, debug))
	case "b":
		fmt.Println("Part b - Distance from best position to 0,0,0:", nanoBotControlB(*fileNamePtr, debug))
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}