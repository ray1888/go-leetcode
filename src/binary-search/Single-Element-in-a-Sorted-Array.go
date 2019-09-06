package binary_search

func singleNonDuplicate(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	i := 0
	length := len(nums)
	j := length - 1
	for i <= j {
		mid := i + (j-i)/2
		if mid < 0 || mid > length-1 {
			break
		}
		if mid-1 >= 0 && nums[mid] == nums[mid-1] {
			mid -= 1
		} else if mid+1 < length && nums[mid] == nums[mid+1] {
		} else {
			return nums[mid]
		}
		if (mid-i)%2 == 1 {
			j = mid - 1
		} else {
			i = mid + 2
		}
	}
	return -1
}
