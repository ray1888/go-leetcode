package numbers

func findDuplicates(nums []int) []int {
	abs := func(i int) int {
		if i < -len(nums) {
			return -i - len(nums)
		} else if i < 0 {
			return -i
		}
		return i
	}

	ret := make([]int, 0)

	for _, num := range nums {
		t := abs(num)
		if nums[t-1] < 0 {
			nums[t-1] -= len(nums)
		} else {
			nums[t-1] = -nums[t-1]
		}
	}

	for i, num := range nums {
		if num < -len(nums) {
			ret = append(ret, i+1)
		}
	}
	return ret
}
