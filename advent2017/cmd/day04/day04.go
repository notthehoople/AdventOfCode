package main

import (
	"aoc/advent2017/utils"
	"fmt"
	"strings"
)

func isAnagram(firstWord string, secondWord string, debug bool) bool {

}

func isValidPassphrase(passphrase string, part byte, debug bool) bool {
	//var wordValue int
	var wordExists bool

	wordsMap := make(map[string]int)
	wordsList := strings.Split(passphrase, " ")
	for _, i := range wordsList {
		if debug {
			fmt.Printf("Word: %s Passphrase: %s\n", i, passphrase)
		}
		//wordsMap[i]++
		_, wordExists = wordsMap[i]
		if wordExists {
			return false
		} else {
			wordsMap[i]++
		}
	}
	// If this is part a then if no words are repeated, this is a valid passcode
	if part == 'a' {
		return true
	}

	/*
		In part b, a valid passphrase must contain no two words
		that are anagrams of each other - that is, a passphrase is invalid if any
		word's letters can be rearranged to form any other word in the passphrase.
	*/
	for wordPos, firstWord := range wordsList {
		fmt.Printf("wordPos: %d word: %s\n", wordPos, firstWord)
		for _, secondWord := range wordsList[wordPos+1:] {
			// To be an anagram the lengths of the words must be the same
			if len(firstWord) == len(secondWord) {
				if secondWord != "" {
					// Loop through the letters in firstWord
					if debug {
						fmt.Println("Testing word:", secondWord)
					}
					if isAnagram(firstWord, secondWord, debug) {
						return false
					}
				}
			}
		}
	}
	return true
}

func countValidPassphrases(filename string, part byte, debug bool) int {
	var numberValid int

	puzzleInput, _ := utils.ReadFile(filename)

	// Loop through passphrases
	// if valid passphrase then count++
	for i := 0; i < len(puzzleInput); i++ {
		if isValidPassphrase(puzzleInput[i], part, debug) {
			numberValid++
		}
	}

	return numberValid
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", countValidPassphrases(filenamePtr, execPart, debug))
	}
}
