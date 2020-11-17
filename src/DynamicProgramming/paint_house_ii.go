package DynamicProgramming

import "math"

func MinCostWithKColors(nums [][]int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	k := len(nums[0])

	d := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		d[i] = make([]int, k)
	}

	for i := 1; i <= n; i++ {
		for j := 0; j < k; j++ {
			minV := math.MaxInt32
			for c := 0; c < k; c++ {
				if c != j {
					minV = min(minV, d[i-1][c])
				}
			}
			d[i][j] = minV + nums[i-1][j]
		}
	}
	minValue := math.MaxInt32
	for i := 0; i < k; i++ {
		minValue = min(minValue, d[n][i])
	}
	return minValue
}
