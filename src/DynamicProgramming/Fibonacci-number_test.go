package DynamicProgramming

import "testing"

func TestFibNumber(t *testing.T) {
	var tests = []struct {
		in       int
		expected int
	}{
		// Add Testcase Here
		{2, 1},
		{3, 2},
		{4, 3},
	}
	for _, tt := range tests {
		actual := fibDP(tt.in)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}

func TestFibNumberRecur(t *testing.T) {
	var tests = []struct {
		in       int
		expected int
	}{
		// Add Testcase Here
		{2, 1},
		{3, 2},
		{4, 3},
	}
	for _, tt := range tests {
		actual := fibRecur(tt.in)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}
