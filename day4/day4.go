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
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d", part2())
}

func overlap(low1 int, up1 int, low2 int, up2 int) bool {
	return (low1 <= low2 && up1 >= up2) || (low1 >= low2 && up1 <= up2)
}

func anyOverlap(low1 int, up1 int, low2 int, up2 int) bool {
	return up1 >= low2 && low1 <= up2
}

func part1() int {
	fileToRead, readError := embedFs.Open(inputFile)

	if readError != nil {
		log.Fatal(readError)
	}

	defer fileToRead.Close()

	scanner := bufio.NewScanner(fileToRead)

	var overlapSum int

	for scanner.Scan() {
		var low1, up1, low2, up2 int

		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &low1, &up1, &low2, &up2)

		if overlap(low1, up1, low2, up2) {
			overlapSum += 1
		}
	}

	return overlapSum
}

func part2() int {
	fileToRead, readError := embedFs.Open(inputFile)

	if readError != nil {
		log.Fatal(readError)
	}

	defer fileToRead.Close()

	scanner := bufio.NewScanner(fileToRead)

	var overlapSum int

	for scanner.Scan() {
		var low1, up1, low2, up2 int

		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &low1, &up1, &low2, &up2)

		if anyOverlap(low1, up1, low2, up2) {
			overlapSum += 1
		}
	}

	return overlapSum
}
