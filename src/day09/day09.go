package main

import (
	"fmt"
	"flag"
	"math"
)

func printMarbleBoard(marbleBoard []int, currentElf int, placedMarble int) {
	// Loop through the array and print each line

	fmt.Printf("[%d] ", currentElf)
	for i:= 0; i < len(marbleBoard); i++ {
		if i == placedMarble {
			fmt.Printf(" (%d) ", marbleBoard[i])
		} else {
			fmt.Printf("  %d  ", marbleBoard[i])
		}
	}
	fmt.Printf("\n")
}

// func recipeMakerA
// PARTA: Handles the proocessing of recipes to generate our answer
// recipesToMake - number of recipes to make prior to the answer being produced
// answersRequired - number of answers to create following the recipesToMake
// part - the section of the puzzle we're on (should be a)
//
// Need to know the position in the array that Elf 1 and Elf 2 has (can i increase this to elf 'n' easily?)
// Need to know each Elf's current recipe
// While the number of recipes created < the recipesToMake
//    add Elf 1 to Elf 2 recipe (elf 'n') to create new recipe. If > 9 then split into 2 digits
//    add the new recipe(s) to the end of the array
//    increase the number of recipes by 1 or 2
//    move each Elf forward by 1 + current recipe value. Loop from end of array to the start if needed
//        NOTE - Elf must move forward on the NEW version of the recipe board including newly added entries
/*func recipeMakerA(recipesToMake int, answersRequired int, printRecipeBoard bool, part byte) string {
	var recipeBoard[]int
	var elf1, elf2 int = 0, 1
	var newrecipe int = 0
	var resultString string
	
	for r := 0; r < recipesToMake + answersRequired; r++ {
		newrecipe = recipeBoard[elf1] + recipeBoard[elf2]

		if newrecipe / 10 >= 1 {
			// we've got a double digit. First mod gets the "tens" number. Second mod gets the "digits"
			// Add an extra 1 to 'r' since we've created 2 recipes
			recipeBoard = append(recipeBoard, int(math.Mod(float64(newrecipe / 10), 10)))
			recipeBoard = append(recipeBoard, int(math.Mod(float64(newrecipe), 10)))

			r++
		} else {
			// single digit so it's easier
			recipeBoard = append(recipeBoard, newrecipe)
		}

		// Divide the amount to move by the length of the current recipe board. The remainder is how much we move
		elf1 = int(math.Mod(float64(recipeBoard[elf1] + 1 + elf1), float64(len(recipeBoard))))
		elf2 = int(math.Mod(float64(recipeBoard[elf2] + 1 + elf2), float64(len(recipeBoard))))
		
	}
}*/

func insertMarbleToBoard(marbleBoard []int, marbleToPlace int, placedMarble int) []int {
	marbleBoard = append(marbleBoard, 0)

	//fmt.Println("In insertMarbleToBoard")
	//printMarbleBoard(marbleBoard, 0, 0)

	if len(marbleBoard) == 2 {
		marbleBoard[1] = marbleToPlace
	} else {
		for i := len(marbleBoard) - 1; i >= 0 ; i-- {
			if i > placedMarble {
				marbleBoard[i] = marbleBoard[i - 1]
			} else {
				if i == placedMarble {
					marbleBoard[i] = marbleToPlace
				} else {
					// Nothing to do
				}
			}
		}
	}

	//printMarbleBoard(marbleBoard, 0, 0)
	//fmt.Println("Exiting insertMarbleToBoard")

	return marbleBoard
}

func removeMarbleFromBoard(marbleBoard []int, removePos int) []int {
	var newMarbleBoard[]int

	newMarbleBoard = make([]int, len(marbleBoard) - 1)

	for i := 0; i < len(newMarbleBoard); i++ {
		if i < removePos {
			newMarbleBoard[i] = marbleBoard[i]
		} else {
			if i >= removePos {
				newMarbleBoard[i] = marbleBoard[i+1]
			}
		}
	}

	return newMarbleBoard
}


