package main

import (
	"AdventOfCode-go/advent2024/utils"
	"fmt"
)

var Debug bool

type blocks struct {
	pos       int
	numBlocks int
}

func printFileMap(fileMap []int) {
	for _, i := range fileMap {
		if i == -1 {
			fmt.Printf("-1 ")
		} else {
			fmt.Printf("%d ", i)
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

// part b - compressFullFileMap
// Attempt to move each file in its entirety. If it can't move as not enough free space, leave it
func compressFullFilesMap(fileMap []int, spaceList []blocks, fileSizes map[int]blocks, maxFileID int) []int {

	// Start at the right hand side. For each file found
	// - calculate the size of the file
	// - loop through the fileMap from the left side looking for enough space
	//		- if enough space found, move the file and zero the original location

	for fileID := maxFileID; fileID >= 0; fileID-- {

		var innerloopDone bool = false
		for i := 0; i < len(spaceList) && !innerloopDone; i++ {
			// Note: we ignore any spaces that are further to the right than the file position
			if spaceList[i].numBlocks >= fileSizes[fileID].numBlocks && spaceList[i].pos < fileSizes[fileID].pos {
				// Found an appropriate place for the size we're looking at
				for j := spaceList[i].pos; j < spaceList[i].pos+fileSizes[fileID].numBlocks; j++ {
					fileMap[j] = fileID
					fileMap[fileSizes[fileID].pos+(j-spaceList[i].pos)] = -1
				}

				// Have we used up all of the available space? If not then we need to record the remaining space
				fileSizes[fileID] = blocks{pos: spaceList[i].pos, numBlocks: fileSizes[fileID].numBlocks}
				spaceList[i].pos += fileSizes[fileID].numBlocks
				spaceList[i].numBlocks -= fileSizes[fileID].numBlocks

				innerloopDone = true
			}
		}
	}

	if Debug {
		printFileMap(fileMap)
	}

	return fileMap
}

func compressMap(fileMap []int) []int {

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
	fileSizes := make(map[int]blocks, 0)
	spaceList := make([]blocks, 0)
	var maxFileID int

	for _, diskItem := range puzzleInput[0] {
		numBlocks := int(diskItem - '0')

		if isFile {
			if Debug {
				fmt.Printf("FileID: %d numBlocks: %d\n", fileID, numBlocks)
			}
			// Needed for part b
			fileSizes[fileID] = blocks{pos: len(fileMap), numBlocks: numBlocks}
			maxFileID = fileID

			for i := 0; i < numBlocks; i++ {
				fileMap = append(fileMap, fileID)
			}

			isFile = false
			fileID++
		} else {
			spaceList = append(spaceList, blocks{pos: len(fileMap), numBlocks: numBlocks})

			for i := 0; i < numBlocks; i++ {
				fileMap = append(fileMap, -1)
			}
			isFile = true
		}
	}

	if part == 'a' {
		// Then compact files by moving files from the end of the disk to the leftmost free space
		// Continue until there are no gaps remaining between file blocks
		fileMap = compressMap(fileMap)
	} else {

		fileMap = compressFullFilesMap(fileMap, spaceList, fileSizes, maxFileID)
	}

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
