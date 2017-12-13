package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//var input = "83,0,193,1,254,237,187,40,88,27,2,255,149,29,42,100"
	var input = "3,4,1,5"
	var knots []int
	strKnots := strings.Split(input, ",")

	for _, n := range strKnots {
		num, _ := strconv.Atoi(n)
		knots = append(knots, num)
	}

	fmt.Println(knotHash(5, knots))
}

func knotHash(max int, knotPoints []int) int {
	var pos = 1
	var skip int
	hash := getSeed(max)
	for i, knot := range knotPoints {
		var chain = hash[i:pos]
		for knot > 0 {

		}
		/*
		if knot+pos > len(hash) {
			chain = hash[knot-(knot+pos%len(hash)):]
			chain = append(chain, hash[:knot+pos%len(hash)])
		} else {
			chain = hash[knot-(knot+pos%len(hash)):]
		}
		*/
	}
	return 1
}

func getSeed(n int) []int {
	var seed []int
	for i := 0; i < n; i++ {
		seed = append(seed, i)
	}
	return seed
}
