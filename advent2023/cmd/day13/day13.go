package main

import (
	"AdventOfCode-go/advent2023/utils"
	"fmt"
)

func flipList(cmpList [][]byte) []string {
	flipped := []string{}

	for j := 0; j < len(cmpList[0]); j++ {
		buffer := ""
		for i := 0; i < len(cmpList); i++ {
			buffer += string(cmpList[i][j])
		}
		flipped = append(flipped, buffer)
	}

	//fmt.Println(flipped)
	return flipped
}

func findReflectionScore(cmpList []string, flipped bool) int {

	for i := 0; i < len(cmpList)-1; i++ {
		if cmpList[i] == cmpList[i+1] {
			//fmt.Println("Found a match:", cmpList[i], i, i+1)

			fmt.Println("In FindReflectionScore:", flipped, i, len(cmpList), len(cmpList)-i+1)
			if i <= (len(cmpList) - i + 1) {
				fmt.Println("It did")
				var matching bool = true
				upwards := i - 1
				downwards := i + 2

				fmt.Printf("i: %d upwards: %d downwards: %d\n", i, upwards, downwards)
				if upwards < 0 || downwards >= len(cmpList) {
					fmt.Println("Exiting as hit the edge")
					if flipped {
						return i + 1
					} else {
						return 100 * (i + 1)
					}
				}
				for matching {
					if cmpList[upwards] != cmpList[downwards] {
						matching = false
						//fmt.Println("Not a match:", upwards, downwards, cmpList[upwards], cmpList[downwards], flipped)
						return 0
					} else {
						//fmt.Println("Still a match:", upwards, downwards, cmpList[upwards], cmpList[downwards], flipped)
					}
					upwards--
					downwards++

					if upwards < 0 || downwards >= len(cmpList) {
						if matching {
							if flipped {
								return i + 1
							} else {
								return 100 * (i + 1)
							}
						} else {
							return 0
						}
					}
				}
			} else {
				if flipped {
					return i + 1
				} else {
					return 100 * (i + 1)
				}
			}
		}
	}
	return 0
}

func copyStringToByte(horizList []string) [][]byte {

	byteArray := make([][]byte, len(horizList))
	for i := 0; i < len(horizList); i++ {
		byteArray[i] = make([]byte, len(horizList[0]))
	}

	for i := 0; i < len(horizList); i++ {
		for j := 0; j < len(horizList[i]); j++ {
			byteArray[i][j] = horizList[i][j]
		}
	}

	return byteArray
}

func day13(filename string, part byte, debug bool) int {
	var result int

	puzzleInput, _ := utils.ReadFile(filename)

	// Make a grid for the starting position to be built in

	horizList := make([]string, 0)

	for _, puzzleLine := range puzzleInput {
		if len(puzzleLine) == 0 {

			// When we find a break we want to work out where the reflection by checking both directions and calculate the score
			var foundScore int = 0

			//utils.PrintArrayString(horizList)

			listToCheck := copyStringToByte(horizList)

			//fmt.Println("======================")
			//utils.Print2DArrayByte(listToCheck)
			//fmt.Println("======================")

			foundScore = findReflectionScore(horizList, false)
			if foundScore == 0 {
				flippedList := flipList(listToCheck)
				//utils.PrintArrayString(flippedList)
				foundScore = findReflectionScore(flippedList, true)
				if debug {
					fmt.Println("Got score after flipped:", foundScore)
				}

			} else {
				if debug {
					fmt.Println("Got score:", foundScore)
				}
			}
			fmt.Println("Score:", foundScore, result)
			//fmt.Println("=====================")
			result += foundScore

			horizList = make([]string, 0)

		} else {
			horizList = append(horizList, puzzleLine)
		}
	}

	return result
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Result is: %d\n", day13(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("Result is: %d\n", day13(filenamePtr, execPart, debug))

	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
