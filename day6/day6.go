package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
)

//go:embed input.txt
var embedFs embed.FS
var inputFile = "input.txt"

func main() {
	inputString := parseInput()
	fmt.Printf("Part 1: %d\n", getMarker(inputString, 4))
	fmt.Printf("Part 2: %d", getMarker(inputString, 14))
}

func parseInput() (input string) {
	fileToRead, readError := embedFs.Open(inputFile)

	if readError != nil {
		log.Fatal(readError)
	}

	defer fileToRead.Close()

	scanner := bufio.NewScanner(fileToRead)

	var inputSlice []string

	for scanner.Scan() {
		inputSlice = append(inputSlice, scanner.Text())
	}

	return inputSlice[0]
}

func getMarker(inputString string, numUnique int) int {
	counter := map[string]int{}

	for x := 0; x <= numUnique-1; x++ {
		counter[string(inputString[x])] += 1
	}

	for x := 0; x < len(inputString); x++ {

		counter[string(inputString[x])] -= 1

		if counter[string(inputString[x])] == 0 {
			delete(counter, string(inputString[x]))
		}

		counter[string(inputString[x+numUnique])] += 1

		if len(counter) == numUnique {
			return x + numUnique + 1
		}
	}

	return 0
}
