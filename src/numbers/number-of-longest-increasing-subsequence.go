package numbers

import "go-leetcode/src/utils"

func findNumberOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length := len(nums)
	count := make([]int, length)
	dp := make([]int, length)

	for i := 0; i < length; i++ {
		count[i] = 1
	}
	maxLength := 0
	for j := 0; j < length; j++ {
		for i := 0; i < j; i++ {
			if nums[j] > nums[i] {
				if dp[i] >= dp[j] {
					dp[j] = dp[i] + 1
					count[j] = count[i]
				} else if dp[i]+1 == dp[j] {
					count[j] += count[i]
				}
			}
			maxLength = utils.Max(maxLength, dp[j])
		}
	}
	sum := 0
	for i := 0; i < length; i++ {
		if dp[i] == maxLength {
			sum += count[i]
		}
	}
	return sum
}
