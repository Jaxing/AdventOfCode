package dec2

import (
	"log"
	"strings"
	"strconv"
	"../helpers"
)

func Part2() int {
	rows, err := helpers.ReadFile(DIR_ROOT)
	var sum int = 0

	if err != nil {
		log.Fatal(err)
	}

	for _, row := range(rows){
		var items []string = strings.Split(row, " ")

		for j := 0; j < len(items); j++ {
			for k := 0; k < len(items); k++ {
				item1, intErr1 := strconv.ParseInt(items[j], 10, 0)
				item2, intErr2 := strconv.ParseInt(items[k], 10, 0)

				if intErr1 != nil || intErr2 != nil {
					log.Fatal("Not an int")
				}

				if j != k && item1 % item2 == 0 {
					sum += int(item1) / int(item2)
				}
			}
		}
	}
	return sum
}
