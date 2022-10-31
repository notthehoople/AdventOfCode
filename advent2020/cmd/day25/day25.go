package main

import (
	"fmt"
	"strconv"
	//"strings"
)

func findLoopSize(publicKey int, debug bool) int {
	// Card transforms the subject number of 7 according to the card's secret loop size. Result: card's public key
	// Door transforms the subject number of 7 according to the door's secret loop size. Result: door's public key
	var loopSize int = 1
	var subjectNumber int = 7
	var calcValue int = 1
	for {
		calcValue *= subjectNumber
		calcValue = calcValue % 20201227

		if calcValue == publicKey {
			break
		}
		loopSize++
	}

	fmt.Printf("publicKey: %d loopSize: %d\n", publicKey, loopSize)

	return loopSize
}

func calcEncryptionKey(publicKey int, loopSize int) int {
	var calcValue int = 1
	for i := 1; i <= loopSize; i++ {
		calcValue = calcValue * publicKey
		calcValue = calcValue % 20201227
	}
	return calcValue
}

func runPartA(filename string, part byte, debug bool) int {
	var cardLoopSize, doorLoopSize int
	var encryptionKey, encryptionKey2 int

	puzzleInput, _ := readFile(filename)
	cardPublicKey, _ := strconv.Atoi(puzzleInput[0])
	doorPublicKey, _ := strconv.Atoi(puzzleInput[1])

	cardLoopSize = findLoopSize(cardPublicKey, debug)
	doorLoopSize = findLoopSize(doorPublicKey, debug)

	encryptionKey = calcEncryptionKey(doorPublicKey, cardLoopSize)
	encryptionKey2 = calcEncryptionKey(cardPublicKey, doorLoopSize)

	fmt.Printf("Encrytion 1: %d 2: %d\n", encryptionKey, encryptionKey2)

	return encryptionKey
}

// Main routine
func main() {
	filenamePtr, execPart, debug, test := catchUserInput()

	if test {
		return
	}

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else if execPart == 'a' {

		fmt.Println("Test Run:", runPartA("testInput.txt", execPart, debug))
		fmt.Println("Live Run:", runPartA("puzzleInput.txt", execPart, debug))

	} else {
		fmt.Println(filenamePtr, execPart, debug)
	}
}
