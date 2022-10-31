package main

import (
	"flag"
	"fmt"
)

func catchUserInput() (string, byte, bool, int, int) {
	var debug bool
	var slopeX, slopeY int

	filenamePtr := flag.String("file", "testInput.txt", "Filename containing the program to run")
	execPartPtr := flag.String("part", "a", "Which part of the puzzle do you want to calc (a or b)")
	flag.BoolVar(&debug, "debug", false, "Turn debug on")
	flag.IntVar(&slopeX, "slopex", 3, "X component of Slope")
	flag.IntVar(&slopeY, "slopey", 1, "Y component of Slope")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		return *filenamePtr, 'a', debug, slopeX, slopeY
	case "b":
		return *filenamePtr, 'b', debug, slopeX, slopeY

	default:
		return *filenamePtr, 'z', debug, slopeX, slopeY
	}
}

func howManyTrees(filename string, part byte, debug bool, slopeX int, slopeY int) int {
	var currentXPos, currentYPos int
	var treeCount int = 0

	puzzleInput, _ := readFile(filename)

	maxX := len(puzzleInput[0])
	maxY := len(puzzleInput)

	if debug {
		fmt.Printf("Puzzle Side is X:%d Y:%d\n", maxX, maxY)
	}

	currentXPos = slopeX
	currentYPos = slopeY

	for ok := true; ok; ok = (currentYPos < maxY) {

		if puzzleInput[currentYPos][currentXPos] == '#' {
			//fmt.Println("Found a tree")
			treeCount++
		}

		if debug {
			fmt.Printf(puzzleInput[currentYPos])
		}

		currentXPos = (currentXPos + slopeX) % maxX
		currentYPos = currentYPos + slopeY
	}

	return treeCount
}

// Main routine
func main() {
	filenamePtr, execPart, debug, slopeX, slopeY := catchUserInput()
	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		if execPart == 'a' {
			fmt.Println("Number of trees: ", howManyTrees(filenamePtr, execPart, debug, slopeX, slopeY))
		} else {
			result := howManyTrees(filenamePtr, execPart, debug, 1, 1)
			result *= howManyTrees(filenamePtr, execPart, debug, 3, 1)
			result *= howManyTrees(filenamePtr, execPart, debug, 5, 1)
			result *= howManyTrees(filenamePtr, execPart, debug, 7, 1)
			result *= howManyTrees(filenamePtr, execPart, debug, 1, 2)
			fmt.Println("Result is: ", result)
		}
	}
}
