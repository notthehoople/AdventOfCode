package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calcSumRecurse(sumRemainder string) int {
	//var sumResult int = 0
	var rightNumber int = 0
	var operator int = 0 // if 1 then add, if 2 then multiply
	var currentResult int = 0
	var skipTillClose bool = false
	// operator order is left to right UNLESS there are parentheses. Those go first

	// Grab the left side, operator and right side of the sum
	// If a number setup the sum as left side, operator, right side
	// If an open bracket then call ourselves with string position + 1
	// If a number with a close bracket then complete sum and RETURN with answer

	sumItems := strings.Split(sumRemainder, " ")

	for _, item := range sumItems {
		if skipTillClose {
			fmt.Println("Skipping:", item)
			if len(item) > 1 {
				if item[1] == ')' {
					skipTillClose = false
					continue
				}
			}
			continue
		}
		if len(item) > 1 {
			// If we get here we have a bracket. Either an open bracket or a number and a close bracket
			fmt.Printf("item '%s' is longer than 1. Must be a bracket somewhere\n", item)
			if item[0] == '(' {
				// Found a bracket. Call ourselves with the remainder of the string
				fmt.Println("sumRemainder is:", sumRemainder)
				fmt.Printf("item is '%s' and recurse with '%s'\n", item, sumRemainder[strings.Index(sumRemainder, item)+1:])

				rightNumber = calcSumRecurse(sumRemainder[strings.Index(sumRemainder, item)+1:])
				fmt.Println("Number back from recursion:", rightNumber, operator, currentResult)
				if operator == 0 {
					currentResult = rightNumber
				} else if operator == 1 {
					fmt.Printf("Adding %d to %d\n", currentResult, rightNumber)
					currentResult += rightNumber
					fmt.Println("Result:", currentResult)
					operator = 0
				} else {
					fmt.Printf("Multiplying %d by %d\n", currentResult, rightNumber)
					currentResult *= rightNumber
					fmt.Println("Result:", currentResult)
					operator = 0
				}

				skipTillClose = true
			} else if item[1] == ')' {
				// closing the sum
				// blah blah and return with the result
				rightNumber, _ = strconv.Atoi(string(item[0]))
				if operator == 1 {
					fmt.Printf("Adding %d to %d\n", currentResult, rightNumber)
					currentResult += rightNumber
					fmt.Println("Result:", currentResult)
					operator = 0 // don't think this is needed
					return currentResult
				} else {
					fmt.Printf("Multiplying %d by %d\n", currentResult, rightNumber)
					currentResult *= rightNumber
					fmt.Println("Result:", currentResult)
					operator = 0 // don't think this is needed
					return currentResult
				}
			}
		} else {
			switch item {
			case "+":
				fmt.Println("+")
				operator = 1
				break
			case "*":
				fmt.Println("*")
				operator = 2
				break
			case "(":
				fmt.Println("Bracket on its own. Probably never happens")
				break
			case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				if operator == 0 {
					currentResult, _ = strconv.Atoi(item)
					fmt.Println("No operator so grab number:", currentResult)
				} else {
					rightNumber, _ = strconv.Atoi(item)
					if operator == 1 {
						fmt.Printf("Adding %d to %d\n", currentResult, rightNumber)
						currentResult += rightNumber
						fmt.Println("Result:", currentResult)
						operator = 0
					} else {
						fmt.Printf("Multiplying %d by %d\n", currentResult, rightNumber)
						currentResult *= rightNumber
						fmt.Println("Result:", currentResult)
						operator = 0
					}
				}
				break
			}
		}
	}
	return currentResult
}

func calcHomeworkAnswers(filename string, part byte, debug bool) int {

	puzzleInput, _ := readFile(filename)

	for lineNum, sumLine := range puzzleInput {
		fmt.Println("--------- Line:", lineNum)
		calcSumRecurse(sumLine)
	}

	// Scan through the sum looking for (
	//   When find one, grab ALL the text until the next ( or )
	//     If ( then grab ALL the text until the next ( or )
	//     if ) then close the sum and calculate it, then return the result

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
		fmt.Println("Part a answer:", calcHomeworkAnswers(filenamePtr, execPart, debug))
	} else {
	}
}
