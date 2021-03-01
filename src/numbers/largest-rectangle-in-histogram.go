package numbers

import "go-leetcode/src/utils"

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	length := len(heights)
	stack := make([]int, 0)
	maxArea := 0
	for i := 0; i < length+1; i++ {
		var h int
		if i < length {
			h = heights[i]
		}
		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			var l int
			if len(stack) > 0 {
				l = stack[len(stack)-1]
			} else {
				l = -1
			}
			area := (i - l - 1) * heights[idx]
			maxArea = utils.Max(maxArea, area)
		}
		stack = append(stack, i)
	}
	return maxArea
}
