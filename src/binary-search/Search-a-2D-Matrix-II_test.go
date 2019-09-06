package binary_search

import "testing"

var TwoDArray [][]int = [][]int{
	{1, 4, 7, 11, 15},
	{2, 5, 8, 12, 19},
	{3, 6, 9, 16, 22},
	{10, 13, 14, 17, 24},
	{18, 21, 23, 26, 30},
}

type TwoDInput struct {
	input   [][]int
	snumber int
}

func TestSearch2DMatrix(t *testing.T) {
	var tests = []struct {
		in       TwoDInput
		expected bool
	}{
		// Add Testcase Here
		{TwoDInput{TwoDArray, 5}, true},
		{TwoDInput{TwoDArray, 20}, false},
	}
	for _, tt := range tests {
		actual := searchMatrix(tt.in.input, tt.in.snumber)
		if actual != tt.expected {
			//t.Fatalf("meet exception, base is %f, times is %d, actual is %f, expect is %f",
			//	tt.in.base, tt.in.times, actual, tt.expected)
			t.Fatalf("error ")
		}
	}
}
