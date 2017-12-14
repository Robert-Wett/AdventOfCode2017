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
	programMap := make(programs)
	for _, line := range strings.Split(input, "\n") {
		addNode(line, &programMap)
	}
	for _, line := range strings.Split(input, "\n") {
		addChildren(line, &programMap)
	}

	var below []*program
	visited := make(map[int]bool)
	traverse(programMap.getProgram(0), &below, &visited, &programMap)

	fmt.Println(len(visited))

}

func (p *programs) getProgram(num int) *program {
	return (*p)[num]
}

func traverse(p *program, numBelow *[]*program, v *map[int]bool, programMap *programs) {
	if !(*v)[p.Value] {
		(*v)[p.Value] = true
		for _, c := range p.Children {
			traverse(c, numBelow, v, programMap)
		}
		(*numBelow) = append((*numBelow), p)
	}
}

type programs map[int]*program

type program struct {
	Value    int
	Parent   *program
	Children []*program
}

func addNode(inputLine string, p *programs) {
	nodeVal, _ := parseLine(inputLine)
	(*p)[nodeVal] = &program{Value: nodeVal}
}

func addChildren(inputLine string, p *programs) {
	nodeVal, children := parseLine(inputLine)
	parent := (*p)[nodeVal]
	for _, intVal := range children {
		child := (*p)[intVal]
		child.Parent = parent
		parent.Children = append(parent.Children, child)
	}
}

func parseLine(inputLine string) (int, []int) {
	line := strings.Split(inputLine, "<->")
	nodeVal, _ := strconv.Atoi(strings.Trim(line[0], " "))
	var children []int
	for _, strVal := range strings.Split(line[1], ", ") {
		childVal, _ := strconv.Atoi(strings.Trim(strVal, " "))
		if childVal != nodeVal {
			children = append(children, childVal)
		}
	}

	return nodeVal, children
}

func contains(in int, coll []int) bool {
	for _, i := range coll {
		if in == i {
			return true
		}
	}
	return false
}
