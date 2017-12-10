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
}

// mfohmwu (344) -> dniumoe, uuimn, ewiugad, vkkds, bqpycy
// fwudfax (27)
func main() {
	input := utils.GetInput("./input.txt")
	partOne(input)
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
	weight, _ := strconv.Atoi((e[strings.Index(e, "(")+1 : strings.Index(e, ")")]))
	name := e[:strings.Index(e, " ")]
	n := &Node{Name: name, Weight: weight}
	if strings.Contains(e, "->") {
		children := strings.Split(e[strings.Index(e, ">")+2:], ", ")
		for _, v := range children {
			n.Children = append(n.Children, &Node{Name: v})
		}
	}
	return n
}

func (n *Node) getChildren() []string {
	var c []string
	for _, n := range n.Children {
		c = append(c, n.Name)
	}
	return c
}

func (n *Node) toString() string {
	return fmt.Sprintf("Name:%s, Weight: %d, Children: %s", n.Name, n.Weight, strings.Join(n.getChildren(), ", "))
}
