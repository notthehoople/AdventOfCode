package main

import (
	"fmt"
	"math"
)

func getAngle(xPos int, yPos int, xxPos int, yyPos int) float64 {
	var dx float64
	var dy float64
	var inDegrees float64

	dx = float64(xPos - xxPos)
	dy = float64(yPos - yyPos)

	inRads := math.Atan2(dy, dx)
	inDegrees = ((inRads / math.Pi) * 180) + 180

	return inDegrees
}

// func: manhattanDistance
// Difference between 2 3D points using Manhattan distance calc
// Returns the distance as an int
func manhattanDistance2D(xCoord1 int, yCoord1 int, xCoord2 int, yCoord2 int) int {
	var distance float64 = 0

	distance = math.Abs(float64(xCoord1-xCoord2)) + math.Abs(float64(yCoord1-yCoord2))

	return int(distance)
}

// Prints the map list
func printMap(tempMap []string) {
	for i := 0; i < len(tempMap); i++ {
		fmt.Printf("%s\n", tempMap[i])
	}
}

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
