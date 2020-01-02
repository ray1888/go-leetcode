package numbers

import "testing"

type Test struct {
	Name   string
	Input  []int
	Output []int
}

func TestPlusOne(t *testing.T) {
	tests := []Test{
		{Name: "normal", Input: []int{1, 2, 3}, Output: []int{1, 2, 4}},
		{Name: "double add one", Input: []int{1, 2, 9, 9}, Output: []int{1, 3, 0, 0}},
		{Name: "single add one", Input: []int{1, 2, 9}, Output: []int{1, 3, 0}},
		{Name: "add one", Input: []int{1, 9, 9, 5}, Output: []int{1, 9, 9, 6}},
		{Name: "all nine add one ", Input: []int{9, 9, 9}, Output: []int{1, 0, 0, 0}},
	}
	for _, test := range tests {
		result := PlusOne(test.Input)
		if len(result) != len(test.Output) {
			t.Fatalf("add result length is wrong")
		}
		for index, value := range test.Output {
			if value != result[index] {
				t.Fatalf("add result error")
			}
		}
	}
}
