package main

import (
	"AdventOfCode-go/advent2023/utils"
	"fmt"
	"strings"
)

type galaxy struct {
	num int // identifier of the galaxy
	x   int // x co-ord of the galaxy
	y   int // y co-ord of the galaxy
}

func noGalaxiesHere(grid [][]byte, column int) bool {
	for y := 0; y < len(grid); y++ {
		if grid[y][column] != '.' {
			return false
		}
	}
	return true
}

func createGalaxies(skyGridExt [][]byte, rowsFound []int, columnsFound []int, increase int) []galaxy {
	var galaxyCount int = 1
	galaxyList := make([]galaxy, 0)

	for y := 0; y < len(skyGridExt); y++ {
		for x := 0; x < len(skyGridExt[y]); x++ {
			if skyGridExt[y][x] == '#' {
				// here we calculate the affect of the expanding galaxy
				// each time x is greater than the columnsFound entries, add increase to the x coord
				// each time y is greater than the rowsFound entries, add increase to the y coord
				var expansionX, expansionY int
				for _, i := range columnsFound {
					if x > i {
						// We're replacing the empty column with our increase, so remove the empty column by subtracting 1
						expansionX += increase - 1
					}
				}
				for _, i := range rowsFound {
					if y > i {
						// We're replacing the empty row with our increase, so remove the empty row by subtracting 1
						expansionY += increase - 1
					}
				}
				galaxyList = append(galaxyList, galaxy{galaxyCount, x + expansionX, y + expansionY})
				galaxyCount++
			}
		}
	}

	return galaxyList
}

//
//

func day11(filename string, part byte, debug bool, increase int) int {

	puzzleInput, _ := utils.ReadFile(filename)

	// Make a grid for the starting position to be built in

	skyGrid := make([][]byte, len(puzzleInput))
	for line, puzzleLine := range puzzleInput {
		skyGrid[line] = make([]byte, len(puzzleLine))
		for key, value := range puzzleLine {
			skyGrid[line][key] = byte(value)
		}
	}

	if debug {
		utils.Print2DArrayByte(skyGrid)
	}

	// Check for row being empty of Galaxies. If found, add to the list of empty rows for later use
	var rowsFound []int
	for i := 0; i < len(skyGrid); i++ {
		if strings.Contains(string(skyGrid[i]), "#") {
		} else {
			rowsFound = append(rowsFound, i)
		}
	}

	// Check for a column being empty of Galaxies. If found, add to the list of empty columns for later use
	var columnsFound []int
	for i := 0; i < len(skyGrid[0]); i++ {
		if noGalaxiesHere(skyGrid, i) {
			columnsFound = append(columnsFound, i)
		}
	}

	// Find the galaxies and build a list with an identifier and coords
	galaxyList := createGalaxies(skyGrid, rowsFound, columnsFound, increase)

	var sumLengths int
	for _, i := range galaxyList {
		for _, j := range galaxyList {
			if i != j {
				sumLengths += utils.ManhattanDistance2D(i.x, i.y, j.x, j.y)
			}
		}
	}

	if debug {
		fmt.Println("skyGrid:", len(skyGrid), len(skyGrid[0]))
	}

	return sumLengths / 2
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Result is: %d\n", day11(filenamePtr, execPart, debug, 2))
	case 'b':
		fmt.Printf("Result is: %d\n", day11(filenamePtr, execPart, debug, 1000000))

	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
