package DynamicProgramming

func maxProduct(nums []int) int {
	curMax := nums[0]
	curMin := nums[0]
	maxValue := nums[0]

	for i := 1; i < len(nums); i++ {
		tmpMax := curMax * nums[i]
		tmpMin := curMin * nums[i]
		c := nums[i]
		curMax = max(tmpMax, max(tmpMin, c))
		curMin = min(tmpMax, min(tmpMin, c))
		maxValue = max(maxValue, curMax)
	}
	return maxValue
}
