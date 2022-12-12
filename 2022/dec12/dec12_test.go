package dec12

import "testing"

func TestFindShortestPathFromStart(t *testing.T) {
	input := []string{
		"Sabqponm",
		"abcryxxl",
		"accszExk",
		"acctuvwj",
		"abdefghi",
	}

	var want = 31
	ans := FindShortestPathFromStart(input)

	if ans != want {
		t.Errorf("expected %d got %d", want, ans)
	}
}

func TestFindShortestPathFromStart2(t *testing.T) {
	input := []string{
		"SabcdefghijklmnopqrstuvwxyzE",
	}

	var want = 27
	ans := FindShortestPathFromStart(input)

	if ans != want {
		t.Errorf("expected %d got %d", want, ans)
	}
}

func TestFindShortestPathFromStart3(t *testing.T) {
	input := []string{
		"Sabcdef",
		"mlkjihg",
		"nopqrst",
		"Ezyxwvu",
	}

	var want = 27
	ans := FindShortestPathFromStart(input)

	if ans != want {
		t.Errorf("expected %d got %d", want, ans)
	}
}

func TestFindShortestPathFromSAnyLowPoint(t *testing.T) {
	input := []string{
		"Sabqponm",
		"abcryxxl",
		"accszExk",
		"acctuvwj",
		"abdefghi",
	}

	var want = 29
	ans := FindShortestPathFromSAnyLowPoint(input)

	if ans != want {
		t.Errorf("expected %d got %d", want, ans)
	}
}
