package main

import (
	"os"
	"./dec1"
	"./dec2"
	"./dec3"
	"./dec4"
	"./dec5"
	"./dec6"
	"./dec7"
	"./dec8"
	"./dec9"
	"./dec10"
	"./dec11"
	"./dec12"
	"./dec13"
	"./dec14"
	"./dec15"
	"./dec16"
	"./dec17"
	"./dec18"
	"fmt"
	"time"
)

type point struct {
	
}
func main() {
	start := time.Now()
	args := os.Args[1:]
	switch args[0] {
		case "dec1_1":
			fmt.Println(dec1.Part1())
		case "dec1_2":
			fmt.Println(dec1.Part2())
		case "dec2_1":
			fmt.Println(dec2.Part1())
		case "dec2_2":
			fmt.Println(dec2.Part2())
		case "dec3_1":
			fmt.Println(dec3.Part1())
		case "dec3_2":
			fmt.Println(dec3.Part2())
		case "dec4_1":
			fmt.Println(dec4.Part1())
		case "dec4_2":
			fmt.Println(dec4.Part2())
		case "dec5_1":
			fmt.Println(dec5.Part1())
		case "dec5_2":
			fmt.Println(dec5.Part2())
		case "dec6_1":
			fmt.Println(dec6.Part1())
		case "dec6_2":
			fmt.Println(dec6.Part2())
		case "dec7_1":
			fmt.Println(dec7.Part1())
		case "dec7_2":
			fmt.Println(dec7.Part2())
		case "dec8_1":
			fmt.Println(dec8.Part1())
		case "dec8_2":
			fmt.Println(dec8.Part2())
		case "dec9_1":
			fmt.Println(dec9.Part1())
		case "dec9_2":
			fmt.Println(dec9.Part2())
		case "dec10_1":
			fmt.Println(dec10.Part1())
		case "dec10_2":
			fmt.Println(dec10.Part2())
		case "dec11_1":
			fmt.Println(dec11.Part1())
		case "dec11_2":
			fmt.Println(dec11.Part2())
		case "dec12_1":
			fmt.Println(dec12.Part1())
		case "dec12_2":
			fmt.Println(dec12.Part2())
		case "dec13_1":
			fmt.Println(dec13.Part1())
		case "dec13_2":
			fmt.Println(dec13.Part2())
		case "dec14_1":
			fmt.Println(dec14.Part1())
		case "dec15_1":
			fmt.Println(dec15.Part1())
		case "dec15_2":
			fmt.Println(dec15.Part2())
		case "dec16_1":
			fmt.Println(dec16.Part1())
		case "dec16_2":
			fmt.Println(dec16.Part2())
		case "dec17_1":
			fmt.Println(dec17.Part1())
		case "dec17_2":
			fmt.Println(dec17.Part2())
		case "dec18_1":
			fmt.Println(dec18.Part1())
		case "dec18_2":
			fmt.Println(dec18.Part2())
		default:
			fmt.Println("No such file")

	}
	fmt.Println(time.Now().Sub(start))
}