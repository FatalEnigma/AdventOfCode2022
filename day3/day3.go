package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"unicode"
)

//go:embed input.txt
var embedFs embed.FS
var inputFile = "input.txt"

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d", part2())
}

func getPriority(char rune) int {
	if unicode.IsUpper(char) {
		return int(char) - 38
	} else {
		return int(char) - 96
	}
}

func part1() int {
	fileToRead, readError := embedFs.Open(inputFile)

	if readError != nil {
		log.Fatal(readError)
	}

	defer fileToRead.Close()

	scanner := bufio.NewScanner(fileToRead)

	var prioritySum int

	for scanner.Scan() {
		row := scanner.Text()
		rucksack1 := row[0 : len(row)/2]
		rucksack2 := row[len(row)/2:]

		charMap := map[rune]bool{}

		for _, char := range rucksack1 {
			charMap[char] = true
		}

		for _, char := range rucksack2 {
			if charMap[char] {
				prioritySum += getPriority(char)
				break
			}
		}

	}

	return prioritySum
}

func part2() int {
	fileToRead, readError := embedFs.Open(inputFile)

	if readError != nil {
		log.Fatal(readError)
	}

	defer fileToRead.Close()

	scanner := bufio.NewScanner(fileToRead)

	var prioritySum int

	charMap1 := map[rune]bool{}
	charMap2 := map[rune]bool{}

	for scanner.Scan() {
		row := scanner.Text()

		if len(charMap1) == 0 {
			for _, char := range row {
				charMap1[char] = true
			}
		} else if len(charMap2) == 0 {
			for _, char := range row {
				if charMap1[char] {
					charMap2[char] = true
				}
			}
		} else {
			for _, char := range row {
				if charMap2[char] {
					prioritySum += getPriority(char)
					break
				}
			}
			charMap1 = make(map[rune]bool)
			charMap2 = make(map[rune]bool)
		}
	}

	return prioritySum
}
