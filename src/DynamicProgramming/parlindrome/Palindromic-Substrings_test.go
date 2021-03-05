package parlindrome

import "testing"

func TestPalindromicSubstrings(t *testing.T) {
	var tests = []struct {
		in       string
		expected int
	}{
		// Add Testcase Here
		{"aba", 4},
		{"abba", 6},
		{"", 0},
		{"abcga", 5},
	}
	for _, tt := range tests {
		actual := CountSubstrings(tt.in)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}

func TestPalindromicSubstringsExpand(t *testing.T) {
	var tests = []struct {
		in       string
		expected int
	}{
		// Add Testcase Here
		{"aba", 4},
		{"abba", 6},
		{"", 0},
		{"abcga", 5},
	}
	for _, tt := range tests {
		actual := CountSubstringsExpand(tt.in)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}
