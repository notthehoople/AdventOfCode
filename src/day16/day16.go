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

// func decodeInstruction
// returns: regA, regB, regC
func decodeInstruction(partAInstruction instructions) (int, int, int) {
	return partAInstruction.a, partAInstruction.b, partAInstruction.c
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

// Check the non-changing registers are the same before and after the sum
// Reference to the changing register is passed in regC
func checkRegisterChanges(beforeRegister registers, afterRegister registers, regC int) bool {
	switch regC {
	case 0: if (beforeRegister.register1 == afterRegister.register1) &&
			   (beforeRegister.register2 == afterRegister.register2) &&
			   (beforeRegister.register3 == afterRegister.register3) {
				   return true
			   }
	case 1: if (beforeRegister.register0 == afterRegister.register0) &&
				(beforeRegister.register2 == afterRegister.register2) &&
				(beforeRegister.register3 == afterRegister.register3) {
					return true
				}
	case 2: if (beforeRegister.register0 == afterRegister.register0) &&
				(beforeRegister.register1 == afterRegister.register1) &&
				(beforeRegister.register3 == afterRegister.register3) {
					return true
			}
	case 3: if (beforeRegister.register0 == afterRegister.register0) &&
				(beforeRegister.register1 == afterRegister.register1) &&
				(beforeRegister.register2 == afterRegister.register2) {
					return true
			}
	}
	return false
}

// **********************************************************************************************************************
// OpCodes
//
func runOpcode(opCode string, beforeRegister registers, partAInstruction instructions, afterRegister registers) bool {
	var regA, regB, regC int = 0, 0, 0
	// var regBeforeC int = 0
	var regBeforeA, regBeforeB int = 0, 0
	//var regAfterA, regAfterB, regAfterC int = 0, 0, 0
	var regAfterC int = 0

	regA, regB, regC = decodeInstruction(partAInstruction)
	
	// check only regC has changed between before and After registers. Otherwise it's not us
	if !checkRegisterChanges(beforeRegister, afterRegister, regC) {
		fmt.Println("Can't be addr as the non-changing registers are different")
		return false
	}

//	regBeforeA, regBeforeB, regBeforeC = decodeRegisters(beforeRegister, regA, regB, regC)

	regBeforeA, regBeforeB, _ = decodeRegisters(beforeRegister, regA, regB, regC)
	_, _, regAfterC = decodeRegisters(afterRegister, regA, regB, regC)

	switch opCode {
		case "addr": 	// CHECKED addr (add register) stores into register C the result of adding register A and register B.
						if (regBeforeA + regBeforeB) == regAfterC {
							return true
						}
		case "addi":	// CHECKED addi (add immediate) stores into register C the result of adding register A and value B.
						if (regBeforeA + regB) == regAfterC {
							return true
						}

		case "mulr":	// CHECKED mulr (multiply register) stores into register C the result of multiplying register A and register B.
						if (regBeforeA * regBeforeB) == regAfterC {
							return true
						}
		case "muli":	// CHECKED muli (multiply immediate) stores into register C the result of multiplying register A and value B.
						if (regBeforeA * regB) == regAfterC {
							return true
						}

		case "banr":	// CHECKED banr (bitwise AND register) stores into register C the result of the bitwise AND of register A and register B.
						if (regBeforeA & regBeforeB) == regAfterC {
							return true
						}

		case "bani":	// bani (bitwise AND immediate) stores into register C the result of the bitwise AND of register A and value B.
						if (regBeforeA & regB) == regAfterC {
							return true
						}

		case "borr":	// borr (bitwise OR register) stores into register C the result of the bitwise OR of register A and register B.
						if (regBeforeA | regBeforeB) == regAfterC {
							return true
						}

		case "bori":	// bori (bitwise OR immediate) stores into register C the result of the bitwise OR of register A and value B.
						if (regBeforeA | regB) == regAfterC {
							return true
						}

		case "setr":	// setr (set register) copies the contents of register A into register C. (Input B is ignored.)
						if regBeforeA == regAfterC {
							return true
						}
		case "seti":	// seti (set immediate) stores value A into register C. (Input B is ignored.)
						if regA == regAfterC {
							return true
						}
		
		case "gtir":	// gtir (greater-than immediate/register) sets register C to 1 if value A is greater than register B. Otherwise, register C is set to 0
						if (regA > regBeforeB) && (regAfterC == 1) {
							return true
						} else {
							if (regA <= regBeforeB) && (regAfterC == 0) {
								return true
							} else {
								return false
							}
						}

		case "gtri":	// gtri (greater-than register/immediate) sets register C to 1 if register A is greater than value B. Otherwise, register C is set to 0
						if (regBeforeA > regB) && (regAfterC == 1) {
							return true
						} else {
							if (regBeforeA <= regB) && (regAfterC == 0) {
								return true
							} else {
								return false
							}
						}

		case "gtrr":	// gtrr (greater-than register/register) sets register C to 1 if register A is greater than register B. Otherwise, register C is set to 0
						if (regBeforeA > regBeforeB) && (regAfterC == 1) {
							return true
						} else {
							if (regBeforeA <= regBeforeB) && (regAfterC == 0) {
								return true
							} else {
								return false
							}
						}

		case "eqir":	// eqir (equal immediate/register) sets register C to 1 if value A is equal to register B. Otherwise, register C is set to 0
						if (regA == regBeforeB) && (regAfterC == 1) {
							return true
						} else {
							if (regA != regBeforeB) && (regAfterC == 0) {
								return true
							} else {
								return false
							}
						}

		case "eqri":	// eqri (equal register/immediate) sets register C to 1 if register A is equal to value B. Otherwise, register C is set to 0
						if (regBeforeA == regB) && (regAfterC == 1) {
							return true
						} else {
							if (regBeforeA != regB) && (regAfterC == 0) {
								return true
							} else {
								return false
							}
						}

		case "eqrr":	// eqrr (equal register/register) sets register C to 1 if register A is equal to register B. Otherwise, register C is set to 0
						if (regBeforeA == regBeforeB) && (regAfterC == 1) {
							return true
						} else {
							if (regBeforeA != regBeforeB) && (regAfterC == 0) {
								return true
							} else {
								return false
							}
						}

	}
	

	return false
}

// **********************************************************************************************************************
// Main Functions
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
	var numPossibleOpCodes, threeOrMore int = 0, 0
	var beforeRegisters []registers
	var partAInstructions []instructions
	var afterRegisters []registers
	var allOpCodes = [16]string{"addr","addi", "mulr", "muli", "banr", "bani", "borr", "bori", "setr", "seti", "gtir", "gtri", "gtrr", "eqir", "eqri", "eqrr"}

	// Read contents of file into a string array
	fileContents, _ := readLines(fileName)

	if part == 'a' {
		beforeRegisters, partAInstructions, afterRegisters = processFirstInstructions(fileContents)
	}

	// Now we need to work through the input data and run it through the operations we have
	for i := 0; i < len(beforeRegisters); i++ {
		for _, tryOpCode := range allOpCodes {
			if runOpcode(tryOpCode, beforeRegisters[i], partAInstructions[i], afterRegisters[i]) {
				numPossibleOpCodes++
			}
		}
		if numPossibleOpCodes > 2 {
			threeOrMore++
		}
		numPossibleOpCodes = 0
	}

	return threeOrMore
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