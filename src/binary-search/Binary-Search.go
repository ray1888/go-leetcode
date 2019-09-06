package binary_search

func _search(nums []int, start, end, target int) int {
	mid := start + (end-start)/2
	if mid < start || mid > end {
		return -1
	}
	if nums[mid] > target {
		return _search(nums, start, mid-1, target)
	} else if nums[mid] < target {
		return _search(nums, mid+1, end, target)
	} else {
		return mid
	}
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	return _search(nums, 0, len(nums)-1, target)
}
