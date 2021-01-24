package numbers

import "testing"

type threeSumTest struct {
	Name  string
	Input struct {
		array  []int
		target int
	}
	Output int
}

func TestThreeSumCloest(t *testing.T) {
	tests := []threeSumTest{
		{
			Name: "add as target", Input: struct {
				array  []int
				target int
			}{
				array:  []int{-5, 1, 3, 4, 5},
				target: 1},
			Output: 1,
		},
		{
			Name: "add not as target", Input: struct {
				array  []int
				target int
			}{
				array:  []int{-3, 1, 3, 4, 5},
				target: -1},
			Output: 0},
		{
			Name: "add in seq", Input: struct {
				array  []int
				target int
			}{
				array:  []int{0, 1, 2},
				target: 3},
			Output: 3},
	}
	for _, test := range tests {
		result := ThreeSumCloest(test.Input.array, test.Input.target)
		if result != test.Output {
			t.Fatalf("test is %s failed", test.Name)
		}
	}
}
