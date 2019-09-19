package DynamicProgramming

import "testing"

func TestTriangle(t *testing.T) {
	var tests = []struct {
		in       [][]int
		expected int
	}{
		// Add Testcase Here
		{[][]int{
			[]int{2},
			[]int{3, 4},
			[]int{6, 5, 7},
			[]int{4, 1, 8, 3},
		}, 11},
		{[][]int{
			[]int{-1},
			[]int{2, 3},
			[]int{1, -1, -3},
		}, -1},
	}
	for _, tt := range tests {
		actual := minimumTotal(tt.in)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}
