package DynamicProgramming

import "testing"

func TestUniPathII(t *testing.T) {
	var tests = []struct {
		in       [][]int
		expected int
	}{
		// Add Testcase Here
		{[][]int{
			[]int{0, 0, 0},
			[]int{0, 1, 0},
			[]int{0, 0, 0},
		}, 2},
		{[][]int{
			[]int{1, 0},
		}, 0},
		{[][]int{
			[]int{0, 0},
			[]int{1, 1},
			[]int{0, 0},
		}, 0},
	}
	for _, tt := range tests {
		actual := uniquePathsWithObstacles(tt.in)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}
