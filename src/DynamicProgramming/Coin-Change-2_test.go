package DynamicProgramming

import "testing"

type Coininput struct {
	amount int
	coins  []int
}

func TestCoinChange2Recv(t *testing.T) {
	var tests = []struct {
		in       Coininput
		expected int
	}{
		// Add Testcase Here
		{in: Coininput{amount: 5, coins: []int{1, 2, 5}}, expected: 4},
		{in: Coininput{amount: 4, coins: []int{1, 2, 4}}, expected: 4},
	}
	for _, tt := range tests {
		actual := ChangeRecursive(tt.in.amount, tt.in.coins)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}

func TestCoinChange2DP(t *testing.T) {
	var tests = []struct {
		in       Coininput
		expected int
	}{
		// Add Testcase Here
		//{in: Coininput{amount: 5, coins: []int{1, 2, 5}}, expected: 4},
		{in: Coininput{amount: 4, coins: []int{1, 2, 4}}, expected: 4},
	}
	for _, tt := range tests {
		actual := ChangeDP(tt.in.amount, tt.in.coins)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}
