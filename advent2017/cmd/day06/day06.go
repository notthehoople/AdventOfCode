package main

import (
	"AdventOfCode-go/advent2017/utils"
	"fmt"
	"strconv"
	"strings"
)

func newMapKey(valueList []int) string {
	var mapKey string

	for _, value := range valueList {
		if len(mapKey) == 0 {
			mapKey = fmt.Sprintf("%d", value)
		} else {
			mapKey = mapKey + "-" + fmt.Sprintf("%d", value)
		}
	}
	return mapKey
}

func debugMemory(filename string, part byte, debug bool) int {
	seenPatternBefore := make(map[string]bool)

	puzzleInput, _ := utils.ReadFile(filename)
	puzzleInputSplit := strings.Fields(puzzleInput[0])

	memoryBanks := make([]int, len(puzzleInputSplit))

	for i := 0; i < len(puzzleInputSplit); i++ {
		memoryBanks[i], _ = strconv.Atoi(puzzleInputSplit[i])
	}

	if debug {
		fmt.Println(memoryBanks)
	}

	var maxPos, maxBlocks, currentPos int
	for counter := 1; ; counter++ {
		maxPos = 0
		maxBlocks = 0

		// Find the memory bank with the most blocks
		for i := 0; i < len(memoryBanks); i++ {
			if memoryBanks[i] > maxBlocks {
				maxPos = i
				maxBlocks = memoryBanks[i]
			}
		}

		if debug {
			fmt.Printf("MaxPos: %d MaxBlocks: %d\n", maxPos, maxBlocks)
		}

		// Distribute the blocks across the memory banks, starting at the next bank
		// following the (former) largest

		memoryBanks[maxPos] = 0
		currentPos = maxPos + 1
		for blocks := maxBlocks; blocks > 0; blocks-- {
			// If we're at the end of the memoryBanks loop back to the beginning
			if currentPos >= len(memoryBanks) {
				currentPos = 0
			}
			memoryBanks[currentPos]++
			currentPos++
		}

		mapKey := newMapKey(memoryBanks)
		_, ok := seenPatternBefore[mapKey]
		if ok {
			// Seen pattern before so exit
			return counter
		} else {
			seenPatternBefore[mapKey] = true
		}
	}
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", debugMemory(filenamePtr, execPart, debug))
	}
}
