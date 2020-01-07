package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
)

// Flawed Frequency Transmission algorithm (FFT)
func applyFFT(sourceSignal []byte, stepsToRun int, debug bool, part byte) string {
	var loopTotalValue int
	var patternValue int = 1
	var newString []byte
	/*
	   As input, FFT takes a list of numbers. In the signal you received (your puzzle input), each number is a single digit: data like 15243 represents the sequence 1, 5, 2, 4, 3.
	   FFT operates in repeated phases. In each phase, a new list is constructed with the same length as the input list. This new list is also used as the input for the next phase.

	   Each element in the new list is built by multiplying every value in the input list by a value in a repeating pattern and then adding up the results.
	   So, if the input list were 9, 8, 7, 6, 5 and the pattern for a given element were 1, 2, 3, the result would be 9*1 + 8*2 + 7*3 + 6*1 + 5*2
	   (with each input element on the left and each value in the repeating pattern on the right of each multiplication). Then, only the "ones" digit is kept: 38 becomes 8,
	   -17 becomes 7, and so on.

	   While each element in the output array uses all of the same input array elements, the actual repeating pattern to use depends on which output element is being calculated.
	   The base pattern is 0, 1, 0, -1. Then, repeat each value in the pattern a number of times equal to the position in the output list being considered. Repeat once for the
	   first element, twice for the second element, three times for the third element, and so on. So, if the third element of the output list is being calculated, repeating the
	   values would produce: 0, 0, 0, 1, 1, 1, 0, 0, 0, -1, -1, -1.

	   When applying the pattern, skip the very first value exactly once. (In other words, offset the whole pattern left by one.) So, for the second element of the output list,
	   the actual pattern used would be: 0, 1, 1, 0, 0, -1, -1, 0, 0, 1, 1, 0, 0, -1, -1, ....

	   After using this process to calculate each element of the output list, the phase is complete, and the output list of this phase is used as the new input list for the next phase,
	   if any.

	   Given the input signal 12345678, below are four phases of FFT. Within each phase, each output digit is calculated on a single line with the result at the far right;
	   each multiplication operation shows the input digit on the left and the pattern value on the right:

	   Input signal: 12345678

	   1*1  + 2*0  + 3*-1 + 4*0  + 5*1  + 6*0  + 7*-1 + 8*0  = 4
	   1*0  + 2*1  + 3*1  + 4*0  + 5*0  + 6*-1 + 7*-1 + 8*0  = 8
	   1*0  + 2*0  + 3*1  + 4*1  + 5*1  + 6*0  + 7*0  + 8*0  = 2
	   1*0  + 2*0  + 3*0  + 4*1  + 5*1  + 6*1  + 7*1  + 8*0  = 2
	   1*0  + 2*0  + 3*0  + 4*0  + 5*1  + 6*1  + 7*1  + 8*1  = 6
	   1*0  + 2*0  + 3*0  + 4*0  + 5*0  + 6*1  + 7*1  + 8*1  = 1
	   1*0  + 2*0  + 3*0  + 4*0  + 5*0  + 6*0  + 7*1  + 8*1  = 5
	   1*0  + 2*0  + 3*0  + 4*0  + 5*0  + 6*0  + 7*0  + 8*1  = 8

	   After 1 phase: 48226158

	   4*1  + 8*0  + 2*-1 + 2*0  + 6*1  + 1*0  + 5*-1 + 8*0  = 3
	   4*0  + 8*1  + 2*1  + 2*0  + 6*0  + 1*-1 + 5*-1 + 8*0  = 4
	   4*0  + 8*0  + 2*1  + 2*1  + 6*1  + 1*0  + 5*0  + 8*0  = 0
	   4*0  + 8*0  + 2*0  + 2*1  + 6*1  + 1*1  + 5*1  + 8*0  = 4
	   4*0  + 8*0  + 2*0  + 2*0  + 6*1  + 1*1  + 5*1  + 8*1  = 0
	   4*0  + 8*0  + 2*0  + 2*0  + 6*0  + 1*1  + 5*1  + 8*1  = 4
	   4*0  + 8*0  + 2*0  + 2*0  + 6*0  + 1*0  + 5*1  + 8*1  = 3
	   4*0  + 8*0  + 2*0  + 2*0  + 6*0  + 1*0  + 5*0  + 8*1  = 8

	   After 2 phases: 34040438

	   3*1  + 4*0  + 0*-1 + 4*0  + 0*1  + 4*0  + 3*-1 + 8*0  = 0
	   3*0  + 4*1  + 0*1  + 4*0  + 0*0  + 4*-1 + 3*-1 + 8*0  = 3
	   3*0  + 4*0  + 0*1  + 4*1  + 0*1  + 4*0  + 3*0  + 8*0  = 4
	   3*0  + 4*0  + 0*0  + 4*1  + 0*1  + 4*1  + 3*1  + 8*0  = 1
	   3*0  + 4*0  + 0*0  + 4*0  + 0*1  + 4*1  + 3*1  + 8*1  = 5
	   3*0  + 4*0  + 0*0  + 4*0  + 0*0  + 4*1  + 3*1  + 8*1  = 5
	   3*0  + 4*0  + 0*0  + 4*0  + 0*0  + 4*0  + 3*1  + 8*1  = 1
	   3*0  + 4*0  + 0*0  + 4*0  + 0*0  + 4*0  + 3*0  + 8*1  = 8

	   After 3 phases: 03415518

	   0*1  + 3*0  + 4*-1 + 1*0  + 5*1  + 5*0  + 1*-1 + 8*0  = 0
	   0*0  + 3*1  + 4*1  + 1*0  + 5*0  + 5*-1 + 1*-1 + 8*0  = 1
	   0*0  + 3*0  + 4*1  + 1*1  + 5*1  + 5*0  + 1*0  + 8*0  = 0
	   0*0  + 3*0  + 4*0  + 1*1  + 5*1  + 5*1  + 1*1  + 8*0  = 2
	   0*0  + 3*0  + 4*0  + 1*0  + 5*1  + 5*1  + 1*1  + 8*1  = 9
	   0*0  + 3*0  + 4*0  + 1*0  + 5*0  + 5*1  + 1*1  + 8*1  = 4
	   0*0  + 3*0  + 4*0  + 1*0  + 5*0  + 5*0  + 1*1  + 8*1  = 9
	   0*0  + 3*0  + 4*0  + 1*0  + 5*0  + 5*0  + 1*0  + 8*1  = 8

	   After 4 phases: 01029498
	   Here are the first eight digits of the final output list after 100 phases for some larger inputs:

	   80871224585914546619083218645595 becomes 24176176.
	   19617804207202209144916044189917 becomes 73745418.
	   69317163492948606335995924319873 becomes 52432133.
	   After 100 phases of FFT, what are the first eight digits in the final output list?
	*/

	for phase := 0; phase < stepsToRun; phase++ {
		loopTotalValue = 0

		// Need something here to reset the pattern to apply to original state

		for signalLoop := 0; signalLoop < len(sourceSignal); signalLoop++ {
			// Need to work out the pattern to apply as we go

			tempSignalValue, _ := strconv.Atoi(string(sourceSignal[signalLoop]))
			loopTotalValue += tempSignalValue * patternValue // Thing I haven't worked out yet
			fmt.Printf("Source Signal char: %c tempSignalValue: %d loopTotalValue: %d\n", sourceSignal[signalLoop], tempSignalValue, loopTotalValue)

			// Set the corresponding digit of the newString to the calculation we've done
			newString[signalLoop] = strconv.Itoa(math.Abs(loopTotalValue % 10))
		}

		fmt.Printf("Old string: %s\n", sourceSignal)
		fmt.Printf("New string: %s\n", newString)
		sourceSignal = newString
		// Completed a loop. We now need to build the replacement string of numbers for the next phase
	}

	return "Empty result"

}

