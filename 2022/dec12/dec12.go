package dec12

import (
	"advent_of_code/2022/helpers"
	"math"
	"strings"

	"github.com/badgerodon/collections"
)

type position struct {
	neighbours     []*position
	value          rune
	isStart, isEnd bool
	depth          int
}

func FindStart(positions []*position) *position {
	for _, pos := range positions {
		if pos.isStart {
			return pos
		}
	}

	panic("")
}

func FindShortestPath(start *position) int {
	queue := collections.NewQueue[*position]()
	nextLevel := make(map[*position]bool)
	visited := make(map[*position]bool)
	visited[start] = true
	queue.Push(start)
	currentLevel := 0

	for ; true; currentLevel++ {
		for pos, ok := queue.Pop(); ok; pos, ok = queue.Pop() {
			if pos.isEnd {
				return currentLevel
			}

			for _, n := range pos.neighbours {
				if _, exists := visited[n]; exists {
					continue
				}
				curr := int(pos.value)
				next := int(n.value)
				if pos.isStart {
					curr = int('a')
				}
				if n.isEnd {
					next = int('z')
				}

				if curr >= next-1 {
					visited[n] = true
					nextLevel[n] = true
				}
			}
		}
		if len(nextLevel) == 0 {
			// panic("Failed to find end")
			return math.MaxInt
		}

		for pos, _ := range nextLevel {
			queue.Push(pos)
		}
		nextLevel = make(map[*position]bool)
	}
	panic("")
}

func ParsePositions(input []string) []*position {
	positions := make([]*position, len(input)*len(input[0]))

	for rowI, row := range input {
		for colI, col := range row {
			pos := &position{
				neighbours: []*position{},
				value:      col,
				isStart:    'S' == col,
				isEnd:      'E' == col,
				depth:      math.MaxInt,
			}

			if rowI > 0 {
				neighbour := positions[colI+(rowI-1)*len(row)]
				pos.neighbours = append(pos.neighbours, neighbour)
				neighbour.neighbours = append(neighbour.neighbours, pos)
			}
			if colI > 0 {
				neighbour := positions[(colI-1)+rowI*len(row)]
				pos.neighbours = append(pos.neighbours, neighbour)
				neighbour.neighbours = append(neighbour.neighbours, pos)
			}

			positions[colI+rowI*len(row)] = pos
		}
	}

	return positions
}

func FindShortestPathFromStart(input []string) int {
	positions := ParsePositions(input)
	start := FindStart(positions)
	return FindShortestPath(start)
}

func FindShortestPathFromSAnyLowPoint(input []string) int {
	positions := ParsePositions(input)
	shortestPath := math.MaxInt

	for _, pos := range positions {
		if pos.isStart || pos.value == 'a' {
			var product = FindShortestPath(pos)
			if product < shortestPath {
				shortestPath = product
			}
		}
	}
	return shortestPath
}

func GetInput() []string {
	raw_input, err := helpers.DownloadInput(2022, 12)
	if err != nil {
		panic("Failed to fetch input")
	}
	raw_input = raw_input[:len(raw_input)-1]

	return strings.Split(raw_input, "\n")
}

func Task1() int {
	input := GetInput()
	return FindShortestPathFromStart(input)
}

func Task2() int {
	input := GetInput()
	return FindShortestPathFromSAnyLowPoint(input)
}
