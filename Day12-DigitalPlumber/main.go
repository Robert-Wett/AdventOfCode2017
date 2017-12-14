package main

import (
	"strconv"
	"strings"
)

func main() {
	var input = `0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5`

	programMap := make(programs)
	for _, line := range strings.Split(input, "\n") {
		addEntry(line, &programMap)
	}

}

func traverse(p *program, numBelow *[]*program, programMap *programs) {
	(*numBelow) = append((*numBelow), p)
	for _, c := range p.Children {
		//childNode := (*prog)
		traverse((*programMap)[c], numBelow, programMap)
	}
}

type programs map[int]program

type program struct {
	Value, Parent int
	Children      []int
}

func addEntry(inputLine string, p *programs) {
	line := strings.Split(inputLine, "<->")
	nodeVal, _ := strconv.Atoi(strings.Trim(line[0], " "))
	var children []int
	for _, strVal := range strings.Split(line[1], ", ") {
		childVal, _ := strconv.Atoi(strings.Trim(strVal, " "))
		if childVal != nodeVal {
			children = append(children, childVal)
		}
	}

	ourNode := program{Value: nodeVal, Children: children}

	if node, ok := (*p)[nodeVal]; ok {
		// Add children, set their parents
		for _, cv := range children {
			if !contains(cv, node.Children) {
				node.Children = append(node.Children, cv)
			}
			if cNode, ok := (*p)[cv]; ok {
				cNode.Parent = node.Value
			}
		}
	} else {
		(*p)[nodeVal] = ourNode
	}
}

func contains(in int, coll []int) bool {
	for _, i := range coll {
		if in == i {
			return true
		}
	}
	return false
}
