package dec11

import (
	"../helpers"
	"log"
	"strings"
	"math"
	"fmt"
)

const inputFile = "dec11/input"

func update(step string, x float64, y float64) (float64, float64) {
		switch step {
		case "n":
			y++
		case "s":
			y--
		case "w":
			x++
		case "e":
			x--
		case "ne":
			y += .5
			x -= .5
		case "nw":
			y += .5
			x += .5
		case "se":
			y -= .5
			x -= .5
		case "sw":
			y -= .5
			x += .5
		}
	return x, y
}

func Part1() int {
	input, err := helpers.ReadLine(inputFile)

	if err != nil {
		log.Fatal("nu such file")
	}

	steps := strings.Split(input, ",")
	
	x, y := 0.0, 0.0
	
	for _, step := range steps {
		x, y = update(step,x ,y)
	}

	fmt.Println(x, y)
	return  int(math.Abs(x) + math.Abs(y))
}

func Part2() int {
	input, err := helpers.ReadLine(inputFile)

	if err != nil {
		log.Fatal("nu such file")
	}

	steps := strings.Split(input, ",")

	x, y := 0.0, 0.0
	max := 0

	for _, step := range steps {
		x, y = update(step,x ,y)

		if int(math.Abs(x) + math.Abs(y)) > max {
			max = int(math.Abs(x) + math.Abs(y))
		}
	}

	fmt.Println(x, y)
	return  max
}