package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var embedFs embed.FS
var inputFile = "input.txt"

func main() {
	allTrees := parseInput()
	fmt.Printf("Part 1: %d\n", part1(allTrees))
	fmt.Printf("Part 2: %d", part2(allTrees))
}

func part1(trees [][]int) int {
	// edge trees
	part1 := 2*(len(trees)) + 2*(len(trees[0])) - 4

	for row := 1; row < len(trees)-1; row++ {
		for col := 1; col < len(trees[row])-1; col++ {
			if isVisible(row, col, trees) {
				part1++
			}
		}
	}
	return part1
}

func part2(trees [][]int) int {
	var highest int

	for row := 1; row < len(trees)-1; row++ {
		for col := 1; col < len(trees[row])-1; col++ {
			var treesUp, treesDown []int

			for idx := 0; idx < row; idx++ {
				treesUp = append(treesUp, trees[idx][col])
			}

			for idx := row + 1; idx < len(trees); idx++ {
				treesDown = append(treesDown, trees[idx][col])
			}

			treeHeight := trees[row][col]

			score := getScenicScore(treeHeight, reverse(trees[row][:col])) *
				getScenicScore(treeHeight, trees[row][col+1:]) *
				getScenicScore(treeHeight, reverse(treesUp)) *
				getScenicScore(treeHeight, treesDown)

			if score > highest {
				highest = score
			}
		}
	}
	return highest
}

func getScenicScore(treeHeight int, treesToSearch []int) int {
	var score int

	for x := 0; x < len(treesToSearch); x++ {
		if treesToSearch[x] < treeHeight {
			score++
		} else {
			score++
			break
		}
	}

	return score
}

func reverse(s []int) []int {
	var r []int
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}

func isVisible(row int, col int, allTrees [][]int) bool {
	treeHeight := allTrees[row][col]

	var treesUp, treesDown []int

	for idx := 0; idx < row; idx++ {
		treesUp = append(treesUp, allTrees[idx][col])
	}

	for idx := row + 1; idx < len(allTrees); idx++ {
		treesDown = append(treesDown, allTrees[idx][col])
	}

	return !hasTreeBlocking(treeHeight, allTrees[row][:col]) ||
		!hasTreeBlocking(treeHeight, allTrees[row][col+1:]) ||
		!hasTreeBlocking(treeHeight, treesUp) ||
		!hasTreeBlocking(treeHeight, treesDown)
}

func hasTreeBlocking(treeHeight int, treesToSearch []int) bool {
	for _, tree := range treesToSearch {
		if tree >= treeHeight {
			return true
		}
	}
	return false
}

func parseInput() (allTrees [][]int) {
	fileToRead, readError := embedFs.Open(inputFile)

	if readError != nil {
		log.Fatal(readError)
	}

	defer fileToRead.Close()

	scanner := bufio.NewScanner(fileToRead)
	count := 0

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		allTrees = append(allTrees, []int{})

		for i := range row {
			num, _ := strconv.Atoi(row[i])
			allTrees[count] = append(allTrees[count], num)
		}
		count++
	}
	return allTrees
}
