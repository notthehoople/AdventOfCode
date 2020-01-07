package main

import (
	"fmt"
	"strconv"
)

//type moonType struct {
//	posX int
//	posY int
//	posZ int
//	velX int
//	velY int
//	velZ int
//}

func findFirstRepeat(filename string, debug bool, part byte) int64 {
	//var potentialEnergy, kineticEnergy, totalEnergy int
	var minuteCount int
	var moonPositions map[string]int
	var keepLooping, ok, firstXFound, firstYFound, firstZFound bool
	var firstXminute, firstYminute, firstZminute int

	// Read the map from the file given
	baseMoonPositions, _ := readLines(filename)
	if debug {
		printMap(baseMoonPositions)
	}

	moonPositions = make(map[string]int)

	// Create an array of moons to track
	moonObjects := make([]moonType, len(baseMoonPositions))
	processMoonFile(moonObjects, baseMoonPositions, debug)

	// Now simulate the moons in time steps using the following rules
	keepLooping = true
	firstXFound = false
	firstYFound = false
	firstZFound = false
	for keepLooping {

		// ==================================================
		// Need to radically update this.
		// Basically find the repeat loop for x, then for y, then for z
		// they should loop back to the initial state at somepoint so use that
		// We're looking for initial x and vel_x being 0
		//
		// Once we've got them all we need to find the point at which the x, y and z loops intersect
		// e.g. using Lowest Common Multiply across the intervals of each of the dimensions
		//      so if x loops every 5, y loops every 6 and z loops every 3, the LCM is 30
		//      Look here for LCM calculations: https://www.calculatorsoup.com/calculators/math/lcm.php
		//==================================================

		// first update the velocity of every moon by applying gravity
		for i := 0; i < len(moonObjects); i++ {
			for j := 0; j < len(moonObjects); j++ {
				if i != j {
					moonObjects[i] = applyGravity(moonObjects[i], moonObjects[j])
				}
			}
		}

		// then update the position of every moon by applying velocity
		applyVelocity(moonObjects, debug)

		if debug {
			printMoonObjects(moonObjects, minuteCount)
		}

		// POSITION X
		positionXKey := "X"
		// now record the positions of the moons and their velocity
		for i := 0; i < len(moonObjects); i++ {
			positionXKey = positionXKey + "|" + strconv.Itoa(moonObjects[i].posX) + "|" + strconv.Itoa(moonObjects[i].velX)
		}

		// compare with previous positions for each moon's posX and velX looking for a match
		_, ok = moonPositions[positionXKey]
		if ok {
			// Found a match!
			if !firstXFound {
				firstXminute = minuteCount
				fmt.Printf("X Found %s at %d\n", positionXKey, firstXminute)
				firstXFound = true
			}
		} else {
			// Not a repeat so record what we've seen
			moonPositions[positionXKey] = minuteCount
		}

		// POSITION Y
		positionYKey := "Y"
		// now record the positions of the moons and their velocity
		for i := 0; i < len(moonObjects); i++ {
			positionYKey = positionYKey + "|" + strconv.Itoa(moonObjects[i].posY) + "|" + strconv.Itoa(moonObjects[i].velY)
		}

		// compare with previous positions for each moon's posX and velX looking for a match
		_, ok = moonPositions[positionYKey]
		if ok {
			// Found a match!
			if !firstYFound {
				firstYminute = minuteCount
				fmt.Printf("Y Found %s at %d\n", positionYKey, firstYminute)
				firstYFound = true
			}
		} else {
			// Not a repeat so record what we've seen
			moonPositions[positionYKey] = minuteCount
		}

		// POSITION Z
		positionZKey := "Z"
		// now record the positions of the moons and their velocity
		for i := 0; i < len(moonObjects); i++ {
			positionZKey = positionZKey + "|" + strconv.Itoa(moonObjects[i].posZ) + "|" + strconv.Itoa(moonObjects[i].velZ)
		}

		// compare with previous positions for each moon's posX and velX looking for a match
		_, ok = moonPositions[positionZKey]
		if ok {
			// Found a match!
			if !firstZFound {
				firstZminute = minuteCount
				fmt.Printf("Z Found %s at %d\n", positionZKey, firstZminute)
				firstZFound = true
			}
		} else {
			// Not a repeat so record what we've seen
			moonPositions[positionZKey] = minuteCount
		}

		// Have we found them all? If so let's quit the loop - we're done!
		if firstXFound && firstYFound && firstZFound {
			return (LCM(int64(firstXminute), int64(firstYminute), int64(firstZminute)))
			//keepLooping = false
		}

		minuteCount++

	}

	return 0
}
