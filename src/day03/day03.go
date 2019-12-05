package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

func printMapArray(tempMapArray [][]byte) {
	for x := 20; x >= 0; x-- {
		for y := 20; y >= 0; y-- {
			fmt.Printf("%c", tempMapArray[x][y])
		}
		fmt.Printf("\n")
	}
}

func createMapArray(debug bool) [][]byte {

	tempMapArray := make([][]byte, 2000)
	for i := 0; i < 2000; i++ {
		tempMapArray[i] = make([]byte, 2000)
	}

	// fill the array with stuff
	for x := 0; x < 20; x++ {
		for y := 0; y < 20; y++ {
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

func drawInstruction(lineRead1 []string, mapArray [][]byte, marker byte, startX int, startY int) {
	mapArray[startX][startY] = 'o'
	for _, currentInstruction := range lineRead1 {
		switch currentInstruction[0] {
		case 'R':
			fmt.Println("Right:", currentInstruction[1:])
		case 'L':
			fmt.Println("Left", currentInstruction[1:])
		case 'U':
			fmt.Println("Up", currentInstruction[1:])
		case 'D':
			fmt.Println("Down", currentInstruction[1:])
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

	// Loop through each instruction of line 1
	//    Draw "1" into each location as per the instruction

	drawInstruction(lineRead1, mapArray, '1', 1, 1)
	drawInstruction(lineRead2, mapArray, '2', 1, 1)

	// Loop through each instruction of line 2
	//    If location in instruction is a "1", draw an "X"
	//    else draw a "2"
	// Scan through the array looking for "X"
	//    Work out manhattan distance to start
	//    If smallest seen so far, record it
	// Output smallest seen

	if debug {
		printMapArray(mapArray)
	}

	return 0
}

// Grab Manhattan Distance function from day06_utils.go

// Main routine
func main() {
	var debug bool

	filenamePtr := flag.String("file", "input.txt", "Filename containing the program to run")
	execPartPtr := flag.String("part", "a", "Which part of day18 do you want to calc (a or b)")
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
