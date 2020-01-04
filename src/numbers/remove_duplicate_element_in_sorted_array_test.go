package numbers

import "testing"

type TestInputArrayOutputInt struct {
	Name   string
	Input  []int
	Output int
}

func TestRemoveDuplicate(t *testing.T) {
	tests := []TestInputArrayOutputInt{
		{Name: "Duplicate In middle", Input: []int{1, 2, 3, 4, 4, 5}, Output: 5},
		{Name: "Duplicate In head", Input: []int{56, 56, 59, 60, 77}, Output: 4},
		{Name: "Duplicate In Tail", Input: []int{53, 54, 55, 56, 56}, Output: 4},
		{Name: "Duplicate In Multi", Input: []int{53, 53, 54, 55, 56, 56}, Output: 4},
		{Name: "Duplicate In Multi", Input: []int{53, 53, 54, 55, 55, 56, 56}, Output: 4},
	}

	for _, test := range tests {
		result := removeDuplicates(test.Input)
		if result != test.Output {
			t.Fatalf("Test %s failed! actual is %d, expected is %d", test.Name, result, test.Output)
		}
	}
}
