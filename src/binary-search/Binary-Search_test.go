package binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
	var tests = []struct {
		in       input
		expected int
	}{
		// Add Testcase Here
		{input{[]int{1, 3, 5, 8}, 2}, -1},
		{input{[]int{1, 3, 5, 6}, 5}, 2},
		{input{[]int{1, 3, 5, 8}, 9}, -1},
		{input{[]int{1, 3, 5, 8}, 8}, 3},
		{input{[]int{2, 3, 5, 6}, 2}, 0},
		{input{[]int{1}, 1}, 0},
	}
	for _, tt := range tests {
		actual := search(tt.in.v, tt.in.iv)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %d",
				tt.in, tt.expected)
		}
	}
}
