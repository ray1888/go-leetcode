package numbers

import "testing"

type BoolTest struct {
	Name   string
	Input  int
	Output bool
}

func TestIsPerfectSquareBinarySearch(t *testing.T) {
	tests := []BoolTest{
		{Name: "normal", Input: 9, Output: true},
		{Name: "double add one", Input: 8, Output: false},
	}
	for _, test := range tests {
		result := isPerfectSquareBinarySearch(test.Input)
		if result != test.Output {
			t.Fatalf("add result length is wrong")
		}
	}
}

func TestIsPerfectSquareNewton(t *testing.T) {
	tests := []BoolTest{
		{Name: "normal", Input: 9, Output: true},
		{Name: "double add one", Input: 8, Output: false},
	}
	for _, test := range tests {
		result := isPerfectSquareNewton(test.Input)
		if result != test.Output {
			t.Fatalf("add result length is wrong")
		}
	}
}
