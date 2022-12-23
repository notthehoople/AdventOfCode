package main

import (
	"AdventOfCode-go/advent2022/utils"
	"fmt"
	"strconv"
)

type monkeyShoutStruct struct {
	number    int
	ready     bool
	needsWork string
}

func buildMonkeyArray(puzzleInput []string, debug bool) map[string]monkeyShoutStruct {

	monkeyShout := make(map[string]monkeyShoutStruct, 0)

	for _, line := range puzzleInput {
		monkey := line[:4]

		if (line[6] >= '0') && (line[6] <= '9') {
			tempNumber, _ := strconv.Atoi(line[6:])
			monkeyShout[monkey] = monkeyShoutStruct{number: tempNumber, ready: true}
		} else {
			monkeyShout[monkey] = monkeyShoutStruct{ready: false, needsWork: line[6:]}
		}

	}

	return monkeyShout
}

func returnMonkeyShout(monkeyShout map[string]monkeyShoutStruct, monkey string) int {

	// If the monkey has shouted a number, return it
	if monkeyShout[monkey].ready {
		return monkeyShout[monkey].number
	}

	// Need to calculate what the given monkey will shout, then return it
	tmpNeedsWork := monkeyShout[monkey].needsWork
	tmpMonkey1 := tmpNeedsWork[:4]
	operation := tmpNeedsWork[5]
	tmpMonkey2 := tmpNeedsWork[7:]

	tmpMonkey1Val := returnMonkeyShout(monkeyShout, tmpMonkey1)
	tmpMonkey2Val := returnMonkeyShout(monkeyShout, tmpMonkey2)

	var tmpResultNumber int

	switch operation {
	case '+':
		tmpResultNumber = tmpMonkey1Val + tmpMonkey2Val
	case '-':
		tmpResultNumber = tmpMonkey1Val - tmpMonkey2Val
	case '/':
		tmpResultNumber = tmpMonkey1Val / tmpMonkey2Val
	case '*':
		tmpResultNumber = tmpMonkey1Val * tmpMonkey2Val
	}

	monkeyShout[monkey] = monkeyShoutStruct{number: tmpResultNumber, needsWork: "", ready: true}
	return monkeyShout[monkey].number
}

func calcMonkeySpeach(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)

	monkeyShout := buildMonkeyArray(puzzleInput, debug)

	return returnMonkeyShout(monkeyShout, "root")
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Monkey root shouts: %d\n", calcMonkeySpeach(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("humn needs to shout: %d\n", calcMonkeySpeachPartB(filenamePtr, execPart, debug))
	case 'z':
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
