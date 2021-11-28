package dec4

import (
	"strings"
	"testing"
)

func TestIsValidPassword(t *testing.T) {
	input := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

	inputs := strings.Split(input, "\n\n")
	validityOfPassport := []bool{true, false, true, false}

	for i, passport := range inputs {
		if validityOfPassport[i] != ContainsAllFields(passport) {
			t.Errorf("Passport number %d was %v but expected %v", i, !validityOfPassport[i], validityOfPassport[i])
		}
	}
}

// func TestParseFields(t *testing.T) {
// 	input := "ecl:gry pid:860033327\neyr:2020 hcl:#fffffd"
// 	result := ParseFields(input, "ecl")
// 	if result != "gry" {
// 		t.Errorf("Expected gry, but found '%v'", result)
// 	}

// 	result = ParseFields(input, "pid")
// 	if result != "860033327" {
// 		t.Errorf("Expected 860033327, but found '%v'", result)
// 	}

// 	result = ParseFields(input, "hcl")
// 	if result != "#fffffd" {
// 		t.Errorf("Expected #fffffd, but found '%v'", result)
// 	}

// 	result = ParseFields(input, "eyr")
// 	if result != "2020" {
// 		t.Errorf("Expected 2020, but found '%v'", result)
// 	}
// }

func TestIsValidPassport(t *testing.T) {
	invalidInput := `eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007`

	for _, passport := range strings.Split(invalidInput, "\n\n") {
		if IsValidPassport(passport) {
			t.Errorf("Passport %v is supposed to be invalid", passport)
		}
	}

	validInput := `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`

	for _, passport := range strings.Split(validInput, "\n\n") {
		if !IsValidPassport(passport) {
			t.Errorf("Passport %v is supposed to be valid", passport)
		}
	}
}
