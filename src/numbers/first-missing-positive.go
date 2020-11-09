package numbers

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	p := 0
	n := len(nums)
	for p < n {
		num := nums[p]
		if num >= 1 && num <= n && nums[p] != nums[num-1] {
			nums[p], nums[num-1] = nums[num-1], nums[p]
		} else {
			p++
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
