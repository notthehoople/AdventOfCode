package main

import (
	"AdventOfCode-go/advent2021/utils"
	"testing"
)

func TestCountDepthIncreases(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"test_input", 150},
	}
	var testsB = []struct {
		input    string
		expected int
	}{
		{"test_input", 900},
	}

	// Test for Part A
	for _, test := range tests {
		puzzleInput, _ := utils.ReadFile(test.input)
		if output := calcDepthAndPosition(puzzleInput); output != test.expected {
			t.Error("Part A Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
	// Test for Part B
	for _, test := range testsB {
		puzzleInput, _ := utils.ReadFile(test.input)
		if output := calcDepthAndPositionUsingAim(puzzleInput); output != test.expected {
			t.Error("Part A Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