func main() {
	var debug bool
	var steps int

	//filenamePtr := flag.String("file", "input.txt", "Filename containing the program to run")
	execPartPtr := flag.String("part", "a", "Which part of day16 do you want to calc (a or b)")
	flag.BoolVar(&debug, "debug", false, "Turn debug on")
	flag.IntVar(&steps, "steps", 10, "Number of steps to run")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - After FFT output list: (test 1):", applyFFT("12345678", 100, debug, 'a'))

		//fmt.Println("Part a - After FFT output list: (test 1):", applyFFT("test-a1.txt", 100, debug, 'a'))
		//fmt.Println("Part a - After FFT output list: (test 2):", applyFFT("test-a2.txt", 100, debug, 'a'))
		//fmt.Println("Part a - After FFT output list: (test 3):", applyFFT("test-a3.txt", 100, debug, 'a'))
		//fmt.Println("Part a - After FFT output list: (test 4):", applyFFT("test-a4.txt", 100, debug, 'a'))
	case "b":
		fmt.Println("Part b - Not implemented yet")
		//fmt.Println("Part b - First repeat (test 1):", findFirstRepeat("test-a1.txt", debug, 'b'))
		//fmt.Println("Part b - First repeat (test 2):", findFirstRepeat("test-a2.txt", debug, 'b'))
		//fmt.Println("Part b - First repeat (test 2):", findFirstRepeat("input.txt", debug, 'b'))

	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
