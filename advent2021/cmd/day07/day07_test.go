package main

import (
	"aoc/advent2021/utils"
	"strconv"
	"strings"
	"testing"
)

func TestCalcLanternfish(t *testing.T) {
	var tests = []struct {
		input    string
		position int
		fuel     int
	}{
		{"test_input", 2, 37},
		{"test_input", 1, 41},
		{"test_input", 3, 39},
		{"test_input", 10, 71},
	}

	var testsB = []struct {
		input    string
		position int
		fuel     int
	}{
		{"test_input", 2, 206},
		{"test_input", 5, 168},
		{"input", 478, 96987874},
	}

	// Test for Part A
	for _, test := range tests {
		var crabPos int
		puzzleInput, _ := utils.ReadFile(test.input)
		puzzleInputSplit := strings.Split(puzzleInput[0], ",")
		crabSubs := make([]int, len(puzzleInputSplit))
		for i := 0; i < len(puzzleInputSplit); i++ {
			crabPos, _ = strconv.Atoi(puzzleInputSplit[i])
			crabSubs[i] = crabPos
		}
		if output := calcFuelUsed(crabSubs, test.position, 'a', false); output != test.fuel {
			t.Error("Part A Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.fuel, output)
		}
	}

	// Test for Part B
	for _, test := range testsB {
		var crabPos int
		puzzleInput, _ := utils.ReadFile(test.input)
		puzzleInputSplit := strings.Split(puzzleInput[0], ",")
		crabSubs := make([]int, len(puzzleInputSplit))
		for i := 0; i < len(puzzleInputSplit); i++ {
			crabPos, _ = strconv.Atoi(puzzleInputSplit[i])
			crabSubs[i] = crabPos
		}
		if output := calcFuelUsed(crabSubs, test.position, 'b', false); output != test.fuel {
			t.Error("Part B Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.fuel, output)
		}
	}

}
