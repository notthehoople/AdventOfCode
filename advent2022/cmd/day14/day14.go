package main

import (
	"AdventOfCode-go/advent2022/utils"
	"fmt"
	"strconv"
	"strings"
)

type Coords struct {
	x int
	y int
}

func printCaves(caves map[Coords]byte) {
	var minX, minY int = 99999, 99999
	var maxX, maxY int

	for tempCoords := range caves {
		if tempCoords.x > maxX {
			maxX = tempCoords.x
		}
		if tempCoords.x < minX {
			minX = tempCoords.x
		}
		if tempCoords.y > maxY {
			maxY = tempCoords.y
		}
		if tempCoords.y < minY {
			minY = tempCoords.y
		}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			value, ok := caves[Coords{x: x, y: y}]
			if ok {
				fmt.Printf("%c", value)
			} else {
				fmt.Printf("%c", '.')
			}
		}
		fmt.Printf("\n")
	}
	println("==")
}

func convertToCoords(s string) Coords {
	var pos Coords

	coords := strings.Split(strings.TrimSpace(s), ",")
	pos.x, _ = strconv.Atoi(coords[0])
	pos.y, _ = strconv.Atoi(coords[1])

	return pos
}

func buildCaveArray(puzzleInput []string, debug bool) map[Coords]byte {

	caves := make(map[Coords]byte, 0)

	caves[Coords{x: 500, y: 0}] = '+'

	for _, line := range puzzleInput {

		locations := strings.Split(line, "->")
		var startPos, endPos Coords
		for i := 1; i < len(locations); i++ {
			startPos = convertToCoords(locations[i-1])
			endPos = convertToCoords(locations[i])

			if debug {
				fmt.Printf("StartPos Location: %s\n", strings.TrimSpace(locations[i-1]))
				fmt.Println("startPos:", startPos)
				fmt.Printf("EndPos Location: %s\n", strings.TrimSpace(locations[i]))
				fmt.Println("endPos:", endPos)
			}

			if startPos.y == endPos.y {
				if endPos.x > startPos.x {
					for j := startPos.x; j <= endPos.x; j++ {
						//fmt.Printf("endPosX > Adding x:%d y:%d Coords %v\n", j, startPos.y, Coords{x: j, y: startPos.y})
						caves[Coords{x: j, y: startPos.y}] = '#'
					}
				} else {
					//fmt.Printf("here endPos.x: %d startPos.x: %d\n", endPos.x, startPos.x)
					for j := startPos.x; j >= endPos.x; j-- {
						//fmt.Printf("endPosX < Adding x:%d y:%d Coords %v\n", j, startPos.y, Coords{x: j, y: startPos.y})
						caves[Coords{x: j, y: startPos.y}] = '#'
					}
				}
			} else {
				if endPos.y > startPos.y {
					for j := startPos.y; j <= endPos.y; j++ {
						//fmt.Printf("endPosY > Adding x:%d y:%d Coords %v\n", startPos.x, j, Coords{x: startPos.x, y: j})

						caves[Coords{x: startPos.x, y: j}] = '#'
					}
				} else {
					for j := startPos.y; j >= endPos.y; j-- {
						//fmt.Printf("endPosY < Adding x:%d y:%d Coords %v\n", j, startPos.y, Coords{x: startPos.x, y: j})

						caves[Coords{x: startPos.x, y: j}] = '#'
					}
				}
			}
		}
	}

	return caves
}

func dropSingleSand(caves map[Coords]byte, startPoint Coords) (Coords, bool) {
	var currentPos Coords
	currentPos = startPoint
	for {
		if value, doesMapContainKey := caves[Coords{x: currentPos.x, y: currentPos.y + 1}]; doesMapContainKey {
			fmt.Printf("Next cell contains: %c\n", value)
			if value == '#' {
				// Need the smarts at this point
				return currentPos, false
			}
		} else {
			fmt.Printf("Map doesn't contain anything\n")
		}
		currentPos = Coords{x: currentPos.x, y: currentPos.y + 1}
	}

	return Coords{x: 0, y: 0}, true
}

func calcUnitsOfSand(filename string, part byte, debug bool) int {
	var startPoint Coords = Coords{x: 500, y: 0}

	puzzleInput, _ := utils.ReadFile(filename)

	caves := buildCaveArray(puzzleInput, debug)
	//fmt.Println("Cave Array:", caves)
	printCaves(caves)

	var sandCoords Coords
	var finished bool
	sandCoords, finished = dropSingleSand(caves, startPoint)
	if !finished {
		caves[sandCoords] = 'o'
	}

	printCaves(caves)

	return 0
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Units of Sand: %d\n", calcUnitsOfSand(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("Not implemented yet\n")
	case 'z':
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
