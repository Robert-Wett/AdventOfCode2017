package main

import (
	"fmt"
	"strconv"
	"strings"

	utils "github.com/Robert-Wett/AdventOfCode2017/helpers"
)

func main() {
	input := utils.GetInput("./input.txt")

	partOne(input)
	partTwo(input)
}

func partTwo(input string) {
	var highestValue int
	values := initRegisters(input)
	for _, line := range strings.Split(input, "\n") {
		symbol, op, amount, condSymbol, condOp, condAmount := parseLine(line)
		if evalCondition(condSymbol, condOp, condAmount, &values) {
			incDec(symbol, op, amount, &values)
			for _, value := range values {
				if value > highestValue {
					highestValue = value
				}
			}
		}
	}

	fmt.Println(highestValue)
}

func partOne(input string) {
	values := initRegisters(input)
	for _, line := range strings.Split(input, "\n") {
		symbol, op, amount, condSymbol, condOp, condAmount := parseLine(line)
		if evalCondition(condSymbol, condOp, condAmount, &values) {
			incDec(symbol, op, amount, &values)
		}
	}

	var highest int
	for _, value := range values {
		if value >= highest {
			highest = value
		}
	}

	fmt.Println(highest)
}

func initRegisters(input string) map[string]int {
	values := make(map[string]int)
	for _, line := range strings.Split(input, "\n") {
		symbol, _, _, _, _, _ := parseLine(line)
		values[symbol] = 0
	}
	return values
}

func incDec(symbol, op string, amount int, values *map[string]int) {
	switch op {
	case "inc":
		(*values)[symbol] += amount
	case "dec":
		(*values)[symbol] -= amount
	}
}

func evalCondition(symbol, op string, amount int, values *map[string]int) bool {
	switch op {
	case ">":
		return (*values)[symbol] > amount
	case ">=":
		return (*values)[symbol] >= amount
	case "<":
		return (*values)[symbol] < amount
	case "<=":
		return (*values)[symbol] <= amount
	case "==":
		return (*values)[symbol] == amount
	case "!=":
		return (*values)[symbol] != amount
	default:
		return false
	}
}

func parseLine(line string) (string, string, int, string, string, int) {
	pieces := strings.Split(line, " ")
	var symbol, op, condSymbol, condOp = pieces[0], pieces[1], pieces[4], pieces[5]
	amt, _ := strconv.Atoi(pieces[2])
	condAmount, _ := strconv.Atoi(pieces[6])
	return symbol, op, amt, condSymbol, condOp, condAmount
}
