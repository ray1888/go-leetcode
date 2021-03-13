package slidingwindow

import "go-leetcode/src/utils"

func maxTurbulenceSize(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	maxLength := 1
	left := 0
	right := 0

	for right < len(arr)-1 {
		if left == right {
			if arr[left] == arr[left+1] {
				left++
			}
			right++
		} else {
			if arr[right-1] < arr[right] && arr[right] > arr[right+1] {
				right++
			} else if arr[right-1] > arr[right] && arr[right] < arr[right+1] {
				right++
			} else {
				left = right
			}
		}
		maxLength = utils.Max(maxLength, right-left+1)
	}
	return maxLength
}
