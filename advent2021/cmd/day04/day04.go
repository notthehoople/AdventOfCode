package main

import (
	"aoc/advent2021/utils"
	"fmt"
	"strconv"
	"strings"
)

type cards struct {
	number int
	marked bool
}

// Count the number of won cards. If we've got 1 left that hasn't won, we've finished
func lastCard(wonCards []bool) bool {
	var countWonCards int

	for _, won := range wonCards {
		if won {
			countWonCards++
		}
	}

	if countWonCards == len(wonCards) {
		return true
	} else {
		return false
	}
}

func calcScore(card []cards) int {
	var unmarkedCardSum int
	for _, value := range card {
		if !value.marked {
			unmarkedCardSum += value.number
		}
	}
	return unmarkedCardSum
}

func playBingo(bingoNumbers []int, bingoCards [][]cards, part byte, debug bool) int {

	wonCards := make([]bool, len(bingoCards))

	for _, calledNumber := range bingoNumbers {
		// Look through all the cards for the next number in the bingoNumbers
		for _, card := range bingoCards {
			for cardPos, cardValue := range card {
				if cardValue.number == calledNumber {
					card[cardPos].marked = true
				}
			}
		}
		// Now check for a winning line in each card
		for cardPos, card := range bingoCards {
			var horizontalPos int = 0
			for horizontalPos < 25 {
				if card[horizontalPos].marked &&
					card[horizontalPos+1].marked &&
					card[horizontalPos+2].marked &&
					card[horizontalPos+3].marked &&
					card[horizontalPos+4].marked {
					if part == 'a' {
						// When found, score is sum of all unmarked numbers * number just called
						return calcScore(card) * calledNumber
					} else {
						wonCards[cardPos] = true
						if lastCard(wonCards) {
							return calcScore(card) * calledNumber
						}
					}
				}
				horizontalPos += 5
			}

			var verticalPos int = 0
			for verticalPos < 5 {
				if card[verticalPos].marked &&
					card[verticalPos+5].marked &&
					card[verticalPos+10].marked &&
					card[verticalPos+15].marked &&
					card[verticalPos+20].marked {
					if part == 'a' {
						// When found, score is sum of all unmarked numbers * number just called
						return calcScore(card) * calledNumber
					} else {
						wonCards[cardPos] = true
						if lastCard(wonCards) {
							return calcScore(card) * calledNumber
						}
					}
				}
				verticalPos++
			}
		}
	}

	return 0
}

func processPuzzleInput(puzzleInput []string, debug bool) (bingoNumbers []int, bingoCards [][]cards) {
	/*
		Bingo input file is a list of comma separated numbers on the first line, followed
		by a blank line, then blocks of 5 x 5 numbers separated by spaces. Each block
		is separated by a blank line.
	*/
	tempNumbers := strings.Split(puzzleInput[0], ",")
	bingoNumbers = make([]int, len(tempNumbers))
	for i := 0; i < len(tempNumbers); i++ {
		bingoNumbers[i], _ = strconv.Atoi(tempNumbers[i])
	}

	// TODO: run a make for bingoCards. Work out how much we need
	bingoCards = make([][]cards, (len(puzzleInput)-2)/6+1)
	for i := 0; i < (len(puzzleInput)-2)/6+1; i++ {
		bingoCards[i] = make([]cards, 25)
	}

	// Blocks of 5x5 cards start at puzzleInput[2]
	var cardNumber int
	for i := 2; i < len(puzzleInput); i += 6 {
		// Let's deal with each card as a block
		var arrayPos int = 0
		for j := i; j < i+5; j++ {
			if debug {
				fmt.Printf("Block: %d Line: %s arrayPos: %d\n", j, puzzleInput[j], arrayPos)
			}
			tempCardLine := strings.Fields(puzzleInput[j])
			for _, value := range tempCardLine {
				bingoCards[cardNumber][arrayPos].number, _ = strconv.Atoi(value)
				arrayPos++
			}
			if debug {
				fmt.Println(bingoCards[cardNumber])
			}
		}
		cardNumber++
	}
	if debug {
		fmt.Println("BingoNumbers:", bingoNumbers)
		fmt.Println("BingoCards:", bingoCards)
	}

	return bingoNumbers, bingoCards
}

func solveDay(filename string, part byte, debug bool) int {
	puzzleInput, _ := utils.ReadFile(filename)
	bingoNumbers, bingoCards := processPuzzleInput(puzzleInput, debug)

	return playBingo(bingoNumbers, bingoCards, part, debug)
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
