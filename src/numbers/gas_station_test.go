package numbers

import "testing"

func TestCanComplteNoram(t *testing.T) {
	var tests = []struct {
		Name   string
		Gas    []int
		Cost   []int
		Result int
	}{
		{Name: "Normal testing ", Gas: []int{1, 2, 3, 4, 5}, Cost: []int{3, 4, 5, 1, 2}, Result: 3},
		{Name: "Normal testing 2", Gas: []int{1, 2, 4, 8}, Cost: []int{4, 8, 1, 2}, Result: 2},
	}

	for _, test := range tests {
		result := canCompleteCircuit(test.Gas, test.Cost)
		if result != test.Result {
			t.Fatalf("test %s failed", test.Name)
		}
	}
}

func TestCanComplteNoramSkip(t *testing.T) {
	var tests = []struct {
		Name   string
		Gas    []int
		Cost   []int
		Result int
	}{
		{Name: "Normal testing ", Gas: []int{1, 2, 3, 4, 5}, Cost: []int{3, 4, 5, 1, 2}, Result: 3},
		{Name: "Normal testing 2", Gas: []int{1, 2, 4, 8}, Cost: []int{4, 8, 1, 2}, Result: 2},
	}

	for _, test := range tests {
		result := canCompleteCircuitSkip(test.Gas, test.Cost)
		if result != test.Result {
			t.Fatalf("test %s failed", test.Name)
		}
	}
}

func TestCanComplteNoramGreedy(t *testing.T) {
	var tests = []struct {
		Name   string
		Gas    []int
		Cost   []int
		Result int
	}{
		{Name: "Normal testing ", Gas: []int{1, 2, 3, 4, 5}, Cost: []int{3, 4, 5, 1, 2}, Result: 3},
		{Name: "Normal testing 2", Gas: []int{1, 2, 4, 8}, Cost: []int{4, 8, 1, 2}, Result: 2},
	}

	for _, test := range tests {
		result := canCompleteCircuitGreedy(test.Gas, test.Cost)
		if result != test.Result {
			t.Fatalf("test %s failed", test.Name)
		}
	}
}
