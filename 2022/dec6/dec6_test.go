package dec6

import (
	"fmt"
	"testing"
)

func TestFindEndOfMarkerTableDriven(t *testing.T) {
	var tests = []struct {
		in             string
		lenght, output int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4, 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 4, 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 4, 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4, 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4, 11},
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14, 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 14, 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 14, 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14, 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14, 26},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.in)
		t.Run(testname, func(t *testing.T) {
			ans := FindEndOfMarker(tt.in, tt.lenght)

			if ans != tt.output {
				t.Errorf("got %d, want %d", ans, tt.output)
			}
		})
	}
}
