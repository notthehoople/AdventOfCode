package main

import (
	"AdventOfCode-go/advent2023/utils"
	"fmt"
	"strings"
)

// File: First line is the movement instructions. Collection of L or R letters to signify movement
//    Rest of the file contains direction nodes which consist of a left element and a right element (BBB, CCC)
//    When running the movement phase the L or R command tells you which element to pick from where you are located

type locNode struct {
	left  string
	right string
}

func day08b(filename string, part byte, debug bool) int {

	locIndex := make(map[string]locNode)
	var currLoc []string

	puzzleInput, _ := utils.ReadFile(filename)

	var movements string
	for line, puzzleLine := range puzzleInput {
		if line == 0 {
			movements = puzzleLine
			continue
		}
		if len(puzzleLine) == 0 {
			continue
		}

		var location, left, right string
		fields := strings.Fields(puzzleLine)
		location = fields[0]
		left = fields[2][1:4]
		right = fields[3][0:3]

		locIndex[location] = locNode{left, right}
		if strings.HasSuffix(location, "A") {
			currLoc = append(currLoc, location)
		}
	}

	if debug {
		fmt.Println("movements:", movements)
		fmt.Println("LocIndex:", locIndex)
		fmt.Println("currLoc:", currLoc)
	}

	if debug {
		fmt.Println("All starting positions:", currLoc)
	}

	var currStep int
	ghostResults := make([]int, len(currLoc))

	// Although the puzzle says we should move all at the same time, we're not going to do that
	// Loops through each of the starting points, looking for when that start reaches a "Z" end point
	// Once we get to the end, record the number of steps it took to get there

	for ghostIndex, ghostLoc := range currLoc {
		currStep = 0
		for while := false; !while; {
			currMove := movements[currStep%len(movements)]
			if debug {
				fmt.Println("currMove:", currMove)
			}

			switch currMove {
			case 'L':
				ghostLoc = locIndex[ghostLoc].left
			case 'R':
				ghostLoc = locIndex[ghostLoc].right
			default:
				panic("Bad move in the movements list")
			}

			if strings.HasSuffix(ghostLoc, "Z") {
				// this one is done. Record it's position for later LCM use
				while = true
			}

			currStep++

		}
		ghostResults[ghostIndex] = currStep
	}

	// Now we have the number of steps needed for each of the starting positions to reach the end
	// position, we use Least Common Multiplier to work out the point at which the results for all converge
	var totalResult int64
	totalResult = int64(ghostResults[0])
	for i := 1; i < len(ghostResults); i++ {
		totalResult = LCM(totalResult, int64(ghostResults[i]))
	}

	return int(totalResult)
}

/*
func day08b(filename string, part byte, debug bool) int {

	locIndex := make(map[string]locNode)
	var currLoc []string

	puzzleInput, _ := utils.ReadFile(filename)

	var movements string
	for line, puzzleLine := range puzzleInput {
		if line == 0 {
			movements = puzzleLine
			continue
		}
		if len(puzzleLine) == 0 {
			continue
		}

		var location, left, right string
		fields := strings.Fields(puzzleLine)
		location = fields[0]
		left = fields[2][1:4]
		right = fields[3][0:3]

		locIndex[location] = locNode{left, right}
		if strings.HasSuffix(location, "A") {
			fmt.Println(location[2:3])
			fmt.Printf("Location %s is starting point\n", location)
			currLoc = append(currLoc, location)
		}
	}

	if debug {
		fmt.Println(movements)
		fmt.Println(locIndex)
	}

	fmt.Println(currLoc)

	var currStep int
	var totalSteps int

	for while := false; !while; {
		currMove := movements[currStep%len(movements)]
		if debug {
			fmt.Println("currLoc:", currLoc)
		}

		switch currMove {
		case 'L':
			for index, ghostLoc := range currLoc {
				currLoc[index] = locIndex[ghostLoc].left
			}
		case 'R':
			for index, ghostLoc := range currLoc {
				currLoc[index] = locIndex[ghostLoc].right
			}
		default:
			panic("Bad move in the movements list")
		}

		totalSteps++
		var allDone int
		for _, ghostLoc := range currLoc {
			if strings.HasSuffix(ghostLoc, "Z") {
				allDone++
			}
		}

		if allDone == len(currLoc) {
			while = true
		}

		currStep++
		if currStep%12599 == 0 {
			fmt.Println(currStep, currLoc)
		}
	}

	return totalSteps
}
*/

//GCD greatest common divisor (GCD) via Euclidean algorithm
// Code lifted from Go Playground: https://play.golang.org/p/SmzvkDjYlb
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

//LCM find Least Common Multiple (LCM) via GCD
// Code lifted from Go Playground: https://play.golang.org/p/SmzvkDjYlb
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func day08a(filename string, part byte, debug bool) int {

	locIndex := make(map[string]locNode)

	puzzleInput, _ := utils.ReadFile(filename)

	var movements string
	for line, puzzleLine := range puzzleInput {
		if line == 0 {
			movements = puzzleLine
			continue
		}
		if len(puzzleLine) == 0 {
			continue
		}

		var location, left, right string
		fields := strings.Fields(puzzleLine)
		location = fields[0]
		left = fields[2][1:4]
		right = fields[3][0:3]

		locIndex[location] = locNode{left, right}
		if debug {
			fmt.Printf("location: %s left: %s right: %s map: %v\n", location, left, right, locIndex[location])
		}
	}

	if debug {
		fmt.Println(movements)
		fmt.Println(locIndex)
	}

	var currStep int
	var currLoc string = "AAA"
	var totalSteps int

	for while := false; !while; {
		currMove := movements[currStep%len(movements)]
		if debug {
			fmt.Println("currLoc:", currLoc)
		}

		switch currMove {
		case 'L':
			currLoc = locIndex[currLoc].left
		case 'R':
			currLoc = locIndex[currLoc].right
		default:
			panic("Bad move in the movements list")
		}

		totalSteps++
		if currLoc == "ZZZ" {
			while = true
		}

		currStep++
	}

	return totalSteps
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Result is: %d\n", day08a(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("Result is: %d\n", day08b(filenamePtr, execPart, debug))
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
