package main

import (
	"fmt"
	"sort"
	"os"
	"bufio"
	"flag"
	"strings"
)

// Each worker has a workItem and a time it will take to complete that workitem
type workerStruct struct {
    timeToComplete  int
    workItem        string
}

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

// Checks if an item already exists in the toDoList. If not, it adds it in there.
func checkWorkList(noPreReqList []string, toDoList map[string]string) []string {
	var reducedNoPreReqList []string

	for i := 0; i < len(noPreReqList); i++ {
		if _, ok := toDoList[noPreReqList[i]]; !ok {
			reducedNoPreReqList = addItemSorted(reducedNoPreReqList, noPreReqList[i])
		}
	}

	return reducedNoPreReqList
}

// Carries out the work for PartA
func goPartA(reducedNoPreReqList []string, toDoList map[string]string) string {
	var solutionOrder string
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

// Carries out the work for PartA
func goPartB(reducedNoPreReqList []string, toDoList map[string]string, timeConst int, numWorkers int) (string, int) {
	var solutionOrder string
	var workToDo bool = true
	var workInProgress int = 0
	var tempWorkItem string
	var ok bool
	var currentTime int = 0

	workers := make([]workerStruct, numWorkers)

	for i := 0; i < len(workers); i++ {
		fmt.Println("workers", i, workers[i])
		if workers[i].workItem == "" {
			fmt.Println("Worker Ready to go:", i)
		}
	}

	fmt.Println("Time Const is:", timeConst)
	fmt.Println("Second   Worker 1   Worker 2   Done")

	for workToDo || workInProgress > 0 {
		// Check if any work has completed
		//    If it has, then add the completed item to solutionOrder and return the Worker to being ready
		//    Any worker found that isn't busy, grab an item from the worklist reducedNoPreReqList (if any)
		//       and give to that worker
		for i := 0; i < len(workers); i++ {
			if workers[i].timeToComplete == currentTime && currentTime > 0 {
				// We have a finisher!
				solutionOrder += workers[i].workItem
				// Remove completed item from toDoList
				toDoList, reducedNoPreReqList = workItemCompleted(workers[i].workItem, toDoList, reducedNoPreReqList)
				// Clear the workItem from the worker
				workers[i].workItem = "-"
				workInProgress--
				// Avoid time counting upwards
			}
			// no else. we want a free worker to be picked up immediately

			if workers[i].workItem == "" || workers[i].workItem == "-" {
				// Worker ready and waiting for orders
				tempWorkItem, reducedNoPreReqList, ok = popTopItem(reducedNoPreReqList)
				if !ok {
					workToDo = false
					// continue
				} else {
					workers[i].workItem = tempWorkItem
					// Time is currentTime + timeConst + number relating to letter (1-26)
					workers[i].timeToComplete = currentTime + int(tempWorkItem[0]) + timeConst - 64
					//fmt.Println("Time to complete for:", tempWorkItem, workers[i].timeToComplete)
					workInProgress++
				}
			}
		}

		if len(workers) > 4 {
			fmt.Printf("%d, %s, %s, %s, %s, %s, %s, %d\n", currentTime, workers[0].workItem, workers[1].workItem, workers[2].workItem, workers[3].workItem, workers[4].workItem, solutionOrder, workInProgress)
		} else {
			if len(workers) == 2 {
				fmt.Printf("%d, %s, %s, %s, %d\n", currentTime, workers[0].workItem, workers[1].workItem, solutionOrder, workInProgress)
			} else {
				fmt.Printf("%d, %s, %s, %d\n", currentTime, workers[0].workItem, solutionOrder, workInProgress)
			}
		}

		currentTime++
	}
	
	return solutionOrder, currentTime - 2
}

func stepOrder(fileName string, part string, timeconst int, numWorkers int) (string, int) {
	var noPreReqList []string
	var reducedNoPreReqList []string
	var toDoList = make(map[string]string)
	var solutionOrder string
	var currentTime int = 0

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

	if part == "a" {
		solutionOrder = goPartA(reducedNoPreReqList, toDoList)
	} else {
		solutionOrder, currentTime = goPartB(reducedNoPreReqList, toDoList, timeconst, numWorkers)
	}

	return solutionOrder, currentTime
}

// Main routine
func main() {
	var solutionOrder string
	var currentTime int
	fileNamePtr := flag.String("file", "input1.txt", "A filename containing input strings")
	execPartPtr := flag.String("part", "a", "Which part of day07 do you want to calc (a or b)")
	timeConstantPtr := flag.Int("const", 0, "Time constant to add to each task for part b")
	numWorkersPtr := flag.Int("workers", 1, "Number of workers to use in part b")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		solutionOrder, currentTime = stepOrder(*fileNamePtr, "a", 0, 1)
		fmt.Println("Part a - Order of steps:", solutionOrder)
	case "b":
		if *numWorkersPtr < 1 {
			*numWorkersPtr = 1
		}
		solutionOrder, currentTime = stepOrder(*fileNamePtr, "b", *timeConstantPtr, *numWorkersPtr)
		fmt.Printf("Part b - Order of steps: %s Time: %d\n", solutionOrder, currentTime)
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}