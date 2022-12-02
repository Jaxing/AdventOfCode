package dec2

import (
	"advent_of_code/2022/helpers"
	"fmt"
	"strings"
)

func LetterToValue(letter string) int {
	switch letter {
	case "A", "X":
		return 1
	case "B", "Y":
		return 2
	case "C", "Z":
		return 3
	default:
		panic("Invalid letter")
	}
}

func CalculateScore(input []string) int {
	var total_score = 0

	for _, round := range input {
		var parts = strings.Split(round, " ")
		// fmt.Println(parts)
		var opponent = parts[0]
		var me = parts[1]

		my_value := LetterToValue(me)
		opponent_value := LetterToValue(opponent)
		total_score = total_score + my_value

		switch my_value - opponent_value {
		case 0:
			// fmt.Println("Draw")
			total_score = total_score + 3
		case 1, -2:
			// fmt.Println("Win")
			total_score = total_score + 6
		case -1, 2:
			// fmt.Println("Loss")
		}
	}

	return total_score
}

func CalculateScore2(input []string) int {
	var total_score = 0

	for _, round := range input {
		var parts = strings.Split(round, " ")
		fmt.Println(parts)
		var opponent = parts[0]
		var outcome = parts[1]
		fmt.Println(outcome)
		var opponent_value = LetterToValue(opponent)
		var my_value = 0

		switch outcome {
		case "X":
			fmt.Println("Lose")
			switch opponent_value {
			case 1:
				my_value = 3
			case 2:
				my_value = 1
			case 3:
				my_value = 2
			}
			total_score = total_score + my_value
			fmt.Println(total_score)
		case "Y":
			fmt.Println("Draw")
			my_value = opponent_value
			total_score = total_score + my_value + 3
			fmt.Println(total_score)
		case "Z":
			fmt.Println("Win")
			my_value = (opponent_value % 3) + 1
			total_score = total_score + my_value + 6
			fmt.Println(total_score)
		}
	}

	return total_score
}

func GetInput() []string {
	raw_input, err := helpers.DownloadInput(2022, 2)
	if err != nil {
		panic("Failed to fetch input")
	}
	raw_input = raw_input[:len(raw_input)-1]

	return strings.Split(raw_input, "\n")
}

func Task1() int {
	input := GetInput()
	var product = CalculateScore(input)
	return product
}

func Task2() int {
	input := GetInput()
	var product = CalculateScore2(input)
	return product
}
