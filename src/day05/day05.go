package main

import (
	"fmt"
	"bufio"
	"os"
	"flag"
//	"sort"
//	"strconv"
	"unicode"
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

// Remove the letter "tempLetter" from the string "tempPolymer" then return the result
func removeUnitType(tempLetter string, tempPolymer string)(string) {

	tempPolymer = strings.Replace(tempPolymer, strings.ToLower(tempLetter), "", -1)
	tempPolymer = strings.Replace(tempPolymer, strings.ToUpper(tempLetter), "", -1)

	return tempPolymer
}

// Carries out the destruction of Polymer units
// Situations to take account of:
//     Adjacent letters of differing capitalisation destroy themselves
//     Adjacent letters of the same capitalisation do nothing
//       In aA, a and A react, leaving nothing behind.
//       In abBA, bB destroys itself, leaving aA. As above, this then destroys itself, leaving nothing.
//       In abAB, no two adjacent units are of the same type, and so nothing happens.
//       In aabAAB, even though aa and AA are of the same type, their polarities match, so nothing happens.
//     After pass 1 destruction there may be new pairings that will again destroy themselves
func actionPolymerDestroy(tempPolymer string)(bool, string) {
	var tempReturnPolymer string
	var didDestroy bool = false
	var treatAsFirstChar bool = true
	var previousChar byte
	var currentChar byte

	for i := 0; i < len(tempPolymer); i++ {

		if treatAsFirstChar {
			treatAsFirstChar = false
			continue
		}
		currentChar = tempPolymer[i]
		previousChar = tempPolymer[i-1]

		// Are previous and current characters the same, ignoring case?
		if unicode.ToLower(rune(currentChar)) != unicode.ToLower(rune(previousChar)) {
			// No, so we can return the previousChar as it won't be destroyed
			tempReturnPolymer += string(previousChar)
		} else {
			// Are they the same *including* case?
			if currentChar == previousChar {
				// Yes, so we can return the previousChar as it won't be destroyed
				tempReturnPolymer += string(previousChar)
			} else {
				// They are different case but the same character so will be destroyed
				treatAsFirstChar = true
				didDestroy = true
				continue
			}
		}
	}
	if !treatAsFirstChar {
		tempReturnPolymer += string(currentChar)
	}

	return didDestroy, tempReturnPolymer
}

// Handles everything needed to work out the polymerLength (Day05 part A)
func polymerLength(fileName string, part string) int {
	var didDestroy bool = true
	var reducedPolymer string
	var polymerReductionResults = make(map[string]int)
	var alphabet string = "abcdefghijklmnopqrstuvwxyz"
	var shortestPolymer int = 0

	// Read contents of file into a string array
	fileContents, _ := readLines(fileName)
	reducedPolymer = fileContents[0]
	shortestPolymer = len(fileContents[0])

	if part == "a" {
		for ; didDestroy; {
			didDestroy, reducedPolymer = actionPolymerDestroy(reducedPolymer)
		}
		return len(reducedPolymer)

	} else {
		for i := 0; i < len(alphabet); i++ {
			reducedPolymer = fileContents[0]
			reducedPolymer = removeUnitType(string(alphabet[i]), reducedPolymer)

			didDestroy = true
			for ; didDestroy; {
				didDestroy, reducedPolymer = actionPolymerDestroy(reducedPolymer)
			}
			polymerReductionResults[string(alphabet[i])] = len(reducedPolymer)
		}

		for _, tempval := range polymerReductionResults {
			if shortestPolymer > tempval {
				shortestPolymer = tempval
			}
		}
		return shortestPolymer

	}

	return 0
}

// Main routine
func main() {
	fileNamePtr := flag.String("file", "input1.txt", "A filename containing input strings")
	execPartPtr := flag.String("part", "a", "Which part of day05 do you want to calc (a or b)")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Length of Polymer:", polymerLength(*fileNamePtr, "a"))
	case "b":
		fmt.Println("Part b - After removing one type, shortest polymer:", polymerLength(*fileNamePtr, "b"))
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}