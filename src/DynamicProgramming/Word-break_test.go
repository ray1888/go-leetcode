package DynamicProgramming

import "testing"

type WordBreakInput struct {
	s        string
	wordDict []string
}

func TestWordBreak(t *testing.T) {
	var tests = []struct {
		in       WordBreakInput
		expected bool
	}{
		// Add Testcase Here
		{WordBreakInput{"leetcode", []string{"leet", "code"}}, true},
		{WordBreakInput{"applepenapple", []string{"apple", "pen"}}, true},
		{WordBreakInput{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}}, false},
	}
	for _, tt := range tests {
		actual := wordBreak(tt.in.s, tt.in.wordDict)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}
