package binary_search

import "testing"

func TestSquareRoot(t *testing.T) {
	var tests = []struct {
		in       int
		expected int
	}{
		// Add Testcase Here
		{8, 2},
		{4, 2},
	}
	for _, tt := range tests {
		actual := mySqrt(tt.in)
		if actual != tt.expected {
			t.Fatalf("get two sum failed, actual is %d, expected is %d",
				actual, tt.expected)
		}
	}
}
