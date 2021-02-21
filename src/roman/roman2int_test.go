package roman

import "testing"

func TestRomanToInt(t *testing.T) {
	var tests = []struct {
		in       string
		expected int
	}{
		// Add Testcase Here
		{"IV", 4},
		{"VI", 6},
		{"XL", 40},
		{"LX", 60},
		{"XC", 90},
		{"CX", 110},
		{"CMDXL", 1440},
	}
	for _, tt := range tests {
		actual := romanToInt(tt.in)
		if actual != tt.expected {
			t.Fatalf("get insert postition wrong, input is %v, expected is %v",
				tt.in, tt.expected)
		}
	}
}
