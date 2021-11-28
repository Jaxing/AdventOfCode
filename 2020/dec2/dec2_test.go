package dec2

import (
	"testing"
)

func TestIsValidPassword(t *testing.T) {
	lowwerBound := 1
	upperBound := 3
	character := "a"
	password := "abcde"

	result := IsValidPassword(lowwerBound, upperBound, character, password)

	if !result {
		t.Errorf("%s is supposed to be valid", password)
	}

	lowwerBound = 1
	upperBound = 3
	character = "b"
	password = "cdefg"

	result = IsValidPassword(lowwerBound, upperBound, character, password)

	if result {
		t.Errorf("%s is supposed to be invalid", password)
	}
	lowwerBound = 2
	upperBound = 9
	character = "c"
	password = "ccccccccc"

	result = IsValidPassword(lowwerBound, upperBound, character, password)

	if !result {
		t.Errorf("%s is supposed to be valid", password)
	}
}
func TestIsValidPassword2(t *testing.T) {
	lowwerBound := 1
	upperBound := 3
	character := "a"
	password := "abcde"

	result := IsValidPassword2(lowwerBound, upperBound, character, password)

	if !result {
		t.Errorf("%s is supposed to be valid", password)
	}

	lowwerBound = 1
	upperBound = 3
	character = "b"
	password = "cdefg"

	result = IsValidPassword2(lowwerBound, upperBound, character, password)

	if result {
		t.Errorf("%s is supposed to be invalid", password)
	}
	lowwerBound = 2
	upperBound = 9
	character = "c"
	password = "ccccccccc"

	result = IsValidPassword2(lowwerBound, upperBound, character, password)

	if result {
		t.Errorf("%s is supposed to be invalid", password)
	}
}
