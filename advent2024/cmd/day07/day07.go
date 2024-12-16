package main

import (
	"AdventOfCode-go/advent2024/utils"
	"fmt"
	"strconv"
	"strings"
)

var Debug bool

func buildSlice(inputString string) []int {
	//var d int
	var numSlice []int

	strSlice := strings.Split(inputString, " ")

	for i := 0; i < len(strSlice); i++ {
		number, _ := strconv.Atoi(strSlice[i])
		numSlice = append(numSlice, number)
	}

	return numSlice
}

func joinNumbers(sumNum int, number int) int {
	wholeString := strconv.Itoa(sumNum) + strconv.Itoa(number)

	wholeNumber, _ := strconv.Atoi(wholeString)
	return wholeNumber
}

func calcSumDetail(workingResults []int, number int, part byte) []int {

	for item, sumNum := range workingResults {
		if item < len(workingResults) {
			workingResults[item] = sumNum + number
		}
		workingResults = append(workingResults, sumNum*number)

		// For part b we have a new operator. || which joins the number to the left with the number to the right e.g. 12 || 34 becomes 1234
		if part == 'b' {
			joinedNumber := joinNumbers(sumNum, number)
			workingResults = append(workingResults, joinedNumber)
		}
	}

	if Debug {
		fmt.Println("working results:", workingResults)
	}

	return workingResults
}

func calcSum(numbers []int, resultLookingFor int, part byte) int {

	workingResults := make([]int, 0)

	for item, number := range numbers {
		if item == 0 {
			if Debug {
				fmt.Println("0 number is:", number)
			}
			workingResults = append(workingResults, number)
		} else {

			// Not the first number in the list. Carry out the sums and build out the list

			if Debug {
				fmt.Println("Other numbers:", number)
				fmt.Println("workingResults at this point is:", workingResults)
			}

			workingResults = calcSumDetail(workingResults, number, part)
		}
	}

	for _, resultItem := range workingResults {
		if resultItem == resultLookingFor {
			return resultItem
		}
	}

	return 0
}

func day07(filename string, part byte) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)

	for _, puzzleLine := range puzzleInput {
		puzzleParts := strings.Split(puzzleLine, ": ")
		resultNum, _ := strconv.Atoi(puzzleParts[0])
		numbers := buildSlice(puzzleParts[1])
		if Debug {
			fmt.Println("resultNum:", resultNum)
			fmt.Println("String:", puzzleParts[1])
			fmt.Println("Number slice:", numbers)
		}

		// Using 2 operators '+' and '*' work out if a combination of operators can create the
		// number given in resultNum. If it can, add to the result

		//numOperands := len(numbers) - 1
		//fmt.Println("numOperands", numOperands, "numbers:", numbers)

		result += calcSum(numbers, resultNum, part)
	}

	return result
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()
	Debug = debug

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", day07(filenamePtr, execPart))
	}
}
