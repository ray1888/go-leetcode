package DynamicProgramming

func findMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	hMap := make(map[int]int)
	hMap[0] = -1
	maxLen := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			count--
		}
		if val, ok := hMap[count]; !ok {
			hMap[count] = i
		} else {
			maxLen = max(maxLen, i-val)
		}
	}
	return maxLen
}
