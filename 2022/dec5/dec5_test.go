package dec5

import (
	"testing"

	"github.com/badgerodon/collections"
)

func TestParseStacks(t *testing.T) {
	var input = []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3",
	}

	stack_one := collections.NewStack[string]()
	stack_two := collections.NewStack[string]()
	stack_three := collections.NewStack[string]()
	stack_one.Push("Z")
	stack_one.Push("N")
	stack_two.Push("M")
	stack_two.Push("C")
	stack_two.Push("D")
	stack_three.Push("P")
	var expectedResult = []collections.Stack[string]{
		stack_one, stack_two, stack_three,
	}
	var result = ParseStacks(input)

	if len(expectedResult) != len(result) {
		t.Errorf("Length not equal,%d is not equal to %d", len(result), len(expectedResult))
		return
	}

	for i := 0; i < 3; i++ {
		expectedStack := expectedResult[i]
		myStack := result[i]

		_, okExpected := expectedStack.Peek()
		_, ok := myStack.Peek()

		for okExpected && ok {
			topExpected, _ := expectedStack.Pop()
			topMy, _ := myStack.Pop()

			if topExpected != topMy {
				t.Errorf("Stack elements differ in stack %d, %s is not %s", i, topMy, topExpected)
				return
			}

			_, okExpected = expectedStack.Peek()
			_, ok = myStack.Peek()
		}

		if okExpected != ok {
			t.Errorf("Stack %d was different length", i)
			return
		}
	}

}

func TestParseInstructions(t *testing.T) {
	var input = []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	expectedResult := []instruction{
		{amount: 1, from: 2, to: 1},
		{amount: 3, from: 1, to: 3},
		{amount: 2, from: 2, to: 1},
		{amount: 1, from: 1, to: 2},
	}
	result := ParseInstructions(input)

	if len(expectedResult) != len(result) {
		t.Errorf("Length not equal,%d is not equal to %d", len(result), len(expectedResult))
		return
	}

	for i, res := range result {
		exp := expectedResult[i]

		if res.from != exp.from {
			t.Errorf("Instructions differ at instruction %d for from, %d not equal to %d", i, res.from, exp.from)
			return
		}
		if res.amount != exp.amount {
			t.Errorf("Instructions differ at instruction %d for amount, %d not equal to %d", i, res.amount, exp.amount)
			return
		}
		if res.to != exp.to {
			t.Errorf("Instructions differ at instruction %d for to, %d not equal to %d", i, res.to, exp.to)
			return
		}
	}
}
