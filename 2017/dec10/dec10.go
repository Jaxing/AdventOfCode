package dec10

import (
	"../helpers"
	"strings"
	"log"
	"strconv"
	"fmt"
	"math"
)

const inputFile  = "dec10/input"

func reverse(list []int, from int, to int) {
	if to >= len(list) {
		to = to % len(list)
	}
	if from >= len(list) {
		from = from % len(list)
	}

	var tmpSlice []int
	if from > to {
		tmpSlice = append(list[from:], list[:to]...)
	} else {
		tmpSlice = list[from:to]
	}

	for i := 0; i < int(math.Ceil(float64(len(tmpSlice)/2))) ;  i++{
		tmpSlice[i], tmpSlice[len(tmpSlice)-1-i] = tmpSlice[len(tmpSlice)-1-i], tmpSlice[i]
	}

	for i,sliceElm := range tmpSlice {
		list[(from+i) % len(list)] = sliceElm
		//fmt.Println(int(math.Mod(float64(from+i), float64(len(list)))), sliceElm)
	}
}

func Part1() int {
	input, err := helpers.ReadLine(inputFile)

	if err != nil {
		log.Fatal("nu such file")
	}

	lengths := strings.Split(input, ",")

	var list = make([]int, 256)
	var current = 0
	var skipSize  = 0

	var val = 0
	for i := 0 ; i < len(list) ; i++ {
		list[i] = val
		val++
	}
	for _, length := range lengths {
		lengthInt, _ := strconv.ParseInt(length, 10, 0)
		reverse(list, current, current + int(lengthInt))
		current += int(lengthInt) + skipSize
		skipSize++
	}
	return list[0] * list[1]
}

func Part2() string {
	input, err := helpers.ReadLine(inputFile)

	if err != nil {
		log.Fatal("nu such file")
	}

	return DensHex(256, input, 16, 16)
}

func DensHex(listLength int, input string, numberOfHex int, hexToBit int) string {
	lengths := append([]byte(input), []byte{17, 31, 73, 47, 23}...)
	var list = make([]int, listLength)
	var current = 0
	var skipSize  = 0

	var val = 0
	for i := 0 ; i < len(list) ; i++ {
		list[i] = val
		val++
	}
	for i := 0; i < 64; i++ {
		for _, length := range lengths{
			reverse(list, current, current + int(length))
			current += int(length) + skipSize
			skipSize++
		}
	}

	var dense = make([]int, numberOfHex, numberOfHex)

	for i:= 0; i < numberOfHex; i++ {
		res := 0
		for j := i*hexToBit; j < (i+1)*hexToBit; j++ {
			res = res ^ list[j]
		}
		dense[i] = res
	}

	var denseHex = make([]string, numberOfHex, numberOfHex)

	for i, hex := range dense {
		denseHex[i] = fmt.Sprintf("%02x", hex)
	}

	return strings.Join(denseHex, "")
}