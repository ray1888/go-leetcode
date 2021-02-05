package houserobber

import "go-leetcode/src/utils"

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) < 2 {
		return utils.Max(nums[0], 0)
	}
	return utils.Max(rob(nums[:len(nums)-1]), rob(nums[1:]))
}
