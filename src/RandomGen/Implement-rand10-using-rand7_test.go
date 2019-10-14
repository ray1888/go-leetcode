package RandomGen

import "testing"

func TestRand10Gen(t *testing.T) {
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
		//actual := isMatch(tt.in[0], tt.in[1])
		actual := true
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v",
				tt.in, tt.expected)
		}
	}
}
