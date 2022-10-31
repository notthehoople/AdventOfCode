package main

import (
	"fmt"
	//"strconv"
	"strings"
)

type rule struct {
	char    string
	subRule [][]string
}

func parseMessage(ruleToUse string, message string, ruleSet map[string]rule, index int) (bool, int) {
	activeRule := ruleSet[ruleToUse]

	if index == len(message) {
		// alread at the end of the message so everything's good
		return true, 0
	}

	if activeRule.char != "" {
		// We have a character rather than a set of rules so we've reached the bottom of this traversal. But does it match?
		if message[index:index+1] == activeRule.char {
			return true, index + 1
		}
		// message fails to match rules
		return false, 0
	}

	var pos int

	for _, ruleGroup := range activeRule.subRule {
		groupResult := true

		pos = index

		for _, rules := range ruleGroup {
			result, newPos := parseMessage(rules, message, ruleSet, pos)
			if !result {
				// that didn't match so bail on this rule and try the next if there is one
				groupResult = false
				break
			}
			pos = newPos
		}
		if groupResult {
			return true, pos
		}
	}
	return false, index
}

func matchMessages(filename string, part byte, debug bool) int {
	var ruleSet map[string]rule
	var processMessages bool = false
	//var matchRuleNum int
	var matchRules []string
	var rawMessages []string

	ruleSet = make(map[string]rule)

	puzzleInput, _ := readFile(filename)
	// Need to build the rule set first, then match the messages

	for inputCounter, inputLine := range puzzleInput {

		if inputLine == "" {
			processMessages = true
			continue
		}

		if !processMessages {
			// Process the rule set
			if debug {
				fmt.Println("----- Processing the rule set -----")
			}
			// [section 0] - some rules of format "<strings>: <number> ..." (1 or 2 numbers)
			//               or "<number>: <number> ... | <number> ..." (1 or 2 numbers in each place)
			//               or "<number>: "char"" (where char is a or b)
			matchRules = strings.Split(inputLine, ": ")
			if debug {
				fmt.Printf("Matched: %s with rules %s\n", matchRules[0], matchRules[1:])
			}
			// Check for single letter enclosed in quotes. If there then set .char to it
			if strings.Contains(matchRules[1], "\"") {
				ruleSet[matchRules[0]] = rule{char: strings.Trim(matchRules[1], "\"")}
			} else {
				// Process the rule parts
				tempRule := rule{subRule: [][]string{}}
				for _, rulePart := range strings.Split(matchRules[1], " | ") {
					tempRule.subRule = append(tempRule.subRule, strings.Split(rulePart, " "))
				}
				ruleSet[matchRules[0]] = tempRule
			}

		} else {
			rawMessages = puzzleInput[inputCounter:]
			break
		}
	}

	goodMessageCount := 0
	for i := 0; i < len(rawMessages); i++ {
		if debug {
			fmt.Printf("Message to match: %s\n", rawMessages[i])
		}
		ok, pos := parseMessage("0", rawMessages[i], ruleSet, 0)
		if ok && pos == len(rawMessages[i]) {
			goodMessageCount++
		}
	}

	if debug {
		fmt.Println("===================================")
		fmt.Println("RuleSet:")
		fmt.Println(ruleSet)
		fmt.Println("Messages:")
		fmt.Println(rawMessages)
	}

	return goodMessageCount
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
		fmt.Println("Part a answer:", matchMessages(filenamePtr, execPart, debug))
	} else {
	}
}
