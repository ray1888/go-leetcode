package DynamicProgramming

import "testing"

func TestMinCostClimbStair(t *testing.T) {
	var tests = []struct {
		in       []int
		expected int
	}{
		// Add Testcase Here
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}
	for _, tt := range tests {
		actual := minCostClimbingStairs(tt.in)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}
