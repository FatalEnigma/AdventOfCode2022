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

var draw = map[string]string{"A": "X", "B": "Y", "C": "Z"}
var win = map[string]string{"A": "Y", "B": "Z", "C": "X"}
var points = map[string]int{"X": 1, "Y": 2, "Z": 3}

type Key struct {
	opp, outcome string
}

var shouldPlay = map[Key]string{
	Key{"A", "X"}: "Z",
	Key{"A", "Y"}: "X",
	Key{"A", "Z"}: "Y",
	Key{"B", "X"}: "X",
	Key{"B", "Y"}: "Y",
	Key{"B", "Z"}: "Z",
	Key{"C", "X"}: "Y",
	Key{"C", "Y"}: "Z",
	Key{"C", "Z"}: "X",
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d", part2())
}

func part1() int {
	fileToRead, readError := embedFs.Open(inputFile)

	if readError != nil {
		log.Fatal(readError)
	}

	defer fileToRead.Close()

	scanner := bufio.NewScanner(fileToRead)

	var total int

	for scanner.Scan() {
		row := scanner.Text()
		opp := string(row[0])
		you := string(row[2])

		total += roundPoints(opp, you)
	}

	return total
}

func roundPoints(opp string, you string) int {

	if draw[opp] == you {
		return 3 + points[you]
	}

	if win[opp] == you {
		return 6 + points[you]
	}

	return points[you]
}

func part2() int {
	fileToRead, readError := embedFs.Open(inputFile)

	if readError != nil {
		log.Fatal(readError)
	}

	defer fileToRead.Close()

	scanner := bufio.NewScanner(fileToRead)

	var total int

	for scanner.Scan() {
		row := scanner.Text()
		opp := string(row[0])
		you := shouldPlay[Key{opp, string(row[2])}]

		total += roundPoints(opp, you)
	}

	return total
}
