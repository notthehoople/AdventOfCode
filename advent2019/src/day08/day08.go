package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
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

// Splits a string into arrays based on the size "n"
func splitSubN(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
}

// Returns: Number of "1" digits multiplied by the number of "2" digits on the layer with the fewest "0" digits
func processImage(filename string, width int, height int, debug bool, part byte) int {
	var lowestZero, lowestOne, lowestTwo int = 50000, 50000, 50000
	var currentZero int

	// Read contents of file into a string array
	fileContents, _ := readLines(filename)

	// The file should be a single line of image data. Break this into layers of width*height
	imageLayers := splitSubN(fileContents[0], width*height)
	for i, subString := range imageLayers {
		currentZero = strings.Count(subString, "0")
		if currentZero < lowestZero {
			lowestZero = currentZero
			lowestOne = strings.Count(subString, "1")
			lowestTwo = strings.Count(subString, "2")
		}

		if debug {
			fmt.Printf("layer: %d number of 0: %d number of 1: %d number of 2: %d\n", i, strings.Count(subString, "0"), strings.Count(subString, "1"), strings.Count(subString, "2"))
		}
	}

	if part == 'b' {
		var decodedImage []byte

		// decoded image = last layer
		// while from the back to the front
		//		loop through the pixels in the layer
		//			switch pixel:
		//				case 0: decoded image pixel = black
		//				case 1: decoded image pixel = white
		//				case 2: decoded image pixel = do nothing

		decodedImage = []byte(imageLayers[len(imageLayers)-1])

		for i := len(imageLayers) - 1; i >= 0; i-- {
			for letterPos, letterValue := range imageLayers[i] {
				switch letterValue {
				case '0': // black
					decodedImage[letterPos] = '0'
				case '1': // white
					decodedImage[letterPos] = '1'
				case '2': // transparanet
				default: // corrupt file
				}
			}
		}
		if debug {
			fmt.Println("Decoded Image:", string(decodedImage))
		}

		displayMessage := splitSubN(string(decodedImage), width)
		for _, subString := range displayMessage {
			fmt.Println(subString)
		}

	}

	if part == 'a' {
		return lowestOne * lowestTwo
	}

	return 0
}

// Main routine
func main() {
	var debug bool

	filenamePtr := flag.String("file", "input.txt", "Filename containing the program to run")
	widthPtr := flag.Int("width", 25, "Width of the image in pixels")
	heightPtr := flag.Int("height", 6, "Height of the image in pixels")
	execPartPtr := flag.String("part", "a", "Which part of day03 do you want to calc (a or b)")
	flag.BoolVar(&debug, "debug", false, "Turn debug on")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - 1 * 2 pixels on least currupt layer:", processImage(*filenamePtr, *widthPtr, *heightPtr, debug, 'a'))
	case "b":
		fmt.Println("Answer is in the above matrix", processImage(*filenamePtr, *widthPtr, *heightPtr, debug, 'b'))

	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
