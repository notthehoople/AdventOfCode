package main

import (
	"fmt"
	//"strconv"
	//"strings"
)

// if blackTile is true, checking as a black tile so return false if there are 0 or more than 2 adjacent black tiles
// if blackTile is false, checking as a white tile so return true if there are exactly 2 adjacent black tiles
func checkTileNeighbours(blackTile bool, checkCoords coords, tilesMap map[coords]bool) bool {
	var blackTilesAdjacent int = 0
	var testCoords coords
	// Check nw(0,-1); ne(+1,-1);e(+1,0);se(0,+1);sw(-1,+1);w(-1,0)

	// nw(0,-1)
	testCoords = coords{qcoord: checkCoords.qcoord, rcoord: checkCoords.rcoord - 1}
	if tilesMap[testCoords] {
		blackTilesAdjacent++
	}

	// ne(+1,-1)
	testCoords = coords{qcoord: checkCoords.qcoord + 1, rcoord: checkCoords.rcoord - 1}
	if tilesMap[testCoords] {
		blackTilesAdjacent++
	}
	// e(+1,0)
	testCoords = coords{qcoord: checkCoords.qcoord + 1, rcoord: checkCoords.rcoord}
	if tilesMap[testCoords] {
		blackTilesAdjacent++
	}
	// se(0,+1)
	testCoords = coords{qcoord: checkCoords.qcoord, rcoord: checkCoords.rcoord + 1}
	if tilesMap[testCoords] {
		blackTilesAdjacent++
	}
	// sw(-1,+1)
	testCoords = coords{qcoord: checkCoords.qcoord - 1, rcoord: checkCoords.rcoord + 1}
	if tilesMap[testCoords] {
		blackTilesAdjacent++
	}
	// w(-1,0)
	testCoords = coords{qcoord: checkCoords.qcoord - 1, rcoord: checkCoords.rcoord}
	if tilesMap[testCoords] {
		blackTilesAdjacent++
	}

	// if blackTile is true we follow rules as a black tile
	if blackTile {
		if blackTilesAdjacent == 0 || blackTilesAdjacent > 2 {
			return false
		}
		return true
	}

	// otherwise we follow rules as a white tile

	if blackTilesAdjacent == 2 {
		return true
	}
	return false
}

func countLivingDisplay(filename string, days int, part byte, debug bool) int {
	var tilesMap map[coords]bool   // true means the tile is black
	var changesMap map[coords]bool // holds the changes

	tilesMap = make(map[coords]bool)
	changesMap = make(map[coords]bool)

	// Setup the day 0 floor tiles
	puzzleInput, _ := readFile(filename)

	for _, line := range puzzleInput {
		tileResult := flipATile(line, debug)
		if blackTile, ok := tilesMap[tileResult]; ok {
			// Already have an entry so we've previously flipped this tile
			if blackTile {
				tilesMap[tileResult] = false
			} else {
				tilesMap[tileResult] = true
			}
		} else {
			tilesMap[tileResult] = true
		}
	}

	var blackTileCount int = 0
	for _, blackTile := range tilesMap {
		if blackTile {
			blackTileCount++
		}
	}
	fmt.Printf("Day 0: %d\n", blackTileCount)

	var whiteCheckCoords coords
	for day := 1; day <= days; day++ {
		// Working from the tilesMap:
		// - any black tile with zero OR more than 2 black tiles adjacent to it flips to white
		// - any white tile with exactly 2 black tiles immediately adjacent to it flips to black
		// - all tile flips happen simultaneously
		for checkCoords := range tilesMap {
			// Check the neighbours of each black tile
			if tilesMap[checkCoords] {
				changesMap[checkCoords] = checkTileNeighbours(true, checkCoords, tilesMap)

				// Now check all around the black tile. If any tile is white, run a check on it for black tiles
				whiteCheckCoords = coords{checkCoords.qcoord, checkCoords.rcoord - 1}
				if !tilesMap[whiteCheckCoords] {
					changesMap[whiteCheckCoords] = checkTileNeighbours(false, whiteCheckCoords, tilesMap)
				}
				whiteCheckCoords = coords{checkCoords.qcoord + 1, checkCoords.rcoord - 1}
				if !tilesMap[whiteCheckCoords] {
					changesMap[whiteCheckCoords] = checkTileNeighbours(false, whiteCheckCoords, tilesMap)
				}
				whiteCheckCoords = coords{checkCoords.qcoord + 1, checkCoords.rcoord}
				if !tilesMap[whiteCheckCoords] {
					changesMap[whiteCheckCoords] = checkTileNeighbours(false, whiteCheckCoords, tilesMap)
				}
				whiteCheckCoords = coords{checkCoords.qcoord, checkCoords.rcoord + 1}
				if !tilesMap[whiteCheckCoords] {
					changesMap[whiteCheckCoords] = checkTileNeighbours(false, whiteCheckCoords, tilesMap)
				}
				whiteCheckCoords = coords{checkCoords.qcoord - 1, checkCoords.rcoord + 1}
				if !tilesMap[whiteCheckCoords] {
					changesMap[whiteCheckCoords] = checkTileNeighbours(false, whiteCheckCoords, tilesMap)
				}
				whiteCheckCoords = coords{checkCoords.qcoord - 1, checkCoords.rcoord}
				if !tilesMap[whiteCheckCoords] {
					changesMap[whiteCheckCoords] = checkTileNeighbours(false, whiteCheckCoords, tilesMap)
				}
			}
		}
		// Make all changes to the tilesMap simultaneously
		for tileKey, newTileValue := range changesMap {
			tilesMap[tileKey] = newTileValue
		}

		blackTileCount = 0
		for _, blackTile := range tilesMap {
			if blackTile {
				blackTileCount++
			}
		}
		fmt.Printf("Day %d: %d\n", day, blackTileCount)
	}

	return blackTileCount
}
