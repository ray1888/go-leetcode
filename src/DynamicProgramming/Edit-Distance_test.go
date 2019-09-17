package DynamicProgramming

import "testing"

func TestEditDistance(t *testing.T) {
	var tests = []struct {
		in       []string
		expected int
	}{
		// Add Testcase Here
		{[]string{"car", "ba"}, 2},
		{[]string{"", ""}, 0},
		{[]string{"abc", "abc"}, 0},
		{[]string{"abd", "abu"}, 1},
	}
	for _, tt := range tests {
		actual := minDistance(tt.in[0], tt.in[1])
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}
