package dec18

import (
	"../helpers"
	"log"
	"strings"
	"fmt"
	"strconv"
)

const inputFile  = "dec18/input"

func Part1() int {
	instruction, err := helpers.ReadFile(inputFile)

	if err != nil {
		log.Fatal("file does not exist")
	}

	var currentInstruction = 0
	var registers = make(map[string]int)
	var lastPlayed = 0

	for currentInstruction < len(instruction) {
		parts := strings.Split(instruction[currentInstruction], " ")

		switch parts[0] {
		case "snd":
			lastPlayed = registers[parts[1]]
			fmt.Println(registers[parts[1]])
		case "set":
			parsedInt, err := strconv.ParseInt(parts[2], 10, 0)
			if err == nil {
				registers[parts[1]] = int(parsedInt)
			} else {
				registers[parts[1]] = registers[parts[2]]
			}
		case "add":
			parsedInt, err := strconv.ParseInt(parts[2], 10, 0)
			if err == nil {
				registers[parts[1]] += int(parsedInt)
			} else {
				registers[parts[1]] += registers[parts[2]]
			}
		case "mul":
			parsedInt, err := strconv.ParseInt(parts[2], 10, 0)
			if err == nil {
				registers[parts[1]] *= int(parsedInt)
			} else {
				registers[parts[1]] *= registers[parts[2]]
			}
		case "mod":
			parsedInt, err := strconv.ParseInt(parts[2], 10, 0)
			if err == nil {
				registers[parts[1]] %= int(parsedInt)
			} else {
				registers[parts[1]] %= registers[parts[2]]
			}
		case "rcv":
			if registers[parts[1]] != 0 {
				return lastPlayed
			}
		case "jgz":
			if registers[parts[1]] != 0 {
				parsedInt, err := strconv.ParseInt(parts[2], 10, 0)
				if err == nil {
					currentInstruction += int(parsedInt)
				} else {
					currentInstruction += registers[parts[2]]
				}
				continue
			}
		}
		currentInstruction++
	}
	return 0
}

func parse(parts []string, registers map[string]int) int {
	parsedInt, err := strconv.ParseInt(parts[len(parts)-1], 10, 0)
	if err == nil {
		return int(parsedInt)
	} else {
		return registers[parts[len(parts)-1]]
	}
}

func goRoutine(instructions []string, registers map[string]int, write chan<- int, read <-chan int, syncWithMain chan<- string) {
	currentInstruction := 0
	prog := registers["p"]
	var queue = new(Queue)

	for currentInstruction < len(instructions) {
		parts := strings.Split(instructions[currentInstruction], " ")
		fmt.Println(queue.head)
		if queue.head != nil {
			select {
			case write <- queue.head.value:
				syncWithMain <- fmt.Sprintf("s%d", prog)
				queue.pop()
			default:
				break
			}
		}

		switch parts[0] {
		case "snd":
			queue.push(parse(parts, registers))
		case "set":
			registers[parts[1]] = parse(parts, registers)
		case "add":
			registers[parts[1]] += parse(parts, registers)
		case "mul":
			registers[parts[1]] *= parse(parts, registers)
		case "mod":
			registers[parts[1]] %= parse(parts, registers)
		case "rcv":
			select {
			case data := <- read:
				syncWithMain <- fmt.Sprintf("r%d", prog)
				registers[parts[1]] = data
			default:
				continue
			}
		case "jgz":
			if registers[parts[1]] > 0 {
				parsedInt, err := strconv.ParseInt(parts[2], 10, 0)
				if err == nil {
					currentInstruction += int(parsedInt)
				} else {
					currentInstruction += registers[parts[2]]
				}
				continue
			}
		}
		currentInstruction++
	}
}

func Part2() int {
	instruction, err := helpers.ReadFile(inputFile)

	if err != nil {
		log.Fatal("file does not exist")
	}

	var zeroToOne = make(chan int)
	var oneToZero = make(chan int)
	var syncWithMain = make(chan string)
	var registerZero = make(map[string]int)
	var registerOne = make(map[string]int)
	registerOne["p"] = 1
	count := 0
	//nbrWaits := 0

	go goRoutine(instruction, registerZero, zeroToOne, oneToZero, syncWithMain)
	go goRoutine(instruction, registerOne, oneToZero, zeroToOne, syncWithMain)

	for v := range syncWithMain {
		switch v {
		case "s1":
			count++
			fmt.Println("Prog1 sends", count)
		case "s0":
			fmt.Println("Prog0 sends")
		case "r1":
			fmt.Println("Prog1 revices")
		case "r0":
			fmt.Println("Prog0 revices")
		}
	}

	return count
}

type Queue struct {
	head *Item
	tail *Item
}

type Item struct {
	value int
	next *Item
	prev *Item
}

func (q *Queue) push(value int) {
	newItem := Item{value: value, next:nil, prev: q.tail}


	if q.head == nil {
		q.head = &newItem
	} else {
		q.tail.next = &newItem
	}
	q.tail = &newItem
}

func (q *Queue) pop() int {
	tmp := q.head.value
	q.head = q.head.next
	return tmp
}