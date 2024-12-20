package main

import (
	"AdventOfCode-go/advent2024/utils"
	"fmt"
	"strconv"
	"strings"
)

// correctPageOrder
func correctPageOrder(pages string, orderRules [100][100]bool) int {
	// Convert all the pages from string to int and store in pageNums
	// For each digit in the pageNums list, check that it is in the correct order based on the orderRules
	// if numbers are in the wrong order, swap the digits and start again

	pageSet := strings.Split(pages, ",")
	pageNums := make([]int, len(pageSet))
	for key, pageStr := range pageSet {
		page, _ := strconv.Atoi(pageStr)
		pageNums[key] = page
	}

	// There must be a better way of doing this, but it's late and I can't think of it. So....messing with loop variable values
	i := 1
	for i < len(pageNums) {
		for j := 0; j < i; j++ {
			if !orderRules[pageNums[j]][pageNums[i]] {
				pageNums[j], pageNums[i] = pageNums[i], pageNums[j]
				j = i
				i = 1
				continue
			}
		}
		i++
	}

	//fmt.Println("Middle number:", len(pageNums)/2)
	return pageNums[len(pageNums)/2]
}

// validatePageOrder
func validatePageOrder(pages string, orderRules [100][100]bool) int {
	// Convert all the pages from string to int and store in pageNums
	// For each digit in the pageNums list, check that it is in the correct order based on the orderRules
	// If all is in the correct order, return the middle digit

	pageSet := strings.Split(pages, ",")
	pageNums := make([]int, len(pageSet))
	for key, pageStr := range pageSet {
		page, _ := strconv.Atoi(pageStr)
		pageNums[key] = page
	}

	for i := 1; i < len(pageNums); i++ {
		for j := 0; j < i; j++ {
			//fmt.Printf("j: %d i: %d orderRule: %t\n", pageNums[j], pageNums[i], orderRules[pageNums[j]][pageNums[i]])
			if !orderRules[pageNums[j]][pageNums[i]] {
				return 0
			}
		}
	}

	//fmt.Println("Middle number:", len(pageNums)/2)
	return pageNums[len(pageNums)/2]
}

func day05(filename string, part byte, debug bool) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)

	// Process page ordering rules. Process until find a blank line

	var processPrint = false
	var orderRules [100][100]bool
	printPages := make([]string, 0)

	// Process the file into 2 slices - the rule set (orderRules) and the pages to print (printPages)
	for _, puzzleLine := range puzzleInput {
		if len(puzzleLine) == 0 {
			processPrint = true
		} else {
			if !processPrint {
				// Build the set of rules we'll use to check the following pages
				var page1, page2 int
				fmt.Sscanf(puzzleLine, "%d|%d", &page1, &page2)
				orderRules[page1][page2] = true
			} else {
				// Build the set of pages to be printed
				printPages = append(printPages, puzzleLine)
			}
		}
	}

	if part == 'a' {
		for _, pages := range printPages {
			result += validatePageOrder(pages, orderRules)
		}

		return result
	}

	// part b. Strangely easier than part a

	for _, pages := range printPages {
		// We only care about the pages that are INVALID. Check for INVALID page sets, then correct them when found
		if validatePageOrder(pages, orderRules) == 0 {
			result += correctPageOrder(pages, orderRules)
		}
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
