package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

type Coords struct {
	x int
	y int
}

// Prints the map list
func printMapList(tempMapList map[Coords]bool) {
	for key := range tempMapList {
		fmt.Printf("x: %d y: %d\n", key.x, key.y)
	}
}

// func: manhattanDistance
// Difference between 2 3D points using Manhattan distance calc
// Returns the distance as an int
func manhattanDistance2D(xCoord1 int, yCoord1 int, xCoord2 int, yCoord2 int) int {
	var distance float64 = 0

	distance = math.Abs(float64(xCoord1-xCoord2)) + math.Abs(float64(yCoord1-yCoord2))

	return int(distance)
}

// Range over the CoordsList of the first line. If a coord exists that's ALSO in the list for Line 2, we have an intersection
func scanForClosestCross(CoordsList1 map[Coords]bool, CoordsList2 map[Coords]bool, debug bool) int {
	var shortestDistance, currentDistance int

	shortestDistance = 50000

	for key := range CoordsList1 {
		if CoordsList2[key] {
			currentDistance = manhattanDistance2D(key.x, key.y, 0, 0)
			if currentDistance < shortestDistance {
				shortestDistance = currentDistance
			}

			if debug {
				fmt.Printf("Found an intersection at x: %d y: %d with distance: %d\n", key.x, key.y, currentDistance)
			}
		}
	}

	return shortestDistance
}

// Processes the draw instructions and adds every coord that a line goes through to the CoordsList map
func drawInstruction(lineRead []string, CoordsList map[Coords]bool, debug bool) {
	var magnitude, currX, currY int

	currX = 0
	currY = 0

	for _, currentInstruction := range lineRead {
		magnitude, _ = strconv.Atoi(currentInstruction[1:])

		switch currentInstruction[0] {
		case 'R':
			for i := currX + 1; i < currX+1+magnitude; i++ {
				CoordsList[Coords{i, currY}] = true
			}
			currX += magnitude
		case 'L':
			for i := currX - 1; i > currX-1-magnitude; i-- {
				CoordsList[Coords{i, currY}] = true
			}
			currX -= magnitude
		case 'U':
			for i := currY + 1; i < currY+1+magnitude; i++ {
				CoordsList[Coords{currX, i}] = true
			}
			currY += magnitude
		case 'D':
			for i := currY - 1; i > currY-1-magnitude; i-- {
				CoordsList[Coords{currX, i}] = true
			}
			currY -= magnitude
		}
	}
}

// Returns: Manhattan Distance of closest intersection to start
func closestIntersection(filename string, debug bool, part byte) int {
	var CoordsList1 map[Coords]bool
	var CoordsList2 map[Coords]bool

	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// 2 lines to process
	lineRead1, err1 := reader.Read()
	if err1 == io.EOF {
		if debug {
			fmt.Println("End of file")
		}
		return 0
	}

	lineRead2, err2 := reader.Read()
	if err2 == io.EOF {
		if debug {
			fmt.Println("End of file")
		}
		return 0
	}
	csvFile.Close()

	// Let's start processing things using a map for each line, and an ordered list attached to each for part 2

	// if we have a map with key of string and a value of int then we can record:
	//    co-ords we've reached in the key e.g. x=100,y=50
	//    the number of

	CoordsList1 = make(map[Coords]bool)
	CoordsList2 = make(map[Coords]bool)

	// Build a map list of all the coords that the lines go through
	drawInstruction(lineRead1, CoordsList1, debug)
	drawInstruction(lineRead2, CoordsList2, debug)

	if debug {
		printMapList(CoordsList1)
	}

	return scanForClosestCross(CoordsList1, CoordsList2, debug)
}

// Main routine
func main() {
	var debug bool

	filenamePtr := flag.String("file", "input.txt", "Filename containing the program to run")
	execPartPtr := flag.String("part", "a", "Which part of day03 do you want to calc (a or b)")
	flag.BoolVar(&debug, "debug", false, "Turn debug on")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Closest Intersection:", closestIntersection(*filenamePtr, debug, 'a'))
	case "b":
		fmt.Println("Part b - Not implemented yet")

	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
