package main

import (
	"fmt"
	"bufio"
	"os"
	"flag"
	"sort"
	"strconv"
	"strings"
)

type GuardSleeping struct {
	date string
	guardID string
	minute [60]byte
}

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

func printGuardMap(tempGuard [2000]GuardSleeping) {
	// Loop through the array and print each line
	for i:= 0; tempGuard[i].date != ""; i++ {
		fmt.Printf("%s %s ", tempGuard[i].date, tempGuard[i].guardID)
		for j := 0; j < len(tempGuard[i].minute); j++ {
			fmt.Printf("%c ", tempGuard[i].minute[j])
		}
		fmt.Println("")
	}
}

// Returns:
//   - string: the ID of the Guard who sleeps most on the chosen Minute
//   - int: the minute of the hour that is the most slept
func mostSleepyMinute(tempguardMap [2000]GuardSleeping) (string, int) {
	minutes := map[string]*[60]int{}

	// Loop through the built guardMap and count the number of sleeping minutes per guard
	for i := 0; tempguardMap[i].date != ""; i++ {
		if minutes[tempguardMap[i].guardID] == nil {
			// Allocate space in the minutes map for the guard, if that guard isn't in the map already
			minutes[tempguardMap[i].guardID] = &[60]int{}
		}
		for j := 0; j < len(tempguardMap[i].minute); j++ {
			if tempguardMap[i].minute[j] == '#' {
				// Add 1 to tempGuardNumSlept[tempguardMap[i].guardID] then lookup Minute
				minutes[tempguardMap[i].guardID][j]++
			}
		}
	}
	
	var overallSleepiestGuard string
	var mostSleptMinute int = 0
	var mostTimesSleptPerMinute int = 0

	for tempGuard := range minutes {
		for i := 0; i < 60; i++ {
			if minutes[tempGuard][i] > mostTimesSleptPerMinute {
				mostTimesSleptPerMinute = minutes[tempGuard][i]
				mostSleptMinute = i
				overallSleepiestGuard = tempGuard
			}
		}
	}

	return overallSleepiestGuard, mostSleptMinute
}

// Returns:
//   - string: the ID of the Sleepiest Guard
//   - int: the minute of the hour that the guard is most likely to be asleep
func mostSleepyGuard(tempguardMap [2000]GuardSleeping) (string, int) {
	var tempMinutes int = 0
	var sleepiestGuard = make(map[string]int)
	var sleepiestMinute = make(map[int]int)
	var mostMinutesSlept int = 0
	var overallSleepiestGuard string

	// Loop through the built guardMap and count the number of sleeping minutes per guard

	for i := 0; tempguardMap[i].date != ""; i++ {
		tempMinutes = 0
		for j := 0; j < len(tempguardMap[i].minute); j++ {
			if tempguardMap[i].minute[j] == '#' {
				tempMinutes++
			}
		}

		sleepiestGuard[tempguardMap[i].guardID] += tempMinutes
	}

	for k, tempval := range sleepiestGuard {
		if tempval > mostMinutesSlept {
			mostMinutesSlept = tempval
			overallSleepiestGuard = k
		}
	}

	// Now let's get the sleepiest minute
	for i := 0; tempguardMap[i].date != ""; i++ {
		if tempguardMap[i].guardID == overallSleepiestGuard {
			for j := 0; j < len(tempguardMap[i].minute); j++ {
				if tempguardMap[i].minute[j] == '#' {
					sleepiestMinute[j]++ 
				}
			}
		}
	}

	// The largest value in the sleepiestGuard map shows the minute he's asleep most
	
	var answerCount int = 0
	var answerMinute int = 0
	//var tempMin string
	//var tempCount int = 0

	for tempMin, tempCount := range sleepiestMinute {
		if tempCount > answerCount {
			answerMinute = tempMin
			answerCount = tempCount
		}
	}

	return overallSleepiestGuard, answerMinute
}


// Handles everything needed to work out the sleepy Guard (Day04 part A)
// Situations to take account of:
//     List needs to be sorted before use
//     A single guard can fall asleep and wake up multiple times during their shift
//     A single guard can stay awake during their shift so have 2 "Guard" input lines next to each other
//     A guard can come on duty at 23:xx hours the previous day
//     No guard ever falls asleep at 23:xx hours, always 00:xx hours
//     Any guard that doesn't sleep isn't important
//     There should never be a "falls asleep" without a "wakes up"
func sleepyGuard(fileName string, part byte) int {

	var guardMap [2000]GuardSleeping			// Our record of sleepy guards
	var dateVar, actionVar, extraVar, currentGuard, wakeUpDate string
	var sleepStartTime int = 0
	var wakeUpTime int = 0
	var storagePoint int = 0
	var chosenGuard string
	var chosenMinute int = 0

	for i := 0; i < len(guardMap); i++ {
		for j := 0; j < len(guardMap[i].minute); j++ {
			guardMap[i].minute[j] = '.'
		}
	}

	// Read contents of file into a string array then sort that array
	fileContents, _ := readLines(fileName)
	sort.Strings(fileContents)

	for i := 0; i < len(fileContents); i++ {
		// Loop through data
		//   read line
		//   if guard then grab id and move on
		//   if falls asleep then grab start time and date and move on
		//   if wakes up then process properly with data gathered
		var tempHour, tempMinute int

		fmt.Sscanf(fileContents[i], "[1518-%s %d:%d] %s %s", &dateVar, &tempHour, &tempMinute, &actionVar, &extraVar)

		switch actionVar {
		case "Guard":
			// Grab ID and move on
			currentGuard = extraVar
		case "falls":
			sleepStartTime = tempMinute
		case "wakes":
			wakeUpTime = tempMinute
			wakeUpDate = dateVar
	
			if storagePoint > 0 {
				if guardMap[storagePoint - 1].date == wakeUpDate {
					storagePoint--
				}
			}

			guardMap[storagePoint].date = wakeUpDate
			guardMap[storagePoint].guardID = currentGuard
			for j := sleepStartTime; j < wakeUpTime; j++ {
				guardMap[storagePoint].minute[j] = '#'
			}

			storagePoint++
		}
	}

	printGuardMap(guardMap)

	if part == 'a' {
		// get the numbers, convert the guard id to int, and return the answer
		chosenGuard, chosenMinute = mostSleepyGuard(guardMap)
	} else {
		chosenGuard, chosenMinute = mostSleepyMinute(guardMap)
	}
	tmpChosenGuard1 := strings.Split(chosenGuard, "#")
	tmpChosenGuard2, _ := strconv.Atoi(tmpChosenGuard1[1])
	fmt.Printf ("Chosen Guard: %d Chosen Minute: %d\n", tmpChosenGuard2, chosenMinute)

	return tmpChosenGuard2 * chosenMinute
}

// Main routine
func main() {
	fileNamePtr := flag.String("file", "input1.txt", "A filename containing input strings")
	execPartPtr := flag.String("part", "a", "Which part of day04 do you want to calc (a or b)")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Guard ID multiplied by minute chosen:", sleepyGuard(*fileNamePtr, 'a'))
	case "b":
		fmt.Println("Part b - Guard ID multiplied by minute chosen:", sleepyGuard(*fileNamePtr, 'b'))
	}
}