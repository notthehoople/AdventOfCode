package main

import (
	"fmt"
	"bufio"
	"os"
//	"os/exec"
	"flag"
)

// Used to build a work list of water sources to be followed
type workListCoords struct {
	done	bool	
	xCoord	int
	yCoord	int
}

type pointStruct struct {
	xCoordStart	int
	xCoordEnd	int
	yCoordStart	int
	yCoordEnd	int
	yCountStart	int
}

// Read the text file passed in by name into a array of strings
// Returns the array as the first return variable
func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
	  return nil, err
	}
	defer file.Close()
  
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	  lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func printStringArray(tempString []string) {
	// Loop through the array and print each line
	for i:= 0; i < len(tempString); i++ {
		fmt.Println(tempString[i])
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

// func countWaterSquares
// Count the '~' and '|' squares in undergroundArea
func countWaterSquares(undergroundArea [][]byte, part byte, yCountStart int) int {
	var count int = 0

	for i := yCountStart; i < len(undergroundArea); i++ {
		for j := 0; j < len(undergroundArea[i]); j++ {
			if part == 'a' {
				if undergroundArea[i][j] == '~' || undergroundArea[i][j] == '|' {
					count++
				}
			} else {
				if undergroundArea[i][j] == '~' {
					count++
				}
			}
		}
	}
	return count
}

func addWorkListItem(workList []workListCoords, workX int, workY int) []workListCoords {
	var checkList bool = true

	for i := 0; i < len(workList); i++ {
		if workList[i].xCoord == workX && workList[i].yCoord == workY {
			checkList = false
		}
	}
	if checkList {
		workList = append(workList, workListCoords{done: false, xCoord: workX, yCoord: workY})
	}

	return workList
}

// func: readInitialState
// takes an array of strings and breaks it into a 2D array of bytes
func readInitialState(coordData []pointStruct, undergroundArea [][]byte, springX int, springY int, minX int, minY int) {

	for i := 0; i < len(undergroundArea); i++ {
		for j := 0; j < len(undergroundArea[i]); j++ {
			undergroundArea[i][j] = '.'
		}
	}

	for i := 0; i < len(coordData); i++ {
		if coordData[i].xCoordStart == coordData[i].xCoordEnd {
			for y := (coordData[i].yCoordStart - minY); y <= (coordData[i].yCoordEnd - minY); y++ {
				undergroundArea[y][coordData[i].xCoordStart-minX] = '#'
			}
		} else {
			// must be 'y'

			for x := (coordData[i].xCoordStart - minX); x < (coordData[i].xCoordEnd - minX); x++ {
				undergroundArea[coordData[i].yCoordStart-minY][x] = '#'
			}
		}
	}

	undergroundArea[springY - minY][springX - minX] = '+'
}

// func scanInputForMaxMins
// Walks through the coordData array and works out the minimum and maximum values of X and Y
func scanInputForMaxMins(coordData []pointStruct, springX int, springY int) (int, int, int, int, int) {
	var minX, maxX, minY, maxY, yCountStart int = 0, 0, 1000, 0, 0

	minX = coordData[0].xCoordStart
	maxX = coordData[0].xCoordEnd
	minY = coordData[0].yCoordStart
	maxY = coordData[0].yCoordEnd

	for i := 0; i < len(coordData); i++ {
		if coordData[i].xCoordStart < minX {
			minX = coordData[i].xCoordStart
		}
		if coordData[i].xCoordEnd > maxX {
			maxX = coordData[i].xCoordEnd + 5
		}
		if coordData[i].yCoordStart < minY {
			minY = coordData[i].yCoordStart
		}
		if coordData[i].yCoordEnd > maxY {
			maxY = coordData[i].yCoordEnd
		}
	}

	yCountStart = minY

	if springX < minX {
		minX = springX
	} else {
		if springX > maxX {
			maxX = springX
		}
	}
	if springY < minY {
		minY = springY
	} else {
		if springY > maxY {
			maxY = springY
		}
	}

	return minX, maxX, minY, maxY, yCountStart
}

// func processInputFile
// Returns a pointStruct array with the processed data in it
// Data is provided in the file as follows:
//   x=495, y=2..7
//   y=13, x=498..504
// So we need to deal with both a single x with multiple ys, and a single y with multiple xs
// In the struct, start and end will be the same for the single digit coord
func processInputFile(fileContents []string) []pointStruct {
	var coordData []pointStruct
	var a, b, c int
	var coord1, coord2 byte
	//var result int

	for i := 0; i < len(fileContents); i++ {
		fmt.Sscanf(fileContents[i], "%c=%d, %c=%d..%d\n", &coord1, &a, &coord2, &b, &c)
		
		if coord1 == 'x' {
			coordData = append(coordData, pointStruct{xCoordStart: a, xCoordEnd: a, yCoordStart: b, yCoordEnd: c})
		} else {
			// must be 'y'
			coordData = append(coordData, pointStruct{xCoordStart: b, xCoordEnd: c, yCoordStart: a, yCoordEnd: a})
		}
	}

	return coordData
}

func fillLine(undergroundArea [][]byte, workList []workListCoords, currentFlowX int, currentFlowY int, maxiMins pointStruct) (bool, []workListCoords) {
	var x int = currentFlowX
	var y int = currentFlowY
	var leftEdge, rightEdge int = currentFlowX, currentFlowX
	var isEnclosed bool = true

	// Look for left extent
	for {
		if undergroundArea[y][x] != '#' && x >= maxiMins.xCoordStart {
			if undergroundArea[y+1][x] == '#' || undergroundArea[y+1][x] == '~' {
				x--
			} else {
				// Fallen off the edge to the left
				isEnclosed = false
				leftEdge = x
				workList = addWorkListItem(workList, x, y)
				break
			}
		} else {
			if undergroundArea[y][x] == '#' {
				// We have found a left edge
				leftEdge = x+1
				break
			} else {
				// No edge found
				isEnclosed = false
				break
			}
		}
	}
	// Look for right extent
	x = currentFlowX
	for {
		if undergroundArea[y][x] != '#' && x <= maxiMins.xCoordEnd {
			if undergroundArea[y+1][x] == '#' || undergroundArea[y+1][x] == '~' {
				x++
			} else {
				// Fallen off the edge to the right
				isEnclosed = false
				rightEdge = x
				workList = addWorkListItem(workList, x, y)
				break
			}
		} else {
			if undergroundArea[y][x] == '#' {
				rightEdge = x-1
				break
			} else {
				// No right edge found
				isEnclosed = false
				break
			}
		}
	}

	//fmt.Printf("Source X: %d Source Y: %d\n", currentFlowX, currentFlowY)
	//fmt.Printf("Left Edge: %d Right Edge: %d\n", leftEdge, rightEdge)

	if isEnclosed {
		for i := leftEdge; i <= rightEdge; i++ {
			undergroundArea[y][i] = '~'
		}
	} else {
		for i := leftEdge; i <= rightEdge; i++ {
			undergroundArea[y][i] = '|'
		}
	}

	return isEnclosed, workList
}

// func letTheWaterFlow
// Handles how the water flows from the source x, y point until it either reaches the end of the undergroundArea or
// it splits into 2 sources itself
func letTheWaterFlow(undergroundArea [][]byte, workList []workListCoords, sourceX int, sourceY int, maxiMins pointStruct) (bool, []workListCoords) {
	var currentFlowX, currentFlowY int = sourceX, sourceY
	var loopThis bool = true
	var firstVisit bool = true
	var result bool

	// forever loop as we're following water. We'll control the loop count ourselves
	for loopThis {

		switch undergroundArea[currentFlowY][currentFlowX] {

		case '+': currentFlowY++

		case '.':
			// Check what's next:
			//    If free area then add '|' and continue
			//    If reached blocker
			//       check if line is fillable
			//          yes -> fill line with '~' and y--
			//          no -> find water sources (spills) and create new water sources in workList
			//             -> fill line with '|'
			//             -> mark this water source as complete
			if currentFlowY >= maxiMins.yCoordEnd {
				// We've reached the end of the grid so we're done
				return true, workList
			}
			if undergroundArea[currentFlowY + 1][currentFlowX] == '.' {
				undergroundArea[currentFlowY][currentFlowX] = '|'
				currentFlowY++
			} else {
				// Make sure fillLine returns false if we have open ends otherwise this does the wrong thing
				if undergroundArea[currentFlowY + 1][currentFlowX] == '#' || undergroundArea[currentFlowY + 1][currentFlowX] == '~' {
					result, workList = fillLine(undergroundArea, workList, currentFlowX, currentFlowY, maxiMins)
					if result {
						currentFlowY--
					} else {
						return true, workList
					}
				} else {
					if undergroundArea[currentFlowY + 1][currentFlowX] == '|' {
						undergroundArea[currentFlowY][currentFlowX] = '|'
						return true, workList
					}
				}
			}

		case '|':
			if currentFlowX == sourceX && currentFlowY == sourceY && firstVisit {
				currentFlowY++
				firstVisit = false
			} else {
				if undergroundArea[currentFlowY + 1][currentFlowX] == '~' {
					result, workList = fillLine(undergroundArea, workList, currentFlowX, currentFlowY, maxiMins)
					if result {
						currentFlowY--
					} else {
						// Found an edge
						loopThis = false
						break
					}
				}
			}

		case '~':
			// This means that we're a water source that has reached the water created by a
			// difference water source. We still have to do something, since we could be on the other side of a blocker and
			// capable of filling where the other water source can't

			if currentFlowX == sourceX && currentFlowY == sourceY {
				currentFlowY++
			} else {
				if undergroundArea[currentFlowY + 1][currentFlowX] == '~' {
					result, workList = fillLine(undergroundArea, workList, currentFlowX, currentFlowY, maxiMins)
					if result {
						currentFlowY--
					} else {
						loopThis = false
						break
					}
				}
			}
		}
	}
	
	return true, workList
}

// func processWaterFlow
// Handles everything needed to work out the water flow (day 17 part A)
func processWaterFlow(fileName string, springX int, springY int, part byte) int {
	var minX, maxX, minY, maxY, gridSizeX, gridSizeY, workX, workY, yCountStart int
	var coordData []pointStruct
	var workList []workListCoords
	var letsLoopThis bool
	var maxiMins pointStruct
	var didWork bool

	// Read contents of file into a string array
	fileContents, _ := readLines(fileName)
	coordData = processInputFile(fileContents)

	minX, maxX, minY, maxY, yCountStart = scanInputForMaxMins(coordData, springX, springY)

	maxiMins.xCoordStart = 0
	maxiMins.xCoordEnd = maxX - minX
	maxiMins.yCoordStart = 0
	maxiMins.yCoordEnd = maxY - minY + 1
	maxiMins.yCountStart = yCountStart

	gridSizeX = (maxX - minX) + 1
	gridSizeY = (maxY - minY) + 2

	undergroundArea := make([][]byte, gridSizeY)
	for i := 0; i < gridSizeY; i++ {
		undergroundArea[i] = make([]byte, gridSizeX)	
	}

	workList = make([]workListCoords, 0)

	readInitialState(coordData, undergroundArea, springX, springY, minX, minY)

	workList = addWorkListItem(workList, springX - minX, springY - minY)
	letsLoopThis = true

	for letsLoopThis {
		didWork = false

		// Loop through the list of work we have. This list is a list of water sources
		for i := 0; i < len(workList); i++ {
			if !workList[i].done {
				workX = workList[i].xCoord
				workY = workList[i].yCoord

				letsLoopThis, workList = letTheWaterFlow(undergroundArea, workList, workX, workY, maxiMins)
				workList[i].done = true
				didWork = true
			}
		}
		if !didWork {
			letsLoopThis = false
		}

	}
	
	// Print final water flow
	print2DSlice(undergroundArea)

	return countWaterSquares(undergroundArea, part, maxiMins.yCountStart)
}

// Main routine
func main() {
	var springX, springY int = 0, 0

	fileNamePtr := flag.String("file", "input1.txt", "A filename containing input strings")
	flag.IntVar(&springX, "springx", 500, "x coord of the spring of water")
	flag.IntVar(&springY, "springy", 0, "y coord of the spring of water")
	execPartPtr := flag.String("part", "a", "Which part of day18 do you want to calc (a or b)")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Number of water tiles:", processWaterFlow(*fileNamePtr, springX, springY, 'a'))
	case "b":
		fmt.Println("Part b - Number of still water tiles:", processWaterFlow(*fileNamePtr, springX, springY, 'b'))
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}