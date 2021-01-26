package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calcSumRecurse(sumRemainder string, debug bool) (int, string) {
	//var sumResult int = 0
	var rightNumber int = 0
	var operator int = 0 // if 1 then add, if 2 then multiply
	var currentResult int = 0
	var completedTheLoop bool = false
	// operator order is left to right UNLESS there are parentheses. Those go first

	// Grab the left side, operator and right side of the sum
	// If a number setup the sum as left side, operator, right side
	// If an open bracket then call ourselves with string position + 1
	// If a number with a close bracket then complete sum and RETURN with answer

	// This needs to be rewritten with a better loop
	// Loop until we're finished
	//   have sumRemainder return the chopped string, removing what it has processed
	//   have a loop inside the main loop which is for _, item := range strings.Split(sumRemainder, " ")
	//     this loop finishes when we come back from a calcSumRecurse call
	//     removes the need for the skipTillClose loop which doesn't work

	keepLooping := true
	for keepLooping {
		if len(sumRemainder) == 0 || completedTheLoop == true {
			break
		}

		sumItems := strings.Split(sumRemainder, " ")
		completedTheLoop = true
		for _, item := range sumItems {

			if len(item) > 1 {
				// If we get here we have a bracket. Either an open bracket or a number and a close bracket
				if debug {
					fmt.Printf("item '%s' is longer than 1. Must be a bracket somewhere\n", item)
				}
				if item[0] == '(' {
					// Found a bracket. Call ourselves with the remainder of the string
					if debug {
						fmt.Println("sumRemainder is:", sumRemainder)
						fmt.Printf("item is '%s' and recurse with '%s'\n", item, sumRemainder[strings.Index(sumRemainder, item)+1:])
					}

					rightNumber, sumRemainder = calcSumRecurse(sumRemainder[strings.Index(sumRemainder, item)+1:], debug)
					if debug {
						fmt.Println("Number back from recursion:", rightNumber, operator, currentResult, sumRemainder)
					}
					if operator == 0 {
						currentResult = rightNumber
					} else if operator == 1 {
						if debug {
							fmt.Printf("Adding %d to %d\n", currentResult, rightNumber)
						}
						currentResult += rightNumber
						if debug {
							fmt.Println("Result:", currentResult)
						}
						operator = 0
					} else {
						if debug {
							fmt.Printf("Multiplying %d by %d\n", currentResult, rightNumber)
						}
						currentResult *= rightNumber
						if debug {
							fmt.Println("Result:", currentResult)
						}
						operator = 0
					}
					completedTheLoop = false
					break
				} else if item[0] == ')' {
					continue
				} else if item[1] == ')' {
					// closing the sum
					// blah blah and return with the result
					rightNumber, _ = strconv.Atoi(string(item[0]))
					if operator == 1 {
						if debug {
							fmt.Printf("Adding %d to %d\n", currentResult, rightNumber)
						}
						currentResult += rightNumber
						if debug {
							fmt.Println("Result:", currentResult)
						}
						operator = 0 // don't think this is needed
						return currentResult, sumRemainder[strings.Index(sumRemainder, item)+1:]
					} else {
						if debug {
							fmt.Printf("Multiplying %d by %d\n", currentResult, rightNumber)
						}
						currentResult *= rightNumber
						if debug {
							fmt.Println("Result:", currentResult)
						}
						operator = 0 // don't think this is needed
						return currentResult, sumRemainder[strings.Index(sumRemainder, item)+1:]
					}
				}
			} else {
				switch item {
				case "+":
					if debug {
						fmt.Println("+")
					}
					operator = 1
					break
				case "*":
					if debug {
						fmt.Println("*")
					}
					operator = 2
					break
				case "(":
					if debug {
						fmt.Println("Bracket on its own. Probably never happens")
					}
					break
				case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
					if operator == 0 {
						currentResult, _ = strconv.Atoi(item)
						if debug {
							fmt.Println("No operator so grab number:", currentResult)
						}
					} else {
						rightNumber, _ = strconv.Atoi(item)
						if operator == 1 {
							if debug {
								fmt.Printf("Adding %d to %d\n", currentResult, rightNumber)
							}
							currentResult += rightNumber
							if debug {
								fmt.Println("Result:", currentResult)
							}
							operator = 0
						} else {
							if debug {
								fmt.Printf("Multiplying %d by %d\n", currentResult, rightNumber)
							}
							currentResult *= rightNumber
							if debug {
								fmt.Println("Result:", currentResult)
							}
							operator = 0
						}
					}
					break
				}
			}
		}
	}
	return currentResult, ""
}

func calcHomeworkAnswers(filename string, part byte, debug bool) int {
	var calculatedResult, tempResult int

	puzzleInput, _ := readFile(filename)

	for lineNum, sumLine := range puzzleInput {
		if debug {
			fmt.Println("--------- Line:", lineNum)
		}
		tempResult, _ = calcSumRecurse(sumLine, debug)
		fmt.Printf("Line: %d Result: %d\n", lineNum, tempResult)
		calculatedResult += tempResult
	}

	return calculatedResult
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
