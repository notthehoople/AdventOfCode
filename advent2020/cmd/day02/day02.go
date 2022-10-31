package main

import (
	"flag"
	"fmt"
)

func catchUserInput() (string, byte, bool) {
	var debug bool

	filenamePtr := flag.String("file", "testInput.txt", "Filename containing the program to run")
	execPartPtr := flag.String("part", "a", "Which part of the puzzle do you want to calc (a or b)")
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

func checkPasswords(filename string, part byte, debug bool) int {
	var minNumber, maxNumber int
	var passwordChar rune
	var password string
	var loopCharCount int
	var correctPasswordCount int = 0

	puzzleInput, _ := readFile(filename)

	for _, passwordLine := range puzzleInput {
		fmt.Sscanf(passwordLine, "%d-%d %c: %s", &minNumber, &maxNumber, &passwordChar, &password)

		if debug {
			fmt.Println(passwordLine)
			fmt.Printf("min: %d max: %d char: %c password: %s\n", minNumber, maxNumber, passwordChar, password)
		}

		if part == 'a' {
			loopCharCount = 0
			for _, loopChar := range password {
				if loopChar == passwordChar {
					loopCharCount++
				}
			}

			if (loopCharCount >= minNumber) && (loopCharCount <= maxNumber) {
				correctPasswordCount++
			}
		} else {
			if (password[minNumber-1] == byte(passwordChar)) && (password[maxNumber-1] != byte(passwordChar)) {
				correctPasswordCount++
			} else if (password[minNumber-1] != byte(passwordChar)) && (password[maxNumber-1] == byte(passwordChar)) {
				correctPasswordCount++
			}
		}

		if debug {
			fmt.Println("Char appeared: ", loopCharCount)
		}
	}

	return correctPasswordCount
}

// Main routine
func main() {
	filenamePtr, execPart, debug := catchUserInput()
	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", checkPasswords(filenamePtr, execPart, debug))
	}
}
