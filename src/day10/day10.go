package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type coords struct {
	x int
	y int
}

// Read the text file passed in by name into a array of strings
// Returns the array as the first return variable
func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func countVisibleAsteroids(baseSpaceMap []string, xPos int, yPos int, debug bool) int {
	var angle float64
	var angleMap map[float64]coords
	var ok bool

	//   For an asteroid, create a new map and loop through it
	tempSpaceMap := make([][]byte, len(baseSpaceMap))
	for i := 0; i < len(baseSpaceMap); i++ {
		tempSpaceMap[i] = make([]byte, len(baseSpaceMap[0]))
	}
	readInitialState(baseSpaceMap, tempSpaceMap)

	// Highlight the asteroid we're currently looking at so we don't count it
	tempSpaceMap[yPos][xPos] = 'P'

	if debug {
		fmt.Printf("Before processing X:%d y:%d\n", xPos, yPos)
		print2DSlice(tempSpaceMap)
	}

	// Use a map of angles to keep note of which asteroids block other ones
	angleMap = make(map[float64]coords)

	for tempY := 0; tempY < len(tempSpaceMap); tempY++ {
		for tempX := 0; tempX < len(tempSpaceMap[tempY]); tempX++ {
			if tempSpaceMap[tempY][tempX] == '#' {

				// Work out the angle from our starting point to this asteroid
				angle = getAngle(xPos, yPos, tempX, tempY)

				_, ok = angleMap[angle]
				if ok {
					// We have found another asteroid at this map
					//		Get the co-ords of the existing asteroid. Work out manhattan distance between start point and existing asteroid
					//		Work out manhattan distance between start point and current asteroid
					//		Which ever is shortest stays and other is set to '.' as is blocked from view

					existingCoords := angleMap[angle]
					existingDistance := manhattanDistance2D(xPos, yPos, existingCoords.x, existingCoords.y)
					currentDistance := manhattanDistance2D(xPos, yPos, tempX, tempY)

					if existingDistance <= currentDistance {
						tempSpaceMap[tempY][tempX] = '.'
					} else {
						tempSpaceMap[existingCoords.y][existingCoords.x] = '.'
						angleMap[angle] = coords{tempX, tempY}
					}

				} else {
					// First time this angle has been seen so record it in the angleMap

					angleMap[angle] = coords{tempX, tempY}
				}
			}
		}
	}

	// Now let's count the number of visible asteroids
	visibleAsteroids := 0
	for y := 0; y < len(tempSpaceMap); y++ {
		for x := 0; x < len(tempSpaceMap[y]); x++ {
			if tempSpaceMap[y][x] == '#' {
				visibleAsteroids++
			}
		}
	}

	if debug {
		fmt.Println("After Processing")
		print2DSlice(tempSpaceMap)
		fmt.Println("=================")
	}

	return visibleAsteroids
}

// func
// Returns:
func bestMonitoringAsteroid(filename string, debug bool, part byte) int {
	var bestVisible, bestXcoord, bestYcoord int
	// Read the map from the file given
	baseSpaceMap, _ := readLines(filename)

	// Coords are X,Y. Top left is 0,0 and the space directly to the right is 1,0
	// A monitoring station can detect any asteroid to which it has direct line of sight - that is, there cannot be another asteroid exactly between them

	// Loop through the whole map one space at a time
	for y := 0; y < len(baseSpaceMap); y++ {
		for x := 0; x < len(baseSpaceMap[y]); x++ {
			//   Is there an asteroid? If no, continue loop
			if baseSpaceMap[y][x] == '#' {
				// We have an asteroid so we need to deal with it
				numberVisibleAsteroids := countVisibleAsteroids(baseSpaceMap, x, y, debug)

				//   If this is the best so far, take note of the asteroid position and the number of asteroids it can see
				if numberVisibleAsteroids > bestVisible {
					bestVisible = numberVisibleAsteroids
					bestXcoord = x
					bestYcoord = y
				}
			}
		}
	}

	// Print out the best asteroid position and the number of asteroids it can see
	if part == 'a' {
		fmt.Printf("Best Asteroid: X:%d Y:%d numVisible:%d\n", bestXcoord, bestYcoord, bestVisible)
		return bestVisible
	}

	return 0
}

// Main routine
func main() {
	var debug bool
	var startXPos, startYPos int

	filenamePtr := flag.String("file", "input.txt", "Filename containing the program to run")
	execPartPtr := flag.String("part", "a", "Which part of day06 do you want to calc (a or b)")
	flag.IntVar(&startXPos, "x", 19, "X Coord of the monitoring asteroid")
	flag.IntVar(&startYPos, "y", 11, "Y Coord of the monitoring asteroid")
	flag.BoolVar(&debug, "debug", false, "Turn debug on")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Number of visible asteroids:", bestMonitoringAsteroid(*filenamePtr, debug, 'a'))
	case "b":
		fmt.Println("Part b - 200th destroyed asteroid", destroyVisibleAsteroids(*filenamePtr, startXPos, startYPos, debug))

	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
