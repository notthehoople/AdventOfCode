package main

import (
	"AdventOfCode-go/advent2021/utils"
	"testing"
)

func TestCalcOverlapPoints(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"test_input", 5},
	}
	var testsB = []struct {
		input    string
		expected int
	}{
		{"test_input", 12},
	}

	// Test for Part A
	for _, test := range tests {
		puzzleInput, _ := utils.ReadFile(test.input)
		if output := calcOverlapPoints(puzzleInput, 'a', false); output != test.expected {
			t.Error("Part A Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
	// Test for Part B

	for _, test := range testsB {
		puzzleInput, _ := utils.ReadFile(test.input)
		if output := calcOverlapPoints(puzzleInput, 'b', false); output != test.expected {
			t.Error("Part A Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
