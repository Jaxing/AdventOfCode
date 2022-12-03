package dec3

import (
	"fmt"
	"testing"
)

func TestFindCommonItemsTableDriven(t *testing.T) {
	var tests = []struct {
		input  string
		output rune
	}{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", 'p'},
		{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", 'L'},
		{"PmmdzqPrVvPwwTWBwg", 'P'},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", 'v'},
		{"ttgJtRGJQctTZtZT", 't'},
		{"CrZsJsPPZsGzwwsLwLmpwMDw", 's'},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.input)
		t.Run(testname, func(t *testing.T) {
			ans := FindCommonItems(tt.input)
			if ans != tt.output {
				t.Errorf("got %s, want %s", string(ans), string(tt.output))
			}
		})
	}
}

func TestSumPriorities(t *testing.T) {
	var input = []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	var expectedResult = 157
	var result = SumPriorities(input)

	if expectedResult != result {
		t.Errorf("%d is not equal to %d", result, expectedResult)
	}
}

func TestFindBadgesTableDriven(t *testing.T) {
	var tests = []struct {
		a, b, c string
		output  rune
	}{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", 'r'},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw", 'Z'},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			ans := FindBadges(tt.a, tt.b, tt.c)
			if ans != tt.output {
				t.Errorf("got %s, want %s", string(ans), string(tt.output))
			}
		})
	}
}

func TestSumBadgePriorities(t *testing.T) {
	var input = []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	var expectedResult = 70
	var result = SumBadgePrios(input)

	if expectedResult != result {
		t.Errorf("%d is not equal to %d", result, expectedResult)
	}
}
