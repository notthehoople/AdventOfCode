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

// func scanInputForMaxMins
// Walks through the coord data array and works out the minimum and maximum values of X, Y and Z
func scanInputForMaxMins(coordData []nanoCoordsStruct) (minX int, maxX int, minY int, maxY int, minZ int, maxZ int) {

	minX = coordData[0].xCoord
	maxX = minX
	minY = coordData[0].yCoord
	maxY = minY
	minZ = coordData[0].zCoord
	maxZ = minZ

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
		if coordData[i].zCoord < minZ {
			minZ = coordData[i].zCoord
		}
		if coordData[i].zCoord > maxZ {
			maxZ = coordData[i].zCoord
		}
	}

	return minX, maxX, minY, maxY, minZ, maxZ
}

// func processInputFile
// Returns a pointStruct array with the processed data in it
// Data is provided in the file as follows:
//   pos=<1,3,1>, r=4 e.g.
//   pos=<X,Y,Z>, r=R
func processInputFile(fileContents []string) ([]nanoCoordsStruct) {
	var tempBotLocation []nanoCoordsStruct
	var x, y, z, r int

	for i := 0; i < len(fileContents); i++ {
		fmt.Sscanf(fileContents[i], "pos=<%d,%d,%d>, r=%d\n", &x, &y, &z, &r)
		
		tempBotLocation = append(tempBotLocation, nanoCoordsStruct{xCoord: x, yCoord: y, zCoord: z, signalRange: r})
	}

	return tempBotLocation
}

// func processWaterFlow
// Handles everything needed to work out the water flow (day 17 part A)
/*func processWaterFlow(fileName string, springX int, springY int, part byte) int {
	var minX, maxX, minY, maxY, gridSizeX, gridSizeY, workX, workY, yCountStart int
	var coordData []pointStruct
	var workList []workListCoords
	var letsLoopThis bool
	var maxiMins pointStruct
	var didWork bool

	// Read contents of file into a string array
	fileContents, _ := readLines(fileName)
	coordData = processInputFile(fileContents)

	minX, maxX, minY, maxY, yCountStart = scanInputForMaxMins(coordData, springX, springY)

	maxiMins.xCoordStart = 0
	maxiMins.xCoordEnd = maxX - minX
	maxiMins.yCoordStart = 0
	maxiMins.yCoordEnd = maxY - minY + 1
	maxiMins.yCountStart = yCountStart

	gridSizeX = (maxX - minX) + 1
	gridSizeY = (maxY - minY) + 2

	undergroundArea := make([][]byte, gridSizeY)
	for i := 0; i < gridSizeY; i++ {
		undergroundArea[i] = make([]byte, gridSizeX)	
	}

	workList = make([]workListCoords, 0)

	readInitialState(coordData, undergroundArea, springX, springY, minX, minY)

	workList = addWorkListItem(workList, springX - minX, springY - minY)
	letsLoopThis = true

	for letsLoopThis {
		didWork = false

		// Loop through the list of work we have. This list is a list of water sources
		for i := 0; i < len(workList); i++ {
			if !workList[i].done {
				workX = workList[i].xCoord
				workY = workList[i].yCoord

				letsLoopThis, workList = letTheWaterFlow(undergroundArea, workList, workX, workY, maxiMins)
				workList[i].done = true
				didWork = true
			}
		}
		if !didWork {
			letsLoopThis = false
		}

	}
	
	// Print final water flow
	print2DSlice(undergroundArea)

	return countWaterSquares(undergroundArea, part, maxiMins.yCountStart)
}*/

// func: findMostPowerful
// Find the most powerful nanobot in the nanoBotLocation array
// Return the array position of the most powerful nanobot
func findMostPowerful(nanoBotLocation []nanoCoordsStruct) (int) {
	var mostPowerfulPos, mostPower int = 0, 0

	for i := 0; i < len(nanoBotLocation); i++ {
		if nanoBotLocation[i].signalRange > mostPower {
			mostPower = nanoBotLocation[i].signalRange
			mostPowerfulPos = i
		}
	}

	return mostPowerfulPos
}

// func: manhattanDistance
// Difference between 2 3D points using Manhattan distance calc
// Returns the distance as an int
func manhattanDistance(firstPoint nanoCoordsStruct, secondPoint nanoCoordsStruct) (int) {
	var distance float64 = 0

	distance = math.Abs(float64(firstPoint.xCoord - secondPoint.xCoord)) + math.Abs(float64(firstPoint.yCoord - secondPoint.yCoord)) + math.Abs(float64(firstPoint.zCoord - secondPoint.zCoord))

	return int(distance)

}

