package DynamicProgramming

func maxSubArrayLen(nums []int, k int) int {

	prefixSum := make([]int, len(nums)+1)
	resultMap := make(map[int]int)
	resultMap[0] = 0
	for i := 1; i <= len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
		if _, ok := resultMap[prefixSum[i]]; !ok {
			resultMap[prefixSum[i]] = i
		}
	}

	maxValue := 0
	for i := len(nums); i >= 0; i-- {
		if val, ok := resultMap[prefixSum[i]-k]; ok {
			maxValue = max(maxValue, i-val)
		}
	}
	return maxValue
}
