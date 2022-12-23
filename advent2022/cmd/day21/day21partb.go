package main

import (
	"AdventOfCode-go/advent2022/utils"
	"fmt"
)

func returnMonkeyShoutPartB(monkeyShout map[string]monkeyShoutStruct, monkey string, humnShouts int) (int, int) {

	// If the monkey has shouted a number, return it
	if monkey == "humn" {
		if humnShouts%10000 == 0 {
			fmt.Println("humn shouts:", humnShouts)
		}
		return humnShouts, 0
	}
	if monkeyShout[monkey].ready {
		return monkeyShout[monkey].number, 0
	}

	// Need to calculate what the given monkey will shout, then return it
	tmpNeedsWork := monkeyShout[monkey].needsWork
	tmpMonkey1 := tmpNeedsWork[:4]
	operation := tmpNeedsWork[5]
	tmpMonkey2 := tmpNeedsWork[7:]

	tmpMonkey1Val, _ := returnMonkeyShoutPartB(monkeyShout, tmpMonkey1, humnShouts)
	tmpMonkey2Val, _ := returnMonkeyShoutPartB(monkeyShout, tmpMonkey2, humnShouts)

	if monkey == "root" {
		return tmpMonkey1Val, tmpMonkey2Val
	}

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
	return monkeyShout[monkey].number, 0
}

func calcMonkeySpeachPartB(filename string, part byte, debug bool) int {
	var humnShouts int = 3769668700000
	var prevHumnShouts int
	//var changeSteps int = 1000000000
	var changeSteps int = 100

	var currentDiff, previousDiff int

	puzzleInput, _ := utils.ReadFile(filename)

	for {
		monkeyShout := buildMonkeyArray(puzzleInput, debug)

		tmpMonkey1Val, tmpMonkey2Val := returnMonkeyShoutPartB(monkeyShout, "root", humnShouts)

		if tmpMonkey1Val == tmpMonkey2Val {
			// We're done
			return humnShouts
		}

		currentDiff = utils.AbsDiff(tmpMonkey1Val, tmpMonkey2Val)
		//fmt.Printf("tmpMonkey1Val: %d tmpMonkey2Val: %d Difference: %d\n", tmpMonkey1Val, tmpMonkey2Val, currentDiff)

		// If the difference has *increased* then we should stop
		//if utils.Abs(previousDiff) < utils.Abs(currentDiff) && (previousDiff != 0) {
		if (previousDiff < currentDiff) && (previousDiff != 0) {
			//	fmt.Printf("Previous diff: %d Current diff: %d Previous was better. Stop!\n", utils.Abs(previousDiff), utils.Abs(currentDiff))
			humnShouts = prevHumnShouts
			currentDiff = previousDiff
			//changeSteps = changeSteps / 10
			changeSteps = 1
		}

		previousDiff = currentDiff
		prevHumnShouts = humnShouts

		humnShouts += changeSteps

	}
}
