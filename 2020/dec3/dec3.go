package dec3

import (
	"strings"

	"../helpers"
)

func GetNextColumn(currentColumn int, steps int, maxColumn int) int {
	if currentColumn+steps < maxColumn {
		return currentColumn + steps
	}
	return currentColumn + steps - maxColumn
}

func CountTrees(input []string, movesRight int, movesDown int) int {
	currentRow, currentColumn := 0, 0
	numTrees := 0

	for currentRow < len(input) {
		if rune(input[currentRow][currentColumn]) == '#' {
			numTrees++
		}
		currentColumn = GetNextColumn(currentColumn, movesRight, len(input[currentRow]))
		currentRow += movesDown
	}

	return numTrees
}

func GetInput() []string {
	raw_input, err := helpers.DownloadInput(2020, 3)
	if err != nil {
		panic("Failed to fetch input")
	}

	stringList := strings.Split(raw_input, "\n")
	stringList = stringList[:len(stringList)-1]

	return stringList
}

func Task1() int {
	input := GetInput()
	return CountTrees(input, 3, 1)
}

func Task2() int {
	inMap := GetInput()
	numTrees1 := CountTrees(inMap, 1, 1)
	numTrees2 := CountTrees(inMap, 3, 1)
	numTrees3 := CountTrees(inMap, 5, 1)
	numTrees4 := CountTrees(inMap, 7, 1)
	numTrees5 := CountTrees(inMap, 1, 2)

	product := numTrees1 * numTrees2 * numTrees3 * numTrees4 * numTrees5
	return product
}
