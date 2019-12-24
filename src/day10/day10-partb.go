package main

import (
	"fmt"
	"sort"
)

// We are given our starting position from part A
// We need to destroy each asteroid in turn, keeping a count
// Start at vertical, then rotate round increase in angle
// Remove the first visible asteroid each time
// Print out details for the 200th removed asteroid
//
// func: destroyVisibleAsteroids
//
// startXPos, startYPos - the position of the monitoring station asteroid
func destroyVisibleAsteroids(filename string, startXPos int, startYPos int, debug bool) int {
	var angle float64
	var ok bool
	var keepLooping bool = true
	var killCount int
	var keepX, keepY int // Need to keep the coords of the 200th asteroid killed

	baseSpaceMap, _ := readLines(filename)

	//   For an asteroid, create a new map and loop through it
	tempSpaceMap := make([][]byte, len(baseSpaceMap))
	for i := 0; i < len(baseSpaceMap); i++ {
		tempSpaceMap[i] = make([]byte, len(baseSpaceMap[0]))
	}
	readInitialState(baseSpaceMap, tempSpaceMap)

	// Highlight the asteroid we're currently looking at so we don't count it
	tempSpaceMap[startYPos][startXPos] = 'P'

	// Go through the whole map
	//   build the angleMap, keeping the CLOSEST to the monitoring asteroid this time
	//   once the map is built, loop through the map
	//     sort the map into a list
	//     destroy the asteroids in order:
	//       270 degrees is straight up
	//       360/0 is horiztonal EAST
	//       90 is straight down

	// Once we've built the "kill" map we need to count through the thing and make sure there's > 0 'K' otherwise we've finished
	// Need a keep a COUNT of the kills done so we can output 200th
	// Need to loop while 'K' count > 0

	// Use a map of angles to keep note of which asteroids block other ones

	for keepLooping {
		angleMap := make(map[float64]coords)

		for tempY := 0; tempY < len(tempSpaceMap); tempY++ {
			for tempX := 0; tempX < len(tempSpaceMap[tempY]); tempX++ {
				if tempSpaceMap[tempY][tempX] == '#' {

					// Work out the angle from our starting point to this asteroid
					angle = getAngle(startXPos, startYPos, tempX, tempY)

					_, ok = angleMap[angle]
					if ok {
						// We have found another asteroid at this map
						//		Get the co-ords of the existing asteroid. Work out manhattan distance between start point and existing asteroid
						//		Work out manhattan distance between start point and current asteroid
						//		Which ever is shortest stays and other is set to '.' as is blocked from view

						existingCoords := angleMap[angle]
						existingDistance := manhattanDistance2D(startXPos, startYPos, existingCoords.x, existingCoords.y)
						currentDistance := manhattanDistance2D(startXPos, startYPos, tempX, tempY)

						if existingDistance > currentDistance {
							// Potential for killing this one
							tempSpaceMap[existingCoords.y][existingCoords.x] = '#'
							tempSpaceMap[tempY][tempX] = 'K'
							angleMap[angle] = coords{tempX, tempY}
						}

					} else {
						// First time this angle has been seen so record it in the angleMap

						angleMap[angle] = coords{tempX, tempY}
						tempSpaceMap[tempY][tempX] = 'K'
					}
				}
			}
		}

		// sort the map of angles and points into a list with lowest angles first
		keys := make([]float64, 0, len(angleMap))
		for k := range angleMap {
			keys = append(keys, k)
		}
		sort.Float64s(keys)

		// destroy the asteroids in order, starting at 0 degrees (vertically up)
		for _, k := range keys {
			killCount++
			tempSpaceMap[angleMap[k].y][angleMap[k].x] = '.'
			if killCount == 200 {
				keepX = angleMap[k].x
				keepY = angleMap[k].y
			}
			if debug {
				fmt.Println("Kill:", killCount, "Key:", k, "Value:", angleMap[k])
			}
		}

		if debug {
			fmt.Println("After Killing")
			print2DSlice(tempSpaceMap)
			fmt.Println("=================")
		}

		// Now let's count the number of asteroids left
		remainingAsteroids := 0
		for y := 0; y < len(tempSpaceMap); y++ {
			for x := 0; x < len(tempSpaceMap[y]); x++ {
				if tempSpaceMap[y][x] == '#' {
					remainingAsteroids++
				}
			}
		}
		if remainingAsteroids == 0 {
			keepLooping = false
		}
	}

	return keepX*100 + keepY
}
