package main

import (
	"AdventOfCode-go/advent2023/utils"
	"fmt"
	"strconv"
	"strings"
)

// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2 is the order of strength of cards
// Five of a kind, where all five cards have the same label: AAAAA
// Four of a kind, where four cards have the same label and one card has a different label: AA8AA
// Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
// Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
// Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
// One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
// High card, where all cards' labels are distinct: 23456
//
// Hands are ordered by type
// If 2 hands have the same type, then compare each card in turn. Hand with the highest
// card from left to right wins
//
// Each hand wins an amount equal to its bid multiplied by its rank, where the weakest
// hand gets rank 1, the second-weakest hand gets rank 2, and so on up to the strongest
// hand. The "rank" is based on the number of hands. If 5 hands then 5 ranks, from 1 to 5
//
// multiplied by hand's rank to get a hand's score. Add up the scores to get result

type handStruct struct {
	cards    string
	bid      int
	handtype string
	five     bool
	four     bool
	three    bool
	two      int
	one      int
	rank     int
}

func rankCards(cardHands []handStruct, typeIndex map[string][]int, numHands int) {
	types := []string{"FIVE", "FOUR", "FULLHOUSE", "THREE", "TWOPAIR", "ONEPAIR", "HIGHCARD"}

	for _, i := range types {
		fmt.Println(i, numHands)
		locations, ok := typeIndex[i]
		if ok {
			if len(locations) == 1 {
				fmt.Printf("only one hand of type %s\n", i)
				cardHands[locations[0]].rank = numHands
				numHands--
			} else {
				for _, loc := range locations {
					// Need to be able to sort the cards. Use a bubble sort rather than comp?
					fmt.Println("Loc:", loc, i)
				}
			}
		} else {
			fmt.Printf("No hands of type %s\n", i)
		}
	}
}

func getHandType(cardHand handStruct) string {
	if cardHand.five {
		return "FIVE"
	}
	if cardHand.four {
		return "FOUR"
	}
	if cardHand.three && cardHand.two != 0 {
		return "FULLHOUSE"
	}
	if cardHand.three && cardHand.one == 2 {
		return "THREE"
	}
	if cardHand.two == 4 {
		return "TWOPAIR"
	}
	if cardHand.two == 2 {
		return "ONEPAIR"
	}
	return "HIGHCARD"
}

func day07(filename string, part byte, debug bool) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)
	cardHands := make([]handStruct, len(puzzleInput))

	typeIndex := make(map[string][]int)

	for i := 0; i < len(puzzleInput); i++ {
		parts := strings.Fields(puzzleInput[i])
		cardHands[i].cards = parts[0]
		cardHands[i].bid, _ = strconv.Atoi(parts[1])
	}
	//fmt.Println(cardHands)

	for i := 0; i < len(cardHands); i++ {
		for _, j := range cardHands[i].cards {
			count := strings.Count(cardHands[i].cards, string(j))
			fmt.Printf("Hand: %s Count char: %c %d\n", cardHands[i].cards, j, count)

			switch count {
			case 5:
				cardHands[i].five = true
			case 4:
				cardHands[i].four = true
			case 3:
				cardHands[i].three = true
			case 2:
				cardHands[i].two++
			case 1:
				cardHands[i].one++
			default:
				panic("Got an unexpected number of cards")
			}
		}
		handType := getHandType(cardHands[i])
		cardHands[i].handtype = handType

		location, ok := typeIndex[handType]
		if ok {
			location = append(location, i)
			typeIndex[handType] = location
		} else {
			typeIndex[handType] = []int{i}
		}

		fmt.Println(cardHands[i])
		fmt.Println(typeIndex)
	}

	rankCards(cardHands, typeIndex, len(cardHands))
	fmt.Println(cardHands)

	return result
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Result is: %d\n", day07(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("Result is: %d\n", day07(filenamePtr, execPart, debug))
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
