package main

// Tried 1573 - too high

import (
	"AdventOfCode-go/advent2024/utils"
	"fmt"
	"sort"
	"strings"
)

var Debug bool

type Connected struct {
	first  string
	second string
	third  string
}

func buildLANMap(puzzleInput []string) map[string]map[string]bool {
	lanMap := make(map[string]map[string]bool)

	for _, puzzleLine := range puzzleInput {
		tempString := strings.Split(puzzleLine, "-")

		if _, ok := lanMap[tempString[0]]; !ok {
			lanMap[tempString[0]] = make(map[string]bool)
		}

		if _, ok := lanMap[tempString[1]]; !ok {
			lanMap[tempString[1]] = make(map[string]bool)
		}

		lanMap[tempString[0]][tempString[1]] = true
		lanMap[tempString[1]][tempString[0]] = true
	}

	return lanMap
}

func calcTComputers(lanMap map[string]map[string]bool) int {
	var count int
	// Find 3 connected computers. Disgard all those without a computer that starts with a 't'

	connectedMap := make(map[Connected]bool)

	// Loop through the lanMap looking for connected computers
	for key, value := range lanMap {
		if len(value) <= 1 { // Is this needed?
			continue
		}

		for key2 := range value {

			for key3 := range value {
				if key2 != key3 {

					if _, ok := lanMap[key2][key3]; ok {
						resString := []string{key, key2, key3}
						sort.Strings(resString)
						connectedMap[Connected{resString[0], resString[1], resString[2]}] = true
					}
				}
			}
		}
	}

	for key, value := range connectedMap {
		if Debug {
			fmt.Println(key, value)
		}
		if key.first[0] == 't' || key.second[0] == 't' || key.third[0] == 't' {
			count++
		}
	}

	return count
}

func day23(filename string, part byte) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)
	lanMap := buildLANMap(puzzleInput)

	//fmt.Println(lanMap)

	result = calcTComputers(lanMap)

	return result
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()
	Debug = debug

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", day23(filenamePtr, execPart))
	}
}
