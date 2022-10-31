package main

import (
	"flag"
	"fmt"
)

func catchUserInput() (string, byte, bool) {
	var debug bool

	filenamePtr := flag.String("file", "testInput.txt", "Filename containing the program to run")
	execPartPtr := flag.String("part", "a", "Which part of the puzzle do you want to calc (a or b)")
	flag.BoolVar(&debug, "debug", false, "Turn debug on")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		return *filenamePtr, 'a', debug
	case "b":
		return *filenamePtr, 'b', debug

	default:
		return *filenamePtr, 'z', debug
	}
}

func processSinglePersonAnswers(answers string, answersStore map[byte]int, part byte, debug bool) {
	// Build a map of the questions

	for _, answer := range answers {
		answersStore[byte(answer)]++
	}
}

func countAllAnswers(answersStore map[byte]int, numOfPeople int) int {
	var groupCount = 0

	for _, value := range answersStore {
		if value >= numOfPeople {
			groupCount++
		}
	}
	return groupCount
}

func processGroupAnswers(filename string, part byte, debug bool) int {
	var answersStore map[byte]int
	var sumOfCounts int = 0
	var numberOfPeople int = 0

	answersStore = make(map[byte]int)

	puzzleInput, _ := readFile(filename)

	for _, singlePersonAnswers := range puzzleInput {
		if debug {
			fmt.Println(singlePersonAnswers)
		}

		if singlePersonAnswers == "" {
			// Onto the next group
			// Sum up the total number of answers we have for this group
			if debug {
				fmt.Println(answersStore)
			}

			if part == 'a' {
				sumOfCounts += countAllAnswers(answersStore, 1)
			} else {
				sumOfCounts += countAllAnswers(answersStore, numberOfPeople)
			}
			// Clear the answers store and start again
			answersStore = make(map[byte]int)
			numberOfPeople = 0
		} else {
			numberOfPeople++
		}

		processSinglePersonAnswers(singlePersonAnswers, answersStore, part, debug)
	}

	// count the last entry in case there wasn't a blank line before EOF
	if part == 'a' {
		sumOfCounts += countAllAnswers(answersStore, 1)
	} else {
		sumOfCounts += countAllAnswers(answersStore, numberOfPeople)
	}

	return sumOfCounts
}

// Main routine
func main() {
	filenamePtr, execPart, debug := catchUserInput()
	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else if execPart == 'a' {
		fmt.Println("Sum of anyone answers:", processGroupAnswers(filenamePtr, execPart, debug))
	} else {
		fmt.Println("Sum of everyone answers:", processGroupAnswers(filenamePtr, execPart, debug))
	}
}
