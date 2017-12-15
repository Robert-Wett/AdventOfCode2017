package main

import (
	"container/ring"
	"fmt"
	"strconv"

	utils "github.com/Robert-Wett/AdventOfCode2017/helpers"
)

func main() {
	input := utils.GetInput("./input.txt")

	partOne(input)
	partTwo(input)
	partTwoRingBuffer(input)
}

func partOne(input string) {
	var total int
	var last = input[len(input)-1:]
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			toAdd, _ := strconv.Atoi(string(input[i]))
			total += toAdd
		}
	}
	if string(input[0]) == last {
		toAdd, _ := strconv.Atoi(last)
		total += toAdd
	}

	fmt.Println(total)
}

func partTwo(input string) {
	var total int
	var end = len(input)
	step := end / 2

	for i := 0; i < end; i++ {
		c, _ := strconv.Atoi(string(input[i]))
		n, _ := strconv.Atoi(string(input[(i+step)%end]))
		if c == n {
			total += c
		}
	}
	fmt.Println(total)
}

func partTwoRingBuffer(input string) {
	var total int
	step := len(input) / 2
	r := ring.New(len(input))

	for i := 0; i < r.Len(); i++ {
		v, _ := strconv.Atoi(string(input[i]))
		r = r.Next()
		r.Value = v
	}

	for i := 0; i < len(input); i++ {
		cur := r.Value.(int)
		r = r.Move(step)
		next := r.Value.(int)
		if cur == next {
			total += cur
		}
		r = r.Move(-(step + 1))
	}
	fmt.Println(total)
}
