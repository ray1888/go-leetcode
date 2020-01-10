package numbers

import "testing"

func TestFactorilTrailingZeros(t *testing.T) {

	tests := []IntergerTest{

		{Name: "has Zero", Input: 4, Output: 0},
		{Name: "has one", Input: 5, Output: 1},
		//{Name:"has Two", Input:25, Output:2},
		//{Name:"has Three", Input:25, Output:2},
	}
	for _, test := range tests {
		result := FactorilTrailingZeros(test.Input)
		if result != test.Output {
			t.Fatalf("test %s failed.", test.Name)
		}
	}

}
