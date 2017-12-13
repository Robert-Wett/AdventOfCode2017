package main

import (
	"fmt"
	"sort"
	"strings"

	utils "github.com/Robert-Wett/AdventOfCode2017/helpers"
)

type runes []rune

func (r runes) Len() int           { return len(r) }
func (r runes) Less(i, j int) bool { return r[i] < r[j] }
func (r runes) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func sortString(w string) string {
	r := []rune(w)
	sort.Sort(runes(r))
	return string(r)
}

func main() {
	input := utils.GetInput("./input")
	partOne(input)
	partTwo(input)
}

func partTwo(input string) {
	var correctAmt int
	for _, v := range strings.Split(input, "\n") {
		if checkValidPartTwo(v) {
			correctAmt += 1
		}
	}
	fmt.Println(fmt.Sprintf("Total number of valid passwords: %d", correctAmt))
}

func checkValidPartTwo(p string) bool {
	w := make(map[string]struct{})
	for _, s := range strings.Split(p, " ") {
		k := sortString(s)
		if _, ok := w[k]; ok {
			return false
		} else {
			w[k] = struct{}{}
		}
	}
	return true
}

func partOne(input string) {
	var correctAmt int
	for _, v := range strings.Split(input, "\n") {
		if checkValidPartOne(v) {
			correctAmt += 1
		}
	}
	fmt.Println(fmt.Sprintf("Total number of valid passwords: %d", correctAmt))
}

func checkValidPartOne(p string) bool {
	w := make(map[string]struct{})
	for _, s := range strings.Split(p, " ") {
		if _, ok := w[s]; ok {
			return false
		} else {
			w[s] = struct{}{}
		}
	}
	return true
}
