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

// CardCompare. Returns TRUE if firstCard is stronger than secondCard
func cardCompare(firstCard string, secondCard string) bool {
	cardStrength := []byte{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
	for f := 0; f < len(firstCard); f++ {
		for _, s := range cardStrength {
			if byte(firstCard[f]) == s && byte(secondCard[f]) != s {
				return true
			}
			if byte(firstCard[f]) != s && byte(secondCard[f]) == s {
				return false
			}
		}
	}
	return false
}

func rankCards(cardHands []handStruct, typeIndex map[string][]int) {
	//types := []string{"FIVE", "FOUR", "FULLHOUSE", "THREE", "TWOPAIR", "ONEPAIR", "HIGHCARD"}
	var currRank int = 1

	types := []string{"HIGHCARD", "ONEPAIR", "TWOPAIR", "THREE", "FULLHOUSE", "FOUR", "FIVE"}

	for _, i := range types {
		locations, ok := typeIndex[i]
		if ok {
			if len(locations) == 1 {
				cardHands[locations[0]].rank = currRank
				currRank++
			} else {
				for _, loc := range locations {
					// Set the rank of each card. It will probably be wrong, but we'll sort and fix it
					cardHands[loc].rank = currRank
					currRank++
				}

				keepSorting := true
				for keepSorting {
					// Assume we stop looping
					keepSorting = false
					for sortPosOuter, locOuter := range locations {
						for sortPosInner, locInner := range locations {
							if sortPosOuter != sortPosInner {
								if cardCompare(cardHands[locOuter].cards, cardHands[locInner].cards) {
									if cardHands[locOuter].rank < cardHands[locInner].rank {
										tempRank := cardHands[locOuter].rank
										cardHands[locOuter].rank = cardHands[locInner].rank
										cardHands[locInner].rank = tempRank
										// We found something to change, so let's keep sorting
										keepSorting = true
									}
								}
							}
						}
					}
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

	for i := 0; i < len(cardHands); i++ {
		for _, j := range cardHands[i].cards {
			count := strings.Count(cardHands[i].cards, string(j))
			if debug {
				fmt.Printf("Hand: %s Count char: %c %d\n", cardHands[i].cards, j, count)
			}

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

		if debug {
			fmt.Println(cardHands[i])
			fmt.Println(typeIndex)
		}
	}

	rankCards(cardHands, typeIndex)

	// Now let's calculate the score of each hand and add them together
	for i := range cardHands {
		result += (cardHands[i].rank * cardHands[i].bid)
	}

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
