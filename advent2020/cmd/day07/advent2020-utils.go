package main

import (
	"bufio"
	"flag"
	"os"
	"strconv"
)

// Read the text file passed in by name into a array of strings
// Returns the array as the first return variable
func readFile(filename string) ([]string, error) {
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

func catchUserInput() (string, byte, bool, bool) {
	var debug bool
	var test bool

	filenamePtr := flag.String("file", "testInput.txt", "Filename containing the program to run")
	execPartPtr := flag.String("part", "a", "Which part of the puzzle do you want to calc (a or b)")
	flag.BoolVar(&debug, "debug", false, "Turn debug on")
	flag.BoolVar(&test, "test", false, "Run tests only")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		return *filenamePtr, 'a', debug, test
	case "b":
		return *filenamePtr, 'b', debug, test

	default:
		return *filenamePtr, 'z', debug, test
	}
}

func convertInputToNumbers(stringList []string) []int {
	var numberList []int

	numberList = make([]int, len(stringList))
	for i := 0; i < len(stringList); i++ {
		numberList[i], _ = strconv.Atoi(stringList[i])
	}
	return numberList
}
