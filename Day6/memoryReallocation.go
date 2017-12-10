package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := `14	0	15	12	11	11	3	5	1	6	8	4	9	1	8	4`

	strBanks := strings.Split(input, "\t")
	var banks []int
	for _, b := range strBanks {
		_b, _ := strconv.Atoi(b)
		banks = append(banks, _b)
	}

	partOne(append([]int(nil), banks...))
	partTwo(append([]int(nil), banks...))
}

type banks []int

func (b banks) getState() string {
	var state string
	for _, bank := range b {
		state += strconv.Itoa(bank)
	}
	return state
}

func (b banks) reAllocate() {
	var highest = b.indexOfHighest()
	memories := b[highest]
	b[highest] = 0
	var index = highest
	for memories > 0 {
		if index+1 >= len(b) {
			index = 0
		} else {
			index += 1
		}
		memories -= 1
		b[index] += 1
	}
}

func (b banks) indexOfHighest() int {
	var idx int
	var highest int
	for i, v := range b {
		if v > highest {
			idx = i
			highest = v
		}
	}
	return idx
}

func partTwo(b banks) {
	states := make(map[string]struct{})
	var moveCount int
	var dup string
	for true {
		s := b.getState()
		if _, ok := states[s]; ok {
			if dup == "" {
				moveCount = 0
				dup = s
			} else if dup == s {
				break
			}
		} else {
			states[s] = struct{}{}
		}

		b.reAllocate()
		moveCount += 1
	}

	fmt.Println(fmt.Sprintf("It took %d steps to reach the duplicate state twice", moveCount))
}

func partOne(b banks) {
	states := make(map[string]struct{})
	var moveCount int
	for true {
		s := b.getState()
		if _, ok := states[s]; ok {
			fmt.Println(fmt.Sprintf("It took %d steps to reach the duplicate state", moveCount))
			break
		} else {
			states[s] = struct{}{}
			b.reAllocate()
			moveCount += 1
		}
	}
}
