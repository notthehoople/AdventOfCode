package main

import (
	"AdventOfCode-go/advent2022/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type monkeyObj struct {
	startingItems   []int64
	operationSize   int64
	operationType   byte
	testDivisor     int64
	trueMonkey      int
	falseMonkey     int
	inspectionCount int
}

func (m monkeyObj) printMonkey(monkey int) {
	fmt.Printf("Monkey %d:\n", monkey)

	fmt.Printf("  Starting items: ")
	for key := range m.startingItems {
		fmt.Printf("%d,", m.startingItems[key])
	}

	fmt.Printf("\n  Operation: new = old ")
	switch m.operationType {
	case '*':
		// Note if operationSize is 0 then the operation is old * old
		if m.operationSize == 0 {
			fmt.Printf("* old\n")
		} else {
			fmt.Printf("* %d\n", m.operationSize)
		}
	case '/':
		fmt.Printf("/ %d\n", m.operationSize)
	case '+':
		fmt.Printf("+ %d\n", m.operationSize)
	case '-':
		fmt.Printf("- %d\n", m.operationSize)
	}
	fmt.Printf("  Test: divisible by %d\n", m.testDivisor)
	fmt.Printf("    If true: throw to monkey %d\n", m.trueMonkey)
	fmt.Printf("    If false: throw to monkey %d\n", m.falseMonkey)
	fmt.Printf("  Inspection Count: %d\n", m.inspectionCount)
}

func buildArrayOfMonkeys(puzzleInput []string) []monkeyObj {
	// Get the number of monkeys so we can build an array of them
	var monkeyCount int
	for _, i := range puzzleInput {
		if len(i) > 5 {
			if i[0:6] == "Monkey" {
				monkeyCount++
			}
		}
	}

	//fmt.Println("Monkeys:", monkeyCount)

	monkeyArray := make([]monkeyObj, monkeyCount)

	var currentMonkey int
	for _, line := range puzzleInput {
		var startingItems string
		fmt.Sscanf(line, "Monkey %d:\n", &currentMonkey)

		if strings.Contains(line, "  Starting items: ") {
			startingItems = strings.TrimPrefix(line, "  Starting items: ")
			startingItemArray := strings.Split(startingItems, ",")
			monkeyArray[currentMonkey].startingItems = make([]int64, len(startingItemArray))
			for key, i := range startingItemArray {
				tmpStartItem, _ := strconv.Atoi(strings.TrimSpace(i))
				monkeyArray[currentMonkey].startingItems[key] = int64(tmpStartItem)
			}

			//fmt.Printf("Monkey %d Starting Items %s\n", currentMonkey, startingItems)
			//fmt.Println(startingItemArray)
		}

		//fmt.Sscanf(line, "  Starting items: %100s\n", &startingItems) // How to do multiples here?
		fmt.Sscanf(line, "  Operation: new = old %c %d\n", &monkeyArray[currentMonkey].operationType, &monkeyArray[currentMonkey].operationSize)
		fmt.Sscanf(line, "  Test: divisible by %d\n", &monkeyArray[currentMonkey].testDivisor)
		fmt.Sscanf(line, "    If true: throw to monkey %d\n", &monkeyArray[currentMonkey].trueMonkey)
		fmt.Sscanf(line, "    If false: throw to monkey %d\n", &monkeyArray[currentMonkey].falseMonkey)

	}

	return monkeyArray
}

func calcWorryLevel(item int64, operationType byte, operationSize int64) int64 {
	var worryLevel int64

	switch operationType {
	case '*':
		// Note if operationSize is 0 then the operation is old * old
		if operationSize == 0 {
			worryLevel = item * item
		} else {
			worryLevel = item * operationSize
		}
	case '/':
		worryLevel = item / operationSize
	case '+':
		worryLevel = item + operationSize
	case '-':
		worryLevel = item - operationSize
	}
	return worryLevel
}

func calcMonkeyBusiness(filename string, part byte, debug bool) int {
	var primeMod int64 = 1

	puzzleInput, _ := utils.ReadFile(filename)

	monkeyArray := buildArrayOfMonkeys(puzzleInput)
	/*
		for key, monkey := range monkeyArray {
			monkey.printMonkey(key)
		}
	*/
	var rounds int

	if part == 'a' {
		rounds = 20
	} else {
		rounds = 10000
		// Loop through the monkeys and extract the divisors. We'll use this later to keep the worry down via mod
		for _, monkey := range monkeyArray {
			primeMod *= monkey.testDivisor
		}
	}

	// TODO: Mod the worrylevel by the product of the test divisors
	// TODO: understand why this works. All the test divisors are prime numbers

	// - count the total number of times each monkey inspects an item over 20 rounds
	for i := 0; i < rounds; i++ {
		// On a single Monkey's turn:
		for monkeyNumber, currMonkey := range monkeyArray {
			// - inspect each item in turn in the order listed:
			for itemKey, item := range currMonkey.startingItems {
				if item != 0 {
					//    - monkey INSPECTS an item in the list
					//    - worry level is item <operationType> <operationSize> e.g. 79 * 6
					worryLevel := calcWorryLevel(item, currMonkey.operationType, currMonkey.operationSize)
					//    - worry level divides by 3 and rounded DOWN to nearest integer
					if part == 'a' {
						worryLevel = worryLevel / 3
					} else {
						worryLevel = worryLevel % primeMod
					}

					//    - TEST the item based on the <testDivisor>
					// - throws to the indicated Monkey which adds the item to the END of their list
					if worryLevel%currMonkey.testDivisor == 0 {
						//       - if TRUE, send the current worrylevel for the item to trueMonkey
						monkeyArray[currMonkey.trueMonkey].startingItems = append(monkeyArray[currMonkey.trueMonkey].startingItems, worryLevel)
						monkeyArray[monkeyNumber].startingItems = monkeyArray[monkeyNumber].startingItems[1:len(monkeyArray[monkeyNumber].startingItems)]
					} else {
						//       - if FALSE, send the current worrylevel for the item to falseMonkey
						monkeyArray[currMonkey.falseMonkey].startingItems = append(monkeyArray[currMonkey.falseMonkey].startingItems, worryLevel)
					}
					currMonkey.startingItems[itemKey] = 0
					monkeyArray[monkeyNumber].inspectionCount++
				}
			}
			// - if a monkey starts its turn with no items it's turn ends immediately
		}
	}

	resultArray := make([]int, len(monkeyArray))
	for key, monkey := range monkeyArray {
		if debug {
			monkey.printMonkey(key)
		}

		resultArray[key] = monkey.inspectionCount
	}
	sort.Sort(sort.Reverse(sort.IntSlice(resultArray)))

	return resultArray[0] * resultArray[1]
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Monkey Business: %d\n", calcMonkeyBusiness(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("Monkey Business: %d\n", calcMonkeyBusiness(filenamePtr, execPart, debug))
	case 'z':
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
