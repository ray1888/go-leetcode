package numbers

func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums) - 1
}

func findPeakElementBinarySearch(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	low := 0
	high := len(nums) - 1
	for low < high {
		mid := low + (high-low)/2
		var left int
		var right int
		if mid-1 >= 0 {
			left = nums[mid-1]
		} else {
			left = -1 * (2 << 30)
		}
		if mid+1 < len(nums) {
			right = nums[mid+1]
		} else {
			right = 2 << 30
		}
		if nums[mid] > left && nums[mid] > right {
			return mid
		} else if left > nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
