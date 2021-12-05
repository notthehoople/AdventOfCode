package main

import (
	"aoc/advent2021/utils"
	"fmt"
)

type coords struct {
	x int
	y int
}

func calcOverlapPoints(ventLines []string, debug bool) int {
	areaDiagram := make(map[coords]int)
	var startCoords, endCoords coords

	for i := 0; i < len(ventLines); i++ {
		fmt.Sscanf(ventLines[i], "%d,%d -> %d,%d\n", &startCoords.x, &startCoords.y, &endCoords.x, &endCoords.y)
		if debug {
			fmt.Println("Input Line:", ventLines[i])
			fmt.Printf("Start x:%d, Start y:%d, End x:%d, End y:%d\n",
				startCoords.x, startCoords.y, endCoords.x, endCoords.y)
		}

		// In part a we only care about vertical or horizontal lines. Check for those

		if startCoords.x == endCoords.x {
			var loopStart, loopEnd int
			if startCoords.y > endCoords.y {
				loopStart = endCoords.y
				loopEnd = startCoords.y
			} else {
				loopStart = startCoords.y
				loopEnd = endCoords.y
			}

			for i := loopStart; i <= loopEnd; i++ {
				if debug {
					fmt.Printf("x:%d y:%d\n", startCoords.x, i)
				}
				areaDiagram[coords{startCoords.x, i}]++
			}
		}

		if startCoords.y == endCoords.y {
			var loopStart, loopEnd int
			if startCoords.x > endCoords.x {
				loopStart = endCoords.x
				loopEnd = startCoords.x
			} else {
				loopStart = startCoords.x
				loopEnd = endCoords.x
			}

			for i := loopStart; i <= loopEnd; i++ {
				if debug {
					fmt.Printf("x:%d y:%d\n", i, startCoords.y)
				}
				areaDiagram[coords{i, startCoords.y}]++
			}
		}
		if debug {
			fmt.Println(areaDiagram)
		}
	}

	var overlapPoints int
	for _, i := range areaDiagram {
		if i > 1 {
			overlapPoints++
		}
	}

	return overlapPoints
}

func solveDay(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)
	if part == 'a' {
		return calcOverlapPoints(puzzleInput, debug)
	} else {
		return 0
	}
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", solveDay(filenamePtr, execPart, debug))
	}
}
