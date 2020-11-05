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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func wiggleMaxLengthDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)
	up := 1
	down := 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}

	return max(up, down)
}
