package main

import (
	"fmt"
	"flag"
)

// Used to build a work list of water sources to be followed
type spaceCoordsStruct struct {	
	xCoord			int
	yCoord			int
}

// func: spaceControlB
//
func spaceControlB(fileName string, totalDistance int, debug bool) (int) {
	var pointLocation []spaceCoordsStruct
	var minX, maxX, minY, maxY int = 0, 0, 0, 0
	var tempDistance int = 0
	var sumOfDistance int = 0
	var maxAreaSize int = 0
	var printChars byte = 'A'

	fileContents, _ := readLines(fileName)
	pointLocation = processInputFile(fileContents)

	if debug {
		fmt.Println(pointLocation)
	}

	minX, maxX, minY, maxY = scanInputForMaxMins(pointLocation, 2)

	if debug {
		fmt.Printf("MinX: %d MaxX: %d\n", minX, maxX)
		fmt.Printf("MinY: %d MaxY: %d\n", minY, maxY)
	}

	areaMap := make([][]byte, maxY)
	for i := 0; i < len(areaMap); i++ {
		areaMap[i] = make([]byte, maxX)
	}

	for i := 0; i < len(pointLocation); i++ {
		areaMap[pointLocation[i].yCoord][pointLocation[i].xCoord] = printChars + byte(i)
	}

	if debug {
		print2DSlice(areaMap)
	}

	// Loop through the entire areaMap
	// Work out the manhattanDistance from our current point to each of the points in pointLocation
	// Add those distances together. When we reach the end of pointLocation, if the sum of distances is <= totalDistance
	// then we have our answer


	for y := 0; y < len(areaMap); y++ {
		for x := 0; x < len(areaMap[y]); x++ {
			sumOfDistance = 0
			for i := 0; i < len(pointLocation); i++ {
				tempDistance = manhattanDistance2D(x, y, pointLocation[i])
				sumOfDistance += tempDistance
			}

			if sumOfDistance < totalDistance {
				areaMap[y][x] = '#'
				maxAreaSize++
			}
		}
	}

	if debug {
		print2DSlice(areaMap)
	}

	return maxAreaSize
}

// func: spaceControlA
// 
func spaceControlA(fileName string, debug bool) (int) {
	var pointLocation []spaceCoordsStruct
	var minX, maxX, minY, maxY int = 0, 0, 0, 0
	var printChars byte = 65
	var minDistance, tempDistance int = 0, 0
	var charToPlace byte
	var maxAreaFinite int = 0

	fileContents, _ := readLines(fileName)
	pointLocation = processInputFile(fileContents)

	if debug {
		fmt.Println(pointLocation)
	}

	minX, maxX, minY, maxY = scanInputForMaxMins(pointLocation, 2)

	if debug {
		fmt.Printf("MinX: %d MaxX: %d\n", minX, maxX)
		fmt.Printf("MinY: %d MaxY: %d\n", minY, maxY)
	}

	areaMap := make([][]byte, maxY)
	for i := 0; i < len(areaMap); i++ {
		areaMap[i] = make([]byte, maxX)
	}

	for i := 0; i < len(pointLocation); i++ {
		areaMap[pointLocation[i].yCoord][pointLocation[i].xCoord] = printChars + byte(i)
	}

	if debug {
		print2DSlice(areaMap)
	}

	// Loop through the entire areaMap
	// Work out the manhattanDistance from our current point to each of the points in pointLocation
	// Whichever has the lowest disatance is marked as belonging to that point
	// If 2 or more have equal distance then mark with a '.'

	countingMap := make(map[byte]int)
	infiniteMap := make(map[byte]bool)

	for y := 0; y < len(areaMap); y++ {
		for x := 0; x < len(areaMap[y]); x++ {
			minDistance = 10000
			for i := 0; i < len(pointLocation); i++ {
				tempDistance = manhattanDistance2D(x, y, pointLocation[i])
				//if debug {
				//	fmt.Printf("At point x: %d y: %d distance from x: %d y: %d is %d\n", x, y, pointLocation[i].xCoord, pointLocation[i].yCoord, tempDistance)
				//}
				if tempDistance < minDistance {
					minDistance = tempDistance
					charToPlace = printChars + byte(i)
				} else {
					if tempDistance == minDistance {
						charToPlace = '.'
					}
				}
			}
			if x == 0 || x == len(areaMap[y]) - 1 {
				infiniteMap[charToPlace] = true
			}
			if y == 0 || y == len(areaMap) - 1 {
				infiniteMap[charToPlace] = true
			}
			areaMap[y][x] = charToPlace
			countingMap[charToPlace]++
		}

	}

	if debug {
		print2DSlice(areaMap)
	}

	// Now need to count the maximum non-infinite area.
	// Use a map to hold the results, keyed on the character that's used at the location in the areaMap
	// Basically need to discount anything that's round the edges of the area

	for mapItem, mapValue := range countingMap {
		if debug {
			fmt.Printf("countingMap %c number %d\n", mapItem, mapValue)
			fmt.Printf("infiniteMap %t\n", infiniteMap[mapItem])
		}
		if !infiniteMap[mapItem] && (mapValue > maxAreaFinite) {
			maxAreaFinite = mapValue
		}
	}

	return maxAreaFinite
}

// Main routine
func main() {
	var debug bool
	var totalDistance int

	fileNamePtr := flag.String("file", "input1.txt", "A filename containing input strings")
	flag.BoolVar(&debug, "debug", false, "turns print debugging on")
	flag.IntVar(&totalDistance, "distance", 10000, "sum of distance to all points")
	execPartPtr := flag.String("part", "a", "Which part of day06 do you want to calc (a or b)")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Largest non-infinite area:", spaceControlA(*fileNamePtr, debug))
	case "b":
		fmt.Printf("Part b - Region size with distance less than %d is: %d\n", totalDistance, spaceControlB(*fileNamePtr, totalDistance, debug))
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}