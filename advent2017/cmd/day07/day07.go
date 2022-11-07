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

func getChildStackWeights(rootProgramName string, tower map[string]program) int {
	rootProgram := tower[rootProgramName]
	rootWeight := rootProgram.weight

	//fmt.Printf("============= Parent: %s ==============\n", rootProgramName)
	var childWeightSum int
	for _, child := range rootProgram.children {
		//fmt.Println("Child: ", child)
		childWeight := getChildStackWeights(child, tower)
		childWeightSum += childWeight
	}

	return rootWeight + childWeightSum
}

func isProgramBalanced(rootProgramName string, tower map[string]program, debug bool) (string, int) {
	topLevelOccurances := make(map[int]int, 0)
	children := make(map[int]string)

	// Count occurance of weight totals for the program's children
	for _, child := range tower[rootProgramName].children {
		topLevelOccurances[getChildStackWeights(child, tower)]++
		children[getChildStackWeights(child, tower)] = child
	}

	// with thanks to tardisman5197 on github
	var result, odd, normal = "", 0, 0
	for k, v := range topLevelOccurances {
		if v == 1 {
			result = children[k]
			odd = k
		} else {
			normal = k
		}
	}
	return result, normal - odd
}

func findWrongWeight(rootProgramName string, tower map[string]program, part byte, debug bool) int {
	// The program that needs its weight to be adjusted is the one that matches these 2 criteria:
	// - it has a different *total* weight (weight of the subtower starting from that program) than total weight of
	//   other programs that are standing on the disc that the program is standing on
	// - it holds a balanced disc (all of its children have the same total weight)
	//
	// with thanks to tardisman5197 on github

	if wrong, diff := isProgramBalanced(rootProgramName, tower, debug); wrong != "" {
		if findWrongWeight(wrong, tower, part, debug) == 0 {
			return tower[wrong].weight + diff
		}
		return findWrongWeight(wrong, tower, part, debug)
	}
	return 0
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		tower := processInputTower(filenamePtr, execPart, debug)
		fmt.Printf("Bottom Program is: %s\n", findBottomProgram(tower, execPart, debug))
	case 'b':
		tower := processInputTower(filenamePtr, execPart, debug)
		rootProgramName := findBottomProgram(tower, 'a', debug)
		fmt.Printf("Revised weight is: %d\n", findWrongWeight(rootProgramName, tower, execPart, debug))
	case 'z':
		if execPart == 'z' {
			fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
		}
	}
}
