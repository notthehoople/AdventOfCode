package main

import (
	"fmt"
	"strings"
)

func readSingleRule(ruleLine string, bagRules map[string]map[string]int, part byte, debug bool) {
	// split on "contain". On the left is the master bag. On the right the remainder
	// check remainder for "no other bags"
	// split the remainder on ","
	//   - process each of the remaining strings as a bag

	//var bagRules map[string]map[string]int
	var topLevelBag string
	//bagRules = make(map[string]map[string]int)

	initialSplit := strings.Split(ruleLine, "contain")
	//fmt.Println("First string:", initialSplit[0])
	//fmt.Println("Remainder:", strings.Trim(initialSplit[1], " ."))

	topLevelBag = strings.Trim(initialSplit[0], " ")
	bagRules[topLevelBag] = make(map[string]int)

	// Note need to either deal with "bags" vs "bag" here or during search

	if strings.Trim(initialSplit[1], " .") == "no other bags" {
		fmt.Println("No other bags so stop processing")
	} else {
		for _, item := range strings.Split(initialSplit[1], ",") {
			var numBags int = 0
			var lowerLevelBag string

			item = strings.Trim(item, " .")

			fmt.Sscanf(item, "%d %s", &numBags, &lowerLevelBag)
			lowerLevelBag = item[strings.IndexByte(item, ' ')+1:]

			bagRules[topLevelBag][lowerLevelBag] = numBags
		}
	}
}

// Process the whole file of bag rules
// Build a map of maps containing the rules
func processAllBagRules(filename string, part byte, debug bool) int {
	var bagRules map[string]map[string]int
	bagRules = make(map[string]map[string]int)

	puzzleInput, _ := readFile(filename)

	for _, singleRule := range puzzleInput {
		//fmt.Println("Single Rule:", singleRule)
		//fmt.Println("bagRules:", bagRules)
		readSingleRule(singleRule, bagRules, part, debug)
	}

	fmt.Println("Bag Rules:", bagRules)

	return 0
}

// Main routine
func main() {
	filenamePtr, execPart, debug := catchUserInput()
	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else if execPart == 'a' {
		fmt.Println("Bag Varieties:", processAllBagRules(filenamePtr, execPart, debug))
	} else {
		fmt.Println("Not implemented yet")
		fmt.Println(filenamePtr, execPart, debug)
	}

	if debug {
		//readSingleRule("light salmon bags contain 5 dark brown bags, 21 dotted coral bags, 5 mirrored turquoise bags.", 'a', true)
		//readSingleRule("drab magenta bags contain 1 vibrant purple bag, 5 dark lime bags, 2 clear silver bags.", 'a', true)
		//readSingleRule("wavy magenta bags contain 1 dotted crimson bag.", 'a', true)
		//readSingleRule("light gold bags contain no other bags.", 'a', true)
	}
}
