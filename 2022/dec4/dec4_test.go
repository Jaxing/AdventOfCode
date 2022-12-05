package dec4

import (
	"fmt"
	"strconv"
	"testing"
)

func TestContainsTableDriven(t *testing.T) {
	var tests = []struct {
		a, b   *section
		output bool
	}{
		{&section{start: 2, end: 4}, &section{start: 6, end: 8}, false},
		{&section{start: 2, end: 3}, &section{start: 4, end: 5}, false},
		{&section{start: 5, end: 7}, &section{start: 7, end: 9}, false},
		{&section{start: 2, end: 8}, &section{start: 3, end: 7}, true},
		{&section{start: 6, end: 6}, &section{start: 4, end: 6}, false},
		{&section{start: 4, end: 6}, &section{start: 6, end: 6}, true},
		{&section{start: 2, end: 6}, &section{start: 4, end: 8}, false},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			ans := tt.a.contains(tt.b)
			if ans != tt.output {
				t.Errorf("got %s, want %s", strconv.FormatBool(ans), strconv.FormatBool(tt.output))
			}
		})
	}
}

func TestOverlapTableDriven(t *testing.T) {
	var tests = []struct {
		a, b   *section
		output bool
	}{
		{&section{start: 2, end: 4}, &section{start: 6, end: 8}, false},
		{&section{start: 2, end: 3}, &section{start: 4, end: 5}, false},
		{&section{start: 5, end: 7}, &section{start: 7, end: 9}, true},
		{&section{start: 2, end: 8}, &section{start: 3, end: 7}, true},
		{&section{start: 6, end: 6}, &section{start: 4, end: 6}, true},
		{&section{start: 4, end: 6}, &section{start: 6, end: 6}, true},
		{&section{start: 2, end: 6}, &section{start: 4, end: 8}, true},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			ans_a := tt.a.overlap(tt.b)
			ans_b := tt.b.overlap(tt.a)

			if ans_a != ans_b {
				t.Errorf("exptected method to be symetric, but a: %s, b: %s", strconv.FormatBool(ans_a), strconv.FormatBool(ans_b))
			}

			if ans_a != tt.output {
				t.Errorf("got %s, want %s", strconv.FormatBool(ans_a), strconv.FormatBool(tt.output))
			}
		})
	}
}

func TestSumPriorities(t *testing.T) {
	var input = []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}

	var expectedResult = 2
	var result = countNumberOfContainedSectionPairs(input)

	if expectedResult != result {
		t.Errorf("%d is not equal to %d", result, expectedResult)
	}
}
