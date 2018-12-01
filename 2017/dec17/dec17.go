package dec17


const stepSize  = 345

func insertAt(list []int, value int, index int) []int {
	list = append(list, list[len(list)-1])
	next := value
	for i := index; i < len(list); i++ {
		tmp := next
		next = list[i]
		list[i] = tmp
	}
	return list
}

func Part1() int {
	var buffer = make([]int, 1, 2018)
	var current  = 0
	var next  = 0


	for i:=1; i < 2018; i++ {
		next = ((current + stepSize) % len(buffer))
		//if i < 10 {
		//	fmt.Println(buffer, next, current, len(buffer))
		//}
		buffer = insertAt(buffer, i, next+1)
		current = next + 1
	}
	return buffer[next+2]
}

func Part2() int {
	var current  = 0
	var res  = 0


	for i:=1; i < 50000000; i++ {
		current = ((current + stepSize) % i) +1
		if current == 1 {
			res = i
		}
	}
	return res
}