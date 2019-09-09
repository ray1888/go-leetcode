package binary_search

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	} else if len(nums) == 1 {
		return nums[0]
	}
	i := 0
	j := len(nums) - 1
	for i < j {
		mid := i + (j-i)/2
		if nums[mid] > nums[j] {
			i = mid + 1
		} else {
			j = mid
		}

	}
	return nums[i]
}
