package dec16

import(
	"../helpers"
	"log"
	"strings"
	"strconv"
)

const inputFile = "dec16/input"

func Part1() string {
	input, err := helpers.ReadLine(inputFile)

	if err != nil {
		log.Fatal("nu such file")
	}

	var programs = "abcdefghijklmnop"

	moves := strings.Split(input, ",")

	for _, move := range moves {
		if move[0] == 's' {
			index, _ := strconv.ParseInt(string(move[1:]), 10, 0)
			programs = spin([]rune(programs), int(index))
		} else {
			indices := strings.Split(move[1:], "/")

			if move[0] == 'x' {
				index1, _ := strconv.ParseInt(indices[0], 10, 0)
				index2, _ := strconv.ParseInt(indices[1], 10, 0)

				programs = exchange([]rune(programs), int(index1), int(index2))
			} else {
				programs = partner([]rune(programs), indices[0], indices[1])
			}
		}
	}

	return programs
}

func spin(programs []rune, pivot int) string {
	part1 := programs[:len(programs) -pivot]
	part2 := programs[len(programs) - pivot:]

	return string(append(part2, part1...))
}

func exchange(programs []rune, index1 int, index2 int) string {
	tmp := programs[index1]
	programs[index1] = programs[index2]
	programs[index2] = tmp
	return string(programs)
}

func partner(programs []rune, prog1 string, prog2 string) string {
	index1 := 0
	index2 := 0

	for i, c := range programs{
		if string(c) == prog1 {
			index1 = i
		} else if string(c) == prog2{
			index2 = i
		}
	}

	return exchange(programs, index1, index2)
}

func Part2() string {
	input, err := helpers.ReadLine(inputFile)

	if err != nil {
		log.Fatal("nu such file")
	}

	var programs = "abcdefghijklmnop"

	moves := strings.Split(input, ",")
	nextState := make(map[string]string)

	for i := 0; i < 1000000000; i++ {
		preState := programs

		if nextState[programs] != "" {
			programs = nextState[programs]
			continue
		}
		for _, move := range moves {
			if move[0] == 's' {
				index, _ := strconv.ParseInt(string(move[1:]), 10, 0)
				programs = spin([]rune(programs), int(index))
			} else {
				indices := strings.Split(move[1:], "/")

				if move[0] == 'x' {
					index1, _ := strconv.ParseInt(indices[0], 10, 0)
					index2, _ := strconv.ParseInt(indices[1], 10, 0)

					programs = exchange([]rune(programs), int(index1), int(index2))
				} else {
					programs = partner([]rune(programs), indices[0], indices[1])
				}
			}
		}
		nextState[preState] = programs
	}

	return programs
}