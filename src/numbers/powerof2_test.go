package numbers

import "testing"

type test struct {
	input  int
	output bool
}

func TestPowerOf2(t *testing.T) {
	tests := []test{
		{1, true},
		{2, true},
		{4, true},
		{8, true},
		{15, false},
		{16, true},
	}
	for _, testexample := range tests {
		result := PowerOf2(testexample.input)
		if result != testexample.output {
			t.Fatalf("Power of 2 meet error")
		}
	}
}

func TestPowerOf2Math(t *testing.T) {
	tests := []test{
		{1, true},
		{2, true},
		{4, true},
		{8, true},
		{15, false},
		{16, true},
	}
	for _, testexample := range tests {
		result := PowerOf2Math(testexample.input)
		if result != testexample.output {
			t.Fatalf("Power of 2 meet error")
		}
	}
}
