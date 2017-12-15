package main

import (
	"fmt"

	utils "github.com/Robert-Wett/AdventOfCode2017/helpers"
)

func main() {
	input := utils.GetInput("./input.txt")

	partOne(input)
	partTwo(input)
}

func partOne(input string) {
	fmt.Println(processStreamOne(input))
}

func partTwo(input string) {
	fmt.Println(processStreamTwo(input))
}

// processStreamTwo counts the number of garbage characters in the stream
func processStreamTwo(input string) int {
	var negate bool
	var inGarbage bool
	var score int
	var stack []string
	for i := 0; i < len(input); i++ {
		switch string(input[i]) {
		case "{":
			if negate {
				negate = !negate
			} else if !inGarbage {
				stack = append(stack, "{")
			} else if inGarbage {
				score++
			}
		case "}":
			if negate {
				negate = !negate
			} else if !inGarbage {
				// Pop off the group and count it (*swish*)
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
			} else if inGarbage {
				score++
			}
		case "<":
			if negate {
				negate = !negate
			} else if !inGarbage {
				inGarbage = true
			} else if inGarbage {
				score++
			}
		case ">":
			if negate {
				negate = !negate
			} else if inGarbage {
				inGarbage = false
			}
		case "!":
			negate = !negate
		default:
			if negate {
				negate = !negate
			} else if inGarbage {
				score++
			}
		}
	}

	return score
}

// processStreamOne counts the number of valid groups are in the given input.
func processStreamOne(input string) int {
	var negate bool
	var inGarbage bool
	var groups int
	var score int
	var stack []string
	for i := 0; i < len(input); i++ {
		switch string(input[i]) {
		case "{":
			if negate {
				negate = !negate
			} else if !inGarbage {
				stack = append(stack, "{")
				groups++
			}
		case "}":
			if negate {
				negate = !negate
			} else if !inGarbage {
				// Pop off the group and count it (*swish*)
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
					score += groups
					groups--
				}
			}
		case "<":
			if negate {
				negate = !negate
			} else if !inGarbage {
				inGarbage = true
			}
		case ">":
			if negate {
				negate = !negate
			} else if inGarbage {
				inGarbage = false
			}
		case "!":
			negate = !negate
		default:
			if negate {
				negate = !negate
			}
		}
	}

	return score
}
