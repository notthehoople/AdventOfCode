package main

import (
	"fmt"
	//"strconv"
	"strings"
)

type rule struct {
	char   string
	groups [][]string
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
			if len(matchRules) == 2 {
				if debug {
					fmt.Printf("Matched: %s with rules %s\n", matchRules[0], matchRules[1:])
				}
				// Check for single letter enclosed in quotes. If there then set .char to it
				// Otherwise process the parts and set the rule parts of the array

				if strings.Contains(matchRules[1], "\"") {
					fmt.Println("before trim:", matchRules[1])
					ruleSet[matchRules[0]] = rule{char: strings.Trim(matchRules[1], "\"")}
					fmt.Printf("char: '%s'\n", strings.Trim(matchRules[1], "\""))
				}
			} else {
				fmt.Printf("Bad input line: %s\n", inputLine)
				return 0
			}
		} else {
			rawMessages = puzzleInput[inputCounter+1:]
			break
		}
	}

	// DELETE
	if debug {
		fmt.Println(ruleSet)
		fmt.Println(rawMessages)
	}
	// DELETE

	if debug {
		fmt.Println("----- Processing messages -----")
	}

	// Need to determine the number of messages that completely match rule 0
	//for message := range rawMessages {
	//
	//	}

	return 0
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
