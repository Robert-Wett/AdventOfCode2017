package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Robert-Wett/AdventOfCode2017/helpers"
)

func main() {
	var inputHash = "hfdlxzhv"
	partOne(inputHash)
	partTwo(inputHash)
}
func partTwo(input string) {
	var groupNo = 1
	var table diskMap
	for i := 0; i < 128; i++ {
		hash := helpers.KnotHash(fmt.Sprintf("%s-%d", input, i))
		var binH string
		for _, c := range hash {
			binH += HexToBin(string(c))
		}
		table = append(table, strings.Replace(binH, "1", "*", -1))
	}

	for i := 0; i < len(table); i++ {
		for ii := 0; i < len(table[i]); ii++ {
			if table.get(i, ii) == "*" {
				// Mark self
				n := table.getNeighborGroup(i, ii)
				if n != -1 {
					table.set(i, ii, strconv.Itoa(n))
					continue
				}
				table.set(i, ii, strconv.Itoa(groupNo))

				// Mark Above
				if i != 0 && table.get(i-1, ii) == "*" {
					table.set(i-1, ii, strconv.Itoa(groupNo))
				}
				// Mark Right
				if ii+1 < len(table[i]) && table.get(i, ii+1) == "*" {
					table.set(i, ii+1, strconv.Itoa(groupNo))
				}
				// Mark Left
				if ii-1 >= 0 && table.get(i, ii-1) == "*" {
					table.set(i, ii-1, strconv.Itoa(groupNo))
				}
				// Mark Below
				if i+1 < len(table) && table.get(i+1, ii) == "*" {
					table.set(i+1, ii, strconv.Itoa(groupNo))
				}
				groupNo++
			}
		}
	}
	fmt.Println(groupNo)
}

func (d *diskMap) getNeighborGroup(i, ii int) int {
	var check = func(s string) bool {
		return !strings.Contains("*0", s)
	}
	var getNum = func(s string) int {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		return num
	}
	// Check Above
	if i != 0 && check(d.get(i-1, ii)) {
		return getNum(d.get(i-1, ii))
	}
	// Check Right
	if ii+1 < len((*d)[i]) && check(d.get(i, ii+1)) {
		return getNum(d.get(i, ii+1))
	}
	// Check Left
	if ii-1 >= 0 && check(d.get(i, ii-1)) {
		return getNum(d.get(i, ii-1))
	}
	// Check Below
	if i+1 < len((*d)) && check(d.get(i+1, ii)) {
		return getNum(d.get(i+1, ii))
	}

	return -1
}

func (d *diskMap) set(rowIdx, lineIdx int, groupNo string) {
	(*d)[rowIdx] = (*d)[rowIdx][:lineIdx] + groupNo + (*d)[rowIdx][lineIdx+1:]
}

func (d *diskMap) get(row, col int) string {
	if row > len((*d)) || col > len((*d)[row]) {
		return ""
	}

	return string((*d)[row][col])
}

func markBitAsGroup(table []string, rowIdx, lineIdx, groupNo int) {
	table[rowIdx] = table[rowIdx][:lineIdx] + strconv.Itoa(groupNo) + table[rowIdx][lineIdx+1:]
}

type diskMap []string

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
	return fmt.Sprintf("%04b", ui)
}
