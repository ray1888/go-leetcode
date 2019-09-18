package DynamicProgramming

import "testing"

func TestDecodeWayRecur(t *testing.T) {
	var tests = []struct {
		in       string
		expected int
	}{
		// Add Testcase Here
		{"12", 2},
		{"124", 3},
	}
	for _, tt := range tests {
		actual := numDecodingsRecur(tt.in)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}

func TestDecodeWayDP(t *testing.T) {
	var tests = []struct {
		in       string
		expected int
	}{
		// Add Testcase Here
		{"12", 2},
		{"124", 3},
	}
	for _, tt := range tests {
		actual := numDecodingsDP(tt.in)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v, actual is %v",
				tt.in, tt.expected, actual)
		}
	}
}
