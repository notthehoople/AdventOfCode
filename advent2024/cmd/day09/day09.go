package main

import (
	"AdventOfCode-go/advent2024/utils"
	"fmt"
)

var Debug bool

func printFileMap(fileMap []int) {
	for _, i := range fileMap {
		if i == -1 {
			fmt.Printf(".")
		} else {
			fmt.Printf("%d", i%10)
		}
	}
	fmt.Printf("\n")
}

func calcChecksum(fileMap []int) int {
	var result int

	// Add up the result of multiplying each of the blocks' position with the file ID number it contains
	// If a block contains free space, skip it

	for pos, item := range fileMap {
		if item != -1 {
			result += pos * item
		}
	}

	return result
}

func compressMap(fileMap []int, part byte) []int {

	rightPos := len(fileMap) - 1

	for leftPos := 0; leftPos < len(fileMap); leftPos++ {
		if fileMap[leftPos] == -1 && rightPos > leftPos+1 {
			// Found a space and we've not met in the middle

			if Debug {
				fmt.Println("rightPos before:", rightPos)
			}
			for ; fileMap[rightPos] == -1; rightPos-- {

			}
			if Debug {
				fmt.Println("rightPos: after", rightPos)
			}
			fileMap[leftPos] = fileMap[rightPos]
			fileMap[rightPos] = -1
		}
	}

	return fileMap
}

func day09(filename string, part byte) int {

	// Our puzzleInput is a single line.
	// - First digit is the number of blocks that a file consists of
	// - Second digit is the number of empty blocks
	// - Repeat
	puzzleInput, _ := utils.ReadFile(filename)

	// Need to build out the filesystem including all the white space
	// Each file has a fileID that starts at 0 and ++ for each file. fileID probably doesn't matter

	var fileID int
	var isFile bool = true

	fileMap := make([]int, 0)

	for _, diskItem := range puzzleInput[0] {
		numBlocks := int(diskItem - '0')

		if isFile {
			if Debug {
				fmt.Printf("FileID: %d numBlocks: %d\n", fileID, numBlocks)
			}

			for i := 0; i < numBlocks; i++ {
				fileMap = append(fileMap, fileID)
			}

			isFile = false
			fileID++
		} else {

			for i := 0; i < numBlocks; i++ {
				fileMap = append(fileMap, -1)
			}
			isFile = true
		}
	}

	// Then compact files by moving files from the end of the disk to the leftmost free space
	// Continue until there are no gaps remaining between file blocks
	fileMap = compressMap(fileMap, part)

	if Debug {
		printFileMap(fileMap)
	}

	return calcChecksum(fileMap)
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()
	Debug = debug

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", day09(filenamePtr, execPart))
	}
}
