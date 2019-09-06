package binary_search

import "testing"

type PowInput struct {
	base  float64
	times int
}

func TestPow(t *testing.T) {
	var tests = []struct {
		in       PowInput
		expected float64
	}{
		// Add Testcase Here
		{PowInput{float64(2), 0}, float64(1)},
		{PowInput{float64(2), 5}, float64(32)},
		{PowInput{float64(2), -5}, float64(1) / float64(32)},
		{PowInput{float64(-2), 5}, float64(-32)},
		{PowInput{float64(-2), 4}, float64(16)},
		{PowInput{float64(-2), -5}, -(float64(1) / float64(32))},
		{PowInput{float64(1), -2147483648}, float64(1)},
		{PowInput{float64(-1), 2147483647}, float64(-1)},
		// 1.00000
		//
	}
	for _, tt := range tests {
		actual := myPow(tt.in.base, tt.in.times)
		if actual != tt.expected {
			t.Fatalf("meet exception, base is %f, times is %d, actual is %f, expect is %f",
				tt.in.base, tt.in.times, actual, tt.expected)
		}
	}
}
