package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func catchUserInput() (string, byte, bool) {
	var debug bool

	filenamePtr := flag.String("file", "input.txt", "Filename containing the program to run")
	execPartPtr := flag.String("part", "a", "Which part of day03 do you want to calc (a or b)")
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

func convertInputToNumbers(stringList []string) []int {
	var numberList []int

	numberList = make([]int, len(stringList))
	for i := 0; i < len(stringList); i++ {
		numberList[i], _ = strconv.Atoi(stringList[i])
	}
	return numberList
}

func checkExpenses(filename string, part byte, debug bool) int {

	puzzleInput, _ := readFile(filename)
	expenses := convertInputToNumbers(puzzleInput)

	if part == 'a' {
		for i := 0; i < len(expenses); i++ {
			for j := 0; j < len(expenses); j++ {
				if i != j { // Ignore when we're looking at the same expenses value
					if expenses[i]+expenses[j] == 2020 {
						fmt.Printf("Got it: %d and %d\n", expenses[i], expenses[j])
						return expenses[i] * expenses[j]
					}
				}
			}
		}
	} else {
		for i := 0; i < len(expenses); i++ {
			for j := 0; j < len(expenses); j++ {
				for k := 0; k < len(expenses); k++ {
					if (i != j) && (i != k) && (j != k) { // Ignore when we're looking at the same expenses value
						if expenses[i]+expenses[j]+expenses[k] == 2020 {
							fmt.Printf("Got it: %d and %d and %d\n", expenses[i], expenses[j], expenses[k])
							return expenses[i] * expenses[j] * expenses[k]
						}
					}
				}
			}
		}
	}

	fmt.Println(puzzleInput)
	return 0
}

// Main routine
func main() {
	filenamePtr, execPart, debug := catchUserInput()
	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", checkExpenses(filenamePtr, execPart, debug))
	}
}
