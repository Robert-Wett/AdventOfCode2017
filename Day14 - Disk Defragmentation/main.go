package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Robert-Wett/AdventOfCode2017/helpers"
)

func main() {
	var inputHash = "hfdlxzhv"
	partOne(inputHash)
}

func partOne(input string) {
	var rows []string
	for i := 0; i < 128; i++ {
		hash := helpers.KnotHash(fmt.Sprintf("%s-%d", input, i))
		var binH string
		for _, c := range hash {
			binH += HexToBin(string(c))
		}
		rows = append(rows, binH)
	}

	var count int
	for _, line := range rows {
		for _, c := range line {
			if string(c) == "1" {
				count++
			}
		}
	}

	fmt.Println(count)
}

func HexToBin(hex string) string {
	ui, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%016b", ui)
}
func countSquares(hash string) int {
	return 1
}

func funcTwo(stub string) int {
	return 1
}
