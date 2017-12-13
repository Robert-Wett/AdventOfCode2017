package main

import (
	"testing"
)

func TestKnotHash(t *testing.T) {
	var cases = []struct {
		max      int
		knots    string
		expected int
	}{
		{5, "3, 4, 1, 5", 12},
	}

	for _, tt := range cases {
		hi := knotHash(tt.max, tt.knots)
		if hi != tt.expected {
			t.Errorf("Expected %d, got %d - case '%s'", tt.expected, hi, tt.input)
		}
	}
}
