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

/*
 ____            _     _____
|  _ \ __ _ _ __| |_  |_   _|_      _____
| |_) / _` | '__| __|   | | \ \ /\ / / _ \
|  __/ (_| | |  | |_    | |  \ V  V / (_) |
|_|   \__,_|_|   \__|   |_|   \_/\_/ \___/
*/
func partTwo(input string) {
	var groupNo int
	table := createTable(input)
	for i := 0; i < len(table); i++ {
		for ii := 0; ii < len(table[i]); ii++ {
			nGroup := table.getNeighborGroup(i, ii)
			if nGroup != -1 {
				table.floodFill(i, ii, nGroup)
			} else if table.getValue(i, ii) == 0 && table.get(i, ii) == "*" {
				groupNo++
				table.floodFill(i, ii, groupNo)
			}
		}
	}

	fmt.Println(groupNo)
}

func createTable(input string) rows {
	var t []string
	for i := 0; i < 128; i++ {
		hash := helpers.KnotHash(fmt.Sprintf("%s-%d", input, i))
		var binH string
		for _, c := range hash {
			binH += HexToBin(string(c))
		}
		t = append(t, strings.Replace(binH, "1", "*", -1))
	}
	var res rows
	for _, line := range t {
		var r row
		for i := 0; i < len(line); i++ {
			r = append(r, cell{
				display: string(line[i]),
				set:     string(line[i]) != "*",
			})
		}
		res = append(res, r)
	}

	return res
}

type rows []row
type row []cell
type cell struct {
	display string
	group   int
	set     bool
}

func (r *rows) set(rowIdx, lineIdx, groupNo int) {
	(*r)[rowIdx][lineIdx].group = groupNo
	(*r)[rowIdx][lineIdx].set = true
}

func (r *rows) getValue(row, col int) int {
	if row > len((*r)) || col >= len((*r)[row]) {
		return -1
	}

	return (*r)[row][col].group
}

func (r *rows) get(row, col int) string {
	if row > len((*r)) || col > len((*r)[row]) {
		return ""
	}

	display := (*r)[row][col].display
	return display
}

func (r *rows) floodFill(row, col, groupNo int) {
	if (*r)[row][col].set {
		return
	}

	var check = func(r *rows, row, col int) bool {
		return !(*r)[row][col].set
	}

	(*r)[row][col].group = groupNo
	(*r)[row][col].set = true

	// Check Below
	if row+1 < len((*r)) && check(r, row+1, col) {
		r.floodFill(row+1, col, groupNo)
	}

	// Check Right
	if col+1 < len((*r)[row]) && check(r, row, col+1) {
		r.floodFill(row, col+1, groupNo)
	}

	// Check Left
	if col-1 >= 0 && check(r, row, col-1) {
		r.floodFill(row, col-1, groupNo)
	}

	// Check Up
	if row-1 >= 0 && check(r, row-1, col) {
		r.floodFill(row-1, col, groupNo)
	}
}

func (r *rows) getNeighborGroup(i, ii int) int {
	var check = func(r *rows, row, col int) bool {
		return r.getValue(row, col) != 0
	}

	// Check Self (We may have been set previously)
	if check(r, i, ii) {
		return r.getValue(i, ii)
	}

	// Check Above
	if i != 0 && check(r, i-1, ii) {
		return r.getValue(i-1, ii)
	}
	// Check Right
	if ii+1 < len((*r)[i]) && check(r, i, ii+1) {
		return r.getValue(i, ii+1)
	}

	// Check Left
	if ii-1 >= 0 && check(r, i, ii-1) {
		return r.getValue(i, ii-1)
	}
	// Check Below
	if i+1 < len((*r)) && check(r, i+1, ii) {
		return r.getValue(i+1, ii)
	}

	return -1
}

/*
 ____            _      ___
|  _ \ __ _ _ __| |_   / _ \ _ __   ___
| |_) / _` | '__| __| | | | | '_ \ / _ \
|  __/ (_| | |  | |_  | |_| | | | |  __/
|_|   \__,_|_|   \__|  \___/|_| |_|\___|
*/
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

type diskMap []string

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

func HexToBin(hex string) string {
	ui, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%04b", ui)
}
