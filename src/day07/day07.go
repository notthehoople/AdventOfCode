package main

import (
	"fmt"
	"sort"
	"os"
	"bufio"
	"flag"
	"strings"
)

// Take the first item out of the list of strings passed in (toDo)
// Returns 3:
//    - string: the item removed from the List
//    - []string: the list of strings with the item removed
//    - bool: true/false on whether an item was available for removal or not
func popTopItem(toDo []string)(string, []string, bool) {
	var popItem string

	if len(toDo) == 0 {
		return "", toDo, false
	}

	popItem = toDo[0]
	toDo = removeItem(toDo, toDo[0])
	return popItem, toDo, true
}

func addItemSorted(toDo []string, itemAdd string) ([]string) {
	var foundIt bool = false

	for i := 0; i < len(toDo) && !foundIt; i++ {
		if toDo[i] == itemAdd {
			foundIt = true
		}	
	}
	if !foundIt {
		toDo = append(toDo, itemAdd)
		sort.Strings(toDo)
	}

	return toDo
}

func removeItem(toDo []string, itemRemove string) ([]string) {
	var toDoReplace [] string
	
	for i := 0; i < len(toDo); i++ {
		if toDo[i] != itemRemove {
			toDoReplace = addItemSorted(toDoReplace, toDo[i])
		}
	}
	return toDoReplace
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

// Loop through toDoList looking for workitem
// Remove workitem from all toDoList preRegs
// Any outcomeSteps with no preReqSteps left are put onto the reducedNoPreReqList
func workItemCompleted(workItem string, toDoList map[string]string, reducedNoPreReqList []string) (map[string]string, []string) {

	for k, _ := range toDoList {
		toDoList[k] = strings.Replace(toDoList[k], workItem, "", -1)

		if toDoList[k] == "" {
			// We've removed the last preReq so can add this to the reducedNoPreReqList
			delete(toDoList, k)
			reducedNoPreReqList = addItemSorted(reducedNoPreReqList, k)
		}
	}

	return toDoList, reducedNoPreReqList
}

func checkWorkList(noPreReqList []string, toDoList map[string]string) []string {
	var reducedNoPreReqList []string

	for i := 0; i < len(noPreReqList); i++ {
		if _, ok := toDoList[noPreReqList[i]]; !ok {
			reducedNoPreReqList = addItemSorted(reducedNoPreReqList, noPreReqList[i])
		}
	}

	return reducedNoPreReqList
}

func stepOrder(fileName string, part string) string {
	var noPreReqList []string
	var reducedNoPreReqList []string
	var toDoList = make(map[string]string)
	var solutionOrder string

	// Read contents of file into a string array then sort that array
	fileContents, _ := readLines(fileName)

	for i := 0; i < len(fileContents); i++ {
		var preReqStep string
		var outcomeStep string
		fmt.Sscanf(fileContents[i], "Step %s must be finished before step %s can begin.", &preReqStep, &outcomeStep)
		toDoList[outcomeStep] += preReqStep

		if toDoList[preReqStep] == "" {
			noPreReqList = addItemSorted(noPreReqList, preReqStep)
		}
	}

	// Double check the Start List (reducedNoPreReqList) is correct
	reducedNoPreReqList = checkWorkList(noPreReqList, toDoList)

	var workToDo bool = true
	var workItem string
	var ok bool

	for workToDo {
		workItem, reducedNoPreReqList, ok = popTopItem(reducedNoPreReqList)
		if !ok {
			workToDo = false
			continue
		} else {
			solutionOrder += workItem

			// Remove completed item from toDoList
			toDoList, reducedNoPreReqList = workItemCompleted(workItem, toDoList, reducedNoPreReqList)
		}
	}

	return solutionOrder
}

// Main routine
func main() {
	fileNamePtr := flag.String("file", "input1.txt", "A filename containing input strings")
	execPartPtr := flag.String("part", "a", "Which part of day07 do you want to calc (a or b)")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Order of steps:", stepOrder(*fileNamePtr, "a"))
	case "b":
		fmt.Println("Part b - Not there yet")
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}