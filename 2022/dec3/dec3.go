package dec3

import (
	"advent_of_code/2022/helpers"
	"strings"
)

func FindCommonItems(input string) rune {
	lenght_of_input := len(input)
	items_types_in_first := make(map[rune]bool)

	for i, item := range input {
		if i < lenght_of_input/2 {
			items_types_in_first[item] = true
			continue
		}
		if _, exists := items_types_in_first[item]; exists {
			return item
		}
	}

	return '0'
}

func FindBadges(rucksack_one string, rucksack_two string, rucksack_three string) rune {
	items_types_in_first := make(map[rune]bool)
	badge_candidates := make(map[rune]bool)

	for _, item := range rucksack_one {
		items_types_in_first[item] = true
	}

	for _, item := range rucksack_two {
		if _, exists := items_types_in_first[item]; exists {
			badge_candidates[item] = true
		}
	}

	for _, item := range rucksack_three {
		if _, exists := badge_candidates[item]; exists {
			return item
		}
	}

	return '0'
}

func CalculatePriority(item rune) int {
	if item >= 'a' {
		return int(item) - 96
	}
	return int(item) - (65 - 27)
}

func SumPriorities(input []string) int {
	var total_prio = 0
	for _, rucksack := range input {
		item := FindCommonItems(rucksack)
		total_prio = total_prio + CalculatePriority(item)
	}
	return total_prio
}

func SumBadgePrios(input []string) int {
	upper_limit := 3
	total_prio := 0

	for ; upper_limit <= len(input); upper_limit = upper_limit + 3 {
		badge := FindBadges(input[upper_limit-3], input[upper_limit-2], input[upper_limit-1])
		total_prio = total_prio + CalculatePriority(badge)
	}

	return total_prio
}

func GetInput() []string {
	raw_input, err := helpers.DownloadInput(2022, 3)
	if err != nil {
		panic("Failed to fetch input")
	}
	raw_input = raw_input[:len(raw_input)-1]

	return strings.Split(raw_input, "\n")
}

func Task1() int {
	input := GetInput()
	var product = SumPriorities(input)
	return product
}

func Task2() int {
	input := GetInput()
	var product = SumBadgePrios(input)
	return product
}
