package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"flag"
)

type pointStruct struct {
	xCoordStart	int
	xCoordEnd	int
	yCoordStart	int
	yCoordEnd	int
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
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()

	for i := 0; i < len(tempSlice); i++ {
		for j := 0; j < len(tempSlice[i]); j++ {
			fmt.Printf("%c", tempSlice[i][j])
		}
		fmt.Printf("\n")
	}
}

// func: readInitialState
// takes an array of strings and breaks it into a 2D array of bytes
func readInitialState(coordData []pointStruct, undergroundArea [][]byte, springX int, springY int, minX int, minY int) {

	fmt.Printf("minX %d minY %d\n", minX, minY)
	for i := 0; i < len(undergroundArea); i++ {
		for j := 0; j < len(undergroundArea[i]); j++ {
			undergroundArea[i][j] = '.'
		}
	}

	for i := 0; i < len(coordData); i++ {
		if coordData[i].xCoordStart == coordData[i].xCoordEnd {
			for y := (coordData[i].yCoordStart - minY); y <= (coordData[i].yCoordEnd - minY); y++ {
				fmt.Printf("y is %d x is %d MODIFIED: y is %d x is %d\n", y, coordData[i].xCoordStart, y, (coordData[i].xCoordStart-minX))
				undergroundArea[y][coordData[i].xCoordStart-minX] = '#'
				fmt.Println("Done")
			}
		} else {
			// must be 'y'

			for x := (coordData[i].xCoordStart - minX); x < (coordData[i].xCoordEnd - minX); x++ {
				fmt.Printf("x is %d y is %d MODIFIED: x is %d y is %d\n", x, coordData[i].yCoordStart, x, (coordData[i].yCoordStart-minY))

				undergroundArea[coordData[i].yCoordStart-minY][x] = '#'
			}
		}
	}

	undergroundArea[springY - minY][springX - minX] = '+'
}

// func scanInputForMaxMins
// Walks through the coordData array and works out the minimum and maximum values of X and Y
func scanInputForMaxMins(coordData []pointStruct, springX int, springY int) (int, int, int, int) {
	var minX, maxX, minY, maxY int = 0, 0, 0, 0

	minX = coordData[0].xCoordStart
	maxX = coordData[0].xCoordEnd
	minY = coordData[0].yCoordStart
	maxY = coordData[0].yCoordEnd

	for i := 0; i < len(coordData); i++ {
		if coordData[i].xCoordStart < minX {
			minX = coordData[i].xCoordStart
		}
		if coordData[i].xCoordEnd > maxX {
			maxX = coordData[i].xCoordEnd
		}
		if coordData[i].yCoordStart < minY {
			minY = coordData[i].yCoordStart
		}
		if coordData[i].yCoordEnd > maxY {
			maxY = coordData[i].yCoordEnd
		}
	}

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

	return minX, maxX, minY, maxY
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
	var result int

	for i := 0; i < len(fileContents); i++ {
		result, _ = fmt.Sscanf(fileContents[i], "%c=%d, %c=%d..%d\n", &coord1, &a, &coord2, &b, &c)
		if result == 5 {
			fmt.Printf("Scan run: %c is %d %c starts: %d ends: %d\n", coord1, a, coord2, b, c)
		}

		if coord1 == 'x' {
			coordData = append(coordData, pointStruct{xCoordStart: a, xCoordEnd: a, yCoordStart: b, yCoordEnd: c})
		} else {
			// must be 'y'
			coordData = append(coordData, pointStruct{xCoordStart: b, xCoordEnd: c, yCoordStart: a, yCoordEnd: a})
		}
	}

	return coordData
}

func fillLine(undergroundArea [][]byte, currentFlowX int, currentFlowY int) bool {
	var x int = currentFlowX
	var y int = currentFlowY

	// This needs to change to a "look all the way left, look all the way right" to decide if we're in an enclosed area or not
	// If not enclosed then we change all '~' or '.' to '|' and spill over the edge.
	// No idea how to handle when we spill over the edge in multiple places
	// Probably need to maintain a work list that contains a list of all the spill areas. We can then loop through the spills
	// Until each of them runs out
	for {
		if undergroundArea[y][x] != '#' {
			if undergroundArea[y+1][x] == '#' || undergroundArea[y+1][x] == '~' {
				undergroundArea[y][x] = '~'
				x--
			}
		} else {
			break
		}
	}
	return true
}

func letTheWaterFlow(undergroundArea [][]byte, springX int, springY int) int {
	var currentFlowX, currentFlowY int = springX, springY
	var loopCount int = 0

	if undergroundArea[currentFlowY][currentFlowX] == '+' {
		fmt.Println("All is good2")
	}

	// forever loop as we're following water. We'll control the loop count ourselves
	for {
		switch undergroundArea[currentFlowY][currentFlowX] {

		case '+': currentFlowY++

		case '.':
			if undergroundArea[currentFlowY + 1][currentFlowX] == '.' {
				undergroundArea[currentFlowY][currentFlowX] = '|'
				currentFlowY++
			} else {
				if undergroundArea[currentFlowY + 1][currentFlowX] == '#' {
					if fillLine(undergroundArea, currentFlowX, currentFlowY) {
						currentFlowY--
					}
				}
			}

		case '|':
			if undergroundArea[currentFlowY + 1][currentFlowX] == '~' {
				if fillLine(undergroundArea, currentFlowX, currentFlowY) {
					currentFlowY--
				}
			}

		case '~':

		}

		fmt.Println("In a loop:", loopCount)
		loopCount++
		if loopCount > 10 {
			break
		}
	}
	
	return 0
}

// func processWaterFlow
// Handles everything needed to work out the water flow (day 17 part A)
func processWaterFlow(fileName string, springX int, springY int, part byte) int {
	var minX, maxX, minY, maxY, gridSizeX, gridSizeY int
	var coordData []pointStruct

	// Read contents of file into a string array
	fileContents, _ := readLines(fileName)
	coordData = processInputFile(fileContents)
	fmt.Println(coordData)

	minX, maxX, minY, maxY = scanInputForMaxMins(coordData, springX, springY)
	fmt.Printf("minX: %d maxX: %d minY: %d maxY: %d\n", minX, maxX, minY, maxY)

	gridSizeX = (maxX - minX) + 1
	gridSizeY = (maxY - minY) + 1
	fmt.Printf("gridSizeX: %d gridSizeY: %d\n", gridSizeX, gridSizeY)

	undergroundArea := make([][]byte, gridSizeY)
	for i := 0; i < gridSizeY; i++ {
		undergroundArea[i] = make([]byte, gridSizeX)	
	}

	readInitialState(coordData, undergroundArea, springX, springY, minX, minY)

	print2DSlice(undergroundArea)

	letTheWaterFlow(undergroundArea, springX - minX, springY - minY)

	print2DSlice(undergroundArea)

	return 0
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
		fmt.Println("Not here yet")
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}