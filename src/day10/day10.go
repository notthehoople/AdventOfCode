package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// func: readInitialState
// takes an array of strings and breaks it into a 2D array of bytes
func readInitialState(tempString []string, tempSlice [][]byte) {
	for i := 0; i < len(tempString); i++ {
		for j := 0; j < len(tempString[i]); j++ {
			tempSlice[i][j] = tempString[i][j]
		}
	}
}

func print2DSlice(tempSlice [][]byte) {
	for i := 0; i < len(tempSlice); i++ {
		for j := 0; j < len(tempSlice[i]); j++ {
			fmt.Printf("%c", tempSlice[i][j])
		}
		fmt.Printf("\n")
	}
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

// Prints the map list
func printMap(tempMap []string) {
	for i := 0; i < len(tempMap); i++ {
		fmt.Printf("%s\n", tempMap[i])
	}
}

func countVisibleAsteroids(baseSpaceMap []string, xPos int, yPos int) int {
	var firstFound bool
	//   For an asteroid, create a new map and loop through it
	tempSpaceMap := make([][]byte, len(baseSpaceMap))
	for i := 0; i < len(baseSpaceMap); i++ {
		tempSpaceMap[i] = make([]byte, len(baseSpaceMap[0]))
	}
	readInitialState(baseSpaceMap, tempSpaceMap)

	tempSpaceMap[yPos][xPos] = 'P'

	fmt.Printf("Before processing X:%d y:%d\n", xPos, yPos)
	print2DSlice(tempSpaceMap)
	// Highly the asteroid we're currently looking at so we don't count it

	// For same X left, look for first asteroid. Then discount any asteroids beyond that one on X
	firstFound = false
	for x := xPos - 1; x >= 0; x-- {
		if x >= 0 {
			if tempSpaceMap[yPos][x] == '#' {
				if !firstFound {
					firstFound = true
				} else {
					tempSpaceMap[yPos][x] = '.'
				}
			}
		}
	}

	// For same X right, look for first asteroid. Then discount any asteroids beyond that one on X
	firstFound = false
	for x := xPos + 1; x < len(tempSpaceMap[yPos]); x++ {
		if tempSpaceMap[yPos][x] == '#' {
			fmt.Println("Found an asteroid at:", x, yPos)
			if !firstFound {
				fmt.Println("This is the first one found")
				firstFound = true
			} else {
				tempSpaceMap[yPos][x] = '.'
			}
		}
	}

	// For same Y up, look for first asteroid. Then discount any asteroids beyond that one on X
	firstFound = false
	for y := yPos - 1; y >= 0; y-- {
		if y >= 0 {
			if tempSpaceMap[y][xPos] == '#' {
				if !firstFound {
					firstFound = true
				} else {
					tempSpaceMap[y][xPos] = '.'
				}
			}
		}
	}

	// For same Y down, look for first asteroid. Then discount any asteroids beyond that one on X
	firstFound = false
	for y := yPos + 1; y < len(tempSpaceMap); y++ {
		if tempSpaceMap[y][xPos] == '#' {
			if !firstFound {
				firstFound = true
			} else {
				tempSpaceMap[y][xPos] = '.'
			}
		}
	}

	// For same X and Y northwest (X-1,Y-1), look for first asteroid. Then discount any asteroids beyond that one on X
	// For same X and Y northeast (X+1,Y-1), look for first asteroid. Then discount any asteroids beyond that one on X
	// For same X and Y southwest (X-1,Y+1), look for first asteroid. Then discount any asteroids beyond that one on X
	// For same X and Y southeast (X+1,Y+1), look for first asteroid. Then discount any asteroids beyond that one on X

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
