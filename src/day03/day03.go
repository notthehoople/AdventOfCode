package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"flag"
	"strings"
	"strconv"
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

//func printIntArray(tempIntArray [][]int) {
//	// Loop through the array and print each line
//	for i:= 0; i < len(tempIntArray); i++ {
//		fmt.Println(tempIntArray[i])
//	}
//}

// Handles everything needed to work out the fabric claim (Day03 part A)
func fabricClaim(fileName string) int {

	// A claim like #123 @ 3,2: 5x4 means that claim ID 123 specifies a rectangle 3 inches from the
	// left edge, 2 inches from the top edge, 5 inches wide, and 4 inches tall

	// We read through the file once, reading in each Elf's fabric claim
	// Each line is #<num> @ <x>,<y>: <a>x<b>
	//   #<num> - claim number
	//   @ - seperator. We don't care about this
	//   <x> - x inches from the left edge
	//   <y> - y inches from the top edge
	//   <a> - a inches wide
	//   <b> - b inches tall

	var resultVar int = 0					// Defining the overall result Variable
	var a, b, x, y int = 0, 0, 0, 0			// Co-ords and size of claim
	var lengthx, lengthy int = 0, 0

	fabricMap := [2000][2000]int{}

	// Read contents of file into a string array
	fileContents, _ := readLines(fileName)

	// Print out what we've read from the file
	// printStringArray(fileContents)

	// Print the fabricMap starting position
	// fmt.Println(fabricMap)

	// Loop through the string array; break into component parts then apply to our fabricMap

	for i := 0; i < len(fileContents); i++ {
		words := strings.Fields(fileContents[i])

		for j := 0; j < len(words); j++ {
			switch j {
			case 0:	// First entry is the claim number
			case 1: // Second entry is the '@'
			case 2: // Third entry is the start co-ordinates <x>,<y>
				tempNumber := strings.Split(words[j], ":")
				numberString := strings.Split(tempNumber[0], ",")
				x, _ = strconv.Atoi(numberString[0])
				y, _ = strconv.Atoi(numberString[1])

			//	fmt.Println("Number <x>:", x)
			//	fmt.Println("Number <y>:", y)
			case 3: // Fourth entry is the size <a>x<b>
				numberString := strings.Split(words[j], "x")
				a, _ = strconv.Atoi(numberString[0])
				b, _ = strconv.Atoi(numberString[1])

			//	fmt.Println("Number <a>:", a)
			//	fmt.Println("Number <b>:", b)
			}
		}

		// Now we have the variables we need it's time to modify the fabricMap
		// Loop through the fabricMap adding 1 to each relevant area found


		// k is our "y" axis in the map
		for k := y; k < y + b; k++ {
			// l is the "x" axis in the map. We need to modify from x through to x+a
			for l := x; l < x + a; l++ {
				fabricMap[l][k]++
			}
		}

		//printIntArray(fabricMap)
		//for i:= 0; i < len(fabricMap); i++ {
		//	fmt.Println("Array entry:",i,fabricMap[i])
		//}
	}	

	// Once complete, our result is:
	//    - loop through the fabricMap
	//    - count the number of entries > 1
	resultVar = 0
	lengthy = len(fabricMap)
	lengthx = len(fabricMap[0])
	for k := 0; k < lengthx; k++ {
		// l is the "x" axis in the map. We need to modify from x through to x+a
		for l := 0; l < lengthy; l++ {
			if fabricMap[l][k] > 1 {
				resultVar++
			}
		}
	}

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
		fmt.Println("Square inches in two or more claims:", fabricClaim(*fileNamePtr))
	} else {
		fmt.Println("Not started yet:")
	}
}