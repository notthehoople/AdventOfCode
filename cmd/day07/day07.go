package main

import (
	"fmt"
	"strings"
)

// Puzzle input is a long list of rules of bags
// A bag can contain many other bags, or none at all
// Each bag inside the top level bag can hold other bags, or none at all
// Need to
// - process the rules file and build a set of rules
// - take the bag type to search for
// - search the set of rules for the bag to search for
//     - Count the different options that can house the bag to search for

func readSingleRule(ruleLine string, part byte, debug bool) {
	//	func readSingleRule(ruleLine string, bagRules map[string]map[string]int, part byte, debug bool) {
	// split on "contain". On the left is the master bag. On the right the remainder
	// check remainder for "no other bags"
	// split the remainder on ","
	//   - process each of the remaining strings as a bag

	// array of maps?
	// map of arrays?

	initialSplit := strings.Split(ruleLine, "contain")
	fmt.Println("First string:", initialSplit[0])
	fmt.Println("Remainder:", strings.Trim(initialSplit[1], " ."))

	// Process First String
	// Process Remainder
	//   - Check for "no other bags"
	//   - range on Split using ','

	if strings.Trim(initialSplit[1], " .") == "no other bags." {
		fmt.Println("No other bags so stop processing")
	} else {
		//bagRules[initialSplit[0]] = make(map[string]int)
		for _, item := range strings.Split(initialSplit[1], ",") {
			var numBags int = 0
			var Bag string

			fmt.Println("One item:", strings.Trim(item, " ."))

			fmt.Sscanf(strings.Trim(item, " ."), "%d %s", &numBags, &Bag)
			fmt.Printf("No Bags: %d Bag: %s\n", numBags, Bag)
		}
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

//func processGroupAnswers(filename string, part byte, debug bool) int {
//	var answersStore map[byte]int
//	var sumOfCounts int = 0
//	var numberOfPeople int = 0
//
//	answersStore = make(map[byte]int)
//
//	puzzleInput, _ := readFile(filename)
//
//	return 0
//}

// Main routine
func main() {
	filenamePtr, execPart, debug := catchUserInput()
	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else if execPart == 'a' {
		fmt.Println("Not implemented yet")
	} else {
		fmt.Println("Not implemented yet")
		fmt.Println(filenamePtr, execPart, debug)
	}

	if debug {
		readSingleRule("light salmon bags contain 5 dark brown bags, 2 dotted coral bags, 5 mirrored turquoise bags.", 'a', true)
		readSingleRule("drab magenta bags contain 1 vibrant purple bag, 5 dark lime bags, 2 clear silver bags.", 'a', true)
		readSingleRule("wavy magenta bags contain 1 dotted crimson bag.", 'a', true)
		readSingleRule("light gold bags contain no other bags.", 'a', true)
	}
}
