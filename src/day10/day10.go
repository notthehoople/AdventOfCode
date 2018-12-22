package main

import (
	"fmt"
	"bufio"
	"os"
	"flag"
	"strconv"
)

type starMeta struct {
    Xcoord	int
    Ycoord	int
	Xvel	int
	Yvel	int
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

func printGridArea(starPoints []starMeta) {
	var minX, minY, maxX, maxY int = 0, 0, 0, 0

	minX = starPoints[0].Xcoord
	minY = starPoints[0].Ycoord

	for i := 0; i < len(starPoints); i++ {
		if starPoints[i].Xcoord < minX {
			minX = starPoints[i].Xcoord
		}
		
		if starPoints[i].Xcoord > maxX {
			maxX = starPoints[i].Xcoord
		}

		if starPoints[i].Ycoord < minY {
			minY = starPoints[i].Ycoord
		}
		
		if starPoints[i].Ycoord > maxY {
			maxY = starPoints[i].Ycoord
		}
	}

	// Now we know the size of the grid, let's build it and print it out
	gridArea := make([][]int, maxY)
	for i := 0; i < len(gridArea); i++ {
		gridArea[i] = make([]int, maxX)
	}

	for i := 0; i < len(gridArea); i++ {
		for j := 0; j < len(gridArea[i]); j++ {
			gridArea[i][j] = '.'
		}
	}

	for i := 0; i < len(starPoints); i++ {
		//gridArea[starPoints[i].Ycoord - minY][starPoints[i].Xcoord - minX] = '#'
		gridArea[starPoints[i].Xcoord - minX][starPoints[i].Ycoord - minY] = '#'
	}

	for i := 0; i < len(gridArea); i++ {
		for j := 0; j < len(gridArea[i]); j++ {
			fmt.Printf("%c", gridArea[i][j])
		}
		fmt.Printf("\n")
	}
}

func detectMessage(tempStarPoints []starMeta) int {
	var bestXCoord, bestYCoord int = 0, 0

	xCoordMap := make(map[int]int)
	yCoordMap := make(map[int]int)

	for i := 0; i < len(tempStarPoints); i++ {
		//fmt.Println("tempStarPoints:", tempStarPoints[i].Xcoord)
		xCoordMap[tempStarPoints[i].Xcoord]++
		yCoordMap[tempStarPoints[i].Ycoord]++
	}

	for _, value := range xCoordMap {
		if value > bestXCoord {
			bestXCoord = value
		}
	}

	for _, value := range yCoordMap {
		if value > bestYCoord {
			bestYCoord = value
		}
	}

	return bestXCoord * bestYCoord
}

func processMessage(fileName string, gridSize int, printSecond int, maxSeconds int, safeSize int, part byte) string {
	var starPoints []starMeta

	// Read contents of file into a string array
	fileContents, _ := readLines(fileName)

	// Read contents of file into an array of "starMeta" struct
	for i := 0; i < len(fileContents); i++ {
		var a, b, c, d int = 0,0,0,0

		fmt.Sscanf(fileContents[i], "position=<%d,%d> velocity=<%d,%d>\n", &a, &b, &c, &d)

		starPoints = append(starPoints, starMeta{Xcoord:a, Ycoord:b, Xvel:c, Yvel:d})
	}
	
	// Run for as many seconds as we've been asked to
	for i := 0; i < maxSeconds; i++ {

		// Apply the movement to each point in starPoints
		for j := 0; j < len(starPoints); j++ {
			starPoints[j].Xcoord += starPoints[j].Xvel
			starPoints[j].Ycoord += starPoints[j].Yvel
		}

		// Print out our estimate of whether the message is going to be there
		fmt.Printf("Time: %d Message likelihood: %d\n", i, detectMessage(starPoints))

		// If we've been asked to print at this second, let's print!
		if printSecond > 0 && i == printSecond {
			fmt.Println("Printing")
			printGridArea(starPoints)
			fmt.Println("Printing Done")
		}
	}

	return "Nothing"
}

// Main routine
func main() {
	var gridSize, printSecond, maxSeconds, safeSize int = 0, 0, 0, 0

	fileNamePtr := flag.String("file", "input1.txt", "A filename containing input strings")
	gridSizePtr := flag.String("grid", "100", "Size of the grid in squares")
	printSecondPtr := flag.String("printsecond", "0", "Print the grid output for time <x> second")
	maxSecondsPtr := flag.String("max", "5", "Maximum seconds to run the message predictor")
	safeSizePtr := flag.String("safe", "1000", "Modifier to make things safer")
	execPartPtr := flag.String("part", "a", "Which part of day10 do you want to calc (a or b)")

// From the file, we grab the starting position AND the velocity of each star
//     Q: how to maintain what the velocity is for each item in the grid (do I use a grid even?)
//     Q: how do I decide how big to make the game board array
// Ideas: if I analyse the number of vert and horiz adjacent characters I can probably detect when the stars align
//        this means I'll need a "print the display at time <x> seconds"
//
// Structure: map of type struct.
//            struct is Xcoord, Ycoord, Xvel, Yvel

	flag.Parse()

	gridSize, _ = strconv.Atoi(*gridSizePtr)
	printSecond, _ = strconv.Atoi(*printSecondPtr)
	maxSeconds, _ = strconv.Atoi(*maxSecondsPtr)
	safeSize, _ = strconv.Atoi(*safeSizePtr)

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Message in the stars:", processMessage(*fileNamePtr, gridSize, printSecond, maxSeconds, safeSize, 'a'))
	case "b":
		fmt.Println("Part b - Not here yet")
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}