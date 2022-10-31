package main

import (
	"fmt"
	"bufio"
	"os"
	"flag"
)

// There are 6 registers, numbered 0 to 5
type registers struct {
	register0	int
	register1	int
	register2	int
	register3	int
	register4	int
	register5	int
}

// The code instructions contain 4 values.
//   - an opcode which is a string
//   - two input values (A and B) and an output (C)
type instructions struct {
	opCode		string
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
	case 4: registerToUpdate.register4 = updateValue
	case 5: registerToUpdate.register5 = updateValue
	}
	
	return registerToUpdate
}

// func decodeInstruction
// returns: opCode, regA, regB, regC
func decodeInstruction(programInstructions instructions) (string, int, int, int) {
	return programInstructions.opCode, programInstructions.a, programInstructions.b, programInstructions.c
}

func decodeSingleRegister(register registers, inputReg int) int {
	switch inputReg {
	case 0: return register.register0
	case 1: return register.register1
	case 2: return register.register2
	case 3: return register.register3
	case 4: return register.register4
	case 5: return register.register5
	}

	return 0
}

// func decodeRegisters
// returns: reg0, reg1, reg2, reg3
func decodeRegisters(register registers, regA int, regB int, regC int) (int, int, int) {
	return decodeSingleRegister(register, regA), decodeSingleRegister(register, regB), decodeSingleRegister(register, regC)
}

// func processInstructions
// reads the input data
// builds beforeRegisters and afterRegisters with blanks ready to hold calculations
// and programInstructions with all the instructions we need
func processInstructions(fileContents []string) (int, []registers, []instructions, []registers) {
	var programInstructions []instructions
	var beforeRegisters []registers
	var afterRegisters []registers
	var iPRegister int
	var opCode string
	var a, b, c int

	fmt.Sscanf(fileContents[0], "#ip %d", &iPRegister)

	for i := 1; i < len(fileContents); i++ {
		// Fill out beforeRegisters and afterRegisters with blanks for later results

		beforeRegisters = append(beforeRegisters, registers{register0:0, register1:0, register2:0, register3:0})

		fmt.Sscanf(fileContents[i], "%s %d %d %d", &opCode, &a, &b, &c)
		programInstructions = append(programInstructions, instructions{opCode: opCode, a:a, b:b, c:c})

		afterRegisters = append(afterRegisters, registers{register0:0, register1:0, register2:0, register3:0})
	}

	return iPRegister, beforeRegisters, programInstructions, afterRegisters
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

// func runProgram
//
func runProgram(fileName string, printInstructions bool, part byte) int {
	var beforeRegisters []registers
	var programInstructions []instructions
	var afterRegisters []registers
	var allOpCodes = [16]string{"bani", "banr", "muli", "setr", "bori", "eqrr", "gtir", "mulr", "gtrr", "seti", "gtri", "eqri", "addi", "borr", "eqir", "addr"}
	var instructionPointer, newInstructionPointer int = 0, 0
	var iPRegister int = 0

	// Read contents of file into a string array
	fileContents, _ := readLines(fileName)

	iPRegister, beforeRegisters, programInstructions, afterRegisters = processInstructions(fileContents)
	if printInstructions {
		fmt.Println("iPRegister:", iPRegister)
		fmt.Println("allOpCodes:", allOpCodes)
		fmt.Println("instructionPointer:", instructionPointer)
	}

	if part == 'b' {
		// Register 0 starts with a value of 1 in part b. Everything else is identical
		beforeRegisters[0] = updateRegisters(beforeRegisters[0], 0, 1)
		fmt.Println("beforeRegisters:", beforeRegisters)
	}
	
	// Now we need to work through the instructions and execute the operations, while dealing with the instruction pointer
	for i := instructionPointer; i < len(programInstructions); i++ {

		beforeRegisters[i] = updateRegisters(beforeRegisters[i], iPRegister, i)

		if printInstructions {
			fmt.Println("Instruction Pointer:", i)
			fmt.Println("BeforeRegisters:", beforeRegisters[i])
			fmt.Println("Instructions:", programInstructions[i])
		}
	
		afterRegisters[i] = runOpcode(programInstructions[i].opCode, beforeRegisters[i], programInstructions[i], afterRegisters[i])
		
		if printInstructions {
			fmt.Println("AfterRegisters:", afterRegisters[i])
			fmt.Println("Instruction Pointer:", i)
		}

		newInstructionPointer = decodeSingleRegister(afterRegisters[i], iPRegister)

		if printInstructions {
			fmt.Println("New Instruction Pointer:", newInstructionPointer)
			fmt.Println("====================")
		}

		if newInstructionPointer < (len(programInstructions) - 1) {
			beforeRegisters[newInstructionPointer+1] = afterRegisters[i]
		} else {
			if (newInstructionPointer + 1) >= len(programInstructions) {
				return decodeSingleRegister(afterRegisters[i], 0)

			}
		}
		i = newInstructionPointer
	}

	if printInstructions {
		fmt.Println(afterRegisters)
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
		fmt.Println("Part a - Value in register 0 after program executes:", runProgram(*fileNamePtr, printInstructions, 'a'))
	case "b":
		fmt.Println("Part b - Value in register 0 after program executes:", runProgram(*fileNamePtr, printInstructions, 'b'))
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}