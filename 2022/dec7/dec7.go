package dec7

import (
	"advent_of_code/2022/helpers"
	"fmt"
	"strconv"
	"strings"

	"github.com/badgerodon/collections"
)

type node struct {
	name     string
	parent   *node
	size     int
	children map[string]*node
}

func (a *node) sumSize() int {
	var items collections.Stack[*node]
	items.Push(a)
	size := 0

	for item, ok := items.Pop(); ok; {
		if len(item.children) == 0 {
			size += item.size
			continue
		}
		for _, child := range item.children {
			items.Push(child)
		}
	}

	return size
}

func ParseTree(input []string) *node {
	root := node{name: "/", size: 0, parent: nil, children: make(map[string]*node)}
	current_node := &root
	for _, line := range input {
		if strings.HasPrefix(line, "$") {
			if line == "$ cd /" {
				current_node = &root
				continue
			}
			if line == "$ cd .." {
				if current_node.parent == nil {
					panic(fmt.Sprintf("Cant cd further up, node name %s", current_node.name))
				}
				current_node = current_node.parent
				continue
			}
			if line == "$ ls" {
				continue
			}
			parts := strings.Split(line, " ")

			if child, exists := current_node.children[parts[len(parts)-1]]; exists {
				current_node = child
			}
			continue
		}
		parts := strings.Split(line, " ")
		size := parts[0]
		name := parts[1]

		switch size {
		case "dir":
			current_node.children[name] = &node{
				name:     name,
				size:     0,
				parent:   current_node,
				children: make(map[string]*node),
			}
		default:
			parsedSize, err := strconv.Atoi(size)

			if err != nil {
				panic(fmt.Sprintf("size not parsable, %s", size))
			}
			current_node.children[name] = &node{
				name:     name,
				size:     parsedSize,
				parent:   current_node,
				children: make(map[string]*node),
			}
		}
	}

	return &root
}

func SumOfDirsOfSize(root *node, maxDirSize int) int {
	sizes := SumOfDirsOfSizeRecursive(root)

	size := 0

	for _, dirSize := range sizes {
		if dirSize <= maxDirSize {
			size += dirSize
		}
	}

	return size
}

func SumOfDirsOfSizeRecursive(root *node) map[*node]int {
	sizes := make(map[*node]int)
	sizes[root] = 0

	for _, child := range root.children {
		if len(child.children) == 0 {
			sizes[root] += child.size
			continue
		}

		tmpSizes := SumOfDirsOfSizeRecursive(child)

		for node, size := range tmpSizes {
			sizes[node] = size
		}
		sizes[root] += sizes[child]
	}

	return sizes
}

func FindSmallestDirToDelete(root *node, sizeNeeded int) int {
	sizes := SumOfDirsOfSizeRecursive(root)

	freeSpace := 70000000 - sizes[root]
	spaceToFreeUp := sizeNeeded - freeSpace

	smallestSize := sizes[root]

	for _, dirSize := range sizes {
		if dirSize < smallestSize && dirSize >= spaceToFreeUp {
			smallestSize = dirSize
		}
	}

	return smallestSize
}

func GetInput() []string {
	raw_input, err := helpers.DownloadInput(2022, 7)
	if err != nil {
		panic("Failed to fetch input")
	}
	raw_input = raw_input[:len(raw_input)-1]

	return strings.Split(raw_input, "\n")
}

func Task1() int {
	input := GetInput()
	var product = SumOfDirsOfSize(ParseTree(input), 100000)
	return product
}

func Task2() int {
	input := GetInput()
	var product = FindSmallestDirToDelete(ParseTree(input), 30000000)
	return product
}
