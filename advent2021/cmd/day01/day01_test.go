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
		{"test_input", 7},
	}

	var testsB = []struct {
		input    string
		expected int
	}{
		{"test_input", 5},
	}

	// Test for Part A
	for _, test := range tests {
		puzzleInput, _ := utils.ReadFile(test.input)
		if output := countDepthIncreases(puzzleInput); output != test.expected {
			t.Error("Part A Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
	// Test for Part B
	for _, testB := range testsB {
		puzzleInput, _ := utils.ReadFile(testB.input)
		if output := countSlidingWindowIncreases(puzzleInput); output != testB.expected {
			t.Error("Part B Test Failed: {} inputted, {} expected, recieved: {}", testB.input, testB.expected, output)
		}
	}

}
