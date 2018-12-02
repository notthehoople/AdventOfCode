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

	//str := "a long string with many repeated characters"
	//numberOfa := strings.Count(str, "a")

	//fmt.Printf("[%v] string has %d of characters of [a] ", str, numberOfa)


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

		fmt.Println("Testing line:", tempString)

		for _, c := range "abcdefghijklmnopqrstuvwxyz" {

			countTemp := strings.Count(tempString, string(c))
//			fmt.Println("Number of: ", string(c), strings.Count(tempString, string(c)))

			if countTemp == 2 && alreadyDoneDouble == 0 {
				alreadyDoneDouble = 1
				doubleCounter++
				fmt.Println("Added one to doubleCounter:", doubleCounter, alreadyDoneDouble)
			}
			if countTemp == 3 && alreadyDoneTriple == 0 {
				alreadyDoneTriple = 1
				tripleCounter++
				fmt.Println("Added one to tripleCounter:", tripleCounter, alreadyDoneTriple)
			}
		}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	resultVar = doubleCounter * tripleCounter

	return resultVar
}

func main() {
	fileNamePtr := flag.String("file", "input1.txt", "A filename containing input strings")
	execPartPtr := flag.String("part", "a", "Which part of day02 do you want to calc (a or b)")

	flag.Parse()

	if *execPartPtr == "a" {
		fmt.Println("Part A - CheckSum is:", checkSum(*fileNamePtr))
	} else {
//		fmt.Println("Part B - Resulting Frequency:", seenTwice(*fileNamePtr))
		fmt.Println("Part B - Resulting Frequency:")

	}
}