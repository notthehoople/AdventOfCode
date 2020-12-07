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

	var topLevelBag string

	initialSplit := strings.Split(ruleLine, "contain")

	topLevelBag = strings.Trim(initialSplit[0], " ")
	topLevelBag = strings.TrimSuffix(topLevelBag, " bags")
	topLevelBag = strings.TrimSuffix(topLevelBag, " bag")
	bagRules[topLevelBag] = make(map[string]int)

	if strings.Trim(initialSplit[1], " .") == "no other bags" {
		if debug {
			fmt.Println("No other bags so stop processing")
		}
	} else {
		for _, item := range strings.Split(initialSplit[1], ",") {
			var numBags int = 0
			var lowerLevelBag string

			item = strings.Trim(item, " .")

			// Grab the number, plus trim spaces, full stops and get rid of bag/bags
			fmt.Sscanf(item, "%d %s", &numBags, &lowerLevelBag)
			lowerLevelBag = item[strings.IndexByte(item, ' ')+1:]
			lowerLevelBag = strings.TrimSuffix(lowerLevelBag, " bags")
			lowerLevelBag = strings.TrimSuffix(lowerLevelBag, " bag")

			bagRules[topLevelBag][lowerLevelBag] = numBags
		}
	}
}

func searchBagsRecurse(bagRules map[string]map[string]int, topLevelBag string, bagToSearch string) int {
	var foundCount int = 0

	for thing := range bagRules[topLevelBag] {

		if thing == bagToSearch {
			//fmt.Println("Got it in searchBagsRecurse")
			return 1
		}

		foundCount = searchBagsRecurse(bagRules, thing, bagToSearch)
		if foundCount > 0 {
			return 1
		}
	}
	return 0
}

func searchBags(bagRules map[string]map[string]int, bagToSearch string) int {
	var foundBagsCount int = 0

	for thing := range bagRules {
		if thing == bagToSearch {
			// If we find it at the top level we ignore it as needs to be INSIDE another bag
			//fmt.Println("Got it in searchBags")
		} else {
			foundBagsCount += searchBagsRecurse(bagRules, thing, bagToSearch)
		}
	}

	return foundBagsCount
}

// Process the whole file of bag rules
// Build a map of maps containing the rules
func processAllBagRules(filename string, bagToSearch string, part byte, debug bool) int {
	var bagRules map[string]map[string]int
	bagRules = make(map[string]map[string]int)

	puzzleInput, _ := readFile(filename)

	for _, singleRule := range puzzleInput {
		readSingleRule(singleRule, bagRules, part, debug)
	}

	bagToSearch = strings.TrimSuffix(bagToSearch, " bags")
	bagToSearch = strings.TrimSuffix(bagToSearch, " bag")
	if debug {
		fmt.Println("Bag to Search for:", bagToSearch)
		fmt.Println("Bag Rules:", bagRules)
	}

	return searchBags(bagRules, bagToSearch)
}

// Main routine
func main() {
	filenamePtr, execPart, debug := catchUserInput()
	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else if execPart == 'a' {
		fmt.Println("Bag Varieties:", processAllBagRules(filenamePtr, "shiny gold bag", execPart, debug))
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
