package main

import (
	"fmt"
)

//type moonType struct {
//	posX int
//	posY int
//	posZ int
//	velX int
//	velY int
//	velZ int
//}

func findFirstRepeat(filename string, debug bool, part byte) int {
	var potentialEnergy, kineticEnergy, totalEnergy int
	var minuteCount int
	var energyMap map[int]int
	var keepLooping, ok bool

	// Read the map from the file given
	baseMoonPositions, _ := readLines(filename)
	if debug {
		printMap(baseMoonPositions)
	}

	energyMap = make(map[int]int)

	// Create an array of moons to track
	moonObjects := make([]moonType, len(baseMoonPositions))
	processMoonFile(moonObjects, baseMoonPositions, debug)

	// Now simulate the moons in time steps using the following rules
	keepLooping = true
	for keepLooping {
		//   - first update the velocity of every moon by applying gravity

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

		for i := 0; i < len(moonObjects); i++ {
			for j := 0; j < len(moonObjects); j++ {
				if i != j {
					moonObjects[i] = applyGravity(moonObjects[i], moonObjects[j])
				}
			}
		}

		minuteCount++

		//   - then update the position of every moon by applying velocity
		applyVelocity(moonObjects, debug)

		if debug {
			printMoonObjects(moonObjects, minuteCount)
		}

		totalEnergy = 0
		for i := 0; i < len(moonObjects); i++ {
			//potentialEnergy = absInt(moonObjects[i].posX) + absInt(moonObjects[i].posY) + absInt(moonObjects[i].posZ)
			//kineticEnergy = absInt(moonObjects[i].velX) + absInt(moonObjects[i].velY) + absInt(moonObjects[i].velZ)
			potentialEnergy = moonObjects[i].posX + (moonObjects[i].posY + 100) + (moonObjects[i].posZ + 10000)
			kineticEnergy = (moonObjects[i].velX + (moonObjects[i].velY + 100) + (moonObjects[i].velZ + 10000)) + 10000000
			totalEnergy += potentialEnergy * kineticEnergy
		}

		//fmt.Printf("Minute; %d Kinetic Energy: %d\n", minuteCount, totalEnergy)
		_, ok = energyMap[totalEnergy]
		if ok {
			fmt.Println("Found the duplicate at:", energyMap[totalEnergy])
			return minuteCount
		} else {
			energyMap[totalEnergy] = minuteCount
		}
	}

	return totalEnergy
}
