package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

// Prints the map list
func printMap(tempMap []string) {
	for i := 0; i < len(tempMap); i++ {
		fmt.Printf("%s\n", tempMap[i])
	}
}

// Part B function. Builts a "distance" map for the supplied orbitList
//		Expectation is that this will have 2 "top level" calls - one for "YOU" and one for "SAN"
func createDistanceMap(orbitor string, orbitList map[string]string, distanceMap map[string]int) {
	// We count 1 each time there is a higher level object to orbit i.e. if the orbitor we're passed is itself orbiting something
	// Most likely we'll stop counting at COM but let's not make that assumption
	thingWeOrbit, moreObjects := orbitList[orbitor]
	if moreObjects {
		// There are more things to count. Take note of the distance of this step, then continue
		distanceMap[thingWeOrbit] = distanceMap[orbitor] + 1
		createDistanceMap(thingWeOrbit, orbitList, distanceMap)
	}
}

func countNumberOrbits(orbitor string, orbitList map[string]string) int {
	// We count 1 each time there is a higher level object to orbit i.e. if the orbitor we're passed is itself orbiting something
	// Most likely we'll stop counting at COM but let's not make that assumption
	thingWeOrbit, moreObjects := orbitList[orbitor]
	if moreObjects {
		// There are more things to count
		return countNumberOrbits(thingWeOrbit, orbitList) + 1
	}

	return 0
}

// func
// Returns:
func bestMonitoringAsteroid(filename string, debug bool, part byte) int {
	// Read contents of file into a string array
	fileContents, _ := readLines(filename)

	// Can I use an array of strings here and reference every byte individually?
	// If so then fileContents can be the baseSpaceMap

	//baseSpaceMap := make( Want byte byte here )

	// Coords are X,Y. Top left is 0,0 and the space directly to the right is 1,0
	// A monitoring station can detect any asteroid to which it has direct line of sight - that is,
	// there cannot be another asteroid exactly between them
	//
	// Looks like I need to detect (and discount) and asteroid that is on the exact same X, the same Y, and the same diagonal in any direction
	// From the examples it looks like everything else can be counted

	// Create map
	// Read the map from the file given
	// Loop through the whole map one space at a time
	//   Is there an asteroid? If no, continue loop
	//   For an asteroid, loop through the map
	//     Create a new map for counting
	//     For same X left, look for first asteroid. Then discount any asteroids beyond that one on X
	//     For same X right, look for first asteroid. Then discount any asteroids beyond that one on X
	//     For same Y up, look for first asteroid. Then discount any asteroids beyond that one on X
	//     For same Y down, look for first asteroid. Then discount any asteroids beyond that one on X
	//     For same X and Y northwest (X-1,Y-1), look for first asteroid. Then discount any asteroids beyond that one on X
	//     For same X and Y northeast (X+1,Y-1), look for first asteroid. Then discount any asteroids beyond that one on X
	//     For same X and Y southwest (X-1,Y+1), look for first asteroid. Then discount any asteroids beyond that one on X
	//     For same X and Y southeast (X+1,Y+1), look for first asteroid. Then discount any asteroids beyond that one on X
	//   EndLoop
	//
	//   Loop through the created map for the asteroid
	//     Count the number of visible asteroids
	//   EndLoop
	//
	//   If this is the best so far, take note of the asteroid position and the number of asteroids it can see
	// EndLoop
	//
	// Print out the best asteroid position and the number of asteroids it can see

	printMap(fileContents)

	if part == 'a' {

	}

	return 0
}

// Main routine
func main() {
	var debug bool

	filenamePtr := flag.String("file", "input.txt", "Filename containing the program to run")
	execPartPtr := flag.String("part", "a", "Which part of day06 do you want to calc (a or b)")
	flag.BoolVar(&debug, "debug", false, "Turn debug on")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Number of orbits:", bestMonitoringAsteroid(*filenamePtr, debug, 'a'))
	case "b":
		fmt.Println("Part b - Not implemented yet")

	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
