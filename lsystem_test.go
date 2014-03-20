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

func TestFunction(t *testing.T) {
	testLindenmayer(t, lindenmayer)
}

func TestChannel(t *testing.T) {
	testLindenmayer(t, coliner)
}

func testLindenmayer(t *testing.T, f func([]string, map[string][]string, int) []string) {
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
		if r := f(c.Start, c.Rules, c.Iterations); !equals(r, c.Expected) {
			t.Errorf("f(%v, %v, %v) != %v (got %v)", c.Start, c.Rules, c.Iterations, c.Expected, r)
		}
	}
}

func BenchmarkFunction(b *testing.B) {
	benchmarkLindenmayer(b, lindenmayer)
}

func BenchmarkChannel(b *testing.B) {
	benchmarkLindenmayer(b, coliner)
}

func benchmarkLindenmayer(b *testing.B, f func([]string, map[string][]string, int) []string) {
	rules := map[string][]string{
		"A": {"A", "B"},
		"B": {"A"},
	}
	start := []string{"A"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(start, rules, 10)
	}
}
