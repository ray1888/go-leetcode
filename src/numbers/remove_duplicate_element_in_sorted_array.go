package numbers

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	p := 1
	length := len(nums)
	for q := 1; q < length; q++ {
		if nums[q] != nums[q-1] {
			nums[p] = nums[q]
			p += 1
		}
	}
	return p
}
