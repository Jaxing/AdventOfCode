package dec1

import (
	"sort"
	"advent_of_code/2022/helpers"
	"strings"
	"fmt"
	"strconv"
)

func ParseElfCalories(input string) ([]int) {
	var elfs []int
	for _, elf := range strings.Split(input, "\n\n") {
		var elf_calories = 0
		for _, calorie := range strings.Split(elf, "\n") {
			new_int, err := strconv.Atoi(calorie)
			// fmt.Println(calorie)
			if (err != nil) {
				panic("Bad input")
			}

			elf_calories = elf_calories + new_int
		} 

		elfs = append(elfs, elf_calories)
	}

	sort.Ints(elfs)

	return elfs 
}

func FindElfWithMostCalories(input string) (int){
	var elfs = ParseElfCalories(input)
	var most_calories = elfs[len(elfs) - 1]

	return most_calories
}

func FindTop3ElfsWithMostCalories(input string) (int){
	var elfs = ParseElfCalories(input)
	var one = elfs[len(elfs) - 1]
	var two = elfs[len(elfs) - 2]
	var three = elfs[len(elfs) - 3]

	return elfs[len(elfs) - 1] +  elfs[len(elfs) - 2]  + elfs[len(elfs) - 3]
}

func GetInput() (string){
	raw_input, err := helpers.DownloadInput(2022, 1)
	if err != nil {
		panic("Failed to fetch input")
	}
	return raw_input[:len(raw_input)-1]
}

func Task1() (int) {
	input := GetInput()
	var product = FindElfWithMostCalories(input)
	return product
}

func Task2() (int) {
	input := GetInput()
	var product = FindTop3ElfsWithMostCalories(input)
	return product
}
