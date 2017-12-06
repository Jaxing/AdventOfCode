package dec5

import (
	"../helpers"
	"log"
	"strconv"
)

const DIR_ROOT = "dec5/input"

func Part1() int {
	rows, err := helpers.ReadFile(DIR_ROOT)

	if err != nil {
		log.Fatal("file does not exist")
	}

	var instruction int = 0
	var steps int = 0
	var intRows = make([]int64, len(rows))

	for i, row := range rows {
		intRows[i], _ = strconv.ParseInt(row, 10, 0)
	}

	for instruction < len(intRows) {
		var offset = int(intRows[instruction])

		intRows[instruction]++
		steps++
		instruction += offset
	}
	return steps
}

func Part2() int {
	rows, err := helpers.ReadFile(DIR_ROOT)

	if err != nil {
		log.Fatal("file does not exist")
	}

	var instruction int = 0
	var steps int = 0
	var intRows = make([]int64, len(rows))

	for i, row := range rows {
		intRows[i], _ = strconv.ParseInt(row, 10, 0)
	}

	for instruction < len(intRows) {
		var offset = int(intRows[instruction])

		if offset >= 3 {
			intRows[instruction]--
		} else {
			intRows[instruction]++
		}
		steps++
		instruction += offset
	}
	return steps
}