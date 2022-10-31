package main

import (
	"AdventOfCode-go/advent2021/utils"
	"fmt"
)

type coords struct {
	x int
	y int
}

func calcOverlapPoints(ventLines []string, part byte, debug bool) int {
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

		if part == 'a' {
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
		} else {
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

			// Part b also has diagonals at 45 degrees. Lets check for them
			if (startCoords.x != endCoords.x) && (startCoords.y != endCoords.y) {
				var loopStartX, loopEndX, loopStartY, loopEndY int
				var backwards bool = false
				if startCoords.x > endCoords.x {
					loopStartX = endCoords.x
					loopEndX = startCoords.x
					loopStartY = endCoords.y
					loopEndY = startCoords.y
					// need to deal with backwards diagonal where y decreases
					if loopStartY > loopEndY {
						backwards = true
					}
				} else {
					loopStartX = startCoords.x
					loopEndX = endCoords.x
					loopStartY = startCoords.y
					loopEndY = endCoords.y
					// need to deal with backwards diagonal where y decreases
					if loopStartY > loopEndY {
						backwards = true
					}
				}

				var currentY = loopStartY
				for currentX := loopStartX; currentX <= loopEndX; currentX++ {
					if debug {
						fmt.Printf("x:%d y:%d\n", currentX, currentY)
					}
					areaDiagram[coords{currentX, currentY}]++
					if backwards {
						currentY--
					} else {
						currentY++
					}
				}
			}
			if debug {
				fmt.Println(areaDiagram)
			}
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
	return calcOverlapPoints(puzzleInput, part, debug)
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
