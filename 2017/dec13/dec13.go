package dec13

import (
	"../helpers"
	"log"
	"strconv"
	"strings"
)

const fileName = "dec13/input"

func penaltyRun(firewalls []int, currentScanPlace []int, moveDown []bool) int {
	var penalty = 0

	for i := range firewalls {
		if currentScanPlace[i] == 0 {
			penalty += i * firewalls[i]
		}

		passOnePico(firewalls, currentScanPlace, moveDown)
	}

	return penalty
}

func passOnePico(firewalls []int, currentScanPlace []int, moveDown []bool) {
	for j := range firewalls {
			if moveDown[j] {
				if currentScanPlace[j] + 1 < firewalls[j] {
					currentScanPlace[j]++
				} else {
					currentScanPlace[j]--
					moveDown[j] = false
				}
			} else {
				if currentScanPlace[j] > 0 {
					currentScanPlace[j]--
				} else  {
					currentScanPlace[j]++
					moveDown[j] = true
				}
			}
		}
}

func Part1() int {
	rows, err := helpers.ReadFile(fileName)

	if err != nil {
		log.Fatal("file does not exist")
	}

	var firewalls = make([]int, 100, 100)
	var currentScanPlace = make([]int, 100, 100)
	var moveDown = make([]bool, 100, 100)

	for _, row := range rows {
		parts := strings.Split(row, ": ")
		part1, _ := strconv.ParseInt(parts[0], 10, 0)
		part2, _ := strconv.ParseInt(parts[1], 10, 0)

		firewalls[int(part1)] = int(part2)
		currentScanPlace[int(part1)] = 0
		moveDown[int(part1)] = true
	}

	return penaltyRun(firewalls, currentScanPlace, moveDown)
}

func Part2() int {
	rows, err := helpers.ReadFile(fileName)

	if err != nil {
		log.Fatal("file does not exist")
	}

	var firewalls = make([]int, 100, 100)
	var currentScanPlace = make([]int, 100, 100)
	var moveDown = make([]bool, 100, 100)

	for _, row := range rows {
		parts := strings.Split(row, ": ")
		part1, _ := strconv.ParseInt(parts[0], 10, 0)
		part2, _ := strconv.ParseInt(parts[1], 10, 0)

		firewalls[int(part1)] = int(part2)
		currentScanPlace[int(part1)] = 0
		moveDown[int(part1)] = true
	}

	var picoSec = 0
	var copyFirewalls = make([]int, 100, 100)
	var copyCurrentScanPlace = make([]int, 100, 100)
	var copyMoveDown = make([]bool, 100, 100)

	//passOnePico(firewalls, currentScanPlace, moveDown)

	var penalty = 1

	for penalty > 0 {
		picoSec++
		passOnePico(firewalls, currentScanPlace, moveDown)
		copy(copyFirewalls, firewalls)
		copy(copyCurrentScanPlace, currentScanPlace)
		copy(copyMoveDown, moveDown)

		if copyCurrentScanPlace[0] == 0 || (copyCurrentScanPlace[1] == 1 && !moveDown[1]){
			continue
		}

		penalty = penaltyRun(copyFirewalls, copyCurrentScanPlace, copyMoveDown)
	}

	return picoSec
}