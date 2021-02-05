package houserobber

import "go-leetcode/src/utils"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) < 2 {
		return utils.Max(nums[0], 0)
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = utils.Max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = utils.Max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}
