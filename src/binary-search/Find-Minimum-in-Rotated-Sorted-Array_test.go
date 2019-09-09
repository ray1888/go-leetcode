package binary_search

import "testing"

func TestFindMinRoutatedSortedArray(t *testing.T) {
	var tests = []struct {
		in       []int
		expected int
	}{
		// Add Testcase Here
		{[]int{1}, 1},
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{2, 1}, 1},
	}
	for _, tt := range tests {
		actual := findMin(tt.in)
		if actual != tt.expected {
			t.Fatalf("get two sum failed, actual is %d, expected is %d",
				actual, tt.expected)
		}
	}
}
