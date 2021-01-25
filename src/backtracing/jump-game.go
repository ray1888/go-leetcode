package backtracing

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	return canJumpRec(0, nums)
}

func canJumpRec(start int, nums []int) bool {
	//
	if start >= len(nums)-1 {
		return true
	}
	var jump int
	if nums[start]+start < len(nums)-1 {
		jump = nums[start] + start
	} else {
		jump = len(nums) - 1
	}
	for i := jump; i > start; i-- {
		if canJumpRec(i, nums) {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canJumpDP(nums []int) bool {
	maxLength := 0
	for i := 0; i < len(nums); i++ {
		if maxLength >= len(nums)-1 {
			return true
		}
		if i > maxLength {
			return false
		}
		maxLength = max(maxLength, i+nums[i])
	}
	return false
}
