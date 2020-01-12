package numbers

import "testing"

func TestFindDuplicateBrutalForce(t *testing.T) {
	tests := []TestInputArrayOutputInt{
		{Name: "has Duplicate", Input: []int{1, 2, 3, 4, 4}, Output: 4},
		{Name: "has Not Duplicate", Input: []int{1, 2, 3, 4}, Output: -1},
	}
	for _, test := range tests {
		result := findDuplicateBrutalForce(test.Input)
		if result != test.Output {
			t.Fatalf("test %s failed", test.Name)
		}
	}

}

func TestFindDuplicateBinarySearch(t *testing.T) {
	tests := []TestInputArrayOutputInt{
		{Name: "has Duplicate", Input: []int{1, 2, 3, 4, 4}, Output: 4},
		{Name: "has Not Duplicate", Input: []int{1, 2, 3, 4}, Output: 3},
	}
	for _, test := range tests {
		result := findDuplicateBinarySearch(test.Input)
		if result != test.Output {
			t.Fatalf("test %s failed, result is %d", test.Name, result)
		}
	}
}

func TestFindDuplicateTwoPointer(t *testing.T) {
	tests := []TestInputArrayOutputInt{
		{Name: "has Duplicate", Input: []int{4, 3, 4, 1, 2, 5}, Output: 4},
		{Name: "has Duplicate", Input: []int{1, 2, 3, 4, 4}, Output: 4},
	}
	for _, test := range tests {
		result := findDuplicateTwoPointer(test.Input)
		if result != test.Output {
			t.Fatalf("test %s failed, result is %d", test.Name, result)
		}
	}
}
