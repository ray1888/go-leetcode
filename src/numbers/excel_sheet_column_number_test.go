package numbers

import "testing"

type StringInputTest struct {
	Name   string
	Input  string
	Output int
}

func TestLeft2Right(t *testing.T) {
	tests := []StringInputTest{
		{Name: "Single Number", Input: "A", Output: 1},
		{Name: "Double Number", Input: "AB", Output: 28},
	}
	for _, test := range tests {
		result := titleToNumberLeft2Right(test.Input)
		if result != test.Output {
			t.Fatalf("test %s is failed", test.Name)
		}
	}
}

func TestRight2Left(t *testing.T) {
	tests := []StringInputTest{
		{Name: "Single Number", Input: "A", Output: 1},
		{Name: "Double Number", Input: "AB", Output: 28},
	}
	for _, test := range tests {
		result := titleToNumberRight2Left(test.Input)
		if result != test.Output {
			t.Fatalf("test %s is failed", test.Name)
		}
	}
}
