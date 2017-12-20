package main

import (
	"testing"
)

func TestDefragPartOne(t *testing.T) {
	var cases = []struct {
		input    string
		expected int
	}{
		{"7ef846a84695f115cbd8840c616f3df7", 8108},
	}

	for _, tt := range cases {
		hi := countSquares(tt.input)
		if hi != tt.expected {
			t.Errorf("Expected %d, got %d - case '%s'", tt.expected, hi, tt.input)
		}
	}
}

func TestDefragPartTwo(t *testing.T) {
	var cases = []struct {
		input    string
		expected int
	}{
		{"<>", 0},
	}

	for _, tt := range cases {
		hi := funcTwo(tt.input)
		if hi != tt.expected {
			t.Errorf("Expected %d, got %d - case '%s'", tt.expected, hi, tt.input)
		}
	}
}
