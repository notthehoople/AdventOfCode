package main

import (
	"AdventOfCode-go/advent2023/utils"
	"fmt"
	"strconv"
	"strings"
)

type rangesStruct struct {
	destinationStart int
	sourceStart      int
	rangeLength      int
}

func findDestination(mapping []rangesStruct, item int) int {
	var difference int
	for i := 0; i < len(mapping); i++ {
		if item >= mapping[i].sourceStart && item <= mapping[i].sourceStart+mapping[i].rangeLength {
			difference = item - mapping[i].sourceStart
			return mapping[i].destinationStart + difference
		}
	}
	// If it doesn't match a range in the map, it maps to the same value
	return item
}

func buildMap(line int, puzzleInput []string) []rangesStruct {
	var lineValues rangesStruct
	var returnMapping []rangesStruct

	for i := line + 1; i < len(puzzleInput) && puzzleInput[i] != ""; i++ {
		fmt.Sscanf(puzzleInput[i], "%d %d %d", &lineValues.destinationStart, &lineValues.sourceStart, &lineValues.rangeLength)
		returnMapping = append(returnMapping, lineValues)
	}
	return returnMapping
}

//
//
//

func day05(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)

	seeds := strings.Split(puzzleInput[0], ":")
	seedNumbersStr := strings.Fields(strings.TrimSpace(seeds[1]))
	seedNumbers := make([]int, len(seedNumbersStr))
	for key, value := range seedNumbersStr {
		seedNumbers[key], _ = strconv.Atoi(value)
	}

	if debug {
		for i, j := range seedNumbers {
			fmt.Printf("Seed: %d Number: %d\n", i, j)
		}
		fmt.Println(seedNumbersStr)
	}

	var seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation []rangesStruct

	for line, puzzleLine := range puzzleInput {
		if line == 0 {
			// Skip the seed line
			continue
		}

		switch puzzleLine {
		case "seed-to-soil map:":
			seedToSoil = buildMap(line, puzzleInput)
		case "soil-to-fertilizer map:":
			soilToFertilizer = buildMap(line, puzzleInput)
		case "fertilizer-to-water map:":
			fertilizerToWater = buildMap(line, puzzleInput)
		case "water-to-light map:":
			waterToLight = buildMap(line, puzzleInput)
		case "light-to-temperature map:":
			lightToTemperature = buildMap(line, puzzleInput)
		case "temperature-to-humidity map:":
			temperatureToHumidity = buildMap(line, puzzleInput)
		case "humidity-to-location map:":
			humidityToLocation = buildMap(line, puzzleInput)
		}

	}

	var result int = 999999999

	if part == 'a' {
		// Now we've built the maps, lets work through each seed
		var destinationFromSource int
		for _, seed := range seedNumbers {
			if debug {
				fmt.Println("Seed:", seed)
			}
			destinationFromSource = findDestination(seedToSoil, seed)
			destinationFromSource = findDestination(soilToFertilizer, destinationFromSource)
			destinationFromSource = findDestination(fertilizerToWater, destinationFromSource)
			destinationFromSource = findDestination(waterToLight, destinationFromSource)
			destinationFromSource = findDestination(lightToTemperature, destinationFromSource)
			destinationFromSource = findDestination(temperatureToHumidity, destinationFromSource)
			destinationFromSource = findDestination(humidityToLocation, destinationFromSource)
			if debug {
				fmt.Println("Location:", destinationFromSource)
			}
			if destinationFromSource < result {
				result = destinationFromSource
			}
		}
	}

	return result
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Result is: %d\n", day05(filenamePtr, execPart, debug))
	case 'b':
		//fmt.Printf("Result is: %d\n", day04b(filenamePtr, execPart, debug))
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
