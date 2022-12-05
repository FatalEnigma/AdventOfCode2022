package main

import (
	"AdventOfCode2022/types"
	"bufio"
	"embed"
	"fmt"
	"log"
	"strings"
)

//go:embed input.txt
var embedFs embed.FS
var inputFile = "input.txt"

func main() {
	parsedStacks, operations := parseInput()
	fmt.Printf("Part 1: %s\n", part1(copyParsedStacks(parsedStacks), operations))
	fmt.Printf("Part 2: %s", part2(copyParsedStacks(parsedStacks), operations))
}

func copyParsedStacks(parsedStacks [9]types.Stack) (stacksCopy [9]types.Stack) {
	for x := range parsedStacks {
		stacksCopy[x] = parsedStacks[x].Copy()
	}
	return stacksCopy
}

func parseInput() (parsedStacks [9]types.Stack, operations []string) {
	fileToRead, readError := embedFs.Open(inputFile)

	if readError != nil {
		log.Fatal(readError)
	}

	defer fileToRead.Close()

	scanner := bufio.NewScanner(fileToRead)

	doneParsing := false

	for scanner.Scan() {
		text := scanner.Text()

		if !doneParsing {
			if text == " 1   2   3   4   5   6   7   8   9" {
				continue
			}

			if text == "" {
				doneParsing = true

				for x := range parsedStacks {
					parsedStacks[x].Reverse()
				}

				continue
			}

			text = strings.ReplaceAll(strings.ReplaceAll(text, "]     ", "] [-] "), "    ", "[-] ")

			for x, s := range strings.Split(text, " ") {
				letter := s[1]
				if string(letter) != "-" {
					parsedStacks[x].Push(s[1])
				}
			}
		} else {
			operations = append(operations, text)
		}
	}
	return parsedStacks, operations
}

func part1(parsedStacks [9]types.Stack, operations []string) string {
	var sb strings.Builder

	for _, line := range operations {
		var numToMove, from, to int
		fmt.Sscanf(line, "move %d from %d to %d", &numToMove, &from, &to)
		parsedStacks[from-1].Transfer(&parsedStacks[to-1], numToMove)
	}
	for x := range parsedStacks {
		item, _ := parsedStacks[x].Pop()
		sb.WriteByte(item.(uint8))
	}
	return sb.String()
}

func part2(parsedStacks [9]types.Stack, operations []string) string {
	var sb strings.Builder

	for _, line := range operations {
		var numToMove, from, to int
		fmt.Sscanf(line, "move %d from %d to %d", &numToMove, &from, &to)

		tempStack := types.Stack{}

		parsedStacks[from-1].Transfer(&tempStack, numToMove)
		tempStack.Transfer(&parsedStacks[to-1], numToMove)
	}
	for x := range parsedStacks {
		item, _ := parsedStacks[x].Pop()
		sb.WriteByte(item.(uint8))
	}
	return sb.String()
}
