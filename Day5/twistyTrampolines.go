package main

import (
	"fmt"
	"strconv"
	"strings"

	utils "github.com/Robert-Wett/AdventOfCode2017/helpers"
)

func main() {
	input := utils.GetInput("./input.txt")
	var m []int
	for _, s := range strings.Split(input, "\n") {
		v, _ := strconv.Atoi(s)
		m = append(m, v)
	}

	partOne(input)
	partTwo(input)
}

/*
Now, the jumps are even stranger: after each jump, if the offset was three or more,
instead decrease it by 1. Otherwise, increase it by 1 as before.
*/
func partTwo(input string) {
	var m []int
	for _, s := range strings.Split(input, "\n") {
		v, _ := strconv.Atoi(s)
		m = append(m, v)
	}
	var steps = 1
	var i int
	for true {
		toDo := m[i]
		if toDo >= 3 {
			m[i] -= 1
		} else {
			m[i] += 1
		}
		if i+toDo < 0 || i+toDo >= len(m) {
			break
		} else {
			i += toDo
			steps += 1
		}
	}
	fmt.Println(steps)
}

func partOne(input string) {
	var m []int
	for _, s := range strings.Split(input, "\n") {
		v, _ := strconv.Atoi(s)
		m = append(m, v)
	}
	var steps = 1
	var i int
	for true {
		toDo := m[i]
		m[i] += 1
		if i+toDo < 0 || i+toDo >= len(m) {
			break
		} else {
			i += toDo
			steps += 1
		}
	}
	fmt.Println(steps)
}
