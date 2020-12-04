package main

import (
	"flag"
	"fmt"
	"strings"
)

type PassportStruct struct {
	byr string // Birth Year
	iyr string // Issue Year
	eyr string // Expiration Year
	hgt string // Height (can include 'cm' or 'in')
	hcl string // Hair Colour
	ecl string // Eye Colour
	pid string // Passport ID
	cid string // Country ID
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

func countValidPassports(passportList [1000]PassportStruct) int {
	var validPassports int = 0

	for _, passport := range passportList {
		var numberGoodFields int = 0
		if passport.byr != "" {
			numberGoodFields++
		}
		if passport.iyr != "" {
			numberGoodFields++
		}
		if passport.eyr != "" {
			numberGoodFields++
		}
		if passport.pid != "" {
			numberGoodFields++
		}
		//we don't need to check cid
		//if passport.cid != 0 { numberGoodFields++ }
		if passport.hgt != "" {
			numberGoodFields++
		}
		if passport.hcl != "" {
			numberGoodFields++
		}
		if passport.ecl != "" {
			numberGoodFields++
		}

		if numberGoodFields == 7 {
			validPassports++
		}
	}

	return validPassports
}

func validatePassports(filename string, part byte, debug bool) int {
	var processingPassport int = 0
	var passportList [1000]PassportStruct

	puzzleInput, _ := readFile(filename)

	for _, passportLine := range puzzleInput {
		if len(passportLine) == 0 {
			processingPassport++
		} else {
			// Do something with the passport line we've found

			for _, item := range strings.Split(passportLine, " ") {
				if debug {
					fmt.Println("Item: ", item)
				}
				fmt.Sscanf(item, "byr:%s", &passportList[processingPassport].byr)
				fmt.Sscanf(item, "iyr:%s", &passportList[processingPassport].iyr)
				fmt.Sscanf(item, "eyr:%s", &passportList[processingPassport].eyr)
				fmt.Sscanf(item, "pid:%s", &passportList[processingPassport].pid)
				fmt.Sscanf(item, "cid:%s", &passportList[processingPassport].cid)
				fmt.Sscanf(item, "hgt:%s", &passportList[processingPassport].hgt)
				fmt.Sscanf(item, "hcl:%s", &passportList[processingPassport].hcl)
				fmt.Sscanf(item, "ecl:%s", &passportList[processingPassport].ecl)
			}

		}
	}

	if debug {
		fmt.Println(passportList)
	}

	return countValidPassports(passportList)
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
