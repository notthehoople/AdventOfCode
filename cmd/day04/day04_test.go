package main

import (
	"testing"
)

func TestIsValidPassphrase(t *testing.T) {
	var testsA = []struct {
		input    string
		expected bool
	}{
		{"aa bb cc dd ee", true},
		{"aa bb cc dd aa", false},
		{"aa bb cc dd aaa", true},
	}

	// Test for Part A
	for _, test := range testsA {
		if output := isValidPassphrase(test.input, false); output != test.expected {
			t.Error("Part A Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
