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

// Handles everything needed to work out the polymerLength (Day05 part A)
func processLicenseFile(fileName string, part string) int {
	var systemLicenseString string
	var currentPosInList int = 0
	var currentTotal int = 0

	// Read contents of file into a string array
	fileContents, _ := readLines(fileName)
	systemLicenseString = fileContents[0]

	systemLicense := strings.Split(systemLicenseString, " ")

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
	if part == "a" {
		currentPosInList, currentTotal = processNode(systemLicense, currentPosInList, currentTotal)
		
		return currentTotal

	} else {
		fmt.Println("Doing nothing. You shouldn't be here yet")
		return 0

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
		fmt.Println("Not ready yet")
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}