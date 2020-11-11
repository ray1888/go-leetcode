package numbers

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	p := 0
	length := len(nums)
	for q := 0; q < length; q++ {
		if nums[q] != val {
			nums[p] = nums[q]
			p++
		}
	}
	return p
}
