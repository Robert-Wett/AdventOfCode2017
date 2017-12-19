package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Robert-Wett/AdventOfCode2017/helpers"
)

func main() {
	partOne(helpers.GetInput("./input.txt"))
	partTwo(helpers.GetInput("./input.txt"))
}

func partTwo(input string) {
	p := program{[]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p'}}
	// This value represents the number of times we ACTUALLY run the commands.
	// This algorithm repeats it's value every 63rd cycle.
	var max = 1000000000 % 63
	for i := 0; i < max; i++ {
		for _, c := range parseInstructions(input) {
			switch c.command {
			case "s":
				p.spin(c.sNum)
			case "p":
				p.partner(c.pFirst, c.pSecond)
			case "x":
				p.exchange(c.eFirst, c.eSecond)
			default:
				log.Fatal("Huh?")
			}
		}

	}
	fmt.Println(p.toString())

}
func partOne(input string) {

	p := program{[]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p'}}
	for _, c := range parseInstructions(input) {
		switch c.command {
		case "s":
			p.spin(c.sNum)
		case "p":
			p.partner(c.pFirst, c.pSecond)
		case "x":
			p.exchange(c.eFirst, c.eSecond)
		default:
			log.Fatal("Huh?")
		}
	}
	fmt.Println(p.toString())
}

type cmd struct {
	command         string
	pFirst, pSecond []byte
	eFirst, eSecond int
	sNum            int
}

func parseInstructions(input string) []cmd {
	var commands []cmd
	for _, line := range strings.Split(input, ",") {
		op := string(line[0])
		var c = cmd{command: op}
		switch op {
		case "s":
			num, _ := strconv.Atoi(string(line[1:]))
			c.sNum = num
			commands = append(commands, c)
		case "p":
			split := strings.Split(line, "/")
			first, second := []byte(split[0][1:]), []byte(split[1])
			c.pFirst, c.pSecond = first, second
			commands = append(commands, c)
		case "x":
			split := strings.Split(line, "/")
			first, _ := strconv.Atoi(split[0][1:])
			second, _ := strconv.Atoi(split[1])
			c.eFirst, c.eSecond = first, second
			commands = append(commands, c)
		default:
			fmt.Println("Whoa....")
		}
	}

	return commands
}

func (p *program) exchange(first, second int) {
	p.Str[first], p.Str[second] = p.Str[second], p.Str[first]
}

func (p *program) partner(first, second []byte) {
	f, s := bytes.Index(p.Str, first), bytes.Index(p.Str, second)
	p.Str[f], p.Str[s] = p.Str[s], p.Str[f]
}

func (p *program) spin(num int) {
	anchor := len(p.Str) - num
	var tmp = p.Str
	if num < len(tmp) {
		tmp = append(tmp[anchor:], tmp[:anchor]...)
	}
	p.Str = tmp
}

type program struct {
	Str []byte
}

func (p *program) toString() string {
	var stringed string
	for _, program := range p.Str {
		stringed += string(program)
	}
	return stringed
}
