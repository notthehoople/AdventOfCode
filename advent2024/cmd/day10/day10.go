package main

import (
	"AdventOfCode-go/advent2024/utils"
	"fmt"
)

var Debug bool

type Coord struct {
	x int
	y int
}

var Directions = []Coord{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1}, // Up, Down, Left, Right
}

// Check if the provided "pos" is a valid point to continue on
func lowEnough(pos Coord, previousHeight int, landscape [][]int, visited [][]bool) bool {
	maxRows := len(landscape)
	maxCols := len(landscape[0])

	if (pos.x >= 0 && pos.x < maxRows) && (pos.y >= 0 && pos.y < maxCols) {
		if !visited[pos.x][pos.y] && landscape[pos.x][pos.y] == previousHeight+1 {
			return true
		}
	}
	return false
}

// Depth First Search
func dfs(pos Coord, landscape [][]int, visited [][]bool, trailEnds *map[Coord]struct{}) {
	visited[pos.x][pos.y] = true
	currentHeight := landscape[pos.x][pos.y]

	if currentHeight == 9 {
		(*trailEnds)[pos] = struct{}{}
	}

	for _, dir := range Directions {
		nextPos := Coord{pos.x + dir.x, pos.y + dir.y}
		if lowEnough(nextPos, currentHeight, landscape, visited) {
			dfs(nextPos, landscape, visited, trailEnds)
		}
	}
}

func calcTrailheadScores(landscape [][]int) int {
	maxRows := len(landscape)
	maxCols := len(landscape[0])
	totalScore := 0

	for i := 0; i < maxRows; i++ {
		for j := 0; j < maxCols; j++ {
			// Check for a trailhead (0)
			if landscape[i][j] == 0 {
				visited := make([][]bool, maxRows)
				for k := range visited {
					visited[k] = make([]bool, maxCols)
				}
				trailEnds := make(map[Coord]struct{})
				dfs(Coord{i, j}, landscape, visited, &trailEnds)
				totalScore += len(trailEnds)
			}
		}
	}

	return totalScore
}

func day10(filename string, part byte) int {

	puzzleInput, _ := utils.ReadFile(filename)

	landscape := make([][]int, len(puzzleInput))

	for key, puzzleLine := range puzzleInput {
		landscape[key] = make([]int, len(puzzleLine))

		for lineKey, lineItem := range puzzleLine {
			landscape[key][lineKey] = int(lineItem - '0')
		}
	}

	if Debug {
		utils.Print2DArrayInt(landscape)
	}

	if part == 'a' {
		return calcTrailheadScores(landscape)
	}

	return calcTrailheadScoresPartB(landscape)
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()
	Debug = debug

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", day10(filenamePtr, execPart))
	}
}
