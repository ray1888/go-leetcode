package heap

import "testing"

func TestFindKthLargestPartition(t *testing.T) {
	tests := []struct {
		input []int
		k     int
		want  int
	}{
		{[]int{2, 1, 5, 6}, 1, 6},
		{[]int{1, 5, 4}, 3, 1},
		{[]int{1, 9, 5, 3, 2, 6}, 3, 5},
		{[]int{}, 3, -1},
		{[]int{1, 9, 5, 3, 2, 6}, 7, -1},
	}
	for _, tt := range tests {
		actual := findKthLargestPartition(tt.input, tt.k)
		if actual != tt.want {
			t.Fatalf("Find Kth Largest number failed, input is %v, want %d, actual %d", tt.input, tt.want, actual)
		}
	}
}

func TestFindKthLargestMinHeap(t *testing.T) {
	tests := []struct {
		input []int
		k     int
		want  int
	}{
		{[]int{2, 1, 5, 6}, 1, 6},
		{[]int{1, 5, 4}, 3, 1},
		{[]int{1, 9, 5, 3, 2, 6}, 3, 5},
		{[]int{}, 3, -1},
		{[]int{1, 9, 5, 3, 2, 6}, 7, -1},
	}
	for _, tt := range tests {
		actual := findKthLargestMinHeap(tt.input, tt.k)
		if actual != tt.want {
			t.Fatalf("Find Kth Largest number failed, input is %v, want %d, actual %d", tt.input, tt.want, actual)
		}
	}
}
