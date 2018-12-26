package main

import (
	"fmt"
	"flag"
	"strconv"
	"math"
)

func printIntegerArray(tempInt []int, elf1 int, elf2 int) {
	// Loop through the array and print each line
	for i:= 0; i < len(tempInt); i++ {
		if i == elf1 {
			fmt.Printf("(%d) ", tempInt[i])
		} else {
			if i == elf2 {
				fmt.Printf("[%d] ", tempInt[i])
			} else {
				fmt.Printf("%d ", tempInt[i])
			}
		}
	}
	fmt.Printf("\n")
}

// func recipeMaker
// Handles the proocessing of recipes to generate our answer
// recipesToMake - number of recipes to make prior to the answer being produced
// answersRequired - number of answers to create following the recipesToMake
// part - the section of the puzzle we're on (a or b)
//
// Need to know the position in the array that Elf 1 and Elf 2 has (can i increase this to elf 'n' easily?)
// Need to know each Elf's current recipe
// While the number of recipes created < the recipesToMake
//    add Elf 1 to Elf 2 recipe (elf 'n') to create new recipe. If > 9 then split into 2 digits
//    add the new recipe(s) to the end of the array
//    increase the number of recipes by 1 or 2
//    move each Elf forward by 1 + current recipe value. Loop from end of array to the start if needed
//        NOTE - Elf must move forward on the NEW version of the recipe board including newly added entries
func recipeMaker(recipesToMake int, answersRequired int, printRecipeBoard bool, part byte) string {
	var recipeBoard[]int
	var elf1, elf2 int = 0, 1
	var newrecipe int = 0
	var resultString string
	
	// creates an array of integers with 2 entries
	recipeBoard = make([]int, 2)
	recipeBoard[0] = 3
	recipeBoard[1] = 7
	
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
		
		if printRecipeBoard {
			printIntegerArray(recipeBoard, elf1, elf2)
		}
	}

	// Build the answer string to return
	for i := recipesToMake; i < recipesToMake + answersRequired; i++ {
		resultString += strconv.Itoa(recipeBoard[i])
	}

	return resultString
}

// Main routine
func main() {
	var recipesToMake, answersRequired int = 0, 0

	recipesToMakePtr := flag.String("recipes", "10", "Number of recipes to make prior to answer")
	answersRequiredPtr := flag.String("answers", "10", "Number of answers to make prior following recipes")
	printRecipeBoardPtr := flag.Bool("print", false, "Print the recipe board as we go")
	execPartPtr := flag.String("part", "a", "Part of the puzzle to work on. a or b")

	flag.Parse()

	recipesToMake, _ = strconv.Atoi(*recipesToMakePtr)
	answersRequired, _ = strconv.Atoi(*answersRequiredPtr)

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Ten recipes following on:", recipeMaker(recipesToMake, answersRequired, *printRecipeBoardPtr, 'a'))
	case "b":
		fmt.Println("Part b - Not there yet")
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}