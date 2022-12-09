package dec9

import (
	"fmt"
	"testing"
)

func TestMoveHeadDirectionTableDriven(t *testing.T) {
	var tests = []struct {
		r   rope
		dir string
		out rope
	}{
		{rope{head: &point{row: 0, column: 0}, tails: []*point{&point{row: 0, column: 0}}}, "U", rope{head: &point{row: 1, column: 0}, tails: []*point{&point{row: 0, column: 0}}}},
		{rope{head: &point{row: 1, column: 1}, tails: []*point{&point{row: 0, column: 0}}}, "R", rope{head: &point{row: 1, column: 2}, tails: []*point{&point{row: 1, column: 1}}}},
		{rope{head: &point{row: 2, column: 1}, tails: []*point{&point{row: 3, column: 1}}}, "D", rope{head: &point{row: 1, column: 1}, tails: []*point{&point{row: 2, column: 1}}}},
		{rope{head: &point{row: 2, column: 1}, tails: []*point{&point{row: 3, column: 1}}}, "L", rope{head: &point{row: 2, column: 0}, tails: []*point{&point{row: 3, column: 1}}}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.dir)
		t.Run(testname, func(t *testing.T) {
			tt.r.MoveHead(tt.dir)

			if tt.r.head == tt.out.head && tt.r.tails[0] == tt.out.tails[0] {
				t.Errorf("test %s failed", testname)
			}
		})
	}
}

func TestCountTailMoves(t *testing.T) {
	input := []string{
		"R 5",
		"U 8",
		"L 8",
		"D 3",
		"R 17",
		"D 10",
		"L 25",
		"U 20",
	}

	expected := 36
	res := CountTailMoves(input, 9)

	if expected != res {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}
