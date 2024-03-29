package main

import (
	"fmt"
	"os"
	"time"

	"advent_of_code/2022/dec1"
	"advent_of_code/2022/dec2"
	"advent_of_code/2022/dec3"
	"advent_of_code/2022/dec4"
	"advent_of_code/2022/dec5"
	"advent_of_code/2022/dec6"
	"advent_of_code/2022/dec7"
	"advent_of_code/2022/dec8"
	"advent_of_code/2022/dec9"
)

func main() {
	start := time.Now()
	args := os.Args[1:]
	switch args[0] {
	case "dec1_1":
		fmt.Println(dec1.Task1())
	case "dec1_2":
		fmt.Println(dec1.Task2())
	case "dec2_1":
		fmt.Println(dec2.Task1())
	case "dec2_2":
		fmt.Println(dec2.Task2())
	case "dec3_1":
		var res = dec3.Task1()
		fmt.Println(res)
		// helpers.Submit(2022, 3, 1, res)
	case "dec3_2":
		fmt.Println(dec3.Task2())
	case "dec4_1":
		fmt.Println(dec4.Task1())
	case "dec4_2":
		fmt.Println(dec4.Task2())
	case "dec5_1":
		fmt.Println(dec5.Task1())
	case "dec5_2":
		fmt.Println(dec5.Task2())
	case "dec6_1":
		fmt.Println(dec6.Task1())
	case "dec6_2":
		fmt.Println(dec6.Task2())
	case "dec7_1":
		fmt.Println(dec7.Task1())
	case "dec7_2":
		fmt.Println(dec7.Task2())
	case "dec8_1":
		fmt.Println(dec8.Task1())
	case "dec8_2":
		fmt.Println(dec8.Task2())
	case "dec9_1":
		fmt.Println(dec9.Task1())
	case "dec9_2":
		fmt.Println(dec9.Task2())
	default:
		fmt.Println("No such file")

	}
	fmt.Println(time.Now().Sub(start))
}
