package main

import (
	"AdventOfCode-go/advent2022/utils"
	"fmt"
)

func printTreeArea(treeArea [][]int, debug bool) {
	for y := 0; y < len(treeArea); y++ {
		for x := 0; x < len(treeArea[y]); x++ {
			fmt.Printf("%d", treeArea[y][x])
		}
		fmt.Printf("\n")
	}
}

// Note: for 2d arrays:
//     treeArea[0][i] results in [0 1 2 3 4] i.e. the 2nd digit is the X axis
//     treeArea[i][0] results in [[0 0 0 0 0] [1 0 0 0 0] [2 0 0 0 0] [3 0 0 0 0] [4 0 0 0 0]]
//			i.e. the first digit is the Y axis

func buildTreeArray(input []string, debug bool) [][]int {

	treeArea := make([][]int, len(input))
	for i := 0; i < len(input); i++ {
		treeArea[i] = make([]int, len(input[0]))
	}

	for y := 0; y < len(treeArea); y++ {
		for x := 0; x < len(treeArea[y]); x++ {
			treeArea[y][x] = int(input[y][x] - '0')
		}
	}

	if debug {
		printTreeArea(treeArea, debug)
	}

	return treeArea
}

func isVisible(treeArea [][]int, y int, x int, debug bool) bool {
	// 2. Check all the trees inside the area:
	//		only check rows and columns, not diagonals
	//		a tree is visible if all the trees to the edge are shorter than it is

	treeToTest := treeArea[y][x]

	if debug {
		fmt.Printf("Testing x: %d y: %d which is %d\n", x, y, treeArea[y][x])
	}

	var visibleUp bool = true
	for checkY := y - 1; checkY >= 0; checkY-- {
		if treeArea[checkY][x] >= treeToTest {
			visibleUp = false
			break
		}
	}
	if visibleUp {
		return true
	}

	var visibleDown bool = true
	for checkY := y + 1; checkY < len(treeArea[y]); checkY++ {
		if treeArea[checkY][x] >= treeToTest {
			visibleDown = false
			break
		}
	}
	if visibleDown {
		return true
	}

	var visibleLeft bool = true
	for checkX := x - 1; checkX >= 0; checkX-- {

		if treeArea[y][checkX] >= treeToTest {
			visibleLeft = false
			break
		}
	}
	if visibleLeft {
		return true
	}

	var visibleRight bool = true
	for checkX := x + 1; checkX < len(treeArea[y]); checkX++ {

		if treeArea[y][checkX] >= treeToTest {
			visibleRight = false
			break
		}
	}
	if visibleRight {
		return true
	}

	return false
}

func countVisibleTrees(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)

	treeArea := buildTreeArray(puzzleInput, debug)

	var visibleTreeCount int

	// To count the trees visible from outside the tree area:
	// 1. Count all the edge trees
	// 2. Check all the trees inside the area:
	//		only check rows and columns, not diagonals
	//		a tree is visible if all the trees to the edge are shorter than it is

	// Ignoring the edge trees, so start loop from 1 and finish before the end edge

	//	isVisible(treeArea, 3, 2)

	for y := 1; y < len(treeArea)-1; y++ {
		for x := 1; x < len(treeArea[y])-1; x++ {
			if isVisible(treeArea, y, x, debug) {
				visibleTreeCount++
			}
		}
	}

	if debug {
		printTreeArea(treeArea, debug)
	}

	// Now count the trees on the edge of the area which are always visible
	visibleTreeCount += len(treeArea[0])
	visibleTreeCount += len(treeArea[0])
	visibleTreeCount += len(treeArea) - 2
	visibleTreeCount += len(treeArea) - 2
	return visibleTreeCount
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Result is: %d\n", countVisibleTrees(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("Result is: %d\n", findMaxScenicScore(filenamePtr, execPart, debug))
	case 'z':
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
