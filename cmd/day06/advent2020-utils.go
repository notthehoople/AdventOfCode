package main

import (
	"bufio"
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

func convertInputToNumbers(stringList []string) []int {
	var numberList []int

	numberList = make([]int, len(stringList))
	for i := 0; i < len(stringList); i++ {
		numberList[i], _ = strconv.Atoi(stringList[i])
	}
	return numberList
}
