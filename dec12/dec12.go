package dec12

import (
	"../helpers"
	"log"
	"strings"
	"fmt"
)

const fileName = "dec12/input"

func Part1() int {
	rows, err := helpers.ReadFile(fileName)

	if err != nil {
		log.Fatal("file does not exist")
	}

	var progToProgs = make(map[string][]string)
	var visitedProgs = make(map[string]bool)

	for _, row := range rows {
		parts := strings.Split(row, " <-> ")
		part1 := parts[0]
		part2 := strings.Split(parts[1], ", ")
		progToProgs[part1] = part2
	}

	return visit("0", progToProgs, visitedProgs)
}

func visit(current string, progToprog map[string][]string, visited map[string]bool) int {
	if current == "" {
		return 0
	}

	visited[current] = true

	var neigbours = progToprog[current]
	var count = 1

	for _, neigbour := range neigbours {
		if !visited[neigbour] {
			count += visit(neigbour, progToprog, visited)
		}
	}

	return count
}

func Part2() int {
	rows, err := helpers.ReadFile(fileName)

	if err != nil {
		log.Fatal("file does not exist")
	}

	var progToProgs = make(map[string][]string)
	var visitedProgs = make(map[string]bool)

	for _, row := range rows {
		parts := strings.Split(row, " <-> ")
		part1 := parts[0]
		part2 := strings.Split(parts[1], ", ")
		progToProgs[part1] = part2
		visitedProgs[part1] = false
	}

	return visitAndCount("0", progToProgs, visitedProgs)
}

func visitAndCount(current string, progToProg map[string][]string, visited map[string]bool) int {
	fmt.Println(current)
	if current == "" {
		return 0
	}

	if !visited[current] {
		visit(current, progToProg, visited)
	}
	var next = ""

	for k, v := range visited {
		if !v {
			next = k
			break
		}
	}

	return visitAndCount(next, progToProg, visited) + 1
}