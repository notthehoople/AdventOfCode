package main

import (
	"AdventOfCode-go/advent2022/utils"
	"fmt"
	"strconv"
)

type file struct {
	name string
	size int64
}

type directory struct {
	name   string
	dirs   []*directory
	files  []file
	parent *directory
	size   int64
}

const maxDirSizePartA = 100000
const totalDiskSpace = 70000000
const minDiskSpace = 30000000

func printDirectoryTree(directoryTree directory, debug bool) {
	fmt.Println("Directory:", directoryTree.name)
	for _, file := range directoryTree.files {
		fmt.Printf("---%s %d\n", file.name, file.size)
	}
	for _, dir := range directoryTree.dirs {
		fmt.Printf("Directory %s\n", dir.name)
		printDirectoryTree(*dir, debug)
	}
}

func buildDirectoryTree(input []string, debug bool) directory {
	var directoryTree directory
	var currDirectory *directory

	directoryTree.name = "/"
	currDirectory = &directoryTree

	for _, line := range input {
		// fmt.Println(line)

		if line[0] == '$' {
			// Command
			var command, action string
			fmt.Sscanf(line, "$ %s %s\n", &command, &action)
			switch command {
			case "cd":
				if debug {
					fmt.Printf("cd seen with %s\n", action)
				}
				if action == ".." {
					currDirectory = currDirectory.parent
				} else {
					// what do we do here?
					// check if we've seen the directory before
					for _, checkDir := range currDirectory.dirs {
						if checkDir.name == action {
							currDirectory = checkDir
						}
					}
				}

			case "ls":
				if debug {
					fmt.Println("ls seen")
				}
				// do we do anything here at all?
			}
		} else {
			var first, second string
			fmt.Sscanf(line, "%s %s\n", &first, &second)
			switch first {
			case "dir":
				// We've seen a directory. Note it's name but that's it at this stage
				// first = "dir", second = <name of directory>
				newDir := directory{name: second, parent: currDirectory}
				currDirectory.dirs = append(currDirectory.dirs, &newDir)
				if debug {
					fmt.Printf("dir seen: %s\n", second)
				}
			default:
				// Everything that isn't a directory is a file
				// first = <size of file>, second = <name of file>
				fileSize, _ := strconv.Atoi(first)
				newFile := file{name: second, size: int64(fileSize)}
				currDirectory.files = append(currDirectory.files, newFile)
				increaseDirSize(currDirectory, newFile.size)
				if debug {
					fmt.Printf("File seen: %s with size %d\n", second, fileSize)
				}
			}

		}
	}
	return directoryTree
}

func increaseDirSize(currDirectory *directory, size int64) {
	currDirectory.size += size
	if currDirectory.parent != nil {
		increaseDirSize(currDirectory.parent, size)
	}
}

// Part A: calculate directories less than maxDirSizePartA (100000)
func calcSmallDirs(dir *directory) int64 {
	var totalSize int64

	if dir.size <= maxDirSizePartA {
		totalSize += dir.size
	}

	for _, subDir := range dir.dirs {
		tempSize := calcSmallDirs(subDir)
		totalSize += tempSize
	}

	return totalSize
}

// Part B: What to delete to save enough space for the update?
func amountToDelete(dir *directory, spaceToSave int64) int64 {

	//	fmt.Println("Current space available:", spaceToSave)
	//	fmt.Println("Space to save:", minDiskSpace-spaceToSave)

	for _, subDir := range dir.dirs {
		if subDir.size > spaceToSave {
			fmt.Printf("Name: %s Size: %d\n", subDir.name, subDir.size)
		}
		amountToDelete(subDir, spaceToSave)
	}

	return 0
}

func processFileSystem(filename string, part byte, debug bool) int64 {

	puzzleInput, _ := utils.ReadFile(filename)

	directoryTree := buildDirectoryTree(puzzleInput, debug)

	if debug {
		printDirectoryTree(directoryTree, debug)
	}

	if part == 'a' {
		return calcSmallDirs(&directoryTree)
	} else {
		var spaceToSave int64 = minDiskSpace - (totalDiskSpace - directoryTree.size)

		return amountToDelete(&directoryTree, spaceToSave)
	}
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Result is: %d\n", processFileSystem(filenamePtr, execPart, debug))
	case 'b':
		fmt.Println("Choose the smallest from the following results:")
		fmt.Printf("Result is: %d\n", processFileSystem(filenamePtr, execPart, debug))
	case 'z':
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
