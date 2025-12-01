package main

import (
	"AdventOfCode-go/advent2025/utils"
	"fmt"
)

func day01(filename string, part byte, debug bool) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)

	direction := 'L'
	steps := 0
	position := 50
	for _, puzzleLine := range puzzleInput {
		fmt.Sscanf(puzzleLine, "%c%d\n", &direction, &steps)
		if debug {
			fmt.Printf("%c %d\n", direction, steps)
		}

		if part == 'a' {
			if direction == 'L' {
				position = (position - steps%100 + 100) % 100
			} else {
				position = (position + steps) % 100
			}

			if position == 0 {
				result++
			}
		} else { // part b
			// Couldn't work out the clever way of doing this so gone with brute force

			var dir, oldposition int
			if direction == 'L' {
				dir = -1
			} else {
				dir = 1
			}

			oldposition = position
			if steps >= 100 {
				result += steps / 100
			}

			position += (steps % 100) * dir
			// if we're going over the top then we've crossed 0. Only count it if we didn't end up on 0 as that's handled later
			if position > 99 {
				position -= 100
				if position != 0 && oldposition != 0 {
					result++
				}
			} else if position < 0 {
				// if we've gone under 0 then we've crossed 0. Only count it if we didn't end up on 0 as that's handled later
				position += 100
				if position != 0 && oldposition != 0 {
					result++
				}
			}
			// If we've landed on 0, increment "crossings"
			if position == 0 {
				result++
			}
		}

		if debug {
			fmt.Println("Position: ", position)
		}
	}

	if part == 'a' {
		// Part 1: Find the number of times the position is set to 0
		return result
	}

	// Part B - find the similarity score between the two lists

	return result
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", day01(filenamePtr, execPart, debug))
	}
}
