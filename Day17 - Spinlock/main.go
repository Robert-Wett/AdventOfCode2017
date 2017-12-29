package main

import (
	"fmt"

	"github.com/Robert-Wett/AdventOfCode2017/helpers"
)

func main() {
	var input = 348
	lock := newLock(input)
	for i := 0; i <= 2017; i++ {
		lock.spin()
	}
	needle := helpers.IndexOf(2017, lock.values)
	fmt.Println(lock.values[needle+1])

}

type lock struct {
	values []int
	pos    int
	step   int
	curVal int
}

func newLock(step int) *lock {
	return &lock{step: step, values: []int{0}, curVal: 1}
}

func (l *lock) spin() {
	if len(l.values) == 1 {
		l.values = append(l.values, l.curVal)
		l.curVal++
		l.pos = 1
		return
	}

	amountOnTheRight := len(l.values) - (l.pos + 1)
	if l.step+l.pos < len(l.values) {
		l.insert(l.step + l.pos + 1)
	} else {
		stuffToUse := l.step - amountOnTheRight
		whereWellBe := (stuffToUse % len(l.values))
		l.insert(whereWellBe)
	}
}

func (l *lock) toString() string {
	var s string
	for i, r := range l.values {
		if i == l.pos {
			s += fmt.Sprintf("(%d)", r)
		} else {
			s += fmt.Sprintf(" %d ", r)
		}
	}
	return s
}

func (l *lock) insert(index int) {
	l.values = append(l.values[:index], append([]int{l.curVal}, l.values[index:]...)...)
	l.curVal++
	l.pos = index
}
