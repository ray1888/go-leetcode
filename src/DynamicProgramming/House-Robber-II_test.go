package DynamicProgramming

import "testing"

func TestHouseRobberII(t *testing.T) {
	var tests = []struct {
		in       []int
		expected int
	}{
		// Add Testcase Here
		{[]int{1, 2, 3, 1}, 4},
		{[]int{8, 1, 9, 3, 5}, 17},
	}
	for _, tt := range tests {
		actual := rob2(tt.in)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}
