package main

import (
	"testing"
)

func TestCalcRowDifference(t *testing.T) {
	var testsA = []struct {
		input    string
		expected int
	}{
		{"5 1 9 5", 8},
		{"7 5 3", 4},
		{"2 4 6 8", 6},
		{"179	2358	5197	867	163	4418	3135	5049	187	166	4682	5080	5541	172	4294	1397", 5378},
	}

	// Test for Part A
	for _, test := range testsA {
		if output := calcRowDifference(test.input, false); output != test.expected {
			t.Error("Part A Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}

func TestSolveChecksum(t *testing.T) {
	var testsA = []struct {
		input    string
		expected int
	}{
		{"test_input", 18},
	}
	// Test for Part A
	for _, test := range testsA {
		if output := solveChecksum(test.input, 'a', false); output != test.expected {
			t.Error("Part A Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
