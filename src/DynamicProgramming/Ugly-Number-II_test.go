package DynamicProgramming

import "testing"

func TestUglyNumber2(t *testing.T) {
	var tests = []struct {
		in       int
		expected int
	}{
		// Add Testcase Here
		{10, 12},
		{6, 6},
	}
	for _, tt := range tests {
		actual := nthUglyNumber(tt.in)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}
