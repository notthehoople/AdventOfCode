package main

import (
	"AdventOfCode-go/advent2024/utils"
	"fmt"
)

var Debug bool

type lockAndKey struct {
	col1, col2, col3, col4, col5 int
}

func processKey(puzzleInput []string, row int) lockAndKey {
	var tempKey lockAndKey

	for i := row + 5; i >= row; i-- {
		if puzzleInput[i][0] == '#' {
			tempKey.col1++
		}
		if puzzleInput[i][1] == '#' {
			tempKey.col2++
		}
		if puzzleInput[i][2] == '#' {
			tempKey.col3++
		}
		if puzzleInput[i][3] == '#' {
			tempKey.col4++
		}
		if puzzleInput[i][4] == '#' {
			tempKey.col5++
		}
	}

	if Debug {
		fmt.Println("Key: ", tempKey)
	}
	return tempKey
}

func processLock(puzzleInput []string, row int) lockAndKey {
	var tempLock lockAndKey

	for i := row + 1; i < row+7; i++ {
		if puzzleInput[i][0] == '#' {
			tempLock.col1++
		}
		if puzzleInput[i][1] == '#' {
			tempLock.col2++
		}
		if puzzleInput[i][2] == '#' {
			tempLock.col3++
		}
		if puzzleInput[i][3] == '#' {
			tempLock.col4++
		}
		if puzzleInput[i][4] == '#' {
			tempLock.col5++
		}

	}

	if Debug {
		fmt.Println("Lock: ", tempLock)
	}
	return tempLock
}

func day25(filename string) int {
	var result int

	//Puzzle Input:
	//- Blocks of 7. Break between blocks of 7 are blank lines
	//- 	Lock: Top row is full '#', bottom row is empty '.'
	//- 	Key: Top row is all empty '.', bottom row is full '#'

	puzzleInput, _ := utils.ReadFile(filename)

	lockList := make([]lockAndKey, 0)
	keyList := make([]lockAndKey, 0)

	for row := 0; row < len(puzzleInput); {
		if len(puzzleInput[row]) == 0 {
			return (0)
		}
		if puzzleInput[row][0] == '#' {
			// We have found a lock
			lockList = append(lockList, processLock(puzzleInput, row))
			row += 8
			continue
		}
		// Must be a key
		keyList = append(keyList, processKey(puzzleInput, row))
		row += 8
	}

	for _, lock := range lockList {
		for _, key := range keyList {
			if lock.col1+key.col1 <= 5 && lock.col2+key.col2 <= 5 &&
				lock.col3+key.col3 <= 5 && lock.col4+key.col4 <= 5 && lock.col5+key.col5 <= 5 {
				result++
			}
			if Debug {
				fmt.Println(lock, key)
			}
		}
	}

	return result
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()
	Debug = debug

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", day25(filenamePtr))
	}
}
