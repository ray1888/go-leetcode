package parlidrome

import "testing"

func TestLongestPalindromicString(t *testing.T) {
	var tests = []struct {
		in       string
		expected int
	}{
		// Add Testcase Here
		{"aba", 3},
		{"abba", 4},
		{"", 0},
		{"abcga", 1},
	}
	for _, tt := range tests {
		actual := longestCountSubstrings(tt.in)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}

func TestLongestPalindromicStringExpand(t *testing.T) {
	var tests = []struct {
		in       string
		expected int
	}{
		// Add Testcase Here
		{"aba", 3},
		{"abba", 4},
		{"", 0},
		{"abcga", 1},
	}
	for _, tt := range tests {
		actual := longestSubstringsExpand(tt.in)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}
