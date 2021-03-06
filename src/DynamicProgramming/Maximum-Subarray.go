package DynamicProgramming

import (
	"errors"
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxSubArray(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("length should be bigger than 0")
	}
	cur := 0
	maxium := -int((math.Pow(2, 32)))
	for i := 0; i < len(nums); i++ {
		if cur > 0 {
			cur += nums[i]
		} else {
			cur = nums[i]
		}
		maxium = max(cur, maxium)

	}
	return maxium, nil
}
