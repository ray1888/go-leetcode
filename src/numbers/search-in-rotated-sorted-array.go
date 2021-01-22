package numbers

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		if target == nums[mid] {
			return mid
		}
		if nums[mid] >= nums[low] {
			if target < nums[mid] && target >= nums[low] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
