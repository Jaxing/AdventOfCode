package dec14

import (
	"../dec10"
	"fmt"
	"strconv"
)

func Part1() int {
	var input = "flqrgnkx"

	count := 0
	for i := 0 ; i < 128; i++ {
		denseHex := dec10.DensHex(128, string(append([]rune(input), []rune(fmt.Sprintf("-%d", i))...)), 32, 4)
		formatedHex := make([]string, 32, 32)

		for j := 0; j < 32; j++ {
			if j + 1 < len(denseHex) {
				formatedHex[j] = string(append([]rune(""), rune(denseHex[j]), rune(denseHex[j+1])))
			}
		}

		for _, hex := range denseHex {
			hexInt, _ := strconv.ParseInt(fmt.Sprintf("%02x", hex), 16, 0)
			binary := fmt.Sprintf("%b", hexInt)
			if i == 0 {
				fmt.Print(binary)
			}
			for _, bit := range binary {
				if string(bit) == "1" {
					count++
				}
			}
			if i == 0 {
				fmt.Println(count)
			}
		}
	}

	return count
}