package main

import (
	"fmt"
	"bufio"
	"os"
	"flag"
//	"sort"
//	"strconv"
	"unicode"
//	"strings"
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

// Carries out the destruction of Polymer units
// Situations to take account of:
//     Adjacent letters of differing capitalisation destroy themselves
//     Adjacent letters of the same capitalisation do nothing
//       In aA, a and A react, leaving nothing behind.
//       In abBA, bB destroys itself, leaving aA. As above, this then destroys itself, leaving nothing.
//       In abAB, no two adjacent units are of the same type, and so nothing happens.
//       In aabAAB, even though aa and AA are of the same type, their polarities match, so nothing happens.
//     After pass 1 destruction there may be new pairings that will again destroy themselves
func actionPolymerDestroy(tempPolymer string)(int, string) {
	var tempReturnPolymer string
	var numDestroyed int = 0
	var j int = 0

	//tempReturnPolymer = tempPolymer

	// Loop through
	// Test for characters being the same
	// Test for upper vs lower or lower vs upper
	// If match then destroy
	// If match then need to ignore future matches, just build temp string
	// Return that we did a destroy

	for i := 0; i < len(tempPolymer); i++ {
		if i > 0 {
			currentChar := tempPolymer[i]
			previousChar := tempPolymer[i-1]

			if unicode.ToLower(rune(currentChar)) != unicode.ToLower(rune(previousChar)) {
				tempReturnPolymer += string(previousChar)
				// Note - this doesn't currently deal with the last character of the string

			} else {
				// Do proper checking here to see if it's a destroy
				i++
			}
		}
	}


	for i, currentChar := range tempPolymer {
		if i > 0 && numDestroyed == 0 {
			previousChar := tempPolymer[i-1]
			fmt.Printf("Loop: %d Char: %s previous Char: %s\n", i, string(currentChar), string(previousChar))

			if unicode.ToLower(rune(currentChar)) == unicode.ToLower(rune(previousChar)) {
				if unicode.IsLower(rune(currentChar)) && unicode.IsUpper(rune(previousChar)) {
					fmt.Printf("Destroy found: %s and %s\n", string(previousChar), string(currentChar))
					numDestroyed++

					// what to do with return string here?
				} else {
					if unicode.IsUpper(rune(currentChar)) && unicode.IsLower(rune(previousChar)) {
						fmt.Printf("Destroy found: %s and %s\n", string(previousChar), string(currentChar))
						numDestroyed++

						// what to do with return string here?
					} else {
						// If previous and current chars are the same but not destroying, add previous to return string
						tempReturnPolymer += string(previousChar)
					}
				}
			} else {
				// there's no match of strings here so add previous char to return string
				tempReturnPolymer += string(previousChar)
			}
		}
		fmt.Println(i, " => ", string(currentChar))
	}

	fmt.Println("tempReturnPolymer: ", tempReturnPolymer)
	return numDestroyed, tempReturnPolymer
}

// Handles everything needed to work out the polymerLength (Day05 part A)
func polymerLength(fileName string, part string) int {
	var numDestroyed int = 0

	// Read contents of file into a string array
	fileContents, _ := readLines(fileName)

	numDestroyed, reducedPolymer := actionPolymerDestroy(fileContents[0])
//	for ; numDestroyed > 0; {
//		numDestroyed, reducedPolymer := actionPolymerDestroys(reducedPolymer)
//		fmt.Printf("numDestroyed: %d reducedPolymer: %s", numDestroyed, reducedPolymer)
//	}

	fmt.Printf("Length of string: %d\n", len(fileContents[0]))
	fmt.Printf("num destroyed: %d reducedPolymer: %s", numDestroyed, reducedPolymer)

	return 1
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
		fmt.Println("Not there yet")
	}
}