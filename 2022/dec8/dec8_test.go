package dec8

import "testing"

func TestCountVisibleTrees(t *testing.T) {
	var input = []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	expectedResult := 21
	result := CountVisibleTrees(ParseTrees(input))

	if expectedResult != result {
		t.Errorf("%d is not equal to %d", result, expectedResult)
		return
	}

}

func TestCalculateHeighestScenicScore(t *testing.T) {
	var input = []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	expectedResult := 8
	result := CalculateHeighestScenicScore(ParseTrees(input))

	if expectedResult != result {
		t.Errorf("%d is not equal to %d", result, expectedResult)
		return
	}

}

func TestCalculateScenicScore(t *testing.T) {
	var input = []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	trees := ParseTrees(input)

	expectedResult := 4
	result := trees[7].ScenicScore()

	if expectedResult != result {
		t.Errorf("%d is not equal to %d", result, expectedResult)
		return
	}

}

func TestCalculateScenicScore2(t *testing.T) {
	var input = []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	trees := ParseTrees(input)

	expectedResult := 8
	result := trees[17].ScenicScore()

	if expectedResult != result {
		t.Errorf("%d is not equal to %d", result, expectedResult)
		return
	}

}
