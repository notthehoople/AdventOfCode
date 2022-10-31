package main

func countNeededBagsRecurse(bagRules map[string]map[string]int, topLevelBag string) int {
	var totalBags int = 0

	for key, value := range bagRules[topLevelBag] {
		// We're digging in to another set of bags. So add the count of the one we've found
		// then we'll loop through that number and look if there's anything inside the bag
		totalBags += value

		for i := 0; i < value; i++ {
			totalBags += countNeededBagsRecurse(bagRules, key)
		}
	}
	return totalBags
}

func countNeededBags(bagRules map[string]map[string]int, bagToSearch string) int {
	var foundBagsCount int = 0

	for key := range bagRules {
		if key == bagToSearch {
			// Find the bag we care about, then dive in and count up all the bags that are contained in it
			foundBagsCount += countNeededBagsRecurse(bagRules, key)
		}
	}

	return foundBagsCount
}
