package main

import (
	"AdventOfCode-go/advent2023/utils"
	"fmt"
	"slices"
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

func extendColumnHere(column int, columnArray []int) bool {
	for _, val := range columnArray {
		if val == column {
			return true
		}
	}
	return false
}

func createGalaxies(skyGridExt [][]byte) []galaxy {
	var galaxyCount int = 1
	galaxyList := make([]galaxy, 0)

	for y := 0; y < len(skyGridExt); y++ {
		for x := 0; x < len(skyGridExt[y]); x++ {
			if skyGridExt[y][x] == '#' {
				galaxyList = append(galaxyList, galaxy{galaxyCount, x, y})
				galaxyCount++
			}
		}
	}

	return galaxyList
}

//
//

func day11(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)

	skyGrid := make([][]byte, len(puzzleInput))
	for line, puzzleLine := range puzzleInput {
		skyGrid[line] = make([]byte, len(puzzleLine))
		for key, value := range puzzleLine {
			skyGrid[line][key] = byte(value)
		}
	}

	// Check for row being empty of Galaxies. Add an extra line if an empty line is found

	if debug {
		utils.Print2DArrayByte(skyGrid)
	}

	for i := 0; i < len(skyGrid); i++ {
		if strings.Contains(string(skyGrid[i]), "#") {
		} else {
			skyGrid = slices.Insert(skyGrid, i, skyGrid[i])
			i++
		}
	}

	// Check for a column being empty of Galaxies

	var columnsFound []int
	for i := 0; i < len(skyGrid[0]); i++ {
		if noGalaxiesHere(skyGrid, i) {
			columnsFound = append(columnsFound, i)
		}
	}

	// Extend the skyGrid with the new columns
	skyGridExt := make([][]byte, len(skyGrid))
	for line, skyGridLine := range skyGrid {
		skyGridExt[line] = make([]byte, len(skyGrid[line])+len(columnsFound))
		var pos int
		for key, value := range skyGridLine {
			if extendColumnHere(key, columnsFound) {
				skyGridExt[line][pos] = '.'
				skyGridExt[line][pos+1] = byte(value)
				pos += 2
			} else {
				skyGridExt[line][pos] = byte(value)
				pos++
			}

		}
	}

	// Find the galaxies and build a list with an identifier and coords
	galaxyList := createGalaxies(skyGridExt)

	var sumLengths int
	for _, i := range galaxyList {
		for _, j := range galaxyList {
			if i != j {
				sumLengths += utils.ManhattanDistance2D(i.x, i.y, j.x, j.y)
			}
		}
	}

	if debug {
		fmt.Println(columnsFound)
		fmt.Println("skyGrid:", len(skyGrid), len(skyGrid[0]))
		fmt.Println("skyGridExt:", len(skyGridExt), len(skyGridExt[0]))
	}

	if debug {
		fmt.Println("====================================")
		utils.Print2DArrayByte(skyGridExt)
	}

	return sumLengths / 2
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Result is: %d\n", day11(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("Result is: %d\n", day11(filenamePtr, execPart, debug))
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
