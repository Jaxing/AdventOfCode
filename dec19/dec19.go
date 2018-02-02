package dec19

import (
	"../helpers"
	"log"
	"unicode"
)

const inputFile = "dec19/input"

func Part1() string {
	diagram, err := helpers.ReadFile(inputFile)

	if err != nil {
		log.Fatal("file does not exist")
	}
}

func pars(diagram []string) [][]string {

	var startIndex = 0

	for startIndex; startIndex < len(diagram[0]); startIndex++ {
		if diagram[0][startIndex] == '|' {
			break
		}
	}

	var letters := make([]string, 0, 100)

	row := 0
	col := startIndex
	direction := "Down"

	for {
		switch direction {
		case "Down":
			if row + 1 < len(diagram) {
				row++
			}
		case "Up":
			if row - 1 >= 0 {
				row--
			}
		case "Left":
			if col - 1 >= 0 {
				col--
			}
		case "Right":
			if col + 1 < len(diagram[0]) {
				col++
			}
		}

		if unicode.IsLetter(diagram[row][col]) {
			letters = append(letters, string(diagram[row][col]))
			break
		}

		if diagram[row][col] == '+' {
			
		}
	}
}