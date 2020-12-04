package main

import (
	"flag"
	"fmt"
	"strconv"
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

func validateNumber(numberToValidate string, part byte, numberDigits int, minValid int, maxValid int) bool {
	if part == 'a' {
		// In part a the validation is for the number to not be empty
		if numberToValidate != "" {
			return true
		}
		return false
	}

	if len(numberToValidate) != numberDigits {
		return false
	}

	number, convError := strconv.Atoi(numberToValidate)

	if minValid == 0 && maxValid == 0 {
		// Value doesn't need to be validated so just check it's a number

		if convError == nil {
			return true
		}
		return false
	}

	if convError == nil && number >= minValid && number <= maxValid {
		return true
	}

	return false
}

func validateHeight(heightToValidate string, part byte) bool {
	if part == 'a' {
		// In part a the validation is for the number to not be empty
		if heightToValidate != "" {
			return true
		}
		return false
	}

	var heightOnly int
	if strings.Index(heightToValidate, "in") > 0 {
		// Inches
		fmt.Sscanf(heightToValidate, "%din", &heightOnly)
		if heightOnly >= 59 && heightOnly <= 76 {
			return true
		}
		return false
	}

	if strings.Index(heightToValidate, "cm") > 0 {
		// cm
		fmt.Sscanf(heightToValidate, "%dcm", &heightOnly)
		if heightOnly >= 150 && heightOnly <= 193 {
			return true
		}
		return false
	}

	return false
}

func validateEyeColour(eyeColour string, part byte) bool {
	if part == 'a' {
		// In part a the validation is for the eye colour to not be empty
		if eyeColour != "" {
			return true
		}
		return false
	}

	switch eyeColour {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		return false
	}
}

func validateHairColour(hairColour string, part byte) bool {
	if part == 'a' {
		// In part a the validation is for the eye colour to not be empty
		if hairColour != "" {
			return true
		}
		return false
	}

	// must start with a # then followed by exactly 6 chars of 0-9 or a-f
	if len(hairColour) != 7 || hairColour[0] != '#' {
		return false
	}

	for i := 1; i < len(hairColour); i++ {
		if (hairColour[i] >= '0' && hairColour[i] <= '9') || (hairColour[i] >= 'a' && hairColour[i] <= 'f') {
			// it's a valid character
		} else {
			return false
		}
	}
	return true
}

func countValidPassports(passportList [1000]PassportStruct, part byte, debug bool) int {
	var validPassports int = 0

	for _, passport := range passportList {

		if debug {
			fmt.Println("=====================")
			fmt.Println(passport)
		}

		var numberGoodFields int = 0
		if validateNumber(passport.byr, part, 4, 1920, 2002) {
			numberGoodFields++
		}
		if validateNumber(passport.iyr, part, 4, 2010, 2020) {
			numberGoodFields++
		}
		if validateNumber(passport.eyr, part, 4, 2020, 2030) {
			numberGoodFields++
		}
		if validateNumber(passport.pid, part, 9, 0, 0) {
			numberGoodFields++
		}
		//we don't need to check cid
		//if passport.cid != 0 { numberGoodFields++ }
		if validateHeight(passport.hgt, part) {
			numberGoodFields++
		}
		if validateHairColour(passport.hcl, part) {
			numberGoodFields++
		}
		if validateEyeColour(passport.ecl, part) {
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

	return countValidPassports(passportList, part, debug)
}

// Main routine
func main() {
	filenamePtr, execPart, debug := catchUserInput()
	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Println("Number of valid passports: ", validatePassports(filenamePtr, execPart, debug))
	}

	if debug {
		fmt.Println("Eye Colour brn: ", validateEyeColour("brn", 'b'))
		fmt.Println("Eye Colour wat: ", validateEyeColour("wat", 'b'))

		fmt.Println("pid 000000001: ", validateNumber("000000001", 'b', 9, 0, 0))
		fmt.Println("pid 0123456789: ", validateNumber("0123456789", 'b', 9, 0, 0))

		fmt.Println("byr 2002: ", validateNumber("2002", 'b', 4, 1920, 2002))
		fmt.Println("byr 2003: ", validateNumber("2003", 'b', 4, 1920, 2002))

		fmt.Println("hcl #123abc: ", validateHairColour("#123abc", 'b'))
		fmt.Println("hcl #123abz: ", validateHairColour("#123abz", 'b'))
		fmt.Println("hcl 123abc: ", validateHairColour("123abc", 'b'))

		fmt.Println("hgt 60in: ", validateHeight("60in", 'b'))
		fmt.Println("hgt 190cm: ", validateHeight("190cm", 'b'))
		fmt.Println("hgt 190in: ", validateHeight("190in", 'b'))
		fmt.Println("hgt 190: ", validateHeight("190", 'b'))
	}
}
