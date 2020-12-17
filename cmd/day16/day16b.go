package main

import (
	"fmt"
	"strconv"
	"strings"
)

func whatRuleSetIsThisValueIn(ruleSet []rule, value int) []int {
	var results []int

	for rulePos := 0; rulePos < len(ruleSet); rulePos++ {
		if (value >= ruleSet[rulePos].lowerLimit1 && value <= ruleSet[rulePos].upperLimit1) ||
			(value >= ruleSet[rulePos].lowerLimit2 && value <= ruleSet[rulePos].upperLimit2) {
			results = append(results, rulePos)
		}
	}

	return results
}

func isTicketValid(ruleSet []rule, ticket string) bool {
	var checkValue int

	// Compare all the values in a ticket to the RuleSet
	// Each value that fails ALL rules is added to the error rate

	ticketValues := strings.Split(ticket, ",")

	for _, value := range ticketValues {
		checkValue, _ = strconv.Atoi(value)

		if !validValueInRuleSet(checkValue, ruleSet) {
			return false
		}
	}

	return true
}

func decodeMyTicket(ruleSet []rule, nearbyTickets []string, myTicket string) int {
	var checkValue int
	var fieldResults map[int]int

	// Will hold the count to find the rule for each field
	fieldResults = make(map[int]int)

	for fieldNumber := 0; fieldNumber < 20; fieldNumber++ {
		var rowChoice []int
		rowChoice = make([]int, len(ruleSet))

		var validTickets = 0
		for _, currentTicket := range nearbyTickets {
			//fmt.Printf("Checking field %d of %s\n", fieldNumber, currentTicket)

			if isTicketValid(ruleSet, currentTicket) {
				//fmt.Println("....is valid")
				validTickets++

				// Check all values in field 1. Then field 2, Then field 3.

				ticketValues := strings.Split(currentTicket, ",")

				//				for _, value := range ticketValues {
				//					checkValue, _ = strconv.Atoi(value)
				checkValue, _ = strconv.Atoi(ticketValues[fieldNumber])
				//fmt.Println(whatRuleSetIsThisValueIn(ruleSet, checkValue))

				for _, i := range whatRuleSetIsThisValueIn(ruleSet, checkValue) {
					rowChoice[i]++
				}
			} else {
				//fmt.Println("....is NOT valid")
			}
		}

		fmt.Println("Results for field:", fieldNumber, rowChoice)

		var topValue int = 0
		var topPos int = 0
		for pos, value := range rowChoice {
			// Edge case: some fields will not choose a winner. They will need the others to eliminate what it can be
			if value > topValue {
				topValue = value
				topPos = pos
			}
		}
		fieldResults[fieldNumber] = topPos
		//fmt.Println("Results: ", fieldResults)
	}
	return 0
}
