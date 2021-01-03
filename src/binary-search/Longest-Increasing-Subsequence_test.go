package binary_search

import (
	"testing"
)

func TestLongIncrSubSeq(t *testing.T) {
	var tests = []struct {
		in       []int
		expected int
	}{
		// Add Testcase Here
		{[]int{1, 8, 2, 6, 4, 5}, 4},
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
	}
	for _, tt := range tests {
		actual := lengthOfLIS(tt.in)
		if actual != tt.expected {
			t.Fatalf("get two sum failed, actual is %d, expected is %d",
				actual, tt.expected)
		}
	}
}

func Test_BinarySearchInsPos(t *testing.T) {
	type args struct {
		d      []int
		end    int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _BinarySearchInsPos(tt.args.d, tt.args.end, tt.args.target); got != tt.want {
				t.Errorf("_BinarySearchInsPos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLISDP(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLISDP(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLISDP() = %v, want %v", got, tt.want)
			}
		})
	}
}
