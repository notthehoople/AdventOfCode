package main

import (
	"AdventOfCode-go/advent2022/utils"
	"fmt"
)

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

		if part == 'a' {

		} else {

		}
	}

	return totalPriority
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", elfRucksack(filenamePtr, execPart, debug))
	}
}
