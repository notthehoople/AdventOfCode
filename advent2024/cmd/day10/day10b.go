package main

var ScorePartB int

// Check if the provided "pos" is a valid point to continue on
func lowEnoughPartB(pos Coord, previousHeight int, landscape [][]int) bool {
	maxRows := len(landscape)
	maxCols := len(landscape[0])

	if (pos.x >= 0 && pos.x < maxRows) && (pos.y >= 0 && pos.y < maxCols) {
		if landscape[pos.x][pos.y] == previousHeight+1 {
			return true
		}
	}
	return false
}

// Depth First Search
func dfsPartB(pos Coord, landscape [][]int, trailEnds *map[Coord]struct{}) {
	currentHeight := landscape[pos.x][pos.y]

	if currentHeight == 9 {
		(*trailEnds)[pos] = struct{}{}
		ScorePartB++
	}

	for _, dir := range Directions {
		nextPos := Coord{pos.x + dir.x, pos.y + dir.y}
		if lowEnoughPartB(nextPos, currentHeight, landscape) {
			dfsPartB(nextPos, landscape, trailEnds)
		}
	}
}

func calcTrailheadScoresPartB(landscape [][]int) int {
	maxRows := len(landscape)
	maxCols := len(landscape[0])

	for i := 0; i < maxRows; i++ {
		for j := 0; j < maxCols; j++ {
			// Check for a trailhead (0)
			if landscape[i][j] == 0 {
				trailEnds := make(map[Coord]struct{})
				dfsPartB(Coord{i, j}, landscape, &trailEnds)
			}
		}
	}

	return ScorePartB
}
