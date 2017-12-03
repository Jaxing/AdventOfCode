package dec3

import (
	"math"
)


func Part1() int {
	const location int = 10
	const level = (location/8 + 1)
	const size = (level * 8) / 4 + 1

	var mem [size][size]int


	var r, c = size/2, size/2

	var value int = 1

	mem[r][c] = value

	//value++
	c++



	for lvl := 1 ; lvl <= level ; lvl++ {
		var sideLength = lvl * 2
		//for i:=0 ; i < size ; i++{
		//	for j:=0 ; j < size ; j++ {
		//		fmt.Print(mem[i][j], "\t")
		//	}
		//	fmt.Println("")
		//}
		//fmt.Println("================================")

		for k := 0 ; k < sideLength - 1 ; k++ {
			if value == location {
				return int(math.Abs(float64(size/2 - r)) + math.Abs(float64(size/2 - c)))
			}
			mem[r][c] = value
			value++
			r--
		}

		for k := 0; k < sideLength ; k++  {
			if value == location {
				return int(math.Abs(float64(size/2 - r)) + math.Abs(float64(size/2 - c)))
			}
			mem[r][c] = value
			value++
			c--
		}
		for k := 0 ; k < sideLength ; k++ {
			if value == location {
				return int(math.Abs(float64(size/2 - r)) + math.Abs(float64(size/2 - c)))
			}
			mem[r][c] = value
			value++
			r++
		}
		for k := 0 ; k <= sideLength ; k++ {
			if value == location {
				return int(math.Abs(float64(size/2 - r)) + math.Abs(float64(size/2 - c)))
			}
			mem[r][c] = value
			value++
			c++

		}
	}


	return int(math.Abs(float64(size/2 - r)) + math.Abs(float64(size/2 - c)))
}
