package dec9

import (
	"../helpers"
	"log"
	"fmt"
)

const input = "dec9/input"

func Part1() int {
	var stream1, err = helpers.ReadLine(input)
	if err != nil {
		log.Fatal("file sucks")
	}

	var pointer = 0
	var openBrackets = 0
	var score = 0
	var ignore = false

	for pointer < len(stream1) {
		fmt.Println(string(stream1[pointer]), ignore)


		switch stream1[pointer] {
		case '<':
			if ignore{
				pointer++
				continue
			}
			fmt.Println("ignore group")
			ignore = true
		case '{':
			if ignore{
				pointer++
				continue
			}
			fmt.Println("open group")
			openBrackets++
		case '!':
			fmt.Println("ignore next")
			pointer++
		case '}':
			if ignore{
				pointer++
				continue
			}
			fmt.Println("close group", openBrackets)
			score += openBrackets
			openBrackets--
		case '>':
			ignore = false
		default:
			fmt.Println("default")
		}
		pointer++
	}
	return score
}

func Part2() int {
	var stream1, err = helpers.ReadLine(input)
	if err != nil {
		log.Fatal("file sucks")
	}

	var pointer = 0
	var count = 0
	var ignore = false

	for pointer < len(stream1) {
		fmt.Println(string(stream1[pointer]), ignore)


		switch stream1[pointer] {
		case '<':
			if ignore {
				pointer++
				count++
				continue
			}
			fmt.Println("ignore group")
			ignore = true
		case '{':
			if ignore {
				pointer++
				count++
				continue
			}
			fmt.Println("open group")
		case '!':
			fmt.Println("ignore next")
			pointer++
		case '}':
			if ignore {
				pointer++
				count++
				continue
			}
			fmt.Println("close group")
		case '>':
			ignore = false
		default:
			if ignore {
				count++
			}
			fmt.Println("default")
		}
		pointer++
	}
	return count
}