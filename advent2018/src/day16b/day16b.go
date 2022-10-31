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

// **********************************************************************************************************************
// Supporting functions for instructions and registers

// func updateRegisters
// returns: updated register
// paramters: the register to update; a pointer to the register to be updated; the value to update the register to
func updateRegisters(registerToUpdate registers, regToUpdate int, updateValue int) registers {
	switch regToUpdate {
	case 0: registerToUpdate.register0 = updateValue
	case 1: registerToUpdate.register1 = updateValue
	case 2: registerToUpdate.register2 = updateValue
	case 3: registerToUpdate.register3 = updateValue
	}
	
	return registerToUpdate
}


// func decodeInstruction
// returns: opCode, regA, regB, regC
func decodeInstruction(partBInstruction instructions) (int, int, int, int) {
	return partBInstruction.opCode, partBInstruction.a, partBInstruction.b, partBInstruction.c
}

func decodeSingleRegister(register registers, inputReg int) int {
	switch inputReg {
	case 0: return register.register0
	case 1: return register.register1
	case 2: return register.register2
	case 3: return register.register3
	}

	return 0
}

// func decodeRegisters
// returns: reg0, reg1, reg2, reg3
func decodeRegisters(register registers, regA int, regB int, regC int) (int, int, int) {
	return decodeSingleRegister(register, regA), decodeSingleRegister(register, regB), decodeSingleRegister(register, regC)
}

// **********************************************************************************************************************
// OpCodes
//
func runOpcode(opCode string, beforeRegister registers, partBInstruction instructions, afterRegister registers) (registers) {
	//var currentOpCode int = 0
	var regA, regB, regC int = 0, 0, 0
	var regBeforeA, regBeforeB int = 0, 0
	var regAfterC int = 0

	//currentOpCode, regA, regB, regC = decodeInstruction(partBInstruction)
	_, regA, regB, regC = decodeInstruction(partBInstruction)

//	regBeforeA, regBeforeB, regBeforeC = decodeRegisters(beforeRegister, regA, regB, regC)

	// Make sure afterRegister is the same as beforeRegister BEFORE we do the operation
	afterRegister = beforeRegister

	regBeforeA, regBeforeB, _ = decodeRegisters(beforeRegister, regA, regB, regC)

	switch opCode {
		case "addr": 	// addr (add register) stores into register C the result of adding register A and register B.
						regAfterC = regBeforeA + regBeforeB
						afterRegister = updateRegisters(afterRegister, regC, regAfterC)
						
		case "addi":	// addi (add immediate) stores into register C the result of adding register A and value B.
						regAfterC = regBeforeA + regB
						afterRegister = updateRegisters(afterRegister, regC, regAfterC)

		case "mulr":	// mulr (multiply register) stores into register C the result of multiplying register A and register B.
						regAfterC = regBeforeA * regBeforeB
						afterRegister = updateRegisters(afterRegister, regC, regAfterC)

		case "muli":	// muli (multiply immediate) stores into register C the result of multiplying register A and value B.
						regAfterC = regBeforeA * regB
						afterRegister = updateRegisters(afterRegister, regC, regAfterC)

		case "banr":	// banr (bitwise AND register) stores into register C the result of the bitwise AND of register A and register B.
						regAfterC = (regBeforeA & regBeforeB)
						afterRegister = updateRegisters(afterRegister, regC, regAfterC)

		case "bani":	// bani (bitwise AND immediate) stores into register C the result of the bitwise AND of register A and value B.
						regAfterC = (regBeforeA & regB)
						afterRegister = updateRegisters(afterRegister, regC, regAfterC)

		case "borr":	// borr (bitwise OR register) stores into register C the result of the bitwise OR of register A and register B.
						regAfterC = (regBeforeA | regBeforeB)
						afterRegister = updateRegisters(afterRegister, regC, regAfterC)

		case "bori":	// bori (bitwise OR immediate) stores into register C the result of the bitwise OR of register A and value B.
						regAfterC = (regBeforeA | regB)
						afterRegister = updateRegisters(afterRegister, regC, regAfterC)

		case "setr":	// setr (set register) copies the contents of register A into register C. (Input B is ignored.)
						regAfterC = regBeforeA
						afterRegister = updateRegisters(afterRegister, regC, regAfterC)

		case "seti":	// seti (set immediate) stores value A into register C. (Input B is ignored.)
						regAfterC = regA
						afterRegister = updateRegisters(afterRegister, regC, regAfterC)
		
		case "gtir":	// gtir (greater-than immediate/register) sets register C to 1 if value A is greater than register B. Otherwise, register C is set to 0
						if (regA > regBeforeB) {
							regAfterC = 1
						} else {
							regAfterC = 0
						}
						afterRegister = updateRegisters(afterRegister, regC, regAfterC)


		case "gtri":	// gtri (greater-than register/immediate) sets register C to 1 if register A is greater than value B. Otherwise, register C is set to 0
						if regBeforeA > regB {
							regAfterC = 1
						} else {
							regAfterC = 0
						}
						afterRegister = updateRegisters(afterRegister, regC, regAfterC)


		case "gtrr":	// gtrr (greater-than register/register) sets register C to 1 if register A is greater than register B. Otherwise, register C is set to 0
						if regBeforeA > regBeforeB {
							regAfterC = 1
						} else {
							regAfterC = 0
						}
						afterRegister = updateRegisters(afterRegister, regC, regAfterC)


		case "eqir":	// eqir (equal immediate/register) sets register C to 1 if value A is equal to register B. Otherwise, register C is set to 0
						if regA == regBeforeB {
							regAfterC = 1
						} else {
							regAfterC = 0
						}
						afterRegister = updateRegisters(afterRegister, regC, regAfterC)


		case "eqri":	// eqri (equal register/immediate) sets register C to 1 if register A is equal to value B. Otherwise, register C is set to 0
						if regBeforeA == regB {
							regAfterC = 1
						} else {
							regAfterC = 0
						}
						afterRegister = updateRegisters(afterRegister, regC, regAfterC)


		case "eqrr":	// eqrr (equal register/register) sets register C to 1 if register A is equal to register B. Otherwise, register C is set to 0
						if regBeforeA == regBeforeB {
							regAfterC = 1
						} else {
							regAfterC = 0
						}
						afterRegister = updateRegisters(afterRegister, regC, regAfterC)


	}
	
	return afterRegister
}

