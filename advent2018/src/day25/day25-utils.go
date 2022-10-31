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

// func processInputFile
// Returns a pointStruct array with the processed data in it
// Data is provided in the file as follows:
//   pos=<1,3,1>, r=4 e.g.
//   pos=<X,Y,Z>, r=R
func processInputFile(fileContents []string) ([]spaceCoordsStruct) {
	var tempPointLocation []spaceCoordsStruct
	var x, y, z, t int

	for i := 0; i < len(fileContents); i++ {
		fmt.Sscanf(fileContents[i], "%d,%d,%d,%d\n", &x, &y, &z, &t)
		
		tempPointLocation = append(tempPointLocation, spaceCoordsStruct{xCoord: x, yCoord: y, zCoord: z, tCoord: t, constellation: 0})
	}

	return tempPointLocation
}

// func: manhattanDistance
// Difference between 2 3D points using Manhattan distance calc
// Returns the distance as an int
func manhattanDistance4D(firstPoint spaceCoordsStruct, secondPoint spaceCoordsStruct) (int) {
	var distance float64 = 0

	distance = math.Abs(float64(firstPoint.xCoord - secondPoint.xCoord)) + 
				math.Abs(float64(firstPoint.yCoord - secondPoint.yCoord)) + 
				math.Abs(float64(firstPoint.zCoord - secondPoint.zCoord)) +
				math.Abs(float64(firstPoint.tCoord - secondPoint.tCoord))

	return int(distance)
}

