package main

import (
	"aoc/advent2021/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func calcAverage(numberList []int) int {
	var totalValue int
	//var average float64

	for _, i := range numberList {
		totalValue += i
	}

	/*
		average = float64(totalValue) / float64(len(numberList))
		fmt.Println("Average", int(math.Round(average)))
		return int(math.Round(average))
	*/

	return totalValue / len(numberList)
}

func calcMedian(numberList []int) int {
	sort.Ints(numberList)

	middleNumber := len(numberList) / 2

	if len(numberList)%2 == 1 {
		return numberList[middleNumber]
	} else {
		return (numberList[middleNumber-1] + numberList[middleNumber]) / 2
	}
}

func calcDifference(firstNum int, secondNum int) int {
	if firstNum > secondNum {
		return firstNum - secondNum
	} else {
		return secondNum - firstNum
	}
}

func calcFuelUsed(crabPos []int, movePosition int, part byte, debug bool) int {
	var totalFuelUsed int

	if part == 'a' {
		for i := 0; i < len(crabPos); i++ {
			totalFuelUsed += calcDifference(movePosition, crabPos[i])
		}
	} else {
		var fuelPerStep int
		for i := 0; i < len(crabPos); i++ {
			fuelPerStep = 1
			moveAmount := calcDifference(movePosition, crabPos[i])
			for j := 1; j <= moveAmount; j++ {
				totalFuelUsed += fuelPerStep
				fuelPerStep++
			}
		}
	}
	return totalFuelUsed
}

func solveDay(filename string, part byte, debug bool) int {
	var crabPos int
	puzzleInput, _ := utils.ReadFile(filename)
	puzzleInputSplit := strings.Split(puzzleInput[0], ",")
	crabSubs := make([]int, len(puzzleInputSplit))
	for i := 0; i < len(puzzleInputSplit); i++ {
		crabPos, _ = strconv.Atoi(puzzleInputSplit[i])
		crabSubs[i] = crabPos
	}

	if part == 'a' {
		median := calcMedian(crabSubs)
		return calcFuelUsed(crabSubs, median, part, debug)
	} else {
		average := calcAverage(crabSubs)
		return calcFuelUsed(crabSubs, average, part, debug)
	}
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", solveDay(filenamePtr, execPart, debug))
	}
}
