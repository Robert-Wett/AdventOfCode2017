package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"

	utils "github.com/Robert-Wett/AdventOfCode2017/helpers"
)

func main() {
	input := utils.GetInput("./input.txt")

	partOne(input)
	partTwo(input)
}

func partOne(input string) {
	var max int
	var min int
	var total int
	for _, r := range strings.Split(input, "\n") {
		max = 0
		min = math.MaxInt32
		for _, i := range strings.Split(r, "\t") {
			num, _ := strconv.Atoi(i)
			if num > max {
				max = num
			}
			if num <= min {
				min = num
			}
		}
		total += (max - min)
	}
	fmt.Println(total)
}

func partTwo(input string) {
	var total int
	for _, r := range strings.Split(input, "\n") {
		curRow := strings.Split(r, "\t")
		for i := 0; i < len(curRow); i++ {
			for idx, ii := range curRow {
				if idx == i {
					continue
				}
				n, _ := strconv.Atoi(ii)
				vv, _ := strconv.Atoi(curRow[i])
				if n % vv == 0 {
					total += n / vv
					break
				}
			}
		}
	}
	fmt.Println(fmt.Sprintf("%d", total))
}