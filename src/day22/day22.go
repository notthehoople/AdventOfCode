package main

import (
	"fmt"
	"bufio"
	"os"
	"flag"
	//"strings"
	"runtime/pprof"
	"log"
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

// The cave is divided into square regions which are either dominantly rocky, narrow, or wet (called its type).
// Each region occupies exactly one coordinate in X,Y format where X and Y are integers and zero or greater.
// (Adjacent regions can be the same type.)

// The scan (your puzzle input) is not very detailed: it only reveals the depth of the cave system and the
// coordinates of the target. However, it does not reveal the type of each region. The mouth of the cave is at 0,0.

// The man explains that due to the unusual geology in the area, there is a method to determine any region's
// type based on its erosion level. The erosion level of a region can be determined from its geologic index.
// The geologic index can be determined using the first rule that applies from the list below:

// The region at 0,0 (the mouth of the cave) has a geologic index of 0.
// The region at the coordinates of the target has a geologic index of 0.
// If the region's Y coordinate is 0, the geologic index is its X coordinate times 16807.
// If the region's X coordinate is 0, the geologic index is its Y coordinate times 48271.
// Otherwise, the region's geologic index is the result of multiplying the erosion levels of the regions at
// X-1,Y and X,Y-1.
// A region's erosion level is its geologic index plus the cave system's depth, all modulo 20183. Then:

// If the erosion level modulo 3 is 0, the region's type is rocky.
// If the erosion level modulo 3 is 1, the region's type is wet.
// If the erosion level modulo 3 is 2, the region's type is narrow.
// For example, suppose the cave system's depth is 510 and the target's coordinates are 10,10. Using % to
// represent the modulo operator, the cavern would look as follows:

// At 0,0, the geologic index is 0. The erosion level is (0 + 510) % 20183 = 510. The type is 510 % 3 = 0, rocky.
// At 1,0, because the Y coordinate is 0, the geologic index is 1 * 16807 = 16807. The erosion level is
// (16807 + 510) % 20183 = 17317. The type is 17317 % 3 = 1, wet.
// At 0,1, because the X coordinate is 0, the geologic index is 1 * 48271 = 48271. The erosion level is
// (48271 + 510) % 20183 = 8415. The type is 8415 % 3 = 0, rocky.
// At 1,1, neither coordinate is 0 and it is not the coordinate of the target, so the geologic index is the
// erosion level of 0,1 (8415) times the erosion level of 1,0 (17317), 8415 * 17317 = 145722555. The erosion
// level is (145722555 + 510) % 20183 = 1805. The type is 1805 % 3 = 2, narrow.
// At 10,10, because they are the target's coordinates, the geologic index is 0. The erosion level is
// (0 + 510) % 20183 = 510. The type is 510 % 3 = 0, rocky.
// Drawing this same cave system with rocky as ., wet as =, narrow as |, the mouth as M, the target as T,
// with 0,0 in the top-left corner, X increasing to the right, and Y increasing downward, the top-left corner
// of the map looks like this:

func calcRiskLevel(fileName string, part string, debug bool) int {
	var totalRiskLevel = 0
	
	if part == "a" {
		// Let's do part a
	} else {
		// part b is where it's at
	}

	return totalRiskLevel
}

// Main routine
func main() {
	var debug bool = false

	fileNamePtr := flag.String("file", "input.txt", "A filename containing input strings")
	execPartPtr := flag.String("part", "a", "Which part of day05 do you want to calc (a or b)")
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to `file`")
	flag.BoolVar(&debug, "debug", false, "Turn debug on or not")

	flag.Parse()

    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal("could not create CPU profile: ", err)
        }
        defer f.Close()
        if err := pprof.StartCPUProfile(f); err != nil {
            log.Fatal("could not start CPU profile: ", err)
        }
        defer pprof.StopCPUProfile()
    }

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Total Risk Level:", calcRiskLevel(*fileNamePtr, "a", debug))
	case "b":
		fmt.Println("Part b - Not implemented yet")
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}