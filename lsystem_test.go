package main

import (
	"testing"
)

func equals(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, va := range a {
		if va != b[i] {
			return false
		}
	}

	return true
}

func TestLindenmayer(t *testing.T) {
	tests := []struct {
		Rules      map[string][]string
		Start      []string
		Iterations int
		Expected   []string
	}{
		{ // simple in/out test
			map[string][]string{
				"A": {"A", "B"},
				"B": {"A"},
			},
			[]string{"A"}, 0,
			[]string{"A"},
		},
		{ // 4 gen tree
			map[string][]string{
				"A": {"A", "B"},
				"B": {"A"},
			},
			[]string{"A"}, 4,
			[]string{"A", "B", "A", "A", "B", "A", "B", "A"},
		},
		{ // dragon curve
			map[string][]string{
				"X": {"X", "+", "Y", "F"},
				"Y": {"F", "X", "-", "Y"},
			},
			[]string{"F", "X"}, 2,
			[]string{"F", "X", "+", "Y", "F", "+", "F", "X", "-", "Y", "F"},
		},
		{ // sierpinski triangle
			map[string][]string{
				"A": {"B", "-", "A", "-", "B"},
				"B": {"A", "+", "B", "+", "A"},
			},
			[]string{"A"}, 2,
			[]string{"A", "+", "B", "+", "A", "-", "B", "-", "A", "-", "B", "-", "A", "+", "B", "+", "A"},
		},
	}

	for _, c := range tests {
		if r := lindenmayer(c.Start, c.Rules, c.Iterations); !equals(r, c.Expected) {
			t.Errorf("lindenmayer(%v, %v, %v) != %v (got %v)", c.Start, c.Rules, c.Iterations, c.Expected, r)
		}
	}
}
