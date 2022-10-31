package main

import (
	"fmt"
	"bufio"
	"os"
	"flag"
	"strconv"
	"strings"
)

// Read the text file passed in by name into a array of strings
// Returns the array as the first return variable
func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
	  return nil, err
	}
	defer file.Close()
  
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	  lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func printStringArray(tempString []string) {
	// Loop through the array and print each line
	for i:= 0; i < len(tempString); i++ {
		fmt.Println(tempString[i])
	}
}

// func buildInitialPots
// File contents comes in through filecontents
// We grab the first line of it, process it, then stick it into tempPlantPots and return the array as first return
// We then deal with the rest of the input, process it, then stick into tempRuleList and return the array as second return
func buildInitialPotsAndRules(filecontents []string) ([]string, []string) {
	var tempPlantPots []string
	var tempRuleList []string

	// Add 5 empty plant pots to the left of the initial state. These may affect the contents of the pots we look at
	// When we examine state, we start from 0 and work our way onwards through the string
	tempPlantPots = append(tempPlantPots, strings.TrimPrefix(filecontents[0], "initial state: "))
	tempPlantPots[0] = "...................." + tempPlantPots[0] + "....................................................................................................................................................................................................................................................................................................................................."

	for i := 2; i < len(filecontents); i++ {
		tempRuleList = append(tempRuleList, filecontents[i])
	}
	
	return tempPlantPots, tempRuleList
}

// Adds up the plant pot numbers with plants in them
// Since 20 empty pots are added at the beginning of the series, the first 0 to 19 count as NEGATIVE in additions
func sumPlantPotsWithPlants(plantPots []string, generations int) int {
	var sumOfAllThings int = 0
	const modifier int = -20

	for i := 0; i < len(plantPots[generations]); i++ {
		if plantPots[generations][i] == '#' {
			sumOfAllThings += i + modifier
		}
	}

	return sumOfAllThings
}


// func plantProcessor
// Handles the generations of plant plots
func plantProcessor(fileName string, generations int, part byte) int {
	var checkPots string
	var ruleListParts [] string
	var changed bool = false
	// The input file contains an initial state indicated by "initial state:" and a set of rules to be followed
	// Rules contain the '=>' directive
	
	// Read contents of file into a string array
	fileContents, _ := readLines(fileName)

	plantPots, ruleList := buildInitialPotsAndRules(fileContents)
	
	printStringArray(ruleList)

	// Loop until we reach the number of generations we need
	//   Loop through each character in the string, grabbing 2 pots before and after like: LLxRR
	//     then loop through each rule and compare it to our LLxRR block
	//     build the next generation in the next string in the array so we keep a copy of every generation

	for g := 0; g <= generations; g++ {
		// Setup a new string for us to build our results into
		plantPots = append(plantPots, "")
		
		//fmt.Println(plantPots[g])
		//fmt.Println(plantPots[g+1])

		plantPots[g+1] = plantPots[g+1] + string(plantPots[g][0])
		plantPots[g+1] = plantPots[g+1] + string(plantPots[g][1])

		for i := 2; i < len(plantPots[g]) - 2; i++ {
			checkPots = string(plantPots[g][i-2]) + string(plantPots[g][i-1]) + string(plantPots[g][i]) + string(plantPots[g][i+1]) + string(plantPots[g][i+2])
			//fmt.Printf("CheckPots: %s\n", checkPots)
			changed = false
			for r := 0; r < len(ruleList); r++ {
				ruleListParts = strings.Fields(ruleList[r])

				//fmt.Printf("rule list: %s\n", ruleListParts[0])
				if checkPots == ruleListParts[0] {
					//fmt.Println("match")
					plantPots[g+1] = plantPots[g+1] + ruleListParts[2]
					//fmt.Println("Changed:", plantPots[g+1])
					changed = true
				}
			}
			if !changed {
				plantPots[g+1] = plantPots[g+1] + "."
				//fmt.Println("LeftAss:", plantPots[g+1])
			}
		}
		plantPots[g+1] = plantPots[g+1] + string(plantPots[g][len(plantPots[g])-2])
		plantPots[g+1] = plantPots[g+1] + string(plantPots[g][len(plantPots[g])-1])
	}

	for i := 0; i < len(plantPots); i++ {
		fmt.Printf("End of generation %d result %s\n", i, plantPots[i])
	}

	return sumPlantPotsWithPlants(plantPots, generations)
}

// Main routine
func main() {
	var generations, resultForTwoHundred, ourBigResult int = 0, 0, 0

	fileNamePtr := flag.String("file", "input1.txt", "A filename containing input strings")
	generationsPtr := flag.String("generations", "10", "Number of generations to plot")
	execPartPtr := flag.String("part", "a", "Which part of day12 do you want to calc (a or b)")

	flag.Parse()

	generations, _ = strconv.Atoi(*generationsPtr)

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Sum of all plants:", plantProcessor(*fileNamePtr, generations, 'a'))
	case "b":
		// Formula is: result for 200 generations + (no of generations wanted - 200)*81
		resultForTwoHundred = plantProcessor(*fileNamePtr, 200, 'b')
		ourBigResult = resultForTwoHundred + (50000000000 - 200) * 81

		fmt.Println("Part b - Sum of all plants for 50000000000:", ourBigResult)
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}