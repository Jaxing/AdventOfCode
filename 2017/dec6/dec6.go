package dec6

import(

)
import (
	"log"
	"strings"
	"../helpers"
	"math"
	"strconv"
	"fmt"
)

const DIR_ROOT = "dec6/input"

type blockStruct struct {
	blocks []int
}

func (blocks blockStruct) toString() string{
	var str = ""

	for _, block := range blocks.blocks {
		str += string(block)
	}

	return str
}

func Part1() int {
	rows, err := helpers.ReadFile(DIR_ROOT)

	if err != nil {
		log.Fatal("file does not exist")
	}

	var intBlocks = make([]int, 0, 10)

	for _, block := range strings.Split(rows[0], "\t") {
		parsedInt, _ := strconv.ParseInt(block, 10, 0)
		intBlocks = append(intBlocks, int(parsedInt))
	}
	fmt.Println(intBlocks)
	blocks  := blockStruct{intBlocks}

	var blockCombos = make(map[string]bool)
	var iteration = 0

	for true  {
		max_index := 0
		max_block := 0

		for i, block := range blocks.blocks {
			if block > max_block {
				max_index = i
				max_block = block
			}
		}

		blockSplit := int(math.Ceil(float64(max_block) / float64(len(blocks.blocks))))
		blocks.blocks[max_index] = 0

		for i := max_index + 1 ; max_block != 0 ; i++ {
			if i >= len(blocks.blocks) {
				i = 0
			}

			if max_block >= blockSplit {
				blocks.blocks[i] += blockSplit
				max_block -= blockSplit
			} else {
				blocks.blocks[i] += max_block
				max_block = 0
			}
		}
		fmt.Println(blocks.blocks, blocks.toString())
		if blockCombos[blocks.toString()] {
			iteration++
			return iteration
		}

		iteration++
		blockCombos[blocks.toString()] = true

	}

	return iteration
}

func Part2() int {
	rows, err := helpers.ReadFile(DIR_ROOT)

	if err != nil {
		log.Fatal("file does not exist")
	}

	var intBlocks = make([]int, 0, 10)

	for _, block := range strings.Split(rows[0], "\t") {
		parsedInt, _ := strconv.ParseInt(block, 10, 0)
		intBlocks = append(intBlocks, int(parsedInt))
	}
	fmt.Println(intBlocks)
	blocks := blockStruct{intBlocks}

	var blockCombos = make(map[string]int)
	var iteration = 0

	for blockCombos[blocks.toString()] <= 0 {
		max_index := 0
		max_block := 0

		for i, block := range blocks.blocks {
			if block > max_block {
				max_index = i
				max_block = block
			}
		}

		blockSplit := int(math.Ceil(float64(max_block) / float64(len(blocks.blocks))))
		blocks.blocks[max_index] = 0

		for i := max_index + 1; max_block != 0; i++ {
			if i >= len(blocks.blocks) {
				i = 0
			}

			if max_block >= blockSplit {
				blocks.blocks[i] += blockSplit
				max_block -= blockSplit
			} else {
				blocks.blocks[i] += max_block
				max_block = 0
			}
		}
		fmt.Println(blocks.blocks, blocks.toString())
		if blockCombos[blocks.toString()] > 0 {
			iteration++
			return iteration - blockCombos[blocks.toString()]
		}

		iteration++
		blockCombos[blocks.toString()] = iteration

	}
	return iteration - blockCombos[blocks.toString()]
}