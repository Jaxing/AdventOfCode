package dec2

import (
	"log"
	"math"
	"strconv"
	"../helpers"
	"strings"
)

const DIR_ROOT string = "dec2/input1"

func Part1() int {
	rows, err := helpers.ReadFile(DIR_ROOT)
	var sum int = 0

	if err != nil {
		log.Fatal(err)
	}

	for _, row := range(rows){
		var min int = math.MaxInt64
		var max int = -math.MaxInt64
		var items []string = strings.Split(row, " ")

		for _, item := range(items) {
			itemInt, intErr := strconv.ParseInt(item, 10, 0)

			if intErr != nil {
				log.Fatal("Not an int")
			}
			max = int(math.Max(float64(max), float64(itemInt)))
			min = int(math.Min(float64(min), float64(itemInt)))
		}
		sum += max - min
	}
	return sum
}



