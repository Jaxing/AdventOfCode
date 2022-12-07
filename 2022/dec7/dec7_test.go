package dec7

import "testing"

func TestParseInstructionsAndSum(t *testing.T) {
	var input = []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"$ ls",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k",
	}

	expectedResult := 95437
	result := SumOfDirsOfSize(ParseTree(input), 100000)

	if expectedResult != result {
		t.Errorf("%d is not equal to %d", result, expectedResult)
		return
	}

}

func TestParseInstructionsAndFind(t *testing.T) {
	var input = []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"$ ls",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k",
	}

	expectedResult := 24933642
	result := FindSmallestDirToDelete(ParseTree(input), 30000000)

	if expectedResult != result {
		t.Errorf("%d is not equal to %d", result, expectedResult)
		return
	}

}
