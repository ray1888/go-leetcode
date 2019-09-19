package DynamicProgramming

import "testing"

func TestPalindromePartition2(t *testing.T) {
	var tests = []struct {
		in       string
		expected int
	}{
		// Add Testcase Here
		{"aab", 1},
		{"abba", 1},
		{"", 0},
		{"abcga", -1},
		{"aa", 0},
		{"abcba", 0},
		{"aabcc", 2},
	}
	for _, tt := range tests {
		actual := minCut(tt.in)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}
