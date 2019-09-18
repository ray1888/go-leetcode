package DynamicProgramming

import "testing"

func TestMaxPrice(t *testing.T) {
	var tests = []struct {
		in       []int
		expected int
	}{
		// Add Testcase Here
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	for _, tt := range tests {
		actual := maxProfit(tt.in)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}
