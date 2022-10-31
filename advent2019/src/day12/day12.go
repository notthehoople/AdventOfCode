package main

import (
	"flag"
	"fmt"
)

type moonType struct {
	posX int
	posY int
	posZ int
	velX int
	velY int
	velZ int
}

func printMoonObjects(moonObjects []moonType, steps int) {
	fmt.Printf("After %d steps:\n", steps)
	for i := 0; i < len(moonObjects); i++ {
		fmt.Printf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>\n", moonObjects[i].posX, moonObjects[i].posY, moonObjects[i].posZ,
			moonObjects[i].velX, moonObjects[i].velY, moonObjects[i].velZ)
	}
}

// Process input lines and store in the moonObjects passed array
// e.g. <x=-1, y=0, z=2>
func processMoonFile(moonObjects []moonType, baseMoonPositions []string, debug bool) {
	for i := 0; i < len(baseMoonPositions); i++ {
		fmt.Sscanf(baseMoonPositions[i], "<x=%d, y=%d, z=%d>", &moonObjects[i].posX, &moonObjects[i].posY, &moonObjects[i].posZ)
	}
}

// Add the corresponding velocity to each of the position axis
func applyVelocity(moonObjects []moonType, debug bool) {
	for i := 0; i < len(moonObjects); i++ {
		moonObjects[i].posX += moonObjects[i].velX
		moonObjects[i].posY += moonObjects[i].velY
		moonObjects[i].posZ += moonObjects[i].velZ
	}
}

func applyGravity(moonOne moonType, moonTwo moonType) moonType {
	var tempMoon moonType
	//  consider every pair of moons. On each axis (x, y, z) the velocity changes by +1 or -1 to pull them together
	//  e.g. the moon with the highest x position gets -1, and the lowest gets +1 (unless it's negative then amounts are reversed)
	//  if the positions on an exis are the same, the velocity doesn't change for that pair of moons

	tempMoon = moonOne

	if moonOne.posX == moonTwo.posX {
		tempMoon.velX += 0
	} else if moonOne.posX > moonTwo.posX {
		tempMoon.velX--
	} else {
		tempMoon.velX++
	}

	if moonOne.posY == moonTwo.posY {
		tempMoon.velY += 0
	} else if moonOne.posY > moonTwo.posY {
		tempMoon.velY--
	} else {
		tempMoon.velY++
	}

	if moonOne.posZ == moonTwo.posZ {
		tempMoon.velZ += 0
	} else if moonOne.posZ > moonTwo.posZ {
		tempMoon.velZ--
	} else {
		tempMoon.velZ++
	}

	return tempMoon
}

// func
// Returns:
func findSystemEnergy(filename string, steps int, debug bool, part byte) int {
	var potentialEnergy, kineticEnergy, totalEnergy int

	// Read the map from the file given
	baseMoonPositions, _ := readLines(filename)
	if debug {
		printMap(baseMoonPositions)
	}

	// Create an array of moons to track
	moonObjects := make([]moonType, len(baseMoonPositions))
	processMoonFile(moonObjects, baseMoonPositions, debug)

	// Now simulate the moons in time steps using the following rules
	for minuteCount := 1; minuteCount <= steps; minuteCount++ {
		//   - first update the velocity of every moon by applying gravity

		for i := 0; i < len(moonObjects); i++ {
			for j := 0; j < len(moonObjects); j++ {
				if i != j {
					moonObjects[i] = applyGravity(moonObjects[i], moonObjects[j])
				}
			}
		}

		//   - then update the position of every moon by applying velocity
		applyVelocity(moonObjects, debug)

		if debug {
			printMoonObjects(moonObjects, minuteCount)
		}
	}

	for i := 0; i < len(moonObjects); i++ {
		potentialEnergy = absInt(moonObjects[i].posX) + absInt(moonObjects[i].posY) + absInt(moonObjects[i].posZ)
		kineticEnergy = absInt(moonObjects[i].velX) + absInt(moonObjects[i].velY) + absInt(moonObjects[i].velZ)
		totalEnergy += potentialEnergy * kineticEnergy
	}

	return totalEnergy
}

// Main routine
func main() {
	var debug bool
	var steps int

	//filenamePtr := flag.String("file", "input.txt", "Filename containing the program to run")
	execPartPtr := flag.String("part", "a", "Which part of day12 do you want to calc (a or b)")
	flag.BoolVar(&debug, "debug", false, "Turn debug on")
	flag.IntVar(&steps, "steps", 10, "Number of steps to run")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Total Energy (test 1):", findSystemEnergy("test-a1.txt", 10, debug, 'a'))
		fmt.Println("Part a - Total Energy (test 2):", findSystemEnergy("test-a2.txt", 100, debug, 'a'))
		fmt.Println("Part a - Total Energy (full):", findSystemEnergy("input.txt", 1000, debug, 'a'))
	case "b":
		fmt.Println("Part b - First repeat (test 1):", findFirstRepeat("test-a1.txt", debug, 'b'))
		fmt.Println("Part b - First repeat (test 2):", findFirstRepeat("test-a2.txt", debug, 'b'))
		fmt.Println("Part b - First repeat (test 2):", findFirstRepeat("input.txt", debug, 'b'))

	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
