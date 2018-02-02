package dec8

import (
	"../helpers"
	"log"
	"strings"
	"strconv"
	"fmt"
	"math"
)

const DIR_ROOT  = "dec8/input"

func executeOrder(instruction string, register string, value int, registers map[string]int){
	switch instruction {
	case "inc":
		registers[register] += value
	case "dec":
		registers[register] -= value
	}
}

func Part1() int {
	rows, err := helpers.ReadFile(DIR_ROOT)

	if err != nil {
		log.Fatal("file does not exist")
	}

	var registers = make(map[string]int)

	for _, row := range rows {
		fmt.Println(row)
		splitStuff := strings.Split(row, " if ")

		instItems := strings.Split(splitStuff[0], " ")
		condItems := strings.Split(splitStuff[1], " ")

		switch condItems[1] {
		case ">":
			lol, _ := strconv.ParseInt(condItems[2], 10, 0)

			if registers[condItems[0]] > int(lol) {
				lol2, _ := strconv.ParseInt(instItems[2], 10, 0)

				executeOrder(instItems[1], instItems[0], int(lol2), registers)
			}
		case "<":
			lol, _ := strconv.ParseInt(condItems[2], 10, 0)

			if registers[condItems[0]] < int(lol) {
				lol2, _ := strconv.ParseInt(instItems[2], 10, 0)

				executeOrder(instItems[1], instItems[0], int(lol2), registers)
			}
		case ">=":
			lol, _ := strconv.ParseInt(condItems[2], 10, 0)

			if registers[condItems[0]] >= int(lol) {
				lol2, _ := strconv.ParseInt(instItems[2], 10, 0)

				executeOrder(instItems[1], instItems[0], int(lol2), registers)
			}
		case "<=":
			lol, _ := strconv.ParseInt(condItems[2], 10, 0)

			if registers[condItems[0]] <= int(lol) {
				lol2, _ := strconv.ParseInt(instItems[2], 10, 0)

				executeOrder(instItems[1], instItems[0], int(lol2), registers)
			}
		case "!=":
			lol, _ := strconv.ParseInt(condItems[2], 10, 0)

			if registers[condItems[0]] != int(lol) {
				lol2, _ := strconv.ParseInt(instItems[2], 10, 0)

				executeOrder(instItems[1], instItems[0], int(lol2), registers)
			}
		case "==":
			lol, _ := strconv.ParseInt(condItems[2], 10, 0)

			if registers[condItems[0]] == int(lol) {
				lol2, _ := strconv.ParseInt(instItems[2], 10, 0)

				executeOrder(instItems[1], instItems[0], int(lol2), registers)
			}

		}
	}

	max_value := 0
	stuff := ""

	for k, v := range registers {
		if v > max_value {
			max_value = v
			stuff = k
		}
	}
	fmt.Println(stuff)
	return max_value
}

func Part2() int {
	rows, err := helpers.ReadFile(DIR_ROOT)

	if err != nil {
		log.Fatal("file does not exist")
	}

	var registers = make(map[string]int)
	var maxReg = int(-math.MaxInt64)


	for _, row := range rows {
		splitStuff := strings.Split(row, " if ")

		instItems := strings.Split(splitStuff[0], " ")
		condItems := strings.Split(splitStuff[1], " ")

		switch condItems[1] {
		case ">":
			lol, _ := strconv.ParseInt(condItems[2], 10, 0)

			if registers[condItems[0]] > int(lol) {
				lol2, _ := strconv.ParseInt(instItems[2], 10, 0)

				executeOrder(instItems[1], instItems[0], int(lol2), registers)

				if registers[instItems[0]] > maxReg {
					maxReg = registers[instItems[0]]
				}
			}

		case "<":
			lol, _ := strconv.ParseInt(condItems[2], 10, 0)

			if registers[condItems[0]] < int(lol) {
				lol2, _ := strconv.ParseInt(instItems[2], 10, 0)

				executeOrder(instItems[1], instItems[0], int(lol2), registers)

				if registers[instItems[0]] > maxReg {
					maxReg = registers[instItems[0]]
				}
			}


		case ">=":
			lol, _ := strconv.ParseInt(condItems[2], 10, 0)

			if registers[condItems[0]] >= int(lol) {
				lol2, _ := strconv.ParseInt(instItems[2], 10, 0)

				executeOrder(instItems[1], instItems[0], int(lol2), registers)

				if registers[instItems[0]] > maxReg {
					maxReg = registers[instItems[0]]
				}
			}

		case "<=":
			lol, _ := strconv.ParseInt(condItems[2], 10, 0)

			if registers[condItems[0]] <= int(lol) {
				lol2, _ := strconv.ParseInt(instItems[2], 10, 0)

				executeOrder(instItems[1], instItems[0], int(lol2), registers)

				if registers[instItems[0]] > maxReg {
					maxReg = registers[instItems[0]]
				}
			}

		case "!=":
			lol, _ := strconv.ParseInt(condItems[2], 10, 0)

			if registers[condItems[0]] != int(lol) {
				lol2, _ := strconv.ParseInt(instItems[2], 10, 0)

				executeOrder(instItems[1], instItems[0], int(lol2), registers)

				if registers[instItems[0]] > maxReg {
					maxReg = registers[instItems[0]]
				}
			}

		case "==":
			lol, _ := strconv.ParseInt(condItems[2], 10, 0)

			if registers[condItems[0]] == int(lol) {
				lol2, _ := strconv.ParseInt(instItems[2], 10, 0)

				executeOrder(instItems[1], instItems[0], int(lol2), registers)

				if registers[instItems[0]] > maxReg {
					maxReg = registers[instItems[0]]
				}
			}
		}
	}
	return maxReg
}