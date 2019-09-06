package binary_search

import "testing"

func TestLongIncrSubSeq(t *testing.T) {
	var tests = []struct {
		in       []int
		expected int
	}{
		// Add Testcase Here
		{[]int{1, 8, 2, 6, 4, 5}, 4},
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
	}
	for _, tt := range tests {
		actual := lengthOfLIS(tt.in)
		if actual != tt.expected {
			t.Fatalf("get two sum failed, actual is %d, expected is %d",
				actual, tt.expected)
		}
	}
}
