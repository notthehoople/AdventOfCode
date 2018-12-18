package main

import (
	"fmt"
	"bufio"
	"os"
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
func checkArea(tempArea [][]byte, yCoord int, xCoord int) (int, int, int) {
	lumberYards = 0
	openAcre = 0
	trees = 0

	// To consider: the edges. So if xCoord / yCoord == 0 or xCoord / yCoord = len(tempArea)

	// x x   x o   o x   x x   x o x   x x x   x x   x x
	// o x   x x   x x   x o   x x x   x o x   o x   x o
	//                                          x x   x x
}


// func playRound
// Play a round of Life!
func playRound(tempArea [][]byte) {
	var lumberYards, openAcre, trees int = 0, 0, 0

	tempCalcSpace := make([][]byte, len(tempArea))

	for i := 0; i < len(tempArea); i++ {
		tempCalcSpace[i] = make([]byte, len(tempArea[i]))	
	}

	for i := 0; i < len(tempArea); i++ {
		for j := 0; j < len(tempArea[i]); j++ {
			lumberYards, openAcre, trees = checkArea(tempArea, i, j)

			// An open acre ('.') will become filled with trees if three or more adjacent acres contained 
			// trees. Otherwise, nothing happens.
			if tempArea[i][j] == '.' && trees > 2 {
				tempCalcSpace[i][j] = '|'
			} else {
				tempCalcSpace[i][j] = '.'
			}

			// An acre filled with trees ('|') will become a lumberyard if three or more adjacent acres were
			//    lumberyards. Otherwise, nothing happens.
			if tempArea[i][j] == '|' && lumberYards > 2 {
				tempCalcSpace[i][j] = '#'
			} else {
				tempCalcSpace[i][j] = '|'
			}

			// An acre containing a lumberyard ('#') will remain a lumberyard if it was adjacent to at least 
			//    one other lumberyard and at least one acre containing trees. Otherwise, it becomes open.
			if tempArea[i][j] == '#' && lumberYards > 0 && trees > 0 {
				tempCalcSpace[i][j] = '#'
			} else {
				tempCalcSpace[i][j] = '.'
			}
		}
	}

	print2DSlice(tempCalcSpace)
}

// func processLumber
// Handles everything needed to work out the size of the lumbar resources (day 18)
func processLumber(fileName string, yardX int, yardY int, minutes int, part byte) int {
	
	// Read contents of file into a string array
	fileContents, _ := readLines(fileName)
	printStringArray(fileContents)

	lumberArea1 := make([][]byte, yardY)
	lumberArea2 := make([][]byte, yardY)

	for i := 0; i < yardY; i++ {
		lumberArea1[i] = make([]byte, yardX)	
		lumberArea2[i] = make([]byte, yardX)
	}

	readInitialState(fileContents, lumberArea1)

	//for t := 0; t < minutes; t++ {
	//
	//}

	playRound(lumberArea1)

	print2DSlice(lumberArea1)

	return 0
}

// Main routine
func main() {
	var yardX, yardY, minutes int = 0, 0, 0

	fileNamePtr := flag.String("file", "input1.txt", "A filename containing input strings")
	yardSizePtr := flag.String("grid", "10", "Size of the grid in squares")
	minutesPtr := flag.String("minutes", "10", "Number of minutes to play")
	execPartPtr := flag.String("part", "a", "Which part of day18 do you want to calc (a or b)")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		yardX, _ = strconv.Atoi(*yardSizePtr)
		yardY = yardX
		minutes, _ = strconv.Atoi(*minutesPtr)
		fmt.Println("Part a - Lumber Resource Value:", processLumber(*fileNamePtr, yardX, yardY, minutes, 'a'))
	case "b":
		fmt.Println("Part b - Not there yet")
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}