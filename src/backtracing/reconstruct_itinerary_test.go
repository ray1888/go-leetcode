package backtracing

import (
	"testing"
)

func TestFindItinerary(t *testing.T) {
	var tests = []struct {
		args     [][]string
		expected []string
	}{
		// {
		// 	args: [][]string{
		// 		[]string{"MUC", "LHR"},
		// 		[]string{"JFK", "MUC"},
		// 		[]string{"SFO", "SJC"},
		// 		[]string{"LHR", "SFO"},
		// 	},
		// 	expected: []string{
		// 		"JFK", "MUC", "LHR", "SFO", "SJC",
		// 	},
		// },
		{
			args: [][]string{
				[]string{"JFK", "SFO"},
				[]string{"JFK", "ATL"},
				[]string{"SFO", "ATL"},
				[]string{"ATL", "JFK"},
				[]string{"ATL", "SFO"},
			},
			expected: []string{
				"JFK", "ATL", "JFK", "SFO", "ATL", "SFO",
			},
		},
	}
	for _, tt := range tests {
		actual := findItinerary(tt.args)
		if len(actual) != len(tt.expected) {
			t.Fatal("len wrong")
		}
		for i := 0; i < len(actual); i++ {
			if actual[i] != tt.expected[i] {
				t.Fatal("route error")
			}
		}
	}
}
