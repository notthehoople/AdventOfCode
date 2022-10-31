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
		if output := isValidPassphrase(test.input, 'a', false); output != test.expected {
			t.Error("Part A Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}

func TestIsValidPassphraseNotAnagram(t *testing.T) {
	var testsB = []struct {
		input    string
		expected bool
	}{
		{"abcde fghij", true},
		{"abcde xyz ecdab", false},
		{"a ab abc abd abf abj", true},
		{"iiii oiii ooii oooi oooo", true},
		{"oiii ioii iioi iiio", false},
	}

	// Test for Part B
	for _, test := range testsB {
		if output := isValidPassphrase(test.input, 'b', false); output != test.expected {
			t.Error("Part B Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
