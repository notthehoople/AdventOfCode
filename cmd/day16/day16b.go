package main

import (
	"fmt"
	//"strconv"
	//"strings"
)

func decodeMyTicket(ruleSet []rule, nearbyTickets [250][40]int, myTicket []int) int {
	var badTickets map[int]bool

	badTickets = make(map[int]bool)

	for ticket := 0; ticket < len(nearbyTickets); ticket++ {
		for valueInTicket := 0; valueInTicket < len(nearbyTickets[ticket])-1; valueInTicket++ {
			fmt.Printf("Testing %d against RuleSet\n", nearbyTickets[ticket][valueInTicket])

			if !validValueInRuleSet(nearbyTickets[ticket][valueInTicket], ruleSet) {
				// Invalid ticket so discard
				badTickets[ticket] = true
			}
		}
	}

	fmt.Println(badTickets)

	return 0
}
