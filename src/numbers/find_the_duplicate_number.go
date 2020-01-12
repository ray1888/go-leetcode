package numbers

func findDuplicateBrutalForce(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	return -1
}

func findDuplicateBinarySearch(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	low := 1
	high := len(nums) - 1
	for low < high {
		mid := low + (high-low)/2
		count := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid {
				count += 1
			}
		}
		if count > mid {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func findDuplicateTwoPointer(nums []int) int {
	// if problem is isexist, it will be panic by index out of range
	slow := nums[0]
	fast := nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	p := 0
	for slow != p {
		slow = nums[slow]
		p = nums[p]
	}
	return slow
}
