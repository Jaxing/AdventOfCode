package dec1

import "testing"

func TestFindTwoEntries(t *testing.T) {
	var input = [6]int{ 1721, 979, 366, 299, 675, 1456}

	var expectedResult = 514579
	var result = findTwoEntriesProduct(input[:], 2020)

	if expectedResult != result {
		t.Errorf("%d is not equal to %d", result, expectedResult)
	}
}

func TestFindThreeEntries(t *testing.T) {
	var input = [6]int{ 1721, 979, 366, 299, 675, 1456}

	var expectedResult = 241861950
	var result = findThreeEntriesProduct(input[:], 2020)

	if expectedResult != result {
		t.Errorf("%d is not equal to %d", result, expectedResult)
	}
}