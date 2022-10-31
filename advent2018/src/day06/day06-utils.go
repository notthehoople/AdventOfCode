package main

import (
	"fmt"
	"bufio"
	"os"
	"math"
)

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

func print2DSlice(tempSlice [][]byte) {
	fmt.Println("==========================")
	for i := 0; i < len(tempSlice); i++ {
		for j := 0; j < len(tempSlice[i]); j++ {
			if tempSlice[i][j] == 0 {
				fmt.Printf(". ")
			} else {
				fmt.Printf("%c ", tempSlice[i][j])
			}
		}
		fmt.Printf("\n")
	}
}

// func scanInputForMaxMins
// Walks through the coord data array and works out the minimum and maximum values of X, Y and Z
// extraBit - the amount to subtract/add to the maximums
func scanInputForMaxMins(coordData []spaceCoordsStruct, extraBit int) (minX int, maxX int, minY int, maxY int) {

	minX = coordData[0].xCoord
	maxX = minX
	minY = coordData[0].yCoord
	maxY = minY
	
	for i := 0; i < len(coordData); i++ {
		if coordData[i].xCoord < minX {
			minX = coordData[i].xCoord
		}
		if coordData[i].xCoord > maxX {
			maxX = coordData[i].xCoord
		}
		if coordData[i].yCoord < minY {
			minY = coordData[i].yCoord
		}
		if coordData[i].yCoord > maxY {
			maxY = coordData[i].yCoord
		}
	}

	if minX >= extraBit {
		minX -= extraBit
	}
	maxX += extraBit

	if minY >= extraBit {
		minY -= extraBit
	}
	maxY += extraBit

	return minX, maxX, minY, maxY
}

// func processInputFile
// Returns a pointStruct array with the processed data in it
// Data is provided in the file as follows:
//   pos=<1,3,1>, r=4 e.g.
//   pos=<X,Y,Z>, r=R
func processInputFile(fileContents []string) ([]spaceCoordsStruct) {
	var tempPointLocation []spaceCoordsStruct
	var x, y int

	for i := 0; i < len(fileContents); i++ {
		fmt.Sscanf(fileContents[i], "%d, %d\n", &x, &y)
		
		tempPointLocation = append(tempPointLocation, spaceCoordsStruct{xCoord: x, yCoord: y})
	}

	return tempPointLocation
}

// func: manhattanDistance
// Difference between 2 3D points using Manhattan distance calc
// Returns the distance as an int
func manhattanDistance2D(xCoord int, yCoord int, secondPoint spaceCoordsStruct) (int) {
	var distance float64 = 0

	distance = math.Abs(float64(xCoord - secondPoint.xCoord)) + 
				math.Abs(float64(yCoord - secondPoint.yCoord))

	return int(distance)
}

