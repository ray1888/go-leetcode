package numbers

import "testing"

func TestNumSquareDp(t *testing.T) {
	tests := []IntergerTest{
		{Name: "has one plus", Input: 1, Output: 1},
		{Name: "has two plus", Input: 5, Output: 2},
		{Name: "has three plus", Input: 12, Output: 3},
		{Name: "has three plus", Input: 15, Output: 4},
	}
	for _, test := range tests {
		result := numSquaresDP(test.Input)
		if result != test.Output {
			t.Fatalf("test %s failed", test.Name)
		}
	}
}

func TestNumSquaresMath(t *testing.T) {
	tests := []IntergerTest{
		{Name: "has one plus", Input: 1, Output: 1},
		{Name: "has two plus", Input: 5, Output: 2},
		{Name: "has three plus", Input: 12, Output: 3},
		{Name: "has three plus", Input: 15, Output: 4},
	}
	for _, test := range tests {
		result := numSquaresMath(test.Input)
		if result != test.Output {
			t.Fatalf("test %s failed", test.Name)
		}
	}
}
