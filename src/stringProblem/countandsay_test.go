package stringProblem

import "testing"

func TestCountAndSay(t *testing.T) {
	type test struct {
		Name   string
		Input  int
		Output string
	}
	tests := []test{
		{Name: "One", Input: 1, Output: "1"},
		{Name: "Two", Input: 2, Output: "11"},
		{Name: "Three", Input: 3, Output: "21"},
		{Name: "Four", Input: 4, Output: "1211"},
		{Name: "Five", Input: 5, Output: "111221"},
	}
	for _, test := range tests {
		result := CountAndSay(test.Input)
		if result != test.Output {
			t.Fatalf("test %s failed", test.Name)
		}
	}
}
