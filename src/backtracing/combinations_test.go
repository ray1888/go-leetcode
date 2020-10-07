package backtracing

import (
	"fmt"
	"testing"
)

func TestCombinations(t *testing.T) {
	type Args struct {
		N int // 1-n取值范围
		K int // 取值个数
	}
	var tests = []struct {
		args     Args
		expected [][]int
	}{
		// Add Testcase Here
		{
			args: Args{N: 4, K: 2},
			expected: [][]int{
				[]int{1, 2},
				[]int{1, 3},
				[]int{1, 4},
				[]int{2, 3},
				[]int{2, 4},
				[]int{3, 4},
			},
		},
	}
	for _, tt := range tests {
		actual := combine(tt.args.N, tt.args.K)
		fmt.Println(actual)
	}
}
