package dec4

import (
	"../helpers"
	"strings"
	"log"
)

const DIR_ROOT string = "dec4/input1"

func Part1() int {
	rows, err := helpers.ReadFile(DIR_ROOT)



	if err != nil {
		log.Fatal("file not exists")
	}
	validKeys := 0

	for _, row := range(rows) {
		var words []string = strings.Split(row, " ")
		var freq map[string]int = make(map[string]int)
		var dup bool = false

		for _, word := range(words) {
			freq[string(word)]++

			if freq[string(word)] > 1 {
				dup = true
				break
			}
		}

		if ! dup {
			validKeys++
		}
	}

	return validKeys
}
