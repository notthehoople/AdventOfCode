package main

import (
	"bufio"
	"fmt"
	"os"
)

// Abs returns the absolute value of x.
func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//GCD greatest common divisor (GCD) via Euclidean algorithm
// Code lifted from Go Playground: https://play.golang.org/p/SmzvkDjYlb
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

//LCM find Least Common Multiple (LCM) via GCD
// Code lifted from Go Playground: https://play.golang.org/p/SmzvkDjYlb
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

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

// Prints the map list
func printMap(tempMap []string) {
	for i := 0; i < len(tempMap); i++ {
		fmt.Printf("%s\n", tempMap[i])
	}
}

// func: readInitialState
// takes an array of strings and breaks it into a 2D array of bytes
func readInitialState(tempString []string, tempSlice [][]byte) {
	for i := 0; i < len(tempString); i++ {
		for j := 0; j < len(tempString[i]); j++ {
			tempSlice[i][j] = tempString[i][j]
		}
	}
}

func print2DSlice(tempSlice [][]byte) {
	for i := 0; i < len(tempSlice); i++ {
		for j := 0; j < len(tempSlice[i]); j++ {
			fmt.Printf("%c", tempSlice[i][j])
		}
		fmt.Printf("\n")
	}
}

// Assumes that destination slice has enough memory allocated to hold the contents of sourceSlice
func copy2DSlice(sourceSlice [][]byte, destinationSlice [][]byte) {
	for i := 0; i < len(sourceSlice); i++ {
		for j := 0; j < len(sourceSlice[i]); j++ {
			destinationSlice[i][j] = sourceSlice[i][j]
		}
	}
}
