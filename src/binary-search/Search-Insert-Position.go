package binary_search

func _searchInsert(nums []int, start, end, target int) int {
	if start < 0 || end < 0 {
		return 0
	}
	mid := start + (end-start)/2
	if mid < start {
		if nums[start] < target {
			return start + 1
		} else {
			return start
		}
	} else if mid > end {
		if nums[end] > target {
			return end
		} else {
			return end + 1
		}
	}
	if nums[mid] > target {
		return _searchInsert(nums, start, mid-1, target)
	} else if nums[mid] < target {
		return _searchInsert(nums, mid+1, end, target)
	} else {
		return mid
	}
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	return _searchInsert(nums, 0, len(nums)-1, target)
}
