package main

import (
	"fmt"
	"sort"
	"strings"
)

func removeListItem(list []string, itemToFind string) (result []string) {
	result = make([]string, 0)

	for _, item := range list {
		if item != itemToFind {
			result = append(result, item)
		}
	}
	return result
}

func itemInList(list []string, itemToFind string) bool {
	for _, item := range list {
		if item == itemToFind {
			return true
		}
	}
	return false
}

func joinedLists(first []string, second []string) (result []string) {
	result = make([]string, 0)

	for _, item := range first {
		if itemInList(second, item) {
			// item is in both lists
			result = append(result, item)
		}
	}

	return result
}

func workOutFoods(foodSource []string, debug bool) (ingredients []string, foundAllergens map[string]string, foundIngredients map[string]string) {
	ingredients = make([]string, 0)
	workingAllergens := make(map[string][]string)
	// Hedging bets for part 2. Let's build a map of allergen to ingredients (foundAllergens) and ingredients to allergens (foundIngredients)
	foundAllergens = make(map[string]string)
	foundIngredients = make(map[string]string)

	for _, line := range foodSource {
		if debug {
			fmt.Println("============================================")
			fmt.Println("Line:", line)
		}
		allergens := strings.SplitAfter(line, "(contains ")
		lineIngredients := strings.Split(strings.TrimSuffix(allergens[0], " (contains "), " ")
		lineAllergens := strings.Split(strings.TrimSuffix(allergens[1], ")"), ", ")

		if debug {
			fmt.Printf("Ingredients: %s\n", lineIngredients)
			fmt.Printf("Allergens: %s\n", lineAllergens)
		}

		for _, allergen := range lineAllergens {
			if debug {
				fmt.Println("Allergen:", allergen)
			}
			// Now store the ingredients against each allergen.
			if _, ok := workingAllergens[allergen]; ok {
				// we've seen this before. Reduce list to the common items between previous and current lines
				workingAllergens[allergen] = joinedLists(workingAllergens[allergen], lineIngredients)
			} else {
				workingAllergens[allergen] = lineIngredients
			}
		}

		for _, i := range lineIngredients {
			ingredients = append(ingredients, i)
		}
	}

	var keepLooping bool = true
	for keepLooping {
		// Work through the workingAllergens looking for matches where an allergen has a single ingredient
		// When we find one, record that in the foundAllergens map and remove from workingAllergens
		for allergen, ingredients := range workingAllergens {
			if len(ingredients) == 1 {
				// If there's only 1 ingredient we've found a match
				foundAllergens[allergen] = ingredients[0]
				foundIngredients[ingredients[0]] = allergen

				// Now need to remove this allergen from workging Allergens
				delete(workingAllergens, allergen)

				for remainingAllergen, remainingIngredients := range workingAllergens {
					workingAllergens[remainingAllergen] = removeListItem(remainingIngredients, ingredients[0])
				}
				break
			}
		}

		if len(workingAllergens) == 0 {
			// Found them all
			keepLooping = false
		}
	}

	if debug {
		fmt.Println("Found allergens:", foundAllergens)
		fmt.Println(workingAllergens)
	}

	return ingredients, foundAllergens, foundIngredients
}

// part b
// Go through the foundAllergens list. Build a list of strings containing the dangerous ingredients
// Sort it, then join all the strings together and return
func canonicalIngredientList(filename string, debug bool) string {
	var tmpAllergensList []string
	var tmpBuildResult []string

	puzzleInput, _ := readFile(filename)

	_, foundAllergens, _ := workOutFoods(puzzleInput, debug)

	if debug {
		fmt.Println("Found allergens map:", foundAllergens)
	}

	for allergen := range foundAllergens {
		tmpAllergensList = append(tmpAllergensList, allergen)
	}

	sort.Strings(tmpAllergensList)

	for _, allergen := range tmpAllergensList {
		tmpBuildResult = append(tmpBuildResult, foundAllergens[allergen])
	}

	return strings.Join(tmpBuildResult, ",")
}

// part a
func countNonAllergens(filename string, debug bool) int {
	var result int

	puzzleInput, _ := readFile(filename)

	ingredients, _, foundIngredients := workOutFoods(puzzleInput, debug)

	if debug {
		fmt.Println("Ingredients array:", ingredients)
		fmt.Println("Found ingredients map:", foundIngredients)
	}

	result = 0
	for _, i := range ingredients {
		if _, ok := foundIngredients[i]; !ok {
			result++
		}
	}

	return result
}

// Main routine
func main() {
	filenamePtr, execPart, debug, test := catchUserInput()

	if test {
		return
	}

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else if execPart == 'a' {
		fmt.Println("Non Allergens appear:", countNonAllergens(filenamePtr, debug))
	} else {
		fmt.Println("Caconical List:", canonicalIngredientList(filenamePtr, debug))
	}
}
