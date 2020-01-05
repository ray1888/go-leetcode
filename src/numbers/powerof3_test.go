package numbers

import "testing"

func TestPowerOf3Brutal(t *testing.T) {
	tests := []BoolTest{
		{Name: "is 3 pow", Input: 9, Output: true},
		{Name: "is not 3 pow", Input: 8, Output: false},
	}
	for _, test := range tests {
		result := PowerOf3Brutal(test.Input)
		if result != test.Output {
			t.Fatalf("Test %s Failed! Acutal is %t, Excepted is %t", test.Name, result, test.Output)
		}
	}
}

func TestPowerOf3Math(t *testing.T) {
	tests := []BoolTest{
		{Name: "is 3 pow", Input: 9, Output: true},
		{Name: "is not 3 pow", Input: 8, Output: false},
	}
	for _, test := range tests {
		result := PowerOf3Math(test.Input)
		if result != test.Output {
			t.Fatalf("Test %s Failed! Acutal is %t, Excepted is %t", test.Name, result, test.Output)
		}
	}
}
