package parlindrome

import "testing"

func TestPalindromePartition(t *testing.T) {
	var tests = []struct {
		in       string
		expected [][]string
	}{
		// Add Testcase Here
		{"aab", [][]string{
			[]string{"a", "a", "b"},
			[]string{"aa", "b"},
		}},
	}
	for _, tt := range tests {
		actual := partition(tt.in)
		if len(tt.expected) != len(actual) {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
		for i := 0; i < len(actual); i++ {
			if len(tt.expected[i]) != len(actual[i]) {
				t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
					tt.in, tt.expected, actual)
			}
			for j := 0; j < len(actual[i]); j++ {
				if actual[i][j] != tt.expected[i][j] {
					t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
						tt.in, tt.expected, actual)
				}
			}
		}
	}
}
