package numbers

import "sort"

func triangleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	length := len(nums)
	count := 0
	for i := 0; i < length-2; i++ {
		k := i + 2
		for j := i + 1; j < length-1 && nums[i] != 0; j++ {
			for k < length && nums[i]+nums[j] > nums[k] {
				k++
			}
			count += (k - j - 1)
		}
	}
	return count
}
