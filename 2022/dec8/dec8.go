package dec8

import (
	"advent_of_code/2022/helpers"
	"strconv"
	"strings"
)

type tree struct {
	height                uint32
	left, right, up, down *tree
}

func (a *tree) IsEdge() bool {
	return a.left == nil ||
		a.right == nil ||
		a.up == nil ||
		a.down == nil
}

func (a *tree) IsVisible() bool {
	if a.IsEdge() {
		return true
	}

	_, isVisibleUp := a.VisibilityInDirection("up")
	_, isVisibleLeft := a.VisibilityInDirection("left")
	_, isVisibleRight := a.VisibilityInDirection("right")
	_, isVisibleDown := a.VisibilityInDirection("down")

	return isVisibleUp || isVisibleDown || isVisibleLeft || isVisibleRight
}

func (a *tree) ScenicScore() int {

	visibilityUp, _ := a.VisibilityInDirection("up")
	visibilityLeft, _ := a.VisibilityInDirection("left")
	visibilityRight, _ := a.VisibilityInDirection("right")
	visibilityDown, _ := a.VisibilityInDirection("down")

	return visibilityUp * visibilityLeft * visibilityRight * visibilityDown
}

func (a *tree) VisibilityInDirection(direction string) (int, bool) {
	nextTree := GetTreeInDirection(a, direction)
	distance := 0

	for nextTree != nil {
		if nextTree.height >= a.height {
			return distance + 1, false
		}
		distance += 1
		nextTree = GetTreeInDirection(nextTree, direction)
	}

	return distance, true
}

func GetTreeInDirection(t *tree, direction string) *tree {
	switch direction {
	case "up":
		return t.up
	case "down":
		return t.down
	case "left":
		return t.left
	case "right":
		return t.right
	default:
		panic("Invalid direction")
	}
}

func CountVisibleTrees(forest []*tree) int {
	count := 0
	for _, tree := range forest {
		if tree.IsVisible() {
			count++
		}
	}
	return count
}

func CalculateHeighestScenicScore(forest []*tree) int {
	maxScenicScore := 0
	for _, tree := range forest {
		scenicScore := tree.ScenicScore()
		if maxScenicScore < scenicScore {
			maxScenicScore = scenicScore
		}
	}
	return maxScenicScore
}

func ParseTrees(input []string) []*tree {
	forest := make([]*tree, len(input)*len(input[0]))
	for rowI, row := range input {
		for colI, treeHeight := range strings.Split(row, "") {
			parsedTreeHeight, err := strconv.Atoi(treeHeight)

			if err != nil {
				panic("Invalid tree height")
			}

			thisTree := &tree{height: uint32(parsedTreeHeight)}

			if rowI > 0 {
				thisTree.up = forest[colI+(rowI-1)*len(row)]
				thisTree.up.down = thisTree
			}
			if colI > 0 {
				thisTree.left = forest[(colI-1)+rowI*len(row)]
				thisTree.left.right = thisTree
			}

			forest[colI+rowI*len(row)] = thisTree
		}
	}

	return forest
}

func GetInput() []string {
	raw_input, err := helpers.DownloadInput(2022, 8)
	if err != nil {
		panic("Failed to fetch input")
	}
	raw_input = raw_input[:len(raw_input)-1]

	return strings.Split(raw_input, "\n")
}

func Task1() int {
	input := GetInput()
	var product = CountVisibleTrees(ParseTrees(input))
	return product
}

func Task2() int {
	input := GetInput()
	var product = CalculateHeighestScenicScore(ParseTrees(input))
	return product
}
