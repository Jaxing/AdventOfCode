package dec4

import (
	"../helpers"
	"strings"
	"log"
)


func isAnagram(word1 string, word2 string) bool {
	var freq1 map[string]int = make(map[string]int)
	var freq2 map[string]int = make(map[string]int)

	if len(word1) != len(word2) {
		return false
	}

	for _, char := range word1 {
		freq1[string(char)]++
	}

	for _, char := range word2 {
		freq2[string(char)]++
	}

	for key := range freq1{
		if freq1[key] != freq2[key] {
			return false
		}
	}
	return true
}

func Part2() int {
	rows, err := helpers.ReadFile(DIR_ROOT)



	if err != nil {
		log.Fatal("file not exists")
	}
	validKeys := 0

	for _, row := range(rows) {
		var words []string = strings.Split(row, " ")
		var dup bool = false

		for i, word1 := range(words) {

			for j, word2 := range words {
				if i != j && isAnagram(word1, word2) {

					dup = true
					break
				}
			}
			if dup {
				break
			}
		}
		if ! dup {
			validKeys++
		}
	}

	return validKeys
}
