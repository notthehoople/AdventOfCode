package main

import (
	"fmt"
	"bufio"
	"os"
	"flag"
	"strconv"
	"strings"
)

// Read the text file passed in by name into a array of strings
// Returns the array as the first return variable
func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
	  return nil, err
	}
	defer file.Close()
  
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	  lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func printStringArray(tempString []string) {
	// Loop through the array and print each line
	for i:= 0; i < len(tempString); i++ {
		fmt.Println(tempString[i])
	}
}

// func processNodes (part A)
//
// We're not going to build a clever structure here. Just walk the license file recursively
// and get the values we need. Bascially:
//
// Function has 3 parameters: list of strings, current position in list, current total
// Function returns 2 parameters: current position in list, current total
//
// - Call function with our current position in the list of strings and the current total
//   - function reads from current position in the list of strings.
//   - if there's a new node, call function with current position and current total
//   - if no new node, use returned current position to read our meta data and add to current total
//     - return current position in list, current total
func processNode(systemLicense []string, currentPosInList int, currentTotal int) (int, int) {
	var numChildren int
	var numMetaData int

	numChildren, _ = strconv.Atoi(systemLicense[currentPosInList])
	numMetaData, _ = strconv.Atoi(systemLicense[currentPosInList+1])
	currentPosInList += 2

	for numChildren > 0 {
		currentPosInList, currentTotal = processNode (systemLicense, currentPosInList, currentTotal)
		numChildren--
	}

	for numMetaData > 0 {
		tempCurrentTotal, _ := strconv.Atoi(systemLicense[currentPosInList])
		currentTotal += tempCurrentTotal
		currentPosInList++
		numMetaData--
	}

	return currentPosInList, currentTotal
}

// func processChild (part B)
//
// We're not going to build a clever structure here. Just walk the license file recursively
// and get the values we need. Bascially:
//
// Function has 3 parameters: list of strings, current position in list, current total
// Function returns 2 parameters: current position in list, current total
//
// - Call function with our current position in the list of strings and the current total
//   - function reads from current position in the list of strings.
//     - if a node has no children, it's value is the sum of its metas
//     - if a node has children, the meta data is pointers to its children that make up its value
//       - if meta is 0 or higher than nodes number of children, skip it
//       - if a child is referenced multiple times, it counts multiple times towards the value
func processChild(systemLicense []string, currentPosInList int, currentTotal int) (int, int) {
	var numChildren int
	var tempNumChildren int
	var tempNumMetaData int
	var numMetaData int
	var childValues[1000]int

	numChildren, _ = strconv.Atoi(systemLicense[currentPosInList])
	numMetaData, _ = strconv.Atoi(systemLicense[currentPosInList+1])
	currentPosInList += 2

	if numChildren == 0 {
		for tempNumMetaData = numMetaData; tempNumMetaData > 0 ; {
			tempCurrentTotal, _ := strconv.Atoi(systemLicense[currentPosInList])
			currentTotal += tempCurrentTotal
			currentPosInList++
			tempNumMetaData--
		}
		return currentPosInList, currentTotal
	}

	for tempNumChildren = 1; tempNumChildren <= numChildren ; {
		currentPosInList, childValues[tempNumChildren] = processChild (systemLicense, currentPosInList, childValues[tempNumChildren])
		tempNumChildren++
	}

	for numMetaData > 0 {
		tempMetaDataItem, _ := strconv.Atoi(systemLicense[currentPosInList])
		if tempMetaDataItem == 0 || tempMetaDataItem > numChildren {
			currentPosInList++
			numMetaData--
		} else {
			currentTotal += childValues[tempMetaDataItem]
			currentPosInList++
			numMetaData--
		}
	}

	return currentPosInList, currentTotal
}

// Handles everything needed to work out the system license file (day 08)
func processLicenseFile(fileName string, part string) int {
	var systemLicenseString string
	var currentPosInList int = 0
	var currentTotal int = 0

	// Read contents of file into a string array
	fileContents, _ := readLines(fileName)
	systemLicenseString = fileContents[0]

	systemLicense := strings.Split(systemLicenseString, " ")

	if part == "a" {
		currentPosInList, currentTotal = processNode(systemLicense, currentPosInList, currentTotal)
		
		return currentTotal

	} else {
		currentPosInList, currentTotal = processChild(systemLicense, currentPosInList, currentTotal)
		
		return currentTotal
	}

	return 0
}

// Main routine
func main() {
	fileNamePtr := flag.String("file", "input1.txt", "A filename containing input strings")
	execPartPtr := flag.String("part", "a", "Which part of day08 do you want to calc (a or b)")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - License File checksum:", processLicenseFile(*fileNamePtr, "a"))
	case "b":
		fmt.Println("Part b - License File checksum:", processLicenseFile(*fileNamePtr, "b"))
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}