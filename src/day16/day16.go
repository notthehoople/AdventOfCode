package main

import (
	"fmt"
	"bufio"
	"os"
	"flag"
//	"strconv"
//	"strings"
)

// There are 4 registers, numbered 0 to 3
type registers struct {
	register0	int
	register1	int
	register2	int
	register3	int
}

// The code instructions contain 4 values.
//   - an opcode, numbered 0 to 15
//   - two input values (A and B) and an output (C)
type instructions struct {
	opCode		int
	a			int
	b			int
	c			int
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

func printStringArray(tempString []string) {
	// Loop through the array and print each line
	for i:= 0; i < len(tempString); i++ {
		fmt.Println(tempString[i])
	}
}

func processFirstInstructions(fileContents []string) ([]registers, []instructions, []registers) {
	var partAInstructions []instructions
	var beforeRegisters []registers
	var afterRegisters []registers
	var result, reg0, reg1, reg2, reg3, opCode, a, b, c int

	for i := 0; i < len(fileContents); i++ {
		// We're reading a block of 3: a "Before" line, a set of intructions, then an "After"
		// If we don't match the beginning of the block, then continue until we do.

		result, _ = fmt.Sscanf(fileContents[i], "Before: [%d, %d, %d, %d]", &reg0, &reg1, &reg2, &reg3)
		if result != 4 {
			// not a new block of Before, Instructions, After so skip until we find a block
			continue
		}
		beforeRegisters = append(beforeRegisters, registers{register0:reg0, register1:reg1, register2:reg2, register3:reg3})

		i++
		result, _ = fmt.Sscanf(fileContents[i], "%d %d %d %d", &opCode, &a, &b, &c)
		partAInstructions = append(partAInstructions, instructions{opCode: opCode, a:a, b:b, c:c})

		i++
		result, _ = fmt.Sscanf(fileContents[i], "After: [%d, %d, %d, %d]", &reg0, &reg1, &reg2, &reg3)
		afterRegisters = append(afterRegisters, registers{register0:reg0, register1:reg1, register2:reg2, register3:reg3})
	}

	return beforeRegisters, partAInstructions, afterRegisters
}

// func processOpcodes
//
func processOpcodes(fileName string, printInstructions bool, part byte) int {

	// Read contents of file into a string array
	fileContents, _ := readLines(fileName)

	if part == 'a' {
		beforeRegisters, partAInstructions, afterRegisters := processFirstInstructions(fileContents)

		fmt.Println(beforeRegisters)
		fmt.Println(partAInstructions)
		fmt.Println(afterRegisters)
	}

	// Now we need to work through the input data and run it through the operations we have

	return 0
}

// Main routine
func main() {
	var printInstructions bool

	fileNamePtr := flag.String("file", "input1.txt", "A filename containing input strings")
	flag.BoolVar(&printInstructions, "print", false, "Print the results as we go")
	execPartPtr := flag.String("part", "a", "Part of the puzzle to work on. a or b")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Number of samples behaving like 3 or more opcodes:", processOpcodes(*fileNamePtr, printInstructions, 'a'))
	case "b":
		fmt.Println("Not here yet")
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}