// **********************************************************************************************************************
// Main Functions
//
// func processSecondInstructions
// reads the 2nd section of the input data
// builds beforeRegisters and afterRegisters with blanks ready to hold calculations
// and partBInstructions with all the instructions we need
func processSecondInstructions(fileContents []string) ([]registers, []instructions, []registers) {
	var partBInstructions []instructions
	var beforeRegisters []registers
	var afterRegisters []registers
	//var result, reg0, reg1, reg2, reg3, opCode, a, b, c int
	var opCode, a, b, c int


	for i := 0; i < len(fileContents); i++ {
		// We're processing the second part of the file which is just a list of instructions
		// Fill out beforeRegisters and afterRegisters with blanks for later results

		beforeRegisters = append(beforeRegisters, registers{register0:0, register1:0, register2:0, register3:0})

		fmt.Sscanf(fileContents[i], "%d %d %d %d", &opCode, &a, &b, &c)
		partBInstructions = append(partBInstructions, instructions{opCode: opCode, a:a, b:b, c:c})

		afterRegisters = append(afterRegisters, registers{register0:0, register1:0, register2:0, register3:0})
	}

	return beforeRegisters, partBInstructions, afterRegisters
}

// func runProgram
//
func runProgram(fileName string, printInstructions bool, part byte) int {
	//var numPossibleOpCodes, threeOrMore, currentOpCode int = 0, 0, 0
	//var opCodeWorked bool
	var beforeRegisters []registers
	var partBInstructions []instructions
	var afterRegisters []registers
	//var tryOpCode string
	//var resultOpCodeNum int
	//var resultOpCodeString string
	var allOpCodes = [16]string{"bani", "banr", "muli", "setr", "bori", "eqrr", "gtir", "mulr", "gtrr", "seti", "gtri", "eqri", "addi", "borr", "eqir", "addr"}

	// Read contents of file into a string array
	fileContents, _ := readLines(fileName)

	if part == 'b' {
		beforeRegisters, partBInstructions, afterRegisters = processSecondInstructions(fileContents)
	}

	// Now we need to work through the instructions and execute the operations
//	for i := 0; i < len(partBInstructions); i++ {
	for i := 0; i < len(partBInstructions); i++ {

		fmt.Println("BeforeRegisters:", beforeRegisters[i])
		fmt.Println("Instructions:", allOpCodes[partBInstructions[i].opCode], partBInstructions[i])
		afterRegisters[i] = runOpcode(allOpCodes[partBInstructions[i].opCode], beforeRegisters[i], partBInstructions[i], afterRegisters[i])
		fmt.Println("AfterRegisters:", afterRegisters[i])
		fmt.Println("====================")

		if i < (len(partBInstructions) - 1) {
			beforeRegisters[i+1] = afterRegisters[i]
		}
	}

	return afterRegisters[len(afterRegisters) - 1].register0
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
		//fmt.Println("Part a - Number of samples behaving like 3 or more opcodes:", processOpcodes(*fileNamePtr, printInstructions, 'a'))
	case "b":
		fmt.Println("Part b - register 0:", runProgram(*fileNamePtr, printInstructions, 'b'))
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}