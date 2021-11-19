package main

import (
	"testing"
)

func TestCalcCaptcha(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
	}
	var testsB = []struct {
		input    string
		expected int
	}{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
	}

	// Test for Part A
	for _, test := range tests {
		if output := calcCaptcha(test.input, 1); output != test.expected {
			t.Error("Part A Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
	// Test for Part B
	for _, testB := range testsB {
		if output := calcCaptcha(testB.input, len(testB.input)/2); output != testB.expected {
			t.Error("Part B Test Failed: {} inputted, {} expected, recieved: {}", testB.input, testB.expected, output)
		}
	}

}
