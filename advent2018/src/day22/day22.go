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

func calcRiskLevel(fileName string, part string, debug bool) int {
	var riskLevel = 0
	var caveDepth, xTarget, yTarget int = 0, 0, 0
	var maxXGrid, maxYGrid int = 0, 0
	
	// Read in the file contents
	fileContents, _ := readLines(fileName)
	_, _ = fmt.Sscanf(fileContents[0], "depth: %d", &caveDepth)
	_, _ = fmt.Sscanf(fileContents[1], "target: %d,%d", &xTarget, &yTarget,)

	maxXGrid = xTarget + 5
	maxYGrid = yTarget + 5

	erosionLevel := make([][]int, maxXGrid)
    for i := 0; i < maxXGrid; i++ {
        erosionLevel[i] = make([]int, maxYGrid)
    }

	fmt.Printf("maxXGrid: %d maxYGrid: %d\n", maxXGrid, maxYGrid)

	if part == "a" {
		// Let's do part a
		if debug {
			fmt.Println("Depth:", caveDepth)
			fmt.Println("X and Y:", xTarget, yTarget)
		}

		for y := 0; y < maxYGrid; y++ {
			for x := 0; x < maxXGrid; x++ {

				if debug {
					fmt.Printf("x: %d y: %d\n", x, y)
				}

				if y == 0 {
					// If the region's Y coordinate is 0, the geologic index is its X coordinate times 16807.
					erosionLevel[x][y] = ((x * 16807) + caveDepth) % 20183
					if debug {
						fmt.Printf("y is zero and x: %d y: %d erosionLevel: %d\n", x, y, erosionLevel[x][y])
					}

				} else {
					if (x == 0 && y == 0) || (x == xTarget && y == yTarget) {
						// The region at 0,0 (the mouth of the cave) has a geologic index of 0.
						// The region at the coordinates of the target has a geologic index of 0.

						erosionLevel[x][y] = (0 + caveDepth) % 20183
						if debug {
							fmt.Printf("x and y are zero and x: %d y: %d erosionLevel: %d\n", x, y, erosionLevel[x][y])
						}

					} else {
						if x == 0 {
							// If the region's X coordinate is 0, the geologic index is its Y coordinate times 48271.
							erosionLevel[x][y] = ((y * 48271) + caveDepth) % 20183
							if debug {
								fmt.Printf("x is zero and x: %d y: %d erosionLevel: %d\n", x, y, erosionLevel[x][y])
							}

						} else {
							// The region's geologic index = multiplying the erosion levels of the regions at X-1,Y and X,Y-1.
							//fmt.Println("erosionLevel x - 1:", erosionLevel[x-1][y])
							//fmt.Println("erosionLevel y - 1:", erosionLevel[x][y-1])

							erosionLevel[x][y] = ((erosionLevel[x-1][y] * erosionLevel[x][y-1]) + caveDepth) % 20183
							if debug {
								fmt.Printf("x is zero and x: %d y: %d erosionLevel: %d\n", x, y, erosionLevel[x][y])
							}
						}
					}
				}
			}
		}

		// Print and Calulate the RiskLevel
		for y := 0; y <= yTarget; y++ {
			for x := 0; x <= xTarget; x++ {
				switch erosionLevel[x][y] % 3 {
				case 0: fmt.Printf(".")
						riskLevel += 0
				case 1: fmt.Printf("=")
						riskLevel += 1
				case 2: fmt.Printf("|")
						riskLevel += 2
				}
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")

				
	} else {
		// part b is where it's at
	}

	return riskLevel
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