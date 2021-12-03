package main

import (
	"aoc/advent2021/utils"
	"testing"
)

func TestCountDepthIncreases(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"test_input", 198},
	}

	// Test for Part A
	for _, test := range tests {
		puzzleInput, _ := utils.ReadFile(test.input)
		if output := calcPowerConsumption(puzzleInput, false); output != test.expected {
			t.Error("Part A Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
	// Test for Part B
	/*
		for _, test := range testsB {
			puzzleInput, _ := utils.ReadFile(test.input)
			if output := calcDepthAndPositionUsingAim(puzzleInput); output != test.expected {
				t.Error("Part A Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
			}
		}*/
}
