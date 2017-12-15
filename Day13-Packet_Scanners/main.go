package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Robert-Wett/AdventOfCode2017/helpers"
)

func main() {
	input := helpers.GetInput("./input.txt")
	partOne(input)
}

func partOne(input string) {
	m := newLayerMap(input)
	p := packet{}
	_ = attemptPass(m, &p)
}

func attemptPass(m *layerMap, p *packet) bool {
	for i := 0; i < m.getMax(); i++ {
		p.tick(m)
		m.tick()
	}
	severity := p.caught.calculateSeverity()
	if severity > 0 {
		fmt.Println(fmt.Sprintf("You were caught %d times, with a severity of %d", len(p.caught), severity))
		return false
	}

	return true
}

type packet struct {
	idx    int
	caught layers
}

func (p *packet) tick(l *layerMap) {
	// If a layer exists at this index
	if curLayer, ok := (*l)[p.idx]; ok {
		// If the layer's scanner is at the top
		if curLayer.idx == 0 {
			p.caught = append(p.caught, curLayer)
		}
	}
	p.idx++
}

type layerMap map[int]*layer

func newLayerMap(input string) *layerMap {
	lm := make(layerMap)
	for _, line := range strings.Split(input, "\n") {
		p, d := parseLine(line)
		lm[p] = &layer{pos: p, depth: d}
	}

	return &lm
}

func (l *layerMap) tick() {
	for _, layer := range *l {
		(*layer).tick()
	}
}

func (l *layerMap) getMax() int {
	var highest int
	for k := range *l {
		if k > highest {
			highest = k
		}
	}

	return highest
}

type layers []*layer

func (l *layers) calculateSeverity() int {
	var total int
	for _, layer := range *l {
		total += layer.pos * layer.depth
	}

	return total
}

type layer struct {
	idx   int
	depth int
	pos   int
	dir   int
}

func (l *layer) tick() {
	// Zero or Single lengthed layers don't change state
	if l.depth-1 <= 0 {
		return
	}

	if l.dir == 0 { // Down
		if l.idx+1 > l.depth-1 {
			// Change direction
			l.dir = 1
			l.idx--
		} else {
			l.idx++
		}
	} else { // Up
		if l.idx-1 < 0 {
			// Change direction
			l.dir = 0
			l.idx++
		} else {
			l.idx--
		}
	}
}

func parseLine(line string) (int, int) {
	layer := strings.Split(line, ": ")
	pos, _ := strconv.Atoi(layer[0])
	depth, _ := strconv.Atoi(layer[1])
	return pos, depth
}
