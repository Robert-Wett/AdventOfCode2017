package main

import (
	"testing"
)

func TestProcessStreamOne(t *testing.T) {
	var cases = []struct {
		input    string
		expected int
	}{
		{"<>", 0},
		{"<random characters>", 0},
		{"<<<<>", 0},
		{"<{!>}>", 0},
		{`<!!>`, 0},
		{"<!!!>>", 0},
		{"<{o\"i!a,<{i<a>", 0},
		{"{}", 1},
		{"{{{}}}", 6},
		{"{{},{}}", 5},
		{"{{{},{},{{}}}}", 16},
		{"{<a>,<a>,<a>,<a>}", 1},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
	}

	for _, tt := range cases {
		hi := processStreamOne(tt.input)
		if hi != tt.expected {
			t.Errorf("Expected %d, got %d - case '%s'", tt.expected, hi, tt.input)
		}
	}
}

func TestProcessStreamTwo(t *testing.T) {
	var cases = []struct {
		input    string
		expected int
	}{
		{"<>", 0},
		{"<random characters>", 17},
		{"<<<<>", 3},
		{"<{!>}>", 2},
		{"<!!>", 0},
		{"<!!!>>", 0},
		{`<{o"i!a,<{i<a>`, 10},
	}

	for _, tt := range cases {
		hi := processStreamTwo(tt.input)
		if hi != tt.expected {
			t.Errorf("Expected %d, got %d - case '%s'", tt.expected, hi, tt.input)
		}
	}
}
