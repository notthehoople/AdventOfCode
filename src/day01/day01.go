package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"flag"
	"strconv"
)

func addAll(fileName string) int {
	var resultVar int = 0		// Defining the overall result Variable
	var tempString string					// Holds the line read from the file

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// fmt.Println("Lets read the file")

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		tempString = scanner.Text()
		tempVal, _ := strconv.Atoi(tempString)
		resultVar += tempVal
	//	fmt.Println("TempString:", tempString)
	//	fmt.Println("TempVal:", tempVal)
	//	fmt.Println("ResultVar:", resultVar)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return resultVar
}

func main() {
	fileNamePtr := flag.String("file", "input1.txt", "A filename containing input numbers")

	flag.Parse()
	
	// fmt.Println("filename:", *fileNamePtr)
	// fmt.Println("tail:", flag.Args())

	fmt.Println("Resulting Frequency:", addAll(*fileNamePtr))

	// fmt.Println(addAll("input.txt"))
}