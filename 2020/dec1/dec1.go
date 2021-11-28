package dec1

import (
	"sort"
	"../helpers"
	"strings"
	// "fmt"
)

func GetLastRelevantIndex(input []int, sum int) (int) {
	sort.SliceStable(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	i := len(input) - 1
	for ;i >= 0; i-- {
		if(input[i] <= sum) {
			break
		}
	}

	return i
}

func findTwoEntriesProduct(input []int, sum int) (int){
	i := GetLastRelevantIndex(input, sum)

	var small_pointer, larger_point = 0, i

	for input[small_pointer] + input[larger_point] != sum {
		if (input[small_pointer] + input[larger_point] > sum) {
			larger_point--
		} else {
			small_pointer++
		}
	}

	return input[small_pointer] * input[larger_point]
}

func findThreeEntriesProduct(input []int, sum int) (int){
	i := GetLastRelevantIndex(input, sum)

	for j:=0; j < i; j++ {
		for k:=j; k < i; k++ {
			for m:=k; m < i; m++ {
				if (input[j] + input[k] + input[m] == sum) {
					return input[j] * input[k] * input[m]
				}
			}
		}
	}

	return -1
}

func GetInput() ([]int){
	raw_input, err := helpers.DownloadInput(2020, 1)
	if err != nil {
		panic("Failed to fetch input")
	}

	stringList := strings.Split(raw_input, "\n")
	stringList = stringList[:len(stringList)-1]
	input, err := helpers.StringListToIntList(stringList)
	if err != nil {
		panic("Failed to parse input")
	}

	return input
}

func Task1() (int) {
	input := GetInput()

	var product = findTwoEntriesProduct(input[:], 2020) 
	return product
}

func Task2() (int) {
	input := GetInput()

	var product = findThreeEntriesProduct(input[:], 2020) 
	return product
}