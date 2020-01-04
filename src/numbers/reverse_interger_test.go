package numbers

import "testing"

type IntergerTest struct {
	Name   string
	Input  int
	Output int
}

func TestReverseInterger(t *testing.T) {
	tests := []IntergerTest{
		{Name: "normal", Input: 42, Output: 24},
		{Name: "negative", Input: -42, Output: -24},
		{Name: "end-with-zero", Input: 420, Output: 24},
		{Name: "MaxValue", Input: 2 << 30, Output: 0},
		{Name: "MinValue", Input: -(2 << 30), Output: 0},
	}

	for _, test := range tests {
		result := reverse(test.Input)
		if result != test.Output {
			t.Fatalf("Test %s failed! Input is %d, actual is %d, output is %d", test.Name, test.Input, result, test.Output)
		}
	}
}
