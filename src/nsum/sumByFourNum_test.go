package nsum

import "testing"

type input struct {
	v  []int
	iv int
}

func TestNSumBy4(t *testing.T) {
	var tests = []struct {
		in       input
		expected [][]int
	}{
		// Add Testcase Here
		{input{[]int{-3, 0, 1, 2}, 0}, [][]int{[]int{-3, 0, 1, 2}}},
		{input{[]int{1, 0, -1, 0, -3, 2}, 0}, [][]int{[]int{-3, 0, 1, 2}, []int{-1, 0, 0, 1}}},
	}
	for _, tt := range tests {
		actual := NSumBy4Nums(tt.in.v, tt.in.iv)
		if len(actual) != len(tt.expected) {
			t.Fatal("length is not match ")
		}
		counter := 0
		for _, val := range actual {
			for i := 0; i < len(actual); i++ {
				result := true
				for j := 0; j < 4; j++ {
					if val[j] != tt.expected[i][j] {
						result = false
					}
				}
				if result == true {
					counter += 1
				}
			}
		}
		if counter != len(actual) {
			t.Fatalf("the match seq is not right")
		}

	}
}
