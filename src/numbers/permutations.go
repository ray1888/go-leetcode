package numbers

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	result := make([][]int, 0)
	permuteRec(nums, 0, &result)
	return result
}

func permuteRec(nums []int, start int, result *[][]int) {
	if start == len(nums) {
		items := make([]int, len(nums))
		copy(items, nums)
		*result = append(*result, items)
		return
	}
	for i := start; i < len(nums); i++ {
		nums[i], nums[start] = nums[start], nums[i]
		permuteRec(nums, start+1, result)
		nums[start], nums[i] = nums[i], nums[start]
	}
}