// func playMarbles
func playMarbles(numPlayers int, lastMarbleValue int, printBoard bool, part byte) int {
	var highestScore int = 0
	var marbleBoard[]int
	var elfScores[]int
	var tempPos int
	var currentElf, currentMarble, placedMarble int = 0, 0, 0

	elfScores = make([]int, numPlayers + 1)
	marbleBoard = make([]int, 1)

	currentElf = 1
	//printMarbleBoard(marbleBoard, 0, 0)


	for marbleToPlace := 0; marbleToPlace <= lastMarbleValue; marbleToPlace++ {
		if marbleToPlace == 0 {
			// Deal with the starting marble
			marbleBoard[0] = 0
			currentMarble = 0
			continue
		}

		if int(math.Mod(float64(marbleToPlace), float64(100000))) == 0 {
			fmt.Println(marbleToPlace)
		}

		if int(math.Mod(float64(marbleToPlace), float64(23))) == 0 {
			// When divisible by 23, different things happen:
			// - current Elf keeps the marbleToPlace, adding it to their elfScore
			// - the marble 7 marbles counter-clockwise (left) from the current marble is REMOVED and added to current Elf's score
			// - finally, the marble immediately clockwise (right) of the removed marble becomes the new current marble

			// Add the marble to place to the current elf's score
			elfScores[currentElf] += marbleToPlace
			tempPos = currentMarble - 7
			if tempPos < 0 {
				// This might be one off. Might need to subtract an extra
				tempPos = len(marbleBoard) + tempPos
			}

			//fmt.Println("Adding to score:", marbleBoard[tempPos])
			elfScores[currentElf] += marbleBoard[tempPos]

			marbleBoard = removeMarbleFromBoard(marbleBoard, tempPos)
			
			// If the removed marble is the last in the marbleBoard this won't work:
			currentMarble = tempPos

		} else {
			// Not divisible by 23, so just a marble to be placed
			// Place the marbleToPlace into the circle between the marbles that are 1 and 2 marbles clockwise of the current marble
			// e.g. in a large circle that means you'll end up with <current-marble> <something-else> <placed-marble>
			//      Need to take care of wrapping off the end of the marbleBoard back to the start again
			//      Also entries need to be inserted into the middle of the marbleBoard

			placedMarble = int(math.Mod(float64(currentMarble + 2), float64(len(marbleBoard))+1))
			//fmt.Printf("currentMarble: %d placedMarble: %d len(MarbleBoard): %d marbleToPlace: %d\n", currentMarble,
			//														placedMarble, len(marbleBoard), marbleToPlace)
			if placedMarble == 0 {
				placedMarble = 1
			}
			marbleBoard = insertMarbleToBoard(marbleBoard, marbleToPlace, placedMarble)
			currentMarble = placedMarble
		}

		if printBoard {
			printMarbleBoard(marbleBoard, currentElf, currentMarble)
		}

		currentElf++
		if currentElf > numPlayers {
			currentElf = 1
		}
	}

	// Find the highest score to return
	for i := 0; i < len(elfScores); i++ {
		if elfScores[i] > highestScore {
			highestScore = elfScores[i]
		}
	}

	return highestScore
}

// Main routine
func main() {
	var numPlayers, lastMarbleValue int

	flag.IntVar(&numPlayers, "players", 9, "Number of players in the game")
	flag.IntVar(&lastMarbleValue, "marble", 25, "Value of last marble in the game")
	printMarbleBoardPtr := flag.Bool("print", false, "Print the marble board as we go")
	execPartPtr := flag.String("part", "a", "Part of the puzzle to work on. a or b")

	flag.Parse()

	if numPlayers < 1 {
		fmt.Println("Minimum players: 1")
		return
	}

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Winning Elf's score:", playMarbles(numPlayers, lastMarbleValue, *printMarbleBoardPtr, 'a'))
	case "b":
		fmt.Println("Part b - Winning Elf's score:", playMarbles(numPlayers, lastMarbleValue, *printMarbleBoardPtr, 'b'))
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}