package main

import (
	"fmt"
	"strconv"
	"strings"
)

type rule struct {
	ruleName    string
	lowerLimit1 int
	upperLimit1 int
	lowerLimit2 int
	upperLimit2 int
}

func validValueInRuleSet(value int, ruleSet []rule) bool {
	var isValid bool = false

	for rulePos := 0; rulePos < len(ruleSet); rulePos++ {
		if ruleSet[rulePos].lowerLimit1 == 0 && ruleSet[rulePos].upperLimit1 == 0 {
			// Stop if we reach the end of the rules
			return isValid
		}

		if (value >= ruleSet[rulePos].lowerLimit1 && value <= ruleSet[rulePos].upperLimit1) ||
			(value >= ruleSet[rulePos].lowerLimit2 && value <= ruleSet[rulePos].upperLimit2) {
			// value is fine by this rule
			return true
		}
	}

	return false
}

func validateTickets(ruleSet []rule, nearbyTickets []string) int {
	var errorRate int = 0
	var checkValue int = 0

	// Compare all the values in a ticket to the RuleSet
	// Each value that fails ALL rules is added to the error rate

	for _, inputLine := range nearbyTickets {
		ticketValues := strings.Split(inputLine, ",")

		for _, value := range ticketValues {
			checkValue, _ = strconv.Atoi(value)

			if !validValueInRuleSet(checkValue, ruleSet) {
				errorRate += checkValue
			}
		}

	}

	return errorRate
}

func calcScanningErrorRate(filename string, part byte, debug bool) int {
	var ruleSet []rule
	var myTicket string
	var processSection int = 0
	var rawNearbyTickets []string

	var grabField string
	var grabFirstLow, grabFirstHigh, grabSecondLow, grabSecondHigh int

	puzzleInput, _ := readFile(filename)
	// NEED TO GET THE RULE SET CREATED CORRECTLY

	for inputCounter, inputLine := range puzzleInput {

		if inputLine == "your ticket:" {
			processSection = 1
			continue
		}

		if inputLine == "nearby tickets:" {
			// The rest of the file is nearby tickets. Grab them then exit the input processing loop
			rawNearbyTickets = puzzleInput[inputCounter+1:]
			break
		}

		// [section 0]  - some rules of format "<strings>: <lower>-<upper> or <lower>-<upper>"
		if processSection == 0 && inputLine != "" {
			//			departure location: 49-239 or 247-960
			grabFields := strings.Split(inputLine, ":")
			grabField = grabFields[0]

			fmt.Sscanf(grabFields[1], " %d-%d or %d-%d", &grabFirstLow, &grabFirstHigh, &grabSecondLow, &grabSecondHigh)

			ruleSet = append(ruleSet, rule{ruleName: grabField, lowerLimit1: grabFirstLow, upperLimit1: grabFirstHigh, lowerLimit2: grabSecondLow, upperLimit2: grabSecondHigh})
			// process rule sets and store in ruleSet map

			if debug {
				fmt.Println("------")
				fmt.Println(ruleSet)
			}
		}

		// [section 1]  - text "your ticket:"

		if processSection == 1 && inputLine != "" {
			myTicket = inputLine
			if debug {
				fmt.Println("My Ticket:", myTicket)
			}
		}
	}

	if part == 'a' {
		return validateTickets(ruleSet, rawNearbyTickets)
	} else {
		return decodeMyTicket(ruleSet, rawNearbyTickets, myTicket)
	}
}

// Main routine
func main() {
	filenamePtr, execPart, debug, test := catchUserInput()

	if test {
		return
	}

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else if execPart == 'a' {
		fmt.Println("Part a answer:", calcScanningErrorRate(filenamePtr, execPart, debug))
	} else {
		fmt.Println("Part b answer:", calcScanningErrorRate(filenamePtr, execPart, debug))
	}
}
