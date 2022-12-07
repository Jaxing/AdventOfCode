package dec5

import (
	"advent_of_code/2022/helpers"
	"fmt"
	"strconv"
	"strings"

	"github.com/badgerodon/collections"
)

type instruction struct {
	amount, from, to int
}

func ParseStacks(input []string) []collections.Stack[string] {
	index := input[len(input)-1]
	input = input[:len(input)-1]
	numCharactersInPosition := 4

	stacks := make([]collections.Stack[string], len(strings.Split(index, "   ")))

	for i, _ := range stacks {
		stacks[i] = collections.NewStack[string]()
	}

	for i := len(input) - 1; i >= 0; i-- {
		stackUpper := numCharactersInPosition - 2
		stackLower := 1
		for _, stack := range stacks {
			crate := input[i][stackLower:stackUpper]
			if crate != " " {
				stack.Push(crate)
			}
			stackLower += numCharactersInPosition
			stackUpper += numCharactersInPosition
		}
	}

	return stacks
}

func ParseInstructions(input []string) []instruction {
	indexOfAmount := 1
	indexOfFrom := 3
	indexOfTo := 5

	instructions := make([]instruction, len(input))

	for i := 0; i < len(input); i++ {
		parts := strings.Split(input[i], " ")
		amount, err1 := strconv.Atoi(parts[indexOfAmount])
		from, err2 := strconv.Atoi(parts[indexOfFrom])
		to, err3 := strconv.Atoi(parts[indexOfTo])

		if err1 != nil {
			panic(fmt.Sprintf("Failed to parse amount for instruction '%s'", input[i]))
		}
		if err2 != nil {
			panic(fmt.Sprintf("Failed to parse from for instruction '%s'", input[i]))
		}
		if err3 != nil {
			panic(fmt.Sprintf("Failed to parse to for instruction '%s'", input[i]))
		}

		instructions[i] = instruction{amount: amount, from: from, to: to}
	}

	return instructions
}

func PerformInstructions(instructions []instruction, stacks []collections.Stack[string], preserveOrder bool) []collections.Stack[string] {
	for j, instruction := range instructions {
		cratesToMove := make([]string, instruction.amount)
		for i := 0; i < instruction.amount; i++ {
			value, ok := stacks[instruction.from-1].Pop()

			if !ok {
				fmt.Println(j)
				fmt.Println(stacks[instruction.from-1].Size())
				panic("could not pop stack")
			}
			if preserveOrder {
				cratesToMove[len(cratesToMove)-(i+1)] = value
			} else {
				cratesToMove[i] = value
			}
		}

		for _, crate := range cratesToMove {
			stacks[instruction.to-1].Push(crate)
		}
	}

	return stacks
}

func GetTheTopAfterInstructions(input string, preserveOrder bool) string {
	parts := strings.Split(input, "\n\n")

	partsStacks := strings.Split(parts[0], "\n")
	partsInstructions := strings.Split(parts[1], "\n")

	stacks := ParseStacks(partsStacks)
	instructions := ParseInstructions(partsInstructions)

	newStacks := PerformInstructions(instructions, stacks, preserveOrder)

	topCrates := ""

	for _, stack := range newStacks {
		crate, _ := stack.Peek()

		topCrates += crate
	}

	return topCrates
}

func GetInput() string {
	raw_input, err := helpers.DownloadInput(2022, 5)
	if err != nil {
		panic("Failed to fetch input")
	}
	raw_input = raw_input[:len(raw_input)-1]

	return raw_input
}

func Task1() string {
	input := GetInput()
	var product = GetTheTopAfterInstructions(input, false)
	return product
}

func Task2() string {
	input := GetInput()
	var product = GetTheTopAfterInstructions(input, true)
	return product
}
