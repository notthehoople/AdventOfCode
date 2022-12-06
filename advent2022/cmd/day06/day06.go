package main

import (
	"AdventOfCode-go/advent2022/utils"
	"fmt"
	"strconv"
	"strings"
)

type Stack []string

// Check if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Append the new value onto the slice
}

// Pop the top element off the stack
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   //Get the inddex of the top most element
		element := (*s)[index] // Index into the slice and obtain the element
		*s = (*s)[:index]      // Remove it from the stack
		return element, true
	}
}

// Print the top element of the stack
func (s *Stack) PrintTopItem() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   //Get the inddex of the top most element
		element := (*s)[index] // Index into the slice and obtain the element
		return element, true
	}
}

// Part A:
func cargoArrangement(filename string, part byte, debug bool) string {

	//var stackSet[100] Stack

	puzzleInput, _ := utils.ReadFile(filename)

	stackSet := make([]Stack, 100)

	var cargoMovements int
	for pos, inputLine := range puzzleInput {

		// Build a stack for each column label
		// Use an array of stacks

		if inputLine[1] == '1' {

			cargoMovements = pos + 2
			stackNames := strings.Fields(inputLine)

			for _, name := range stackNames {
				if debug {
					fmt.Println("======================")
					fmt.Println("Printing stack for", name)
				}

				stackPos := strings.Index(inputLine, name)
				namePos, _ := strconv.Atoi(name)

				for i := pos - 1; i >= 0; i-- {
					if puzzleInput[i][stackPos] != ' ' {
						stackSet[namePos].Push(string(puzzleInput[i][stackPos]))
						if debug {
							fmt.Println(stackSet[namePos])
						}
					}
				}

				if debug {
					fmt.Printf("%s is at pos %d\n", name, strings.Index(inputLine, name))
				}
			}

			break
		}
	}

	if debug {
		fmt.Println("======= Starting =======")
		fmt.Println(stackSet[1])
		fmt.Println(stackSet[2])
		fmt.Println(stackSet[3])
		fmt.Println("======= On we go =======")
	}

	var numToMove, start, destination int
	for i := cargoMovements; i < len(puzzleInput); i++ {
		fmt.Sscanf(puzzleInput[i], "move %d from %d to %d\n", &numToMove, &start, &destination)

		if debug {
			fmt.Printf("Move %d from %d to %d\n", numToMove, start, destination)
		}

		// Carry out the movement instructions
		if part == 'a' {
			var item string
			for move := numToMove; move > 0; move-- {
				item, _ = stackSet[start].Pop()
				stackSet[destination].Push(item)
			}
		} else {
			// In part b we need to keep the order of items the same. So build a temporary list to hold the items
			moveItems := make([]string, numToMove+1)
			for move := numToMove; move > 0; move-- {
				moveItems[move], _ = stackSet[start].Pop()
			}
			// Now add the items in the reverse that they came off the stack so we maintain the order
			for move := 1; move < numToMove+1; move++ {
				stackSet[destination].Push(moveItems[move])
			}
		}
	}

	// Now build the string to return
	var resultString string
	for i := 1; i < len(stackSet); i++ {
		topItem, result := stackSet[i].PrintTopItem()
		if result {
			resultString = resultString + topItem
		}
	}

	return resultString
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Result is: %s\n", cargoArrangement(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("Result is: %s\n", cargoArrangement(filenamePtr, execPart, debug))
	case 'z':
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
