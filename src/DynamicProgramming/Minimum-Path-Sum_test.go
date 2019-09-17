package DynamicProgramming

import "testing"

func TestMinPathSum(t *testing.T) {
	var tests = []struct {
		in       [][]int
		expected int
	}{
		// Add Testcase Here
		{[][]int{
			[]int{1, 3, 1},
			[]int{1, 5, 1},
			[]int{4, 2, 1},
		}, 7},
		{[][]int{
			[]int{1, 2, 1},
			[]int{8, 4, 1},
			[]int{-8, 2, 1},
		}, 4},
	}
	for _, tt := range tests {
		actual := minPathSum(tt.in)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}
