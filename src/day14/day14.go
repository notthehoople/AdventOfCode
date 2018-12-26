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

// func recipeMakerB
//
// Coming at it from the other direction. resultRequired contains the recipeBoard string we want to find.
// When we find that, we count the number of recipes BEFORE the recipeBoard string found. That's our answer to return
func recipeMakerB(resultRequired string, printRecipeBoard bool, part byte) int {
	var recipeBoard[]int
	var elf1, elf2 int = 0, 1
	var newrecipe int = 0
	var numberOfRecipes int = 0
	var areWeLooping bool = true
	var compareIsGood bool = true
	var resultIndex, firstMatch int = 0, 0
	//var doubleRecipeAdded int = 0
	
	// creates an array of integers with 2 entries
	recipeBoard = make([]int, 2)
	recipeBoard[0] = 3
	recipeBoard[1] = 7
	
	for areWeLooping {
		//doubleRecipeAdded = 0
		newrecipe = recipeBoard[elf1] + recipeBoard[elf2]

		if newrecipe / 10 >= 1 {
			// we've got a double digit. First mod gets the "tens" number. Second mod gets the "digits"
			// Add an extra 1 to 'r' since we've created 2 recipes
			recipeBoard = append(recipeBoard, int(math.Mod(float64(newrecipe / 10), 10)))
			recipeBoard = append(recipeBoard, int(math.Mod(float64(newrecipe), 10)))
			//doubleRecipeAdded = 1

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

		// Is the recipe board big enough to check for our required result?
		if len(recipeBoard) > len(resultRequired) + 5 {
			// Yes it is. Let's start looking

			compareIsGood = true
			resultIndex = 0
			for i := (len(recipeBoard) - len(resultRequired)) - 5; i < len(recipeBoard) && compareIsGood; i++ {					

				if strconv.Itoa(recipeBoard[i]) == string(resultRequired[resultIndex]) {
					if resultIndex == 0 {
						// first match so grab the position in recipeBoard
						firstMatch = i
					}
					//if resultIndex == 4 {
						//fmt.Printf("Character match! i is %d, recipeBoard: %s, resultIndex: %d, resultRequired: %c\n", i-4, strconv.Itoa(recipeBoard[i-4]), resultIndex-4, resultRequired[resultIndex-4])
						//fmt.Printf("Character match! i is %d, recipeBoard: %s, resultIndex: %d, resultRequired: %c\n", i-3, strconv.Itoa(recipeBoard[i-3]), resultIndex-3, resultRequired[resultIndex-3])
						//fmt.Printf("Character match! i is %d, recipeBoard: %s, resultIndex: %d, resultRequired: %c\n", i-2, strconv.Itoa(recipeBoard[i-2]), resultIndex-2, resultRequired[resultIndex-2])
						//fmt.Printf("Character match! i is %d, recipeBoard: %s, resultIndex: %d, resultRequired: %c\n", i-1, strconv.Itoa(recipeBoard[i-1]), resultIndex-1, resultRequired[resultIndex-1])
						//fmt.Printf("Character match! i is %d, recipeBoard: %s, resultIndex: %d, resultRequired: %c\n", i, strconv.Itoa(recipeBoard[i]), resultIndex, resultRequired[resultIndex])
						//fmt.Printf("Character match! i is %d, recipeBoard: %s, resultIndex: %d, resultRequired: %c\n", i+1, strconv.Itoa(recipeBoard[i+1]), resultIndex, resultRequired[resultIndex])

					//}
					resultIndex++
					if resultIndex == len(resultRequired) {
						//fmt.Println("We are DONE!")
						compareIsGood = false
						areWeLooping = false
						numberOfRecipes = firstMatch
						break
					}
				} else {
					//fmt.Printf("No match. i is %d, recipeBoard: %s, resultIndex: %d, resultRequired: %c\n", i, strconv.Itoa(recipeBoard[i]), resultIndex, resultRequired[resultIndex])
					compareIsGood = false
				}
			} 
		}
	}

	// Build the answer string to return
	//for i := recipesToMake; i < recipesToMake + answersRequired; i++ {
	//	resultString += strconv.Itoa(recipeBoard[i])
	//}

	//printIntegerArray(recipeBoard, 0, 0)

	return numberOfRecipes
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
func recipeMakerA(recipesToMake int, answersRequired int, printRecipeBoard bool, part byte) string {
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

	printIntegerArray(recipeBoard, 0, 0)

	return resultString
}

// Main routine
func main() {
	var recipesToMake, answersRequired int = 0, 0

	recipesToMakePtr := flag.String("recipes", "10", "Number of recipes to make prior to answer")
	answersRequiredPtr := flag.String("answers", "10", "Number of answers to make prior following recipes")
	resultRequiredPtr := flag.String("result", "51589", "part b ONLY: this is the answer to look for")
	printRecipeBoardPtr := flag.Bool("print", false, "Print the recipe board as we go")
	execPartPtr := flag.String("part", "a", "Part of the puzzle to work on. a or b")

	flag.Parse()

	recipesToMake, _ = strconv.Atoi(*recipesToMakePtr)
	answersRequired, _ = strconv.Atoi(*answersRequiredPtr)

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Ten recipes following on:", recipeMakerA(recipesToMake, answersRequired, *printRecipeBoardPtr, 'a'))
	case "b":
		fmt.Println("Part b - Recipes before given answer:", recipeMakerB(*resultRequiredPtr, *printRecipeBoardPtr, 'b'))
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}