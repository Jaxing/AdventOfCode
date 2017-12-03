package main

import (
	"os"
	"./dec1"
	"./dec2"
	"./dec3"
	"fmt"
)

type point struct {
	
}
func main() {
	args := os.Args[1:]
	switch args[0] {
		case "dec1_1":
			fmt.Println(dec1.Part1())
		case "dec1_2":
			fmt.Println(dec1.Part2())
		case "dec2_1":
			fmt.Println(dec2.Part1())
		case "dec3_1":
			fmt.Println(dec3.Part1())
		case "dec3_2":
			fmt.Println(dec3.Part2())
		default:
			fmt.Println("No such file")

	}
}