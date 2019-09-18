package DynamicProgramming

import "testing"

func TestCoinChangeDP(t *testing.T) {
	var tests = []struct {
		in       Coininput
		expected int
	}{
		// Add Testcase Here
		{in: Coininput{amount: 4, coins: []int{1, 2, 4}}, expected: 1},
		{in: Coininput{amount: 4, coins: []int{1, 2}}, expected: 2},
		{in: Coininput{amount: 11, coins: []int{1, 2, 5}}, expected: 3},
		{in: Coininput{amount: 3, coins: []int{2}}, expected: -1},
	}
	for _, tt := range tests {
		actual := coinChange(tt.in.coins, tt.in.amount)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}
