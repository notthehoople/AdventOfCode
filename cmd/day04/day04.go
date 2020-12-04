package main

import (
	"flag"
	"fmt"
)

type PassportStruct struct {
	byr int    // Birth Year
	iyr int    // Issue Year
	eyr int    // Expiration Year
	hgt string // Height (can include 'cm' or 'in')
	hcl string // Hair Colour
	ecl string // Eye Colour
	pid int    // Passport ID
	cid int    // Country ID
}

func catchUserInput() (string, byte, bool) {
	var debug bool

	filenamePtr := flag.String("file", "testInput.txt", "Filename containing the program to run")
	execPartPtr := flag.String("part", "a", "Which part of the puzzle do you want to calc (a or b)")
	flag.BoolVar(&debug, "debug", false, "Turn debug on")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		return *filenamePtr, 'a', debug
	case "b":
		return *filenamePtr, 'b', debug

	default:
		return *filenamePtr, 'z', debug
	}
}

func validatePassports(filename string, part byte, debug bool) int {
	var numValidPassports int = 0
	var processingPassport int = 0
	var passportList [10]PassportStruct

	puzzleInput, _ := readFile(filename)

	// When we start processing it's the first passport
	// We'll finish processing the last passport and not see a finish break
	for _, passportLine := range puzzleInput {
		if len(passportLine) == 0 {
			processingPassport++
		} else {
			// Do something with the passport line we've found
			fmt.Sscanf(passportLine, "eyr:%d", passportList[processingPassport].eyr)
		}
		//fmt.Sscanf(passwordLine, "%d-%d %c: %s", &minNumber, &maxNumber, &passwordChar, &password)

	}

	fmt.Println(passportList)
	return numValidPassports
}

// Main routine
func main() {
	filenamePtr, execPart, debug := catchUserInput()
	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		if execPart == 'a' {
			fmt.Println("Number of valid passports: ", validatePassports(filenamePtr, execPart, debug))
		} else {
			fmt.Println("Not implemented yet")
		}
	}
}
