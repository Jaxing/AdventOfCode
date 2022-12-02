package dec2

import (
	"fmt"
	"testing"
)

func TestCalculateScore(t *testing.T) {
	fmt.Println("TestCalculateScore")
	var input = []string{"A Y", "B X", "C Z"}
	var expectedResult = 15
	var result = CalculateScore(input)

	if expectedResult != result {
		t.Errorf("%d is not equal to %d", result, expectedResult)
	}
}

func TestCalculateScore2(t *testing.T) {
	fmt.Println("TestCalculateScore2")
	var input = []string{"A Y", "B X", "C Z"}
	var expectedResult = 12
	var result = CalculateScore2(input)

	if expectedResult != result {
		t.Errorf("%d is not equal to %d", result, expectedResult)
	}
}

func TestCalculateScore2_1(t *testing.T) {
	fmt.Println("TestCalculateScore2_1")
	var input = []string{"A Y", "B Y", "C Y"}
	var expectedResult = 1 + 3 + 2 + 3 + 3 + 3
	var result = CalculateScore2(input)

	if expectedResult != result {
		t.Errorf("%d is not equal to %d", result, expectedResult)
	}
}
func TestCalculateScore2_2(t *testing.T) {
	fmt.Println("TestCalculateScore2_2")
	var input = []string{"A X", "B X", "C X"}
	var expectedResult = 3 + 0 + 1 + 0 + 2 + 0
	var result = CalculateScore2(input)

	if expectedResult != result {
		t.Errorf("%d is not equal to %d", result, expectedResult)
	}
}

func TestCalculateScore2_3(t *testing.T) {
	fmt.Println("TestCalculateScore2_3")
	var input = []string{"A Z", "B Z", "C Z"}
	var expectedResult = 2 + 6 + 3 + 6 + 1 + 6
	var result = CalculateScore2(input)

	if expectedResult != result {
		t.Errorf("%d is not equal to %d", result, expectedResult)
	}
}
