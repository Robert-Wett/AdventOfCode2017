package main

import (
	"fmt"
	"strconv"
	"strings"

	utils "github.com/Robert-Wett/AdventOfCode2017/helpers"
)

type Node struct {
	Name     string
	Weight   int
	Children []*Node
	Parent   *Node
}
type Graph map[string]*Node

func main() {
	input := utils.GetInput("./input.txt")
	partOne(input)
	partTwo(input)
}

func partTwo(input string) {
	g := make(Graph)
	for _, line := range strings.Split(input, "\n") {
		n := newNodeNoKids(line)
		g[n.Name] = n
	}
	addChildren(input, &g)
	//head := g.getHead()
	//for _, child := range head.Children {
	//for _, child := range g["ggxgmci"].Children {
	//for _, child := range g["anygv"].Children {
	// Note: This is a mess, basically manually ran this until I found
	// that 'fabacam' was the mis-weighted node
	for _, child := range g["fabacam"].Children {

		fmt.Println(child.Name, child.Weight)
		var hi []*Node
		traverse(child, map[string]bool{}, &hi)
		total := 0
		for _, x := range hi {
			total += x.Weight
		}
		fmt.Println(total)
	}
}

func traverse(n *Node, visited map[string]bool, todo *[]*Node) {
	(*todo) = append((*todo), n)
	for _, c := range n.Children {
		traverse(c, visited, todo)
	}
}

func (g *Graph) getHead() *Node {
	for _, v := range *g {
		if v.Parent == nil {
			return v
		}
	}
	return nil
}

func parseLine(line string) (string, int, []string) {
	weight, _ := strconv.Atoi((line[strings.Index(line, "(")+1 : strings.Index(line, ")")]))
	name := line[:strings.Index(line, " ")]
	var children []string
	if strings.Contains(line, "->") {
		c := strings.Split(line[strings.Index(line, ">")+2:], ", ")
		for _, v := range c {
			children = append(children, v)
		}
	}
	return name, weight, children
}

func addChildren(input string, g *Graph) {
	for _, line := range strings.Split(input, "\n") {
		name, _, children := parseLine(line)
		if len(children) > 0 {
			for _, child := range children {
				(*g)[name].Children = append((*g)[name].Children, (*g)[child])
				(*g)[child].Parent = (*g)[name]
			}
		}
	}
}

func newNodeNoKids(e string) *Node {
	name, weight, _ := parseLine(e)
	n := &Node{Name: name, Weight: weight}
	return n
}

func partOne(input string) {
	isHead := make(map[string]bool)
	var nodes []Node
	for _, line := range strings.Split(input, "\n") {
		n := newNode(line)
		if len(n.Children) > 0 {
			nodes = append(nodes, *n)
			isHead[n.Name] = true
		}
	}

	for _, n := range nodes {
		for _, child := range n.Children {
			isHead[child.Name] = false
		}
	}

	for k, v := range isHead {
		if v == true {
			fmt.Println(k)
		}
	}
}

func newNode(e string) *Node {
	name, weight, children := parseLine(e)
	n := &Node{Name: name, Weight: weight}
	if len(children) > 0 {
		for _, c := range children {
			n.Children = append(n.Children, &Node{Name: c})
		}
	}
	return n
}
