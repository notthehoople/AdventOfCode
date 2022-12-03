package main

import (
	"AdventOfCode-go/advent2022/utils"
	"fmt"
)

func findTheElfBadge(filename string, part byte, debug bool) int {
	var totalPriority int

	puzzleInput, _ := utils.ReadFile(filename)

	for elfBag := 0; elfBag < len(puzzleInput); {
		var firstElfBag = puzzleInput[elfBag]
		var secondElfBag = puzzleInput[elfBag+1]
		var thirdElfBag = puzzleInput[elfBag+2]

		firstBagUniqueItem := make(map[byte]bool)
		for _, firstBagItem := range firstElfBag {
			firstBagUniqueItem[byte(firstBagItem)] = true
		}

		secondBagUniqueItem := make(map[byte]bool)
		for _, secondBagItem := range secondElfBag {
			secondBagUniqueItem[byte(secondBagItem)] = true
		}

		thirdBagUniqueItem := make(map[byte]bool)
		for _, thirdBagItem := range thirdElfBag {
			thirdBagUniqueItem[byte(thirdBagItem)] = true
		}

		var priority int
		for checkBagItem := range firstBagUniqueItem {
			if _, ok := secondBagUniqueItem[byte(checkBagItem)]; ok {
				if _, secondOk := thirdBagUniqueItem[byte(checkBagItem)]; secondOk {

					if checkBagItem >= 'A' && checkBagItem <= 'Z' {
						priority = int(checkBagItem - 38)
					} else {
						priority = int(checkBagItem - 96)
					}

					totalPriority += priority
					if debug {
						fmt.Printf("Item %c in all 3 bags, priority :%d\n", checkBagItem, priority)
					}

					break
				}
			}
		}
		elfBag += 3
	}

	return totalPriority
}

func elfRucksack(filename string, part byte, debug bool) int {

	var totalPriority int

	puzzleInput, _ := utils.ReadFile(filename)

	for _, bagContents := range puzzleInput {
		splitPoint := len(bagContents) / 2

		firstBagPocket := bagContents[:splitPoint]
		secondBagPocket := bagContents[splitPoint:]

		if debug {
			fmt.Println("1st bag", firstBagPocket)
			fmt.Println("2nd bag", secondBagPocket)
		}

		uniqueItem := make(map[byte]bool)
		for _, firstBagItem := range firstBagPocket {
			if _, ok := uniqueItem[byte(firstBagItem)]; !ok {
				uniqueItem[byte(firstBagItem)] = true
			}
		}

		var priority int
		for _, secondBagItem := range secondBagPocket {
			if _, ok := uniqueItem[byte(secondBagItem)]; ok {
				if secondBagItem >= 'A' && secondBagItem <= 'Z' {
					priority = int(secondBagItem - 38)
				} else {
					priority = int(secondBagItem - 96)
				}

				totalPriority += priority
				if debug {
					fmt.Printf("Item %c in both pockets, priority :%d\n", secondBagItem, priority)
					fmt.Printf("Thing %d %d\n", 'a'-96, 'A'-38)
				}

				break
			}
		}
	}

	return totalPriority
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Result is: %d\n", elfRucksack(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("Result is: %d\n", findTheElfBadge(filenamePtr, execPart, debug))
	case 'z':
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
