package dec3




type point struct {
	r int
	c int
}

func Part2() int {
	const location int = 325489
	const level = (location/8 + 1)
	const size = (level * 8) / 4 + 1



	var mem map[point]int = make(map[point]int)


	var r, c = size/2, size/2

	var value int = 1
	current := point{r, c}
	mem[current] = value

	//value++
	c++



	for lvl := 1 ; lvl <= level ; lvl++ {
		//fmt.Println(lvl)
		var sideLength = lvl * 2
		//for i:=0 ; i < size ; i++{
		//	for j:=0 ; j < size ; j++ {
		//		fmt.Print(mem[i][j], "\t")
		//	}
		//	fmt.Println("")
		//}
		//fmt.Println("================================")

		for k := 0 ; k < sideLength - 1 ; k++ {
			current = point{r, c}
			leftup, up, rightup, left, right, leftdown, down, rightdown :=
				point{r-1,c-1}, point{r-1,c}, point{r-1,c+1}, point{r, c-1},
				point{r, c+1}, point{r+1, c-1}, point{r+1,c}, point{r+1,c+1}
			val := 0
			val += mem[leftup]
			val += mem[up]
			val += mem[rightup]
			val += mem[left]
			val += mem[right]
			val += mem[leftdown]
			val += mem[down]
			val += mem[rightdown]

			if val > location {
				return val
			}
			mem[current] = val
			r--
		}

		for k := 0; k < sideLength ; k++  {
			current = point{r, c}
			leftup, up, rightup, left, right, leftdown, down, rightdown :=
				point{r-1,c-1}, point{r-1,c}, point{r-1,c+1}, point{r, c-1},
				point{r, c+1}, point{r+1, c-1}, point{r+1,c}, point{r+1,c+1}
			val := 0
			val += mem[leftup]
			val += mem[up]
			val += mem[rightup]
			val += mem[left]
			val += mem[right]
			val += mem[leftdown]
			val += mem[down]
			val += mem[rightdown]
			if val > location {
				return val
			}
			mem[current] = val
			c--
		}
		for k := 0 ; k < sideLength ; k++ {
			current = point{r, c}
			leftup, up, rightup, left, right, leftdown, down, rightdown :=
				point{r-1,c-1}, point{r-1,c}, point{r-1,c+1}, point{r, c-1},
				point{r, c+1}, point{r+1, c-1}, point{r+1,c}, point{r+1,c+1}
			val := 0
			val += mem[leftup]
			val += mem[up]
			val += mem[rightup]
			val += mem[left]
			val += mem[right]
			val += mem[leftdown]
			val += mem[down]
			val += mem[rightdown]
			if val > location {
				return val
			}
			mem[current] = val
			r++
		}
		for k := 0 ; k <= sideLength ; k++ {
			current = point{r, c}
			leftup, up, rightup, left, right, leftdown, down, rightdown :=
				point{r-1,c-1}, point{r-1,c}, point{r-1,c+1}, point{r, c-1},
				point{r, c+1}, point{r+1, c-1}, point{r+1,c}, point{r+1,c+1}
			val := 0
			val += mem[leftup]
			val += mem[up]
			val += mem[rightup]
			val += mem[left]
			val += mem[right]
			val += mem[leftdown]
			val += mem[down]
			val += mem[rightdown]
			if val > location {
				return val
			}
			mem[current] = val
			c++

		}
	}
	current = point{r, c}
	leftup, up, rightup, left, right, leftdown, down, rightdown :=
		point{r-1,c-1}, point{r-1,c}, point{r-1,c+1}, point{r, c-1},
		point{r, c+1}, point{r+1, c-1}, point{r+1,c}, point{r+1,c+1}
	val := 0
	val += mem[leftup]
	val += mem[up]
	val += mem[rightup]
	val += mem[left]
	val += mem[right]
	val += mem[leftdown]
	val += mem[down]
	val += mem[rightdown]

	return val
}
