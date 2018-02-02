package dec15

import (
	//"time"
	//"fmt"
)

const genAStart = 65
const genBStart = 8921

func generate(prev int, factor int) int {
	tmp := (prev * factor) % 2147483647

	return tmp
}

func Part1() int {
	const divder = 2147483647
	const genAFactor  = 16807
	const genBFactor  = 48271

	genA := genAStart
	genB := genBStart
	count := 0

	for i := 0 ; i < 40000000 ; i++ {
		sixteenBits := 65535

		genA = generate(genA, genAFactor)
		genB = generate(genB, genBFactor)

		binaryA := genA & sixteenBits//fmt.Sprintf("%032b", genA)
		binaryB := genB & sixteenBits//fmt.Sprintf("%032b", genB)

		if binaryA == binaryB {
			count++
		}
	}

	return count
}

func Part2() int {
	const genAFactor  = 16807
	const genBFactor  = 48271
	const sixteenBits = 65535
	genA := genAStart
	genB := genBStart
	count := 0


	for i := 0; i < 5000000; i++ {
		//start := time.Now()
		aEven := false
		bEven := false

		for !aEven || !bEven {
			if !aEven {
				genA = generate(genA, genAFactor)
				aEven = 0 == genA % 4
			}
			if !bEven {
				genB = generate(genB, genBFactor)
				bEven = 0 == genB % 8
			}
		}

		binaryA := genA & sixteenBits
		binaryB := genB & sixteenBits

		if binaryA == binaryB {
			count++
		}
	}

	return count
}

