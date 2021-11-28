package dec4

import (
	"regexp"
	"strings"

	"../helpers"
)

var requiredFields [7]string = [7]string{
	"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid",
}
var validationRegex [7]string = [7]string{
	`byr:(19[2-9]\d|200[0-2])( |$|\n)`,
	`iyr:(201\d|2020)( |$|\n)`,
	`eyr:(202\d|2030)( |$|\n)`,
	`hgt:(59|6\d|7[0-6])in|(1[5-8]\d|19[0-3])cm( |$|\n)`,
	`hcl:#[\da-f]{6}( |$|\n)`,
	`ecl:amb|blu|brn|gry|grn|hzl|oth( |$|\n)`,
	`pid:\d{9}( |$|\n)`,
}

func ContainsAllFields(passport string) bool {
	for _, field := range requiredFields {
		if !strings.Contains(passport, field) {
			return false
		}
	}
	return true
}

func IsValidPassport(passport string) bool {
	for _, re := range validationRegex {
		matched, err := regexp.Match(re, []byte(passport))
		if err != nil || !matched {
			return false
		}
	}

	return true
}

func Task1() int {
	input, err := helpers.DownloadInput(2020, 4)

	if err != nil {
		panic("Failed to download input")
	}

	inputs := strings.Split(input, "\n\n")
	numValidPassport := 0

	for _, passport := range inputs {
		if ContainsAllFields(passport) {
			numValidPassport++
		}
	}

	return numValidPassport
}

func Task2() int {
	input, err := helpers.DownloadInput(2020, 4)

	if err != nil {
		panic("Failed to download input")
	}

	inputs := strings.Split(input, "\n\n")
	numValidPassport := 0

	for _, passport := range inputs {
		if IsValidPassport(passport) {
			numValidPassport++
		}
	}

	return numValidPassport
}
