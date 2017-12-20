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
	var rows []string
	for i := 0; i < 128; i++ {
		hash := helpers.KnotHash(fmt.Sprintf("%s-%d", input, i))
		var binH string
		for _, c := range hash {
			binH += HexToBin(string(c))
		}
		rows = append(rows, strings.Replace(binH, "1", "*", -1))
	}
	table := diskMap(rows)
	var groupNo = 1
	//for i, row := range rows {
	for i := 0; i < 128; i++ {
		for ii := 0; ii < 128; ii++ {
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
				if ii+1 < len(rows[i]) && table.get(i, ii+1) == "*" {
					table.set(i, ii+1, strconv.Itoa(groupNo))
				}
				// Mark Left
				if ii-1 >= 0 && table.get(i, ii-1) == "*" {
					table.set(i, ii-1, strconv.Itoa(groupNo))
				}
				// Mark Below
				if i+1 < len(rows) && table.get(i+1, ii) == "*" {
					table.set(i+1, ii, strconv.Itoa(groupNo))
				}
				groupNo++
			}
		}
	}

	fmt.Println(groupNo)
}

func (d *diskMap) getNeighborGroup(i, ii int) int {
	// Check Above
	if i != 0 && !strings.Contains("*0", d.get(i-1, ii)) {
		num, err := strconv.Atoi(d.get(i-1, ii))
		if err != nil {
			log.Fatal(err)
		}
		return num
	}
	// Check Right
	if ii+1 < len((*d)[i]) && !strings.Contains("*0", d.get(i, ii+1)) {
		num, err := strconv.Atoi(d.get(i, ii+1))
		if err != nil {
			log.Fatal(err)
		}
		return num
	}
	// Check Left
	if ii-1 >= 0 && !strings.Contains("*0", d.get(i, ii-1)) {
		num, err := strconv.Atoi(d.get(i, ii-1))
		if err != nil {
			log.Fatal(err)
		}
		return num
	}
	// Check Below
	if i+1 < len((*d)) && !strings.Contains("*0", d.get(i+1, ii)) {
		num, err := strconv.Atoi(d.get(i+1, ii))
		if err != nil {
			log.Fatal(err)
		}
		return num
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
	return fmt.Sprintf("%016b", ui)
}
