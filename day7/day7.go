package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var embedFs embed.FS
var inputFile = "input.txt"

type dir struct {
	parent *dir
	name   string
	size   int64
}

func main() {
	allDirs := parseInput()
	fmt.Printf("Part 1: %d\n", part1(allDirs))
	fmt.Printf("Part 2: %d", part2(allDirs))
}

func part1(allDirs []*dir) int64 {
	var part1 int64

	for _, dir := range allDirs {
		if dir.size <= 100000 {
			part1 += dir.size
		}
	}
	return part1
}

func part2(allDirs []*dir) int64 {
	var part2 int64

	for _, dir := range allDirs {
		if dir.size >= 30000000-(70000000-allDirs[0].size) && dir.size < part2 {
			part2 = dir.size
		}
	}
	return part2
}

func addSize(dir *dir, size int64) {
	dir.size += size
	if dir.parent != nil {
		addSize(dir.parent, size)
	}
}

func parseInput() (allDirs []*dir) {
	fileToRead, readError := embedFs.Open(inputFile)

	if readError != nil {
		log.Fatal(readError)
	}

	defer fileToRead.Close()

	scanner := bufio.NewScanner(fileToRead)

	var currentDir *dir

	var digitCheck = regexp.MustCompile(`^[0-9]+\s`)

	for scanner.Scan() {
		switch line := scanner.Text(); {

		case strings.Contains(line, "cd "):
			splits := strings.Split(line, " ")
			name := splits[len(splits)-1]

			if name == ".." {
				currentDir = currentDir.parent
			} else {
				var newDir dir
				if currentDir == nil {
					newDir = dir{parent: nil, name: splits[len(splits)-1]}
				} else {
					newDir = dir{parent: currentDir, name: splits[len(splits)-1]}
				}
				allDirs = append(allDirs, &newDir)
				currentDir = &newDir
			}

		case digitCheck.MatchString(line):
			splits := strings.Split(line, " ")
			size, _ := strconv.ParseInt(splits[0], 10, 64)
			addSize(currentDir, size)
		}
	}
	return allDirs
}
