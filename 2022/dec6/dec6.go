package dec6

import (
	"advent_of_code/2022/helpers"

	"github.com/badgerodon/collections"
)

func FindEndOfMarker(input string, lengthOfMarker int) int {
	i := lengthOfMarker
	for ; i < len(input); i++ {
		marker := collections.NewSet[rune]()
		for _, m := range input[i-lengthOfMarker : i] {
			marker.Add(m)
		}

		if marker.Size() == lengthOfMarker {
			return i
		}
	}

	return i
}

func GetInput() string {
	raw_input, err := helpers.DownloadInput(2022, 6)
	if err != nil {
		panic("Failed to fetch input")
	}
	raw_input = raw_input[:len(raw_input)-1]

	return raw_input
}

func Task1() int {
	input := GetInput()
	var product = FindEndOfMarker(input, 4)
	return product
}

func Task2() int {
	input := GetInput()
	var product = FindEndOfMarker(input, 14)
	return product
}
