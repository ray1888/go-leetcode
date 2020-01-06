package numbers

import "testing"

func TestHappyNumberTwoPointer(t *testing.T) {
	tests := []BoolTest{
		{Name: "is happy number", Input: 28, Output: true},
		{Name: "is not happy number", Input: 2, Output: false},
	}
	for _, test := range tests {
		result := HappyNubmerTwoPointer(test.Input)
		if result != test.Output {
			t.Fatalf("Test %s failed", test.Name)
		}

	}
}
