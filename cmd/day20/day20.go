package main

import (
	"fmt"
)

func getAllEdges(picture string) [8]string {
	var returnEdges [8]string
	var forwardPos int
	var flippedPos int

	// Top Edge
	// Top Edge flipped
	returnEdges[0] = picture[0:10]
	flippedTopEdge := []byte("..........")
	flippedPos = 0
	for i := 9; i >= 0; i-- {
		flippedTopEdge[flippedPos] = picture[i]
		flippedPos++
	}
	returnEdges[1] = string(flippedTopEdge)

	// Bottom Edge
	returnEdges[2] = picture[len(picture)-10 : len(picture)]
	// Bottom Edge flipped
	flippedBottomEdge := []byte("..........")
	flippedPos = 9
	for i := len(picture) - 10; i <= len(picture)-1; i++ {
		flippedBottomEdge[flippedPos] = picture[i]
		flippedPos--
	}
	returnEdges[3] = string(flippedBottomEdge)

	// Left Edge
	// Left Edge flipped
	forwardLeftEdge := []byte("..........")
	flippedLeftEdge := []byte("..........")

	forwardPos = 0
	flippedPos = 9
	for i := 0; i < len(picture); i += 10 {
		forwardLeftEdge[forwardPos] = picture[i]
		flippedLeftEdge[flippedPos] = picture[i]
		forwardPos++
		flippedPos--
	}
	returnEdges[4] = string(forwardLeftEdge)
	returnEdges[5] = string(flippedLeftEdge)

	// Right Edge
	// Right Edge flipped
	forwardRightEdge := []byte("..........")
	flippedRightEdge := []byte("..........")

	forwardPos = 0
	flippedPos = 9
	for i := 9; i < len(picture); i += 10 {
		forwardRightEdge[forwardPos] = picture[i]
		flippedRightEdge[flippedPos] = picture[i]
		forwardPos++
		flippedPos--
	}
	returnEdges[6] = string(forwardRightEdge)
	returnEdges[7] = string(flippedRightEdge)

	return returnEdges
}

func rotatePictures(filename string, part byte, debug bool) int {
	var tileNumber, matches int
	var picture string
	var edges map[string]int
	var pictures map[int]string
	var result int = 0

	puzzleInput, _ := readFile(filename)

	edges = make(map[string]int)
	pictures = make(map[int]string)

	// Pictures are 10 x 10 in size
	// Pictures are in a random order
	// Pictures are rotated randomly
	// 144 pictures so should be 12 x 12 arrangement
	//
	// Need to find a square arrangement
	// Each picture that fits with another has a shared border line that's the same
	// Need to rotate each picture to find those that match
	// Build a larger picture composes of smaller pictures
	//
	// Result is the multiplication of the IDs of the four corner tiles

	for _, line := range puzzleInput {
		matches, _ = fmt.Sscanf(line, "Tile %d:", &tileNumber)
		if len(line) == 0 {
			pictures[tileNumber] = picture
			picture = ""
			continue
		}

		if matches == 1 {
			if debug {
				fmt.Println("=========================")
				fmt.Println("Processing Tile:", tileNumber)
			}
		} else {
			// Processing the picture itself
			picture += line
		}
	}

	// Count matching edges of each tile in the edge map. Note that each tile will also match itself 4 times
	for _, tilePicture := range pictures {
		// Build the map of Edges
		var tmpok bool
		for _, checkEdge := range getAllEdges(tilePicture) {
			_, tmpok = edges[checkEdge]
			if !tmpok {
				edges[checkEdge] = 0
			} else {
				edges[checkEdge]++
			}

		}
	}

	var tileMatches int
	for tile, tilePicture := range pictures {
		// Count matches per tile

		tileMatches = 0
		var tmpok bool
		var tmpval int
		for _, checkEdge := range getAllEdges(tilePicture) {
			tmpval, tmpok = edges[checkEdge]
			if tmpok {
				tileMatches += tmpval
			}
		}
		// If the tile matches exactly 4 edges then we're in a corner so multiply the results together
		if tileMatches == 4 {
			if result == 0 {
				result = tile
			} else {
				result *= tile
			}
		}
	}
	return result
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
		fmt.Println("Part a answer:", rotatePictures(filenamePtr, execPart, debug))
	} else {
	}
}
