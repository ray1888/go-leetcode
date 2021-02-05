package houserobber

import "testing"

func TestHouseRobberI(t *testing.T) {
	var tests = []struct {
		in       []int
		expected int
	}{
		// Add Testcase Here
		{[]int{8, 1, 9, 3}, 17},
		{[]int{8, 1, -9, 3}, 11},
		{[]int{8, 1, 9, 3, 5}, 22},
	}
	for _, tt := range tests {
		actual := rob(tt.in)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}
