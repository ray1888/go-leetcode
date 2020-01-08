package numbers

import "testing"

func TestSumOfSqaureNumberByDiff(t *testing.T) {
	tests := []BoolTest{
		{Name: "add up to a square number", Input: 8, Output: true},
		{Name: "can't add up to a square number", Input: 7, Output: false},
	}

	for _, test := range tests {
		result := SumOfSqaureNumberByDiff(test.Input)
		if result != test.Output {
			t.Fatalf("Test %s failed ", test.Name)
		}
	}
}

func TestSumOfSqaureNumberMap(t *testing.T) {
	tests := []BoolTest{
		{Name: "add up to a square number", Input: 8, Output: true},
		{Name: "can't add up to a square number", Input: 7, Output: false},
	}

	for _, test := range tests {
		result := SumOfSquareNumberByMap(test.Input)
		if result != test.Output {
			t.Fatalf("Test %s failed ", test.Name)
		}
	}
}

func TestSumOfSqaureNumberTwoPointer(t *testing.T) {
	tests := []BoolTest{
		{Name: "add up to a square number", Input: 8, Output: true},
		{Name: "can't add up to a square number", Input: 7, Output: false},
	}

	for _, test := range tests {
		result := SumOfSquareNumberTwoPointer(test.Input)
		if result != test.Output {
			t.Fatalf("Test %s failed ", test.Name)
		}
	}
}
