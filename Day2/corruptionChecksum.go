package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
	"math"
)

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var input = string(content)

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
				}
			}
		}
	}
	fmt.Println(fmt.Sprintf("%d", total))
}