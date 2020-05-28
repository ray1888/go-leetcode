package stringProblem

import "testing"

func TestRotateString(t *testing.T) {
	var tests = []struct {
		in       []string
		expected bool
	}{
		// Add Testcase Here
		{[]string{"abcde", "cdeab"}, true},
		{[]string{"abcde", "cedab"}, false},
		{[]string{"abcde", "ceab"}, false},
	}
	for _, test := range tests {
		result := rotateString(test.in[0], test.in[1])
		if result != test.expected {
			t.Fatalf("test failed, input is %v", test.in)
		}
	}

}
