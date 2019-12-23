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

func countVisibleAsteroids(baseSpaceMap []string, xPos int, yPos int) int {
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

	fmt.Printf("Before processing X:%d y:%d\n", xPos, yPos)
	print2DSlice(tempSpaceMap)

	// Loop through the whole Map
	// 		Work out the angle from our asteroid to others
	// 		If another asteroid is on the same angle, only keep the closest asteroid for counting (manhattan distance)
	// (Use a map to hold this, with the angle as the index and the co-ords as the value)
	// Count the visible asteroids and return

	angleMap = make(map[float64]coords)

	// Create a MAP here

	for tempY := 0; tempY < len(tempSpaceMap); tempY++ {
		for tempX := 0; tempX < len(tempSpaceMap[tempY]); tempX++ {
			if tempSpaceMap[tempY][tempX] == '#' {

				// Work out the angle from our starting point to this asteroid
				angle = getAngle(xPos, yPos, tempX, tempY)

				_, ok = angleMap[angle]
				if ok {
					// Found another asteroid at this map
					//		Get the co-ords of the existing asteroid
					//		Work out manhattan distance between start point and existing asteroid
					//		Work out manhattan distance between start point and current asteroid
					//		Which ever is shortest stays
					//		Other one is set to '.'
					//fmt.Println("Found another asteroid on angle", angle)

					existingCoords := angleMap[angle]
					existingDistance := manhattanDistance2D(xPos, yPos, existingCoords.x, existingCoords.y)
					//fmt.Printf("Existing x: %d Existing y: %d Manhattan: %d\n", existingCoords.x, existingCoords.y, existingDistance)
					currentDistance := manhattanDistance2D(xPos, yPos, tempX, tempY)

					if existingDistance <= currentDistance {
						tempSpaceMap[tempY][tempX] = '.'
					} else {
						tempSpaceMap[existingCoords.y][existingCoords.x] = '.'
						angleMap[angle] = coords{tempX, tempY}
					}

				} else {
					// First time this angle has been seen
					//fmt.Println("First time we've seen angle", angle)

					angleMap[angle] = coords{tempX, tempY}
				}
				// Work out the manhattan distance from starting point to this asteroid
				// Look in temp map to see if this angle has been used before.
				// 		If not, store co-ords against angle
				//		If yes:

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

	fmt.Println("After Processing")
	print2DSlice(tempSpaceMap)
	fmt.Println("=================")
	//   Loop through the created map for the asteroid
	//     Count the number of visible asteroids
	//   EndLoop
	return visibleAsteroids
}

// func
// Returns:
func bestMonitoringAsteroid(filename string, debug bool, part byte) int {
	var bestVisible, bestXcoord, bestYcoord int
	// Read the map from the file given
	baseSpaceMap, _ := readLines(filename)

	// Coords are X,Y. Top left is 0,0 and the space directly to the right is 1,0
	// A monitoring station can detect any asteroid to which it has direct line of sight - that is,
	// there cannot be another asteroid exactly between them

	// Loop through the whole map one space at a time
	for y := 0; y < len(baseSpaceMap); y++ {
		for x := 0; x < len(baseSpaceMap[y]); x++ {
			//   Is there an asteroid? If no, continue loop
			if baseSpaceMap[y][x] == '#' {
				// We have an asteroid so we need to deal with it
				// Looks like I need to detect (and discount) and asteroid that is on the exact same X,
				// the same Y, and the same diagonal in any direction
				// From the examples it looks like everything else can be counted

				numberVisibleAsteroids := countVisibleAsteroids(baseSpaceMap, x, y)
				//   If this is the best so far, take note of the asteroid position and the number of asteroids it can see
				if numberVisibleAsteroids > bestVisible {
					bestVisible = numberVisibleAsteroids
					bestXcoord = x
					bestYcoord = y
				}

			}
			//fmt.Printf("X:%d Y:%d Char: %c\n", x, y, baseSpaceMap[y][x])
		}
	}

	printMap(baseSpaceMap)

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

	filenamePtr := flag.String("file", "input.txt", "Filename containing the program to run")
	execPartPtr := flag.String("part", "a", "Which part of day06 do you want to calc (a or b)")
	flag.BoolVar(&debug, "debug", false, "Turn debug on")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Number of orbits:", bestMonitoringAsteroid(*filenamePtr, debug, 'a'))
	case "b":
		fmt.Println("Part b - Not implemented yet")

	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
