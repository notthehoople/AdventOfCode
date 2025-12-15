package main

import (
	"AdventOfCode-go/advent2025/utils"
	"fmt"
	"strconv"
	"strings"
)

type freshRange struct {
	start int
	end   int
}

func smallest(first int, second int) int {
	if first < second {
		return first
	}
	return second
}

func largest(first int, second int) int {
	if first > second {
		return first
	}
	return second
}

func overlappingRange(first freshRange, second freshRange) bool {
	if ((first.end >= second.start) && (first.end <= second.end)) ||
		((first.start >= second.start) && (first.start <= second.end)) {
		return true
	}
	return false
}

func day05(filename string, part byte, debug bool) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)

	// Get the size of the fresh ingredients array that we need
	var countArraySize int
	for _, puzzleLine := range puzzleInput {
		if len(puzzleLine) == 0 {
			break
		}
		countArraySize++
	}
	if debug {
		fmt.Println("Array Size:", countArraySize)
	}

	ingredientDB := make([]freshRange, countArraySize)

	var checkIngredients bool = false
	for i, puzzleLine := range puzzleInput {
		if !checkIngredients {
			// Process fresh ingredient ID ranges until we see the blank line

			if len(puzzleLine) == 0 {
				if debug {
					fmt.Println("found the blank line")
				}
				checkIngredients = true
			} else {
				ingredientRange := strings.Split(puzzleLine, "-")
				ingredientDB[i].start, _ = strconv.Atoi(ingredientRange[0])
				ingredientDB[i].end, _ = strconv.Atoi(ingredientRange[1])

			}
		} else {

			if part == 'a' {
				// Check each ingredient in the fresh ingredient DB we previously built
				ingredient, _ := strconv.Atoi(puzzleLine)
				for _, freshStuff := range ingredientDB {
					if ingredient >= freshStuff.start && ingredient <= freshStuff.end {
						result++
						break
					}
				}
			} else {
				break
			}
		}
	}

	if part == 'a' {
		return result
	}

	// part b. Don't process the ingredient list. Focus on the DB

	// Need to work out how many unique fresh IDs there are
	// Use the ranges in the DB
	// look for overlapping ranges and COMBINE them
	// once everything has been combined, add up the remaining IDs

	expandedDB := make(map[freshRange]bool, 0)

	for _, ingredientRange := range ingredientDB {
		var didOverlap bool = false

		for expandedRange := range expandedDB {

			if overlappingRange(ingredientRange, expandedRange) {
				// Combine the ranges
				didOverlap = true

				overlapRange := freshRange{smallest(ingredientRange.start, expandedRange.start), largest(ingredientRange.end, expandedRange.end)}
				expandedDB[overlapRange] = true

				if overlapRange != ingredientRange {
					delete(expandedDB, ingredientRange)
				}

				if overlapRange != expandedRange {
					delete(expandedDB, expandedRange)
				}
			}
		}

		// If no overlap then just add that ingredient range to our list
		if !didOverlap {
			expandedDB[ingredientRange] = true
		}
	}

	// Now repeatedly loop through our expanded range looking for ranges that might overlap. Combine them, then look again
	var madeAChange bool = true
	for madeAChange {
		madeAChange = false
		for firstItem := range expandedDB {
			for secondItem := range expandedDB {
				if firstItem != secondItem {
					if overlappingRange(firstItem, secondItem) {
						// Combine the ranges

						madeAChange = true

						overlapRange := freshRange{smallest(firstItem.start, secondItem.start), largest(firstItem.end, secondItem.end)}
						expandedDB[overlapRange] = true

						if overlapRange != firstItem {
							delete(expandedDB, firstItem)
						}

						if overlapRange != secondItem {
							delete(expandedDB, secondItem)
						}

					}
				}
			}
		}
	}

	// Sum up all the fresh ingredients in all the ranges
	for i := range expandedDB {
		result += i.end - i.start + 1
	}

	return result
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", day05(filenamePtr, execPart, debug))
	}
}

/*

List 1 - original list of stuff
List 2 - new, expanded list of stuff

For each entry in List 1
  check through List 2
    if nothing overlapping, append to List 2
	if overlapping
	  change overlapping entry to include List 1 entry
	    check through all of List 2
		  doesn't anything else now overlap?
		  if overlap
		    combine entries and remove 1 of them

*/
