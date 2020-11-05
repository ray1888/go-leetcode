package numbers

func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)
	i := 1
	length := 1
	for i < n && nums[i] == nums[i-1] {
		i++
	}
	for i < n {
		if nums[i] > nums[i-1] {
			length++

		}
		for i < n && nums[i] >= nums[i-1] {
			i++
		}
		if i < n {
			length++
		}
		for i < n && nums[i] <= nums[i-1] {
			i++
		}
	}
	return length
}
