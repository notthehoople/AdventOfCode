package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"flag"
	"strconv"
	//"strings"
)

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
			fmt.Printf("%c ", tempSlice[i][j])
		}
		fmt.Printf("\n")
	}
}

// func: readInitialState
// takes an array of strings and breaks it into a 2D array of bytes
func readInitialState(tempString []string, tempSlice [][]byte) {
	for i := 0; i < len(tempString); i++ {
		for j := 0; j < len(tempString[i]); j++ {
			tempSlice[i][j] = tempString[i][j]
		}
	}
}

// func checkArea
// returns: lumberYards, openAcre, trees
func checkArea(tempArea [][]byte, yCoord int, xCoord int) (int, int) {
	var lumberYards int = 0
	var trees int = 0
	var startX, startY, maxX, maxY int

	// To consider: the edges. So if xCoord / yCoord == 0 or xCoord / yCoord = len(tempArea)

	// x x   x o   o x   x x   x o x   x x x   x x   x x
	// o x   x x   x x   x o   x x x   x o x   o x   x o
	//                                         x x   x x

	//fmt.Println("Printing 2D Slice in checkArea")
	//print2DSlice(tempArea)

	startX = xCoord - 1
	startY = yCoord - 1
	maxX = xCoord + 1
	maxY = yCoord + 1

	if xCoord == 0 {
		startX = 0
	} else {
		if xCoord == len(tempArea[yCoord])-1 {
			maxX = xCoord
		}
	}
	if yCoord == 0 {
		startY = 0
	} else {
		if yCoord == len(tempArea)-1 {
			maxY = yCoord
		}
	}

	//fmt.Println("================")
	//fmt.Printf("xCoord: %d yCoord: %d MaxX: %d MaxY: %d startX: %d startY: %d\n", xCoord, yCoord, maxX, maxY, startX, startY)

	for i := startY; i <= maxY; i++ {
		for j := startX; j <= maxX; j++ {
			if i == yCoord && j == xCoord {
				// Do nothing
			} else {
				if tempArea[i][j] == '|' {
					trees++
				} else {
					if tempArea[i][j] == '#' {
						lumberYards++
					}
				}
			}
		}
	}

	return lumberYards, trees
}


// func playRound
// Play a round of Life!
func playRound(tempArea [][]byte, scratchArea [][]byte) {
	var lumberYards, trees int = 0, 0
	var lenTempArea, lenTempAreaI int = 0, 0

	//tempCalcSpace := make([][]byte, len(tempArea))
	//
	//for i := 0; i < len(tempArea); i++ {
	//	tempCalcSpace[i] = make([]byte, len(tempArea[i]))	
	//}

	lenTempArea = len(tempArea)
	for i := 0; i < lenTempArea; i++ {
		lenTempAreaI = len(tempArea[i])
		for j := 0; j < lenTempAreaI; j++ {
			lumberYards, trees = checkArea(tempArea, i, j)

			// An open acre ('.') will become filled with trees if three or more adjacent acres contained 
			// trees. Otherwise, nothing happens.
			if tempArea[i][j] == '.' {
				if trees > 2 {
					scratchArea[i][j] = '|'
				} else {
					scratchArea[i][j] = '.'
				}
			}

			// An acre filled with trees ('|') will become a lumberyard if three or more adjacent acres were
			//    lumberyards. Otherwise, nothing happens.
			if tempArea[i][j] == '|' {
				if lumberYards > 2 {
					scratchArea[i][j] = '#'
				} else {
					scratchArea[i][j] = '|'
				}
			}

			// An acre containing a lumberyard ('#') will remain a lumberyard if it was adjacent to at least 
			//    one other lumberyard and at least one acre containing trees. Otherwise, it becomes open.
			if tempArea[i][j] == '#' {
				if lumberYards > 0 && trees > 0 {
					scratchArea[i][j] = '#'
				} else {
					scratchArea[i][j] = '.'
				}
			}
		}
	}

	// Now copy it back to the original so all the changes happen at the same time
	for i := 0; i < lenTempArea; i++ {
		for j := 0; j < lenTempAreaI; j++ {
			tempArea[i][j] = scratchArea[i][j]
		}
	}
}

// func processLumber
// Handles everything needed to work out the size of the lumbar resources (day 18)
func processLumber(fileName string, yardX int, yardY int, minutes int, part byte) int {
	var woodedArea, lumberYards int = 0, 0
	var fingerPrint int = 0
	var openAreaFinger, woodedAreaFinger, lumberYardsFinger int = 0, 0, 0

	// Read contents of file into a string array
	fileContents, _ := readLines(fileName)
	//printStringArray(fileContents)

	lumberArea1 := make([][]byte, yardY)
	lumberArea2 := make([][]byte, yardY)
	fingerPrintRecord := make(map[int]int)

	for i := 0; i < yardY; i++ {
		lumberArea1[i] = make([]byte, yardX)	
		lumberArea2[i] = make([]byte, yardX)
	}

	readInitialState(fileContents, lumberArea1)

	// Loop through a round of life
	for i := 0; i < minutes; i++ {
		playRound(lumberArea1, lumberArea2)

		if part == 'b' {
			// Need to record the fingerprint in a map (openArea x trees x lumberYard?)
			// If we find that we have already seen that fingerprint, print it out then update
			// we can view the repeats easily then. Run over 10,000 iterations, work out the repeat
			// then scale up to the number of minutes we've been asked for

			// Generate fingerprint
			openAreaFinger = 0
			woodedAreaFinger = 0
			lumberYardsFinger = 0

			for i := 0; i < len(lumberArea1); i++ {
				for j := 0; j < len(lumberArea1[i]); j++ {
					if lumberArea1[i][j] == '|' {
						woodedAreaFinger++
					} else {
						if lumberArea1[i][j] == '#' {
							lumberYardsFinger++
						} else {
							if lumberArea1[i][j] == '.' {
								openAreaFinger++
							}
						}
					}
				}
			}

			fingerPrint = openAreaFinger * lumberYardsFinger * woodedAreaFinger

			result := fingerPrintRecord[fingerPrint]
			if result > 0 {
				fmt.Println("We have a match at:", i, result)
				fingerPrintRecord[fingerPrint] = i
			} else {
				fingerPrintRecord[fingerPrint] = i
			}

			// 1000 is 201341
			// Repeats every 28 cycles
			// Need for 1000000000
			// Need to work this out automatically but done for now

		}

	}

	for i := 0; i < len(lumberArea1); i++ {
		for j := 0; j < len(lumberArea1[i]); j++ {
			if lumberArea1[i][j] == '|' {
				woodedArea++
			} else {
				if lumberArea1[i][j] == '#' {
					lumberYards++
				}
			}
		}
	}

	return woodedArea * lumberYards
}

// Main routine
func main() {
	var yardX, yardY, minutes int = 0, 0, 0

	fileNamePtr := flag.String("file", "input1.txt", "A filename containing input strings")
	yardSizePtr := flag.String("grid", "10", "Size of the grid in squares")
	minutesPtr := flag.String("minutes", "10", "Number of minutes to play")
	execPartPtr := flag.String("part", "a", "Which part of day18 do you want to calc (a or b)")

	flag.Parse()

	yardX, _ = strconv.Atoi(*yardSizePtr)
	yardY = yardX
	minutes, _ = strconv.Atoi(*minutesPtr)

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Lumber Resource Value:", processLumber(*fileNamePtr, yardX, yardY, minutes, 'a'))
	case "b":
		fmt.Println("Part b - Lumber Resource Value:", processLumber(*fileNamePtr, yardX, yardY, minutes, 'b'))
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}