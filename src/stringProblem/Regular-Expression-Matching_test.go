package stringProblem

import "testing"

func TestRegExp(t *testing.T) {
	var tests = []struct {
		in       []string
		expected bool
	}{
		// Add Testcase Here
		{[]string{"aa", "a*"}, true},
		{[]string{"", ".*"}, true},
		{[]string{"aa", "a"}, false},
		{[]string{"ab", ".*"}, true},
	}
	for _, tt := range tests {
		actual := isMatch(tt.in[0], tt.in[1])
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v",
				tt.in, tt.expected)
		}
	}
}
