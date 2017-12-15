package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Robert-Wett/AdventOfCode2017/helpers"
)

func main() {
	partOne(helpers.GetInput("./input.txt"))
	partTwoBrute(helpers.GetInput("./input.txt"))
}

func partTwoBrute(input string) {
	var i int
	m := newLayerMap(input)
	for true {
		m.zeroMap()
		p := packet{}
		if attemptPassWithWait(m, &p, i) {
			fmt.Println(fmt.Sprintf("You had to wait %d picoseconds before crossing successfuly", i))
			break
		}
		i++
	}
}

func attemptPassWithWait(m *layerMap, p *packet, numPreTicks int) bool {
	m.tick(numPreTicks)
	max := m.getMax()
	for i := 0; i < max; i++ {
		p.tick(m)
		m.tick(1)
	}

	return len(p.caught) == 0
}

func partOne(input string) {
	m := newLayerMap(input)
	p := packet{}
	attemptPass(m, &p)
}

func attemptPass(m *layerMap, p *packet) {
	max := m.getMax()
	for i := 0; i < max; i++ {
		p.tick(m)
		m.tick(1)
	}
	severity := p.caught.calculateSeverity()
	if severity > 0 {
		fmt.Println(fmt.Sprintf("You were caught %d times, with a severity of %d", len(p.caught), severity))
	}
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

func (l *layerMap) zeroMap() {
	for _, layer := range *l {
		layer.idx = 0
		layer.dir = 0
	}
}

func (l *layerMap) tick(num int) {
	if num <= 0 {
		return
	}
	for _, layer := range *l {
		(*layer).tick(num)
	}
}

func (l *layerMap) getMax() int {
	var highest int
	for k := range *l {
		if k > highest {
			highest = k
		}
	}

	return highest + 1
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

func (l *layer) tick(num int) {
	for i := 0; i < num; i++ {
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
}

func parseLine(line string) (int, int) {
	layer := strings.Split(line, ": ")
	pos, _ := strconv.Atoi(layer[0])
	depth, _ := strconv.Atoi(layer[1])
	return pos, depth
}
