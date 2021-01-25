package DynamicProgramming

func jump(nums []int) int {
	maxSteps := make([]int, len(nums))
	maxLength := 0
	for i := 0; i < len(nums); i++ {
		if maxLength >= len(nums)-1 {
			return maxSteps[len(nums)-1]
		}
		// 已经不可能再前进了
		if i > maxLength {
			return -1
		}
		maxLength = max(maxLength, i+nums[i])
		last := min(len(nums)-1, maxLength)
		for j := last; j > i && maxSteps[j] == 0; j-- {
			maxSteps[j] = maxSteps[i] + 1
		}
	}
	return -1
}
