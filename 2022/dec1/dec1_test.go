package dec1

import "testing"

func TestFindElfWithMostCalories(t *testing.T) {
	var input = "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000"
	var expectedResult = 24000
	var result = FindElfWithMostCalories(input)

	if expectedResult != result {
		t.Errorf("%d is not equal to %d", result, expectedResult)
	}
}


func TestFindTop3ElfsWithMostCalories(t *testing.T) {
	var input = "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000"
	var expectedResult = 45000
	var result = FindTop3ElfsWithMostCalories(input)

	if expectedResult != result {
		t.Errorf("%d is not equal to %d", result, expectedResult)
	}
}
