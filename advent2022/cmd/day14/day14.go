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

func calcMaximums(caves map[Coords]byte) (int, int, int, int) {
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

	return minX, maxX, minY, maxY
}

func printCaves(caves map[Coords]byte) {

	minX, maxX, minY, maxY := calcMaximums(caves)

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
						caves[Coords{x: j, y: startPos.y}] = '#'
					}
				} else {
					for j := startPos.x; j >= endPos.x; j-- {
						caves[Coords{x: j, y: startPos.y}] = '#'
					}
				}
			} else {
				if endPos.y > startPos.y {
					for j := startPos.y; j <= endPos.y; j++ {

						caves[Coords{x: startPos.x, y: j}] = '#'
					}
				} else {
					for j := startPos.y; j >= endPos.y; j-- {

						caves[Coords{x: startPos.x, y: j}] = '#'
					}
				}
			}
		}
	}

	return caves
}

func dropSingleSandPartB(caves map[Coords]byte, startPoint Coords, maxY int) (Coords, bool) {
	var currentPos Coords

	currentPos = startPoint
	for currentPos.y < maxY+1 {
		if value, doesMapContainKey := caves[Coords{x: currentPos.x, y: currentPos.y + 1}]; doesMapContainKey {
			if value == '#' || value == 'o' {

				if valueLeft, doesMapContainKeyLeft := caves[Coords{x: currentPos.x - 1, y: currentPos.y + 1}]; doesMapContainKeyLeft {
					if valueLeft == '#' || valueLeft == 'o' {
						if valueRight, doesMapContainKeyRight := caves[Coords{x: currentPos.x + 1, y: currentPos.y + 1}]; doesMapContainKeyRight {
							if valueRight == '#' || valueRight == 'o' {
								return currentPos, false
							} else {
								currentPos = Coords{x: currentPos.x + 1, y: currentPos.y + 1}
							}
						} else {
							// Doesn't exist in the map
							currentPos = Coords{x: currentPos.x + 1, y: currentPos.y + 1}
						}
					} else {
						// Exists in the map but isn't a # or o
						fmt.Println("Got here???")
						//currentPos = Coords{x: currentPos.x - 1, y: currentPos.y + 1}
					}
				} else {
					// Doesn't exist in the map
					currentPos = Coords{x: currentPos.x - 1, y: currentPos.y + 1}
				}
			} else if value == '+' {
				return Coords{x: currentPos.x, y: currentPos.y + 1}, true
			}
		} else {
			currentPos = Coords{x: currentPos.x, y: currentPos.y + 1}
		}
	}
	if currentPos.y == maxY+1 {
		return currentPos, false
	}

	return Coords{x: 0, y: 0}, true
}

func dropSingleSand(caves map[Coords]byte, startPoint Coords, part byte) (Coords, bool) {
	var currentPos Coords

	// We need the maximum Y of the walls so we know when we've gone off the end of the cave
	_, _, _, maxY := calcMaximums(caves)

	currentPos = startPoint
	for currentPos.y <= maxY {
		if value, doesMapContainKey := caves[Coords{x: currentPos.x, y: currentPos.y + 1}]; doesMapContainKey {
			if value == '#' || value == 'o' {

				if valueLeft, doesMapContainKeyLeft := caves[Coords{x: currentPos.x - 1, y: currentPos.y + 1}]; doesMapContainKeyLeft {
					if valueLeft == '#' || valueLeft == 'o' {
						if valueRight, doesMapContainKeyRight := caves[Coords{x: currentPos.x + 1, y: currentPos.y + 1}]; doesMapContainKeyRight {
							if valueRight == '#' || valueRight == 'o' {
								return currentPos, false
							} else {
								currentPos = Coords{x: currentPos.x + 1, y: currentPos.y + 1}
							}
						} else {
							// Doesn't exist in the map
							currentPos = Coords{x: currentPos.x + 1, y: currentPos.y + 1}
						}
					} else {
						// Exists in the map but isn't a # or o
						fmt.Println("Got here???")
						//currentPos = Coords{x: currentPos.x - 1, y: currentPos.y + 1}
					}
				} else {
					// Doesn't exist in the map
					currentPos = Coords{x: currentPos.x - 1, y: currentPos.y + 1}
				}
			} else if value == '+' {
				return Coords{x: currentPos.x, y: currentPos.y + 1}, true
			}
		} else {
			currentPos = Coords{x: currentPos.x, y: currentPos.y + 1}
		}
	}

	return Coords{x: 0, y: 0}, true
}

func calcUnitsOfSand(filename string, part byte, debug bool) int {
	var startPoint Coords = Coords{x: 500, y: 0}
	var maxY int

	puzzleInput, _ := utils.ReadFile(filename)

	caves := buildCaveArray(puzzleInput, debug)
	if debug {
		printCaves(caves)
	}

	if part == 'b' {
		// We need the maximum Y of the walls so we know when we've gone off the end of the cave
		_, _, _, maxY = calcMaximums(caves)
	}

	for {
		var sandCoords Coords
		var finished bool

		if part == 'a' {
			sandCoords, finished = dropSingleSand(caves, startPoint, part)
			if !finished {
				caves[sandCoords] = 'o'
			} else {
				printCaves(caves)
				break
			}
		} else {
			// part b
			sandCoords, finished = dropSingleSandPartB(caves, startPoint, maxY)

			if !finished {
				if value, check := caves[sandCoords]; check {
					if value == '+' {
						fmt.Println("Finished")
						caves[sandCoords] = 'o'
						printCaves(caves)
						break
					}
				}
				caves[sandCoords] = 'o'
			} else {
				printCaves(caves)
				break
			}
		}
	}

	var sandCount int
	for _, value := range caves {
		if value == 'o' {
			sandCount++
		}
	}

	return sandCount
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Units of Sand: %d\n", calcUnitsOfSand(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("Units of Sand: %d\n", calcUnitsOfSand(filenamePtr, execPart, debug))
	case 'z':
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
