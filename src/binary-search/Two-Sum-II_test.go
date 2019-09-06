package binary_search

import "testing"

func TestTwoSum2(t *testing.T) {
	var tests = []struct {
		in       input
		expected []int
	}{
		// Add Testcase Here

		{input{[]int{1, 3, 5, 8}, 2}, []int{-1, -1}},
		{input{[]int{1, 3, 5, 6}, 4}, []int{1, 2}},
		{input{[]int{1, 3, 5, 6}, 5}, []int{-1, -1}},
		{input{[]int{1, 3, 5, 8}, 9}, []int{1, 4}},
	}
	for _, tt := range tests {
		actual := twoSum(tt.in.v, tt.in.iv)
		if len(actual) != len(tt.expected) {
			t.Fatalf("get two sum failed, actual len != expected len")
		}
		for i := 0; i < len(actual); i++ {
			if actual[i] != tt.expected[i] {
				t.Fatalf("value not match,input is %v,  ai is %d, ei is %d, i is %d",
					tt.in, actual[i], tt.expected[i], i)
			}
		}
	}
}
