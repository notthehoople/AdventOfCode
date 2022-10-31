package main

import (
	"fmt"
	"math"
	"strconv"
)

// Flawed Frequency Transmission algorithm (FFT)
func applyFFTpartb(inputSignal []byte, stepsToRun int, debug bool, part byte) string {
	var loopTotalValue int
	var patternValue int = 1
	var newString []byte
	var sourceSignal []byte
	var pattern = []int{0, 1, 0, -1}

	fmt.Println("Dumb make starting...")
	sourceSignal = make([]byte, len(inputSignal)*10000)
	fmt.Println("Dumb make done")
	// Build the source array, which is 10000 times the inputSignal
	for i := 0; i < 10000; i++ {
		for j := 0; j < len(inputSignal); j++ {
			sourceSignal[i+j] = inputSignal[j]
		}
	}
	fmt.Println("Dumb big string built")

	newString = make([]byte, len(sourceSignal))

	for phase := 0; phase < stepsToRun; phase++ {
		fmt.Println("Phase:", phase)
		loopTotalValue = 0

		// Need something here to reset the pattern to apply to original state

		for signalLoop := 0; signalLoop < len(sourceSignal); signalLoop++ {
			//fmt.Println("signalLoop:", signalLoop)
			// Need to work out the pattern to apply as we go

			tempSignalValue, _ := strconv.Atoi(string(sourceSignal[signalLoop]))
			loopTotalValue += tempSignalValue * patternValue // Thing I haven't worked out yet

			loopTotalValue = calcSignalPatternVal(sourceSignal, pattern, signalLoop+1, debug)
			if debug {
				fmt.Printf("Source Signal char: %c tempSignalValue: %d loopTotalValue: %d\n", sourceSignal[signalLoop], tempSignalValue, loopTotalValue)
			}

			// Set the corresponding digit of the newString to the calculation we've done
			tempString := strconv.Itoa(int(math.Abs(float64(loopTotalValue % 10))))

			newString[signalLoop] = tempString[0]
		}

		if debug {
			fmt.Printf("Old string: %s\n", sourceSignal)
			fmt.Printf("New string: %s\n", newString)
		}
		// Use the newly built string as the input for the next stage
		copy(sourceSignal, newString)
	}

	return string(newString[0:8])

}
