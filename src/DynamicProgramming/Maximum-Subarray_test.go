package DynamicProgramming

import "testing"

func TestMaxSubarray(t *testing.T) {
	var tests = []struct {
		in       []int
		expected int
	}{
		// Add Testcase Here
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{1, -2, 3, -4, 5}, 5},
		{[]int{1, 2, -3, 4, 5}, 9},
		{[]int{1, 2, -2, -4, 5}, 5},
		{[]int{-3, -2, 0, -1}, 0},
		{[]int{-1}, -1},
		{[]int{-2147483647}, -2147483647},
	}
	for _, tt := range tests {
		actual := maxSubArray(tt.in)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}
