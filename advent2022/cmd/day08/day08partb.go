package main

import (
	"AdventOfCode-go/advent2022/utils"
	"fmt"
)

func numberOfTreesVisible(treeArea [][]int, y int, x int, debug bool) int {
	// 2. Check all the trees inside the area:
	//		only check rows and columns, not diagonals
	//		a tree is visible if all the trees to the edge are shorter than it is

	treeToTest := treeArea[y][x]

	if debug {
		fmt.Printf("Testing x: %d y: %d which is %d\n", x, y, treeArea[y][x])
	}

	var upTrees int
	for checkY := y - 1; checkY >= 0; checkY-- {
		if treeArea[checkY][x] < treeToTest {
			upTrees++
		} else {
			upTrees++
			break
		}
	}

	var downTrees int
	for checkY := y + 1; checkY < len(treeArea[y]); checkY++ {
		if treeArea[checkY][x] < treeToTest {
			downTrees++
		} else {
			downTrees++
			break
		}
	}

	var leftTrees int
	for checkX := x - 1; checkX >= 0; checkX-- {
		if treeArea[y][checkX] < treeToTest {
			leftTrees++
		} else {
			leftTrees++
			break
		}
	}

	var rightTrees int
	for checkX := x + 1; checkX < len(treeArea[y]); checkX++ {
		if treeArea[y][checkX] < treeToTest {
			rightTrees++
		} else {
			rightTrees++
			break
		}
	}

	if debug {
		fmt.Printf("up: %d down: %d left: %d right: %d\n", upTrees, downTrees, leftTrees, rightTrees)
	}

	return upTrees * downTrees * leftTrees * rightTrees
}

func findMaxScenicScore(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)

	treeArea := buildTreeArray(puzzleInput, debug)

	var maxScore int

	for y := 1; y < len(treeArea)-1; y++ {
		for x := 1; x < len(treeArea[y])-1; x++ {
			treeScore := numberOfTreesVisible(treeArea, y, x, debug)

			if treeScore > maxScore {
				maxScore = treeScore
			}
		}
	}

	if debug {
		printTreeArea(treeArea, debug)
	}

	return maxScore
}
