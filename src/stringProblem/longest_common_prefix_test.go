package stringProblem

import "testing"

type Test struct {
	Name   string
	Input  []string
	Output string
}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []Test{
		{Name: "has some common prefix", Input: []string{"car", "care", "cat"}, Output: "ca"},
		{Name: "has no common prefix", Input: []string{"car", "bus", "boot"}, Output: ""},
		{Name: "has all common prefix", Input: []string{"car", "car", "car"}, Output: "car"},
	}
	for _, test := range tests {
		result := longestCommonPrefix(test.Input)
		if result != test.Output {
			t.Fatalf("Test %s failed! acutual is %s, expected is %s", test.Name, test.Input, test.Output)
		}
	}
}
