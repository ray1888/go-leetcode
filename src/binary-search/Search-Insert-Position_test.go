package binary_search

import (
	"testing"
)

type input struct {
	v  []int
	iv int
}

func TestSearchInsertPosition(t *testing.T) {
	var tests = []struct {
		in       input
		expected int
	}{
		// Add Testcase Here
		{input{[]int{1, 3, 5, 8}, 2}, 1},
		{input{[]int{1, 3, 5, 6}, 5}, 2},
		{input{[]int{1, 3, 5, 8}, 9}, 4},
		{input{[]int{1, 3, 5, 8}, 6}, 3},
		{input{[]int{2, 3, 5, 6}, 1}, 0},
		{input{[]int{1}, 1}, 0},
	}
	for _, tt := range tests {
		actual := searchInsert(tt.in.v, tt.in.iv)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %d",
				tt.in, tt.expected)
		}
	}
}
