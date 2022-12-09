package dec9

import (
	"advent_of_code/2022/helpers"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type point struct {
	row, column int
}

type rope struct {
	head  *point
	tails []*point
}

func (a *rope) MoveHead(direction string) {
	switch direction {
	case "U":
		a.head.row++
	case "D":
		a.head.row--
	case "L":
		a.head.column--
	case "R":
		a.head.column++
	default:
		panic("Invalid direction")
	}

	for i := 0; i < len(a.tails); i++ {
		tail := a.tails[i]
		var previousKnot *point
		if i == 0 {
			previousKnot = a.head
		} else {
			previousKnot = a.tails[i-1]
		}
		if math.Abs(float64(previousKnot.column-tail.column)) <= 1 &&
			math.Abs(float64(previousKnot.row-tail.row)) <= 1 {
			continue
		}

		if previousKnot.column == tail.column {
			if previousKnot.row < tail.row {
				tail.row--
				continue
			}
			tail.row++
			continue
		}
		if previousKnot.row == tail.row {
			if previousKnot.column < tail.column {
				tail.column--
				continue
			}
			tail.column++
			continue
		}
		if previousKnot.row < tail.row {
			if previousKnot.column < tail.column {
				tail.column--
				tail.row--
				continue
			}
			tail.column++
			tail.row--
			continue
		}
		if previousKnot.column > tail.column {
			tail.column++
			tail.row++
			continue
		}
		tail.column--
		tail.row++
	}

}

func CountTailMoves(input []string, numTails int) int {
	tailPositions := make(map[point]bool)
	theRope := rope{
		head:  &point{row: 0, column: 0},
		tails: make([]*point, numTails),
	}

	for i, _ := range theRope.tails {
		tail := &point{row: 0, column: 0}
		theRope.tails[i] = tail
	}

	tailPositions[*theRope.tails[len(theRope.tails)-1]] = true

	for _, move := range input {
		parts := strings.Split(move, " ")
		numSteps, err := strconv.Atoi(parts[1])

		if err != nil {
			panic(fmt.Sprintf("Invalid input %s", move))
		}

		for i := 0; i < numSteps; i++ {
			theRope.MoveHead(parts[0])
			tailPositions[*theRope.tails[len(theRope.tails)-1]] = true
		}
	}

	return len(tailPositions)
}

func GetInput() []string {
	raw_input, err := helpers.DownloadInput(2022, 9)
	if err != nil {
		panic("Failed to fetch input")
	}
	raw_input = raw_input[:len(raw_input)-1]

	return strings.Split(raw_input, "\n")
}

func Task1() int {
	input := GetInput()
	var product = CountTailMoves(input, 1)
	return product
}

func Task2() int {
	input := GetInput()
	var product = CountTailMoves(input, 9)
	return product
}
