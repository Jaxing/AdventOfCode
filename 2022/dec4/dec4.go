package dec4

import (
	"advent_of_code/2022/helpers"
	"fmt"
	"strconv"
	"strings"
)

type section struct {
	start, end int
}

func (a *section) contains(b *section) bool {
	return a.start <= b.start && b.end <= a.end
}

func (a *section) overlap(b *section) bool {
	return a.start <= b.start && b.start <= a.end ||
		a.start <= b.end && b.end <= a.end ||
		b.contains(a)
}

func newSection(input string) *section {
	parts := strings.Split(input, "-")

	if len(parts) != 2 {
		panic(fmt.Sprintf("Invalid section %s", input))
	}

	start, err_start := strconv.Atoi(parts[0])
	end, err_end := strconv.Atoi(parts[1])

	if err_start != nil || err_end != nil {
		panic(fmt.Sprintf("Invalid section, not integers, %s", input))
	}

	return &section{start: start, end: end}
}

func ParseSectionPairs(pairs []string) [][]*section {
	sectionPairs := make([][]*section, len(pairs))
	for i, pair := range pairs {
		parts := strings.Split(pair, ",")
		raw_sect_1 := parts[0]
		raw_sect_2 := parts[1]
		sectionPairs[i] = []*section{newSection(raw_sect_1), newSection(raw_sect_2)}
	}

	return sectionPairs
}

func countNumberOfContainedSectionPairs(input []string) int {
	numberOfContainedSectionPairs := 0

	pairs := ParseSectionPairs(input)

	for _, pair := range pairs {
		if pair[0].contains(pair[1]) || pair[1].contains(pair[0]) {
			numberOfContainedSectionPairs += 1
		}
	}

	return numberOfContainedSectionPairs
}

func countNumberOfOverlapSectionPairs(input []string) int {
	numberOfOverlapSectionPairs := 0

	pairs := ParseSectionPairs(input)

	for _, pair := range pairs {
		if pair[0].overlap(pair[1]) {
			numberOfOverlapSectionPairs += 1
		}
	}

	return numberOfOverlapSectionPairs
}

func GetInput() []string {
	raw_input, err := helpers.DownloadInput(2022, 4)
	if err != nil {
		panic("Failed to fetch input")
	}
	raw_input = raw_input[:len(raw_input)-1]

	return strings.Split(raw_input, "\n")
}

func Task1() int {
	input := GetInput()
	var product = countNumberOfContainedSectionPairs(input)
	return product
}

func Task2() int {
	input := GetInput()
	var product = countNumberOfOverlapSectionPairs(input)
	return product
}
