package DynamicProgramming

import "testing"

func TestUniPathI(t *testing.T) {
	var tests = []struct {
		in       []int
		expected int
	}{
		// Add Testcase Here
		{[]int{3, 2}, 3},
		{[]int{7, 3}, 28},
	}
	for _, tt := range tests {
		actual := uniquePathsDP(tt.in[0], tt.in[1])
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}
