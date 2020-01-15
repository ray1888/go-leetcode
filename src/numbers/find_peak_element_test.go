package numbers

import "testing"

func TestPeakElementSeqSearch(t *testing.T) {
	tests := []TestInputArrayOutputInt{
		{Name: "has one Peak", Input: []int{1, 2, 3, 2}, Output: 2},
		{Name: "has two Peak", Input: []int{1, 2, 3, 2, 4, 2}, Output: 2},
		{Name: "has no Peak", Input: []int{1, 2, 3, 4}, Output: 3},
	}
	for _, test := range tests {
		result := findPeakElement(test.Input)
		if result != test.Output {
			t.Fatalf("test %s failed", test.Name)
		}
	}
}

func TestPeakElementBinarySearch(t *testing.T) {
	tests := []TestInputArrayOutputInt{
		{Name: "has one Peak", Input: []int{1, 2, 3, 2}, Output: 2},
		{Name: "has two Peak", Input: []int{1, 2, 3, 2, 4, 2}, Output: 2},
		{Name: "has no Peak", Input: []int{1, 2, 3, 4}, Output: 3},
	}
	for _, test := range tests {
		result := findPeakElementBinarySearch(test.Input)
		if result != test.Output {
			t.Fatalf("test %s failed", test.Name)
		}
	}
}
