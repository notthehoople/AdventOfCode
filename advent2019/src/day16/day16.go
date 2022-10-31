package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
)

// sourceSignal - the signal list to work on
// arrayPos - the position in the array we're calculating a value for
// debug - whether debug is on or not
//
// While each element in the output array uses all of the same input array elements, the actual repeating pattern to use depends on which output element is being calculated.
// The base pattern is 0, 1, 0, -1. Then, repeat each value in the pattern a number of times equal to the position in the output list being considered. Repeat once for the
// first element, twice for the second element, three times for the third element, and so on. So, if the third element of the output list is being calculated, repeating the
// values would produce: 0, 0, 0, 1, 1, 1, 0, 0, 0, -1, -1, -1.
func calcSignalPatternVal(sourceSignal []byte, pattern []int, arrayPos int, debug bool) int {
	var patternPos, tempLoopTotal, patternRepeat int

	patternPos = 0               // start at the beginning
	patternRepeat = arrayPos - 1 // counts the number of repeat multiplications we do per character. We start with one less than the position of the digit we're working with
	if patternRepeat == 0 {      // Special case for first digit
		patternPos++
		patternRepeat = arrayPos
	}

	// Loop through the sourceSignal and the pattern list
	for signalLoop := 0; signalLoop < len(sourceSignal); signalLoop++ {
		if debug {
			fmt.Printf("Signal digit: %c ArrayPos: %d Pattern: %d patternRepeat: %d\n", sourceSignal[signalLoop], arrayPos, pattern[patternPos], patternRepeat)
		}

		// We need to apply the pattern digit the same number of times as the position in the array (arrayPos)
		// We also skip the FIRST digit in the pattern (e.g. for the second digit it should be 0,0,1,1,0,0,-1,-1 but instead it's 0,1,1,0,0,-1,-1 then we loop to start again)

		//   Covert the signal byte to an int we can work with
		tempSignalValue, _ := strconv.Atoi(string(sourceSignal[signalLoop]))

		// ******* Need to rework the repeat here as I'm getting it wrong
		// ******* The repeat applies to the pattern used, not to the digit itself
		// ******* So when we're looping through the sourceSignal, we decide with pattern value to use based on the number of times we've already used that value
		// ******* No need for a sub loop

		//   Multiply signal int and pattern
		tempLoopTotal = tempLoopTotal + (tempSignalValue * pattern[patternPos])

		patternRepeat--
		if patternRepeat == 0 { // We're done with this pattern digit so move on
			patternRepeat = arrayPos
			patternPos++
			if patternPos >= len(pattern) { // If we've got to the end of the pattern list, loop
				patternPos = 0
			}
		}
	}

	// Return the numbers digit (% 10)
	return int(math.Abs(float64(tempLoopTotal % 10)))
}

// Flawed Frequency Transmission algorithm (FFT)
func applyFFT(sourceSignal []byte, stepsToRun int, debug bool, part byte) string {
	var loopTotalValue int
	var patternValue int = 1
	var newString []byte
	var pattern = []int{0, 1, 0, -1}

	newString = make([]byte, len(sourceSignal))

	for phase := 0; phase < stepsToRun; phase++ {
		loopTotalValue = 0

		// Need something here to reset the pattern to apply to original state

		for signalLoop := 0; signalLoop < len(sourceSignal); signalLoop++ {
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

func main() {
	var debug bool

	execPartPtr := flag.String("part", "a", "Which part of day16 do you want to calc (a or b)")
	flag.BoolVar(&debug, "debug", false, "Turn debug on")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - After FFT output list: (test 1):", applyFFT([]byte("12345678"), 4, debug, 'a'))
		fmt.Println("Part a - After FFT output list: (test 2):", applyFFT([]byte("80871224585914546619083218645595"), 100, debug, 'a'))
		fmt.Println("Part a - After FFT output list: (test 3):", applyFFT([]byte("19617804207202209144916044189917"), 100, debug, 'a'))
		fmt.Println("Part a - After FFT output list: (test 4):", applyFFT([]byte("69317163492948606335995924319873"), 100, debug, 'a'))
		fmt.Println("Part a - After FFT output list: (production):", applyFFT([]byte("59756772370948995765943195844952640015210703313486295362653878290009098923609769261473534009395188480864325959786470084762607666312503091505466258796062230652769633818282653497853018108281567627899722548602257463608530331299936274116326038606007040084159138769832784921878333830514041948066594667152593945159170816779820264758715101494739244533095696039336070510975612190417391067896410262310835830006544632083421447385542256916141256383813360662952845638955872442636455511906111157861890394133454959320174572270568292972621253460895625862616228998147301670850340831993043617316938748361984714845874270986989103792418940945322846146634931990046966552"), 100, debug, 'a'))
	case "b":
		fmt.Println("Part a - After FFT output list: (test 2):", applyFFTpartb([]byte("80871224585914546619083218645595"), 1, debug, 'b'))
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
