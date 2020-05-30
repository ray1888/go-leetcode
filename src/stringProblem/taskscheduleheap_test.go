package stringProblem

import "testing"

type TestCases []struct {
	Name   string
	Input  []byte
	N      int
	Output int
}

func TestLeastInterval(t *testing.T) {
	tests := TestCases{
		//{Name: "normal", Input: []byte("AAABBC"), N: 2, Output: 7},
		//{Name: "normal", Input: []byte("AAABBB"), N: 2, Output: 8},
		{Name: "normal", Input: []byte("AAAAAABCDEFG"), N: 2, Output: 16},
	}
	for _, test := range tests {
		out := LeastInterval(test.Input, test.N)
		if out != test.Output {
			t.Fatalf("test %s failed, output is %d, expected output is %d", test.Name, out, test.Output)
		}
	}
}
