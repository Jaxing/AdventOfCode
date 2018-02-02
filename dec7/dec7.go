package dec7

import (
	"log"
	"../helpers"
	"strings"
	"strconv"
	"fmt"
	"math"
)

const DIR_ROOT  = "dec7/input"

type program struct {
	name string
	weight int
	nextProgramNames []string
}

func parseProgram(prog string) program {
	splitByArrow := strings.Split(prog, " -> ")
	splitNameWeight := strings.Split(splitByArrow[0], " ")

	weight, _ := strconv.ParseInt(splitNameWeight[1][1:len(splitNameWeight[1])-1], 10, 0)

	var newProgram = program{name: splitNameWeight[0],
				weight: int(weight)}

	if len(splitByArrow) > 1 {
		listOfPrograms := strings.Split(splitByArrow[1], ", ")
		newProgram.nextProgramNames = listOfPrograms
	}

	return newProgram
}

func Part1() string {
	rows, err := helpers.ReadFile(DIR_ROOT)

	if err != nil {
		log.Fatal("file does not exist")
	}

	var programs = make([]program, 0, len(rows))
	var programNames = make(map[string]bool)

	for _, row := range rows {
		newProg := parseProgram(row)
		programs = append(programs, newProg)
		programNames[newProg.name] = true
		for _, progName := range newProg.nextProgramNames{
			programNames[progName] = true
		}
	}

	for _, prog := range programs {
		for _, names := range prog.nextProgramNames {
			delete(programNames, names)
		}
	}

	for k := range programNames {
		return k
	}

	return "sup"
}

func diffWeight(prog program, programs map[string]program) int {
	var subTowerWeights = make([]int, len(prog.nextProgramNames), len(prog.nextProgramNames))
	for i, name := range prog.nextProgramNames {
		tmpProg := programs[name]

		diff := diffWeight(tmpProg, programs)

		if diff > 0 {
			return diff
		}
		subTowerWeights[i] += subTowerWeight(tmpProg, programs)
	}

	var correctWeight = 0

	if len(subTowerWeights) > 0 {
		correctWeight = subTowerWeights[0]
	} else {
		return correctWeight
	}
	fmt.Println(prog.name, prog.nextProgramNames, subTowerWeights)
	for i, subTowerWeight := range subTowerWeights {
		if correctWeight != subTowerWeight {
			fmt.Println(subTowerWeights, i)
			if (i + 1 < len(subTowerWeights) && subTowerWeights[i+1] == correctWeight )  || subTowerWeights[i-1] == correctWeight {
				fmt.Println(programs[prog.nextProgramNames[i]].name, programs[prog.nextProgramNames[i]].weight, subTowerWeight)
				return programs[prog.nextProgramNames[i]].weight - int((math.Abs(float64(correctWeight - subTowerWeight))))
			}
			fmt.Println(programs[prog.nextProgramNames[0]].name, programs[prog.nextProgramNames[0]].weight, subTowerWeights[0])
			return programs[prog.nextProgramNames[0]].weight - int((math.Abs(float64(correctWeight - subTowerWeight))))

		}
	}

	return 0
}

func subTowerWeight(prog program, programs map[string]program) int {
	stw := prog.weight

	for _, subProgName := range prog.nextProgramNames {
		stw += subTowerWeight(programs[subProgName], programs)
	}
	return stw
}

func Part2() int {
	rows, err := helpers.ReadFile(DIR_ROOT)

	if err != nil {
		log.Fatal("file does not exist")
	}

	var programs = make(map[string]program)

	for _, row := range rows {
		var newProg = parseProgram(row)

		programs[newProg.name] = newProg
	}

	startProg := programs[Part1()]

	return diffWeight(startProg, programs)
}