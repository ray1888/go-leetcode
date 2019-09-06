package binary_search

import "testing"

func TestSingleElement(t *testing.T) {
	var tests = []struct {
		in       []int
		expected int
	}{
		// Add Testcase Here
		{[]int{2, 2, 3, 3, 6, 6, 8}, 8},
		{[]int{2, 2, 3, 4, 4}, 3},
		{[]int{5, 5, 7, 7, 8, 9, 9}, 8},
	}
	for _, tt := range tests {
		actual := singleNonDuplicate(tt.in)
		if actual != tt.expected {
			t.Fatalf("get two sum failed, actual is %d, expected is %d",
				actual, tt.expected)
		}
	}
}
