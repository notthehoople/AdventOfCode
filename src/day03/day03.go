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

// Prints the array with 0,0 at the bottom left
func printMapArray(tempMapArray [][]byte) {
	for y := 49999; y >= 0; y-- {
		for x := 0; x < 50000; x++ {
			fmt.Printf("%c", tempMapArray[y][x])
		}
		fmt.Printf("\n")
	}
}

func createMapArray(debug bool) [][]byte {
	tempMapArray := make([][]byte, 50000)
	for i := 0; i < 50000; i++ {
		tempMapArray[i] = make([]byte, 50000)
	}

	// fill the array with stuff
	for x := 0; x < 50000; x++ {
		for y := 0; y < 50000; y++ {
			tempMapArray[y][x] = '.'
		}
	}

	return tempMapArray
}

func processCSVLine(record []string, debug bool) []string {
	tempLineProcess := make([]string, len(record))
	for i := 0; i < len(record); i++ {
		tempLineProcess[i] = record[i]
	}

	return tempLineProcess
}

// func: manhattanDistance
// Difference between 2 3D points using Manhattan distance calc
// Returns the distance as an int
func manhattanDistance2D(xCoord1 int, yCoord1 int, xCoord2 int, yCoord2 int) int {
	var distance float64 = 0

	distance = math.Abs(float64(xCoord1-xCoord2)) + math.Abs(float64(yCoord1-yCoord2))

	return int(distance)
}

func scanForClosestCross(mapArray [][]byte, startX int, startY int) int {
	var shortestDistance, currentDistance int

	shortestDistance = 50000

	for y := 0; y < 50000; y++ {
		for x := 0; x < 50000; x++ {
			if mapArray[y][x] == 'X' {
				// Found a cross point
				currentDistance = manhattanDistance2D(x, y, startX, startY)
				if currentDistance < shortestDistance {
					shortestDistance = currentDistance
				}
				fmt.Printf("Manhattan Distance of x:%d y:%d is: %d\n", x, y, currentDistance)
			}
		}
	}

	return shortestDistance
}

func drawInstruction(lineRead1 []string, mapArray [][]byte, marker byte, startX int, startY int) {
	var magnitude, currX, currY int

	mapArray[startX][startY] = 'o'
	currX = startX
	currY = startY

	for _, currentInstruction := range lineRead1 {
		magnitude, _ = strconv.Atoi(currentInstruction[1:])

		switch currentInstruction[0] {
		case 'R':
			fmt.Println("Right:", magnitude)
			for i := currX + 1; i < currX+1+magnitude; i++ {
				if mapArray[currY][i] == '.' {
					mapArray[currY][i] = marker
				} else {
					mapArray[currY][i] = 'X'
				}
			}
			currX += magnitude
		case 'L':
			fmt.Println("Left", magnitude)
			for i := currX - 1; i > currX-1-magnitude; i-- {
				if mapArray[currY][i] == '.' {
					mapArray[currY][i] = marker
				} else {
					mapArray[currY][i] = 'X'
				}
			}
			currX -= magnitude
		case 'U':
			fmt.Println("Up", magnitude)
			for i := currY + 1; i < currY+1+magnitude; i++ {
				if mapArray[i][currX] == '.' {
					mapArray[i][currX] = marker
				} else {
					mapArray[i][currX] = 'X'
				}
			}
			currY += magnitude
		case 'D':
			fmt.Println("Down", magnitude)
			for i := currY - 1; i > currY-1-magnitude; i-- {
				if mapArray[i][currX] == '.' {
					mapArray[i][currX] = marker
				} else {
					mapArray[i][currX] = 'X'
				}
			}
			currY -= magnitude
		}
	}
}

// Returns: Manhattan Distance of closest intersection to start
func closestIntersection(filename string, debug bool, part byte) int {
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
	firstLine := processCSVLine(lineRead1, debug)
	fmt.Println("First line:", firstLine)

	lineRead2, err2 := reader.Read()
	if err2 == io.EOF {
		if debug {
			fmt.Println("End of file")
		}
		return 0
	}
	secondLine := processCSVLine(lineRead2, debug)
	fmt.Println("Second line:", secondLine)

	csvFile.Close()

	mapArray := createMapArray(debug)

	// Draw in the mapArray based on the instructions read from the file
	drawInstruction(lineRead1, mapArray, '1', 25000, 25000)
	drawInstruction(lineRead2, mapArray, '2', 25000, 25000)

	if debug {
		printMapArray(mapArray)
	}

	return scanForClosestCross(mapArray, 25000, 25000)
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
