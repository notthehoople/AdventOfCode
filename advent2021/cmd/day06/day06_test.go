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
		days     int
		expected int
	}{
		{"test_input", 18, 26},
		{"test_input", 80, 5934},
	}
	/*
		var testsB = []struct {
			input    string
			expected int
		}{
			{"test_input", 12},
		}
	*/

	// Test for Part A
	for _, test := range tests {
		var startFish int
		puzzleInput, _ := utils.ReadFile(test.input)
		puzzleInputSplit := strings.Split(puzzleInput[0], ",")
		lanternfish := make([]int, 10)
		for i := 0; i < len(puzzleInputSplit); i++ {
			startFish, _ = strconv.Atoi(puzzleInputSplit[i])
			lanternfish[startFish]++
		}
		if output := calcLanternfish(lanternfish, test.days, 'a', false); output != test.expected {
			t.Error("Part A Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}

	/*
		// Test for Part B
		for _, test := range testsB {
			puzzleInput, _ := utils.ReadFile(test.input)
			if output := calcOverlapPoints(puzzleInput, 'b', false); output != test.expected {
				t.Error("Part A Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
			}
		}
	*/
}
