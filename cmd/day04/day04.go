package main

import (
	"aoc/advent2017/utils"
	"fmt"
	"strings"
)

func isValidPassphrase(passphrase string, debug bool) bool {
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
	return true
}

func countValidPassphrases(filename string, part byte, debug bool) int {
	var numberValid int

	puzzleInput, _ := utils.ReadFile(filename)
	if part == 'a' {
		// Loop through passphrases
		// if valid passphrase then count++
		for i := 0; i < len(puzzleInput); i++ {
			if isValidPassphrase(puzzleInput[i], debug) {
				numberValid++
			}
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
