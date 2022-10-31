package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"flag"
	"strings"
)

func checkSum(fileName string) int {
	var resultVar int = 0					// Defining the overall result Variable
	var tempString string					// Holds the line read from the file
	var doubleCounter int = 0				// Counter of number of double chars seen
	var tripleCounter int = 0				// Counter of number of triple chars seen
	var alreadyDoneDouble int = 0
	var alreadyDoneTriple int = 0

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// We read through the file once, working out the checksum component of each line
	// This is then multiplied to the running total to build our checksum

	// Need to COUNT the number of times a string has a character repeated exactly TWICE
	// Need to COUNT the number of times a string has a character repeated exactly THREE TIMES
	// If a line has more than one double, or more than one triple, it only counts once
	// We then multiple the COUNTs together to get our checksum

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		alreadyDoneDouble = 0
		alreadyDoneTriple = 0

		tempString = scanner.Text()

//		fmt.Println("Testing line:", tempString)

		for _, c := range "abcdefghijklmnopqrstuvwxyz" {

			countTemp := strings.Count(tempString, string(c))
//			fmt.Println("Number of: ", string(c), strings.Count(tempString, string(c)))

			if countTemp == 2 && alreadyDoneDouble == 0 {
				alreadyDoneDouble = 1
				doubleCounter++
//				fmt.Println("Added one to doubleCounter:", doubleCounter, alreadyDoneDouble)
			}
			if countTemp == 3 && alreadyDoneTriple == 0 {
				alreadyDoneTriple = 1
				tripleCounter++
//				fmt.Println("Added one to tripleCounter:", tripleCounter, alreadyDoneTriple)
			}
		}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	resultVar = doubleCounter * tripleCounter

	return resultVar
}

func closeIDs(fileName string) (string, string) {
	var lineA string						// Holds the line read from the file
	var lineB string						// Holds the comparison line read from the 2nd loop of the file
	var differencesCount int = 0			// Count of differences between current line and line in file

	// Confident that your list of box IDs is complete, you're ready to find the boxes full of prototype fabric.
	// The boxes will have IDs which differ by exactly one character at the same position in both strings.

	// Loop through file to get lineA
	//    Loop through file to get line B
	//        For each character in lineA, compare against line B until we get more than 1 difference
	//        Move on to next line in file
	// At anytime when we find a line that has exactly 1 difference we can quit

	fileA, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fileA.Close()

	scanner := bufio.NewScanner(fileA)
	
	// Top level loop starting
    for scanner.Scan() {

		lineA = scanner.Text()

		fileB, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		scannerB := bufio.NewScanner(fileB)

		for scannerB.Scan() {
			differencesCount = 0

			lineB = scannerB.Text()

			// Make sure we don't compare a line with itself
			if lineA != lineB {
				// Loop through the characters in lineA and compare with the character at the same pos in lineB
				// If it's different, count it. If we only get 1 difference in the whole line we've found it.
				
				for i, c := range lineA {
					if string(c) != string(lineB[i]) {
						differencesCount++
					}
				}
			}

			if differencesCount == 1 {
				break
			}
		}

		if differencesCount == 1 {
			break
		}		

		fileB.Close()
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	if differencesCount == 1 {
		return lineA, lineB
	} else {
		return "No IDs Found", "No IDs Found"
	}
}

func main() {
	fileNamePtr := flag.String("file", "input1.txt", "A filename containing input strings")
	execPartPtr := flag.String("part", "a", "Which part of day02 do you want to calc (a or b)")

	flag.Parse()

	if *execPartPtr == "a" {
		fmt.Println("Part A - CheckSum is:", checkSum(*fileNamePtr))
	} else {
		firstBoxID, secondBoxID := closeIDs(*fileNamePtr)
		fmt.Println("Part B - Prototype Clothing is in:", firstBoxID, secondBoxID)

	}
}