package main

import (
	"AdventOfCode-go/advent2024/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var Debug bool

type GateStruct struct {
	operation  string
	inputWire1 string
	inputWire2 string
	outputWire string
}

func canOperate(gateLine GateStruct, wires map[string]int) bool {
	if _, ok1 := wires[gateLine.inputWire1]; ok1 {
		if _, ok2 := wires[gateLine.inputWire2]; ok2 {
			return true
		}
	}
	return false
}

func combineWires(gateLine GateStruct, wires map[string]int) int {
	switch gateLine.operation {
	case "AND":
		if (wires[gateLine.inputWire1] == 1) && (wires[gateLine.inputWire2] == 1) {
			return 1
		}
		return 0

	case "OR":
		if (wires[gateLine.inputWire1] == 1) || (wires[gateLine.inputWire2] == 1) {
			return 1
		}
		return 0

	case "XOR":
		if wires[gateLine.inputWire1] != wires[gateLine.inputWire2] {
			return 1
		}
		return 0
	}
	return 0
}

func day24(filename string, part byte) int {
	var result int

	//Puzzle Input:
	//- line x: Start value for an x or y wire <x|y><wire-num (2 digits)>: <value (1 digit)>
	//- line y: <blank line<
	//- line Y+1 onwards: wire combinations and destination
	//       e.g. tnw OR fst -> frj
	// 3 different operators: AND, OR, XOR

	wires := make(map[string]int)        // Hold the current value of calculated wires
	gates := make(map[string]GateStruct) // Hold the instructions for calculating wires

	puzzleInput, _ := utils.ReadFile(filename)

	// Process the input file
	var startValueSection bool = true
	for _, i := range puzzleInput {
		if len(i) == 0 {
			startValueSection = false
			continue
		}
		if startValueSection {
			// Start values for wires

			tempStrings := strings.Split(i, ": ")
			tempWire := tempStrings[0]
			tempValue, _ := strconv.Atoi(strings.TrimSpace(tempStrings[1]))

			wires[tempWire] = tempValue
		} else {
			// wire instructions section

			tempInstructions := strings.Split(i, " -> ")
			tempLeft := strings.Split(tempInstructions[0], " ")

			gates[tempInstructions[1]] = GateStruct{operation: tempLeft[1],
				inputWire1: tempLeft[0],
				inputWire2: tempLeft[2]}
		}
	}

	if Debug {
		fmt.Println(wires)
		fmt.Println(gates)
	}

	// How do I decide if a gate instruction has already been used? Bool in the struct?
	// Or delete the entry from the map?

	for len(gates) > 0 {
		for currWire, gateLine := range gates {
			if canOperate(gateLine, wires) {
				wires[currWire] = combineWires(gateLine, wires)
				delete(gates, currWire)
			}
		}
	}

	// Now need to collect ONLY the Zxx wires for our result
	resultWires := make([]string, 0)
	for wire := range wires {
		if wire[0] == 'z' {
			resultWires = append(resultWires, wire)
		}
	}
	sort.Strings(resultWires)
	if Debug {
		fmt.Println(resultWires)
	}

	for i := len(resultWires) - 1; i >= 0; i-- {
		result = (result << 1) + wires[resultWires[i]]
	}

	return result
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()
	Debug = debug

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", day24(filenamePtr, execPart))
	}
}
