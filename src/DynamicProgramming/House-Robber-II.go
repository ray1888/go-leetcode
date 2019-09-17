package DynamicProgramming

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) < 2 {
		return max(nums[0], 0)
	}
	return max(rob(nums[:len(nums)-1]), rob(nums[1:]))
}
