package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"sort"
	"strconv"
)

//go:embed input.txt
var embedFs embed.FS
var inputFile = "input.txt"

func main() {

	fmt.Println(fmt.Sprintf("Part 1: %d", part1()))
	fmt.Println(fmt.Sprintf("Part 2: %d", part2()))
}

func part1() int {
	fileToRead, readError := embedFs.Open(inputFile)

	if readError != nil {
		log.Fatal(readError)
	}

	defer fileToRead.Close()

	scanner := bufio.NewScanner(fileToRead)

	var biggest int
	var current int

	for scanner.Scan() {
		if scanner.Text() == "" {
			if current > biggest {
				biggest = current
			}
			current = 0
		} else {
			calorie, err := strconv.Atoi(scanner.Text())

			if err != nil {
				fmt.Println("Error during conversion")
				log.Fatal(err)
			}
			current += calorie
		}
	}

	return biggest
}

func part2() int {
	fileToRead, readError := embedFs.Open(inputFile)

	if readError != nil {
		log.Fatal(readError)
	}

	defer fileToRead.Close()

	scanner := bufio.NewScanner(fileToRead)

	var sortedCounts []int
	var current int

	for scanner.Scan() {
		if scanner.Text() == "" {
			i := sort.SearchInts(sortedCounts, current)
			sortedCounts = append(sortedCounts, 0)     // make room
			copy(sortedCounts[i+1:], sortedCounts[i:]) // move elements down
			sortedCounts[i] = current
			current = 0
		} else {
			calorie, err := strconv.Atoi(scanner.Text())

			if err != nil {
				fmt.Println("Error during conversion")
				log.Fatal(err)
			}
			current += calorie
		}
	}

	return sortedCounts[len(sortedCounts)-1] + sortedCounts[len(sortedCounts)-2] + sortedCounts[len(sortedCounts)-3]
}
