package dec2

import (
	"fmt"
	"strconv"
	"strings"

	"../helpers"
)

func IsValidPassword(lowerBound int, upperBound int, character string, password string) bool {
	if len(password) < lowerBound {
		return false
	}

	res := strings.Split(password, character)
	return len(res)-1 >= lowerBound && len(res)-1 <= upperBound
}

func IsValidPassword2(firstPosition int, secondPosition int, character string, password string) bool {
	return string(password[firstPosition-1]) == character && string(password[secondPosition-1]) != character || string(password[secondPosition-1]) == character && string(password[firstPosition-1]) != character
}

func GetInput() []string {
	raw_input, err := helpers.DownloadInput(2020, 2)
	if err != nil {
		panic("Failed to fetch input")
	}

	stringList := strings.Split(raw_input, "\n")
	stringList = stringList[:len(stringList)-1]

	return stringList
}

func GetPartsOfPasswordPolicy(input string) (int, int, string, string) {
	parts := strings.Split(input, "-")
	// fmt.Printf("%v\n", parts)
	lowerBound, err1 := strconv.Atoi(parts[0])
	parts = strings.Split(parts[1], " ")
	upperBound, err2 := strconv.Atoi(parts[0])
	character := strings.Split(parts[1], ":")[0]
	password := parts[2]

	if err1 != nil || err2 != nil {
		panic(fmt.Sprintf("Failed to parse input, err1: %s, err2: %s", err1, err2))
	}

	return lowerBound, upperBound, character, password
}

func Task1() int {
	input := GetInput()
	numValid := 0

	for _, in := range input {
		lowerBound, upperBound, character, password := GetPartsOfPasswordPolicy(in)
		if IsValidPassword(lowerBound, upperBound, character, password) {
			numValid++
		}
	}
	return numValid
}


func Task2() int {
	input := GetInput()
	numValid := 0

	for _, in := range input {
		lowerBound, upperBound, character, password := GetPartsOfPasswordPolicy(in)
		if IsValidPassword2(lowerBound, upperBound, character, password) {
			numValid++
		}
	}
	return numValid
}
