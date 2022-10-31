package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"flag"
	"strconv"
)

func addAll(fileName string) int {
	var resultVar int = 0					// Defining the overall result Variable
	var tempString string					// Holds the line read from the file

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// We read through the file once, adding up all the numbers.
	// Since the numbers come through as text we need to convert them to integers

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		tempString = scanner.Text()
		tempVal, _ := strconv.Atoi(tempString)
		resultVar += tempVal
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return resultVar
}

func seenTwice(fileName string) int {
	var resultVar int = 0					// Defining the overall result Variable
	var tempString string					// Holds the line read from the file
	var seenBefore int = 0					// Value of the first value seen twice

	m := make(map[int]string)				// A map containing all the seen values

	// We're going to read the file again and again until we see something twice
	// This could take a long time.
	for seenBefore == 0 {

		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(file)
		
    	for scanner.Scan() {
			tempString = scanner.Text()
			tempVal, _ := strconv.Atoi(tempString)
			resultVar += tempVal

			_, ok := m[resultVar]
			if ok == false {
				m[resultVar] = "Yes"
				// fmt.Println("Elements:", resultVar, m[resultVar])
			} else {
				fmt.Println("I've seen this before:", resultVar, m[resultVar])
				seenBefore = resultVar
				break
			}
		}

    	if err := scanner.Err(); err != nil {
        	log.Fatal(err)
		}
		file.Close()
	}

	return seenBefore
}

func main() {
	fileNamePtr := flag.String("file", "input1.txt", "A filename containing input numbers")
	execPartPtr := flag.String("part", "a", "Which part of day01 do you want to calc (a or b)")

	flag.Parse()

	if *execPartPtr == "a" {
		fmt.Println("Part A - Resulting Frequency:", addAll(*fileNamePtr))
	} else {
		fmt.Println("Part B - Resulting Frequency:", seenTwice(*fileNamePtr))
	}
}