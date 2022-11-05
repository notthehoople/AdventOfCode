package main

import (
	"AdventOfCode-go/advent2017/utils"
	"fmt"
	"strings"
)

type program struct {
	name     string
	weight   int
	parent   string // Note this could be a pointer to an entry in the map
	children []string
}

func processInputTower(filename string, part byte, debug bool) map[string]program {

	puzzleInput, _ := utils.ReadFile(filename)
	tower := make(map[string]program, 0)

	var programName string
	var programWeight int
	for i := 0; i < len(puzzleInput); i++ {
		var currentProgram program

		components := strings.Split(puzzleInput[i], "->")
		fmt.Sscanf(components[0], "%s (%d)", &programName, &programWeight)

		currentProgram.name = programName
		currentProgram.weight = programWeight

		if debug {
			fmt.Println("First component:", components[0])
			fmt.Printf("Program Name: %s Program Weight: %d\n", programName, programWeight)
		}

		if len(components) > 1 {
			if debug {
				fmt.Println("We have children to process")
			}

			children := strings.Fields(components[1])
			for _, j := range children {
				child := strings.Trim(j, ",")
				if childRecord, ok := tower[child]; ok {
					// This child already exists in the tower. Set the parent of this child to the record being processed
					childRecord.parent = currentProgram.name
					tower[child] = childRecord
				} else {
					// Child didn't already exist so let's create it with the current program as the parent
					tower[child] = program{name: child, parent: currentProgram.name}
				}
				if debug {
					fmt.Printf("Child:%s\n", child)
				}
				currentProgram.children = append(currentProgram.children, child)
			}
		}
		// Check if a record already exists. If it does, maintain the parent record
		if existingRecord, ok := tower[currentProgram.name]; ok {
			existingRecord.weight = currentProgram.weight
			existingRecord.children = currentProgram.children
			tower[existingRecord.name] = existingRecord
		} else {
			tower[currentProgram.name] = currentProgram
		}
	}

	if debug {
		fmt.Println(tower)
	}

	return tower
}

func findBottomProgram(tower map[string]program, part byte, debug bool) string {
	for _, i := range tower {
		if i.parent == "" {
			return i.name
		}
	}
	return "Not found"
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		tower := processInputTower(filenamePtr, execPart, debug)
		fmt.Printf("Bottom Program is: %s\n", findBottomProgram(tower, execPart, debug))
	case 'b':
		fmt.Printf("Not implemented yet\n")
	case 'z':
		if execPart == 'z' {
			fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
		}
	}
}
