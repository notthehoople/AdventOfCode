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

func validateTickets(ruleSet []rule, nearbyTickets [250][40]int) int {
	var errorRate int = 0

	// Compare all the values in a ticket to the RuleSet
	// Each value that fails ALL rules is added to the error rate

	for ticket := 0; ticket < len(nearbyTickets); ticket++ {
		for valueInTicket := 0; valueInTicket < len(nearbyTickets[ticket])-1; valueInTicket++ {
			if !validValueInRuleSet(nearbyTickets[ticket][valueInTicket], ruleSet) {
				errorRate += nearbyTickets[ticket][valueInTicket]
			}
		}
	}

	return errorRate
}

func calcScanningErrorRate(filename string, part byte, debug bool) int {
	var ruleSet []rule
	var ruleNumber int
	var myTicket []int
	var nearbyTickets [250][40]int
	var processSection int = 0
	var nearbyCounter int

	var grabField string
	var grabFirstLow, grabFirstHigh, grabSecondLow, grabSecondHigh int

	puzzleInput, _ := readFile(filename)
	ruleSet = make([]rule, 30)

	for _, inputLine := range puzzleInput {

		if inputLine == "your ticket:" {
			processSection = 1
			continue
		}

		if inputLine == "nearby tickets:" {
			processSection = 2
			nearbyCounter = 0
			continue
		}

		// [section 0]  - some rules of format "<strings>: <lower>-<upper> or <lower>-<upper>"
		if processSection == 0 && inputLine != "" {
			//			departure location: 49-239 or 247-960
			grabFields := strings.Split(inputLine, ":")
			grabField = grabFields[0]

			fmt.Sscanf(grabFields[1], " %d-%d or %d-%d", &grabFirstLow, &grabFirstHigh, &grabSecondLow, &grabSecondHigh)

			ruleSet[ruleNumber].ruleName = grabField
			ruleSet[ruleNumber].lowerLimit1 = grabFirstLow
			ruleSet[ruleNumber].upperLimit1 = grabFirstHigh
			ruleSet[ruleNumber].lowerLimit2 = grabSecondLow
			ruleSet[ruleNumber].upperLimit2 = grabSecondHigh

			ruleNumber++
			// process rule sets and store in ruleSet map
		}

		// [section 1]  - text "your ticket:"

		if processSection == 1 && inputLine != "" {
			ticketValues := strings.Split(inputLine, ",")

			myTicket = make([]int, len(ticketValues))

			var ticketCounter int = 0
			for _, value := range ticketValues {
				myTicket[ticketCounter], _ = strconv.Atoi(value)
				ticketCounter++

			}
		}

		// [section 2]  - text "nearby tickets:"

		//   - a series of lines with comma separated numbers

		if processSection == 2 {
			ticketValues := strings.Split(inputLine, ",")

			//nearbyTickets[nearbyCounter] = make([]int, len(ticketValues))

			var oneTicketCounter int = 0
			for _, value := range ticketValues {
				nearbyTickets[nearbyCounter][oneTicketCounter], _ = strconv.Atoi(value)
				oneTicketCounter++
			}

			nearbyCounter++
		}
	}

	return validateTickets(ruleSet, nearbyTickets)
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
	}
}
