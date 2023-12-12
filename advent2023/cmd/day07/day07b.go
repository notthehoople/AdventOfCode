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

// CardCompare. Returns TRUE if firstCard is stronger than secondCard
func cardCompareB(firstCard string, secondCard string) bool {
	cardStrength := []byte{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}
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

func rankCardsB(cardHands []handStruct, typeIndex map[string][]int) {
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
								if cardCompareB(cardHands[locOuter].cards, cardHands[locInner].cards) {
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

func day07b(filename string, part byte, debug bool) int {
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
		joker := 0
		for _, j := range cardHands[i].cards {
			count := strings.Count(cardHands[i].cards, string(j))
			// This is where part 2 comes in. We need to test out any Jokers here
			// and change the card TYPE based on what works best.

			if j == 'J' {
				joker = count
			}

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

		if joker > 0 {
			if debug {
				fmt.Println("Found joker", joker, cardHands[i].cards, cardHands[i].handtype)
			}

			switch cardHands[i].handtype {
			case "FOUR":
				cardHands[i].handtype = "FIVE"
			case "FULLHOUSE":
				cardHands[i].handtype = "FIVE"
			case "THREE":
				switch joker {
				case 1:
					cardHands[i].handtype = "FOUR"
				case 2:
					cardHands[i].handtype = "FIVE"
				case 3: // If jokers=3 and handtype == "THREE" then the jokers are the THREE so becomes a FOUR
					cardHands[i].handtype = "FOUR"
				default:
					panic("wrong number of jokers in THREE")
				}
			case "TWOPAIR": // **** No!
				switch joker {
				case 1:
					cardHands[i].handtype = "FULLHOUSE"
				case 2:
					cardHands[i].handtype = "FOUR"
				default:
					panic("Wrong count of jokers in TWOPAIR")
				}
			case "ONEPAIR":
				switch joker {
				case 1:
					cardHands[i].handtype = "THREE"
				case 2:
					// Jokers == 2 and handtype == "ONEPAIR" means the jokers are the pair
					cardHands[i].handtype = "THREE"
				default:
					panic("Wrong count of jokers in ONEPAIR")
				}
			case "HIGHCARD":
				switch joker {
				case 1:
					cardHands[i].handtype = "ONEPAIR"
				default:
					panic("Wrong count of jokers in HIGHCARD")
				}
			}
			if debug {
				fmt.Println("After joker change", joker, cardHands[i].cards, cardHands[i].handtype)
			}
		}

		location, ok := typeIndex[cardHands[i].handtype]
		if ok {
			location = append(location, i)
			typeIndex[cardHands[i].handtype] = location
		} else {
			typeIndex[cardHands[i].handtype] = []int{i}
		}

		if debug {
			fmt.Println(cardHands[i])
			fmt.Println(typeIndex)
		}
	}

	rankCardsB(cardHands, typeIndex)

	// Now let's calculate the score of each hand and add them together
	for i := range cardHands {
		result += (cardHands[i].rank * cardHands[i].bid)
	}

	return result
}
