package main

import (
	"AdventOfCode-go/advent2024/utils"
	"fmt"
	"strconv"
	"strings"
)

var Debug bool

func getCombo(operand int, regA int, regB int, regC int) int {
	switch operand {
	case 4:
		return regA
	case 5:
		return regB
	case 6:
		return regC
	default:
		return operand
	}
}

func day17(filename string, part byte) int {
	//Puzzle Input:
	//- line 1: Register A: <num>
	//- line 2: Register B: 0
	//- line 3: Register C: 0
	//- line 4:
	//- line 5: Program: <list of digits separated by commas

	puzzleInput, _ := utils.ReadFile(filename)

	var regA, regB, regC int
	fmt.Sscanf(puzzleInput[0], "Register A: %d\n", &regA)
	fmt.Sscanf(puzzleInput[1], "Register B: %d\n", &regB)
	fmt.Sscanf(puzzleInput[2], "Register C: %d\n", &regC)

	tempString := strings.Split(puzzleInput[4], ": ")
	tempNumbers := strings.Split(tempString[1], ",")

	// convert the string of values to an array of integers
	program := make([]int, len(tempNumbers))
	for key, val := range tempNumbers {
		program[key], _ = strconv.Atoi(val)
	}

	if Debug {
		fmt.Printf("regA: %d regB: %d regC: %d\n", regA, regB, regC)
	}

	instPtr := 0 // Instruction Pointer - where are we in the program?
	var instruction, operand int

	//fmt.Println("Program:", program)

	for instPtr < len(program) {
		// Need to bound these reads and make sure we're not at the end of the program

		instruction = program[instPtr]
		operand = program[instPtr+1]
		instPtr += 2

		switch instruction {
		case 0: // adv - division. Numberator is value in regA; denominator is 2^combo operand
			//fmt.Println("adv")
			combo := getCombo(operand, regA, regB, regC)
			regA = regA >> combo

		case 1: // bxl - bitwise XOR of register B and the instructions literal operand. Stores in regB
			//fmt.Println("bxl")
			regB = regB ^ operand

		case 2: // bst - calc value of its combo operand modulo 8. Writes value to regB
			//fmt.Println("bst")
			combo := getCombo(operand, regA, regB, regC)
			regB = combo % 8

		case 3: // jnz - does nothing if the A register is 0
			//fmt.Println("jnz")
			if regA != 0 {
				instPtr = operand
			}

		case 4: // bxc
			//fmt.Println("bxc")
			regB = regB ^ regC

		case 5: // out
			//fmt.Println("out")
			combo := getCombo(operand, regA, regB, regC) % 8
			fmt.Printf(",%d", combo)

		case 6: // bdv
			//fmt.Println("bdv")
			combo := getCombo(operand, regA, regB, regC)
			regB = regA >> combo

		case 7: // cdv
			//fmt.Println("cdv")
			combo := getCombo(operand, regA, regB, regC)
			regC = regA >> combo
		}
	}

	fmt.Printf("\nregA: %d regB: %d regC: %d\n", regA, regB, regC)

	return 0
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()
	Debug = debug

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", day17(filenamePtr, execPart))
	}
}
