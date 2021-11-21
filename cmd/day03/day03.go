package main

import (
	"aoc/advent2017/utils"
	"fmt"
	"strconv"
	"strings"
)

func calcSteps(inputData int, debug bool) int {
	/*
		Each square on the grid is allocated in a spiral pattern starting at a location
		marked 1 and then counting up while spiraling outward. For example, the first
		few squares are allocated like this:

		17  16  15  14  13
		18   5   4   3  12
		19   6   1   2  11
		20   7   8   9  10
		21  22  23---> ...

		While this is very space-efficient (no squares are skipped), requested data must
		be carried back to square 1 (the location of the only access port for this memory
		system) by programs that can only move up, down, left, or right. They always
		take the shortest path: the Manhattan Distance between the location of the data
		and square 1.
	*/

	/*
		Use a map to build the list. Each key is the number, then alter the coords as
		appropriate to spiral outwards
		-or-
		Use a 2d array to hold the positions
	*/
	var 
	var minValue int = 9999999
	var maxValue int = 0
	var currValue int

	for _, j := range strings.Fields(row) {
		currValue, _ = strconv.Atoi(j)
		if currValue < minValue {
			minValue = currValue
		}
		if currValue > maxValue {
			maxValue = currValue
		}
	}

	return maxValue - minValue
}

func solveSteps(inputData int, part byte, debug bool) int {
	var checksum int

	if part == 'a' {
		checksum = calcSteps(inputData, debug)
	}

	return checksum
}

// Main routine
func main() {
	_, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", solveSteps(368078, execPart, debug))
	}
}
