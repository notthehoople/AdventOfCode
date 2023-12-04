package main

import (
	"AdventOfCode-go/advent2023/utils"
	"fmt"
)

func isNumber(item byte) bool {
	if item >= '0' && item <= '9' {
		return true
	}
	return false
}

func isSymbol(item byte) bool {
	if (item >= 35 && item <= 45) || item == 47 || item == 61 || item == 64 {
		return true
	}
	return false
}

func isGear(item byte) bool {
	return item == 42
}

func expandNumber(engine [][]byte, y int, x int) int {
	var number int

	if isNumber(engine[y][x]) {
		if isNumber(engine[y][x+1]) {
			if isNumber(engine[y][x+2]) {
				number = int(engine[y][x+2] - '0')
				number += int(engine[y][x+1]-'0') * 10
				number += int(engine[y][x]-'0') * 100

				return number
			} else {
				if isNumber(engine[y][x-1]) {
					number = int(engine[y][x+1] - '0')
					number += int(engine[y][x]-'0') * 10
					number += int(engine[y][x-1]-'0') * 100

					return number
				} else {
					number = int(engine[y][x+1] - '0')
					number += int(engine[y][x]-'0') * 10

					return number
				}
			}
		} else {
			if isNumber(engine[y][x-1]) {
				if isNumber(engine[y][x-2]) {
					number = int(engine[y][x-2]-'0') * 100
					number += int(engine[y][x-1]-'0') * 10
					number += int(engine[y][x] - '0')

					return number
				} else {
					number = int(engine[y][x-1]-'0') * 10
					number += int(engine[y][x] - '0')

					return number
				}
			}
		}
		number += int(engine[y][x] - '0')

		return number
	}

	return number
}

func checkForNumber(item byte, engine [][]byte, y int, x int, part byte) int {

	var number int
	var result int
	var foundNumbers []int
	var count int // in part b, we only care about gears that have two numbers next to a '*'

	// For UP: check y-1, x:
	// - If it's a '.' then call TWICE, once for y-1, x-1 and once for y-1, x+1;
	// - If it's a number then call ONCE for y-1, x
	if isNumber(engine[y-1][x]) {
		result = expandNumber(engine, y-1, x)
		if result > 0 {
			number += result
			foundNumbers = append(foundNumbers, result)
			count++
		}
	} else {
		result = 0
		result = expandNumber(engine, y-1, x-1)
		if result > 0 {
			number += result
			foundNumbers = append(foundNumbers, result)
			count++
		}

		result = 0
		result = expandNumber(engine, y-1, x+1)
		if result > 0 {
			number += result
			foundNumbers = append(foundNumbers, result)
			count++
		}

	}

	// For LEFT: check y, x-1:
	// - If number then it's a match. Build number out to the left until find another '.' then return number
	if isNumber(engine[y][x-1]) {
		result = 0
		result = expandNumber(engine, y, x-1)
		if result > 0 {
			number += result
			foundNumbers = append(foundNumbers, result)
			count++
		}

	}

	// For RIGHT: check y, x+1.
	// - If number then it's a match. Build number out to the right until find another '.' then return number
	if isNumber(engine[y][x+1]) {
		result = 0
		result = expandNumber(engine, y, x+1)
		if result > 0 {
			number += result
			foundNumbers = append(foundNumbers, result)
			count++
		}
	}

	// For DOWN: check y+1, x:
	// - If it's a '.' then call TWICE, once for y+1, x-1 and once for y+1, x+1;
	// - If it's a number then call ONCE for y+1, x
	if isNumber(engine[y+1][x]) {
		result = 0
		result = expandNumber(engine, y+1, x)
		if result > 0 {
			number += result
			foundNumbers = append(foundNumbers, result)
			count++
		}
	} else {
		result = 0
		result = expandNumber(engine, y+1, x-1)
		if result > 0 {
			number += result
			foundNumbers = append(foundNumbers, result)
			count++
		}

		result = 0
		result = expandNumber(engine, y+1, x+1)
		if result > 0 {
			number += result
			foundNumbers = append(foundNumbers, result)
			count++
		}
	}

	if part == 'a' {
		return number
	}

	if count == 2 {
		return foundNumbers[0] * foundNumbers[1]
	}

	return 0
}

// Any number adjacent to a symbol, even diagonally, is a "part number" and should be
// included in the sum. Periods '.' do not count as a symbol.
// Numbers can be 1, 2 or 3 digits and are 0 to 9 (ASCII 48 to 57)
// Symbols: # $ % & * + - / @ = (ASCII 35 to 47 excluding .(46), 64, 61)
// %s;[#$%&*+\-@=/];;g (NOTE: - removes numbers so needs escape)

func day03(filename string, part byte, debug bool) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)

	engine := make([][]byte, len(puzzleInput))
	for i := range engine {
		engine[i] = make([]byte, len(puzzleInput[0]))
	}

	for line, puzzleLine := range puzzleInput {
		engine[line] = []byte(puzzleLine)
	}

	if debug {
		utils.Print2DArrayByte(engine)
	}

	padding := make([][]byte, len(puzzleInput)+2)
	for i := range padding {
		padding[i] = make([]byte, len(puzzleInput[0])+2)
	}

	for j := 0; j < len(padding); j++ {
		for i := 0; i < len(padding[j]); i++ {
			if j == 0 || j == len(padding)-1 {
				padding[j][i] = '.'
			} else if i == 0 || i == len(padding[j])-1 {
				padding[j][i] = '.'
			} else {
				padding[j][i] = engine[j-1][i-1]
			}
		}

	}

	if debug {
		fmt.Println("==================================================================")
		utils.Print2DArrayByte(padding)
	}

	for y, engineLine := range padding {

		for x := 0; x < len(engineLine); x++ {
			item := engineLine[x]

			if part == 'a' {
				if isSymbol(item) {
					//fmt.Println("Found symbol:", item)
					// Now check for number above, to the left, to the right and below
					// If number found then add the number to the result
					// Edge Case: how do we deal with numbers that are next to TWO or more symbols?
					// NOTE: numbers appear at least twice in the input file so can't use a map
					result += checkForNumber(item, padding, y, x, part)
				}
			} else {
				if isGear(item) {
					result += checkForNumber(item, padding, y, x, part)
				}
			}
		}
	}

	return result
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Result is: %d\n", day03(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("Result is: %d\n", day03(filenamePtr, execPart, debug))
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
