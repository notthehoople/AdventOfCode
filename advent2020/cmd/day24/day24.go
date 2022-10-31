package main

import (
	"fmt"
	//"strconv"
	//"strings"
)

// Using Axial coordinates (from https://www.redblobgames.com/grids/hexagons/)
type coords struct {
	qcoord int
	rcoord int
}

func flipATile(routeToTile string, debug bool) (flippedTile coords) {
	var qcoord, rcoord int
	// instructions come in a non-delimited series of directions
	// - e, se, sw, w, nw, ne
	// - all instructions given from a central reference tile
	// - each line is directions to a tile to be flipped

	for i := 0; i < len(routeToTile); {
		switch routeToTile[i] {
		case 'e':
			qcoord++
			i++
			break
		case 's':
			if routeToTile[i+1] == 'e' {
				rcoord++
			} else {
				// must be 'sw'
				qcoord--
				rcoord++
			}
			i += 2
			break
		case 'w':
			qcoord--
			i++
			break
		case 'n':
			if routeToTile[i+1] == 'e' {
				qcoord++
				rcoord--
			} else {
				// must be 'nw'
				rcoord--
			}
			i += 2
			break
		default:
			panic("Corrupt input in flipTile")
		}

	}
	flippedTile.qcoord = qcoord
	flippedTile.rcoord = rcoord

	return flippedTile
}

func countFlippedTiles(filename string, part byte, debug bool) int {
	var tilesMap map[coords]bool // true means the tile is black

	tilesMap = make(map[coords]bool)
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
	return blackTileCount
}

// Main routine
func main() {
	filenamePtr, execPart, debug, test := catchUserInput()

	if test {
		return
	}

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else if execPart == 'a' {
		if debug {
			fmt.Println("----------------")
			fmt.Println("esenee:", flipATile("esenee", debug))
			fmt.Println("----------------")
			fmt.Println("esew:", flipATile("esew", debug))
			fmt.Println("----------------")
			fmt.Println("nwwswee:", flipATile("nwwswee", debug))
			fmt.Println("----------------")
			fmt.Println("sesenwnenenewseeswwswswwnenewsewsw:", flipATile("sesenwnenenewseeswwswswwnenewsewsw", debug))
		}
		fmt.Println("Flipped black tiles:", countFlippedTiles(filenamePtr, execPart, debug))

	} else {
		fmt.Println("Flipped black tiles:", countLivingDisplay(filenamePtr, 100, execPart, debug))
	}
}
