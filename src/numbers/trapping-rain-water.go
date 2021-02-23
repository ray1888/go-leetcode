package numbers

import "go-leetcode/src/utils"

func trapSpaceOn(height []int) int {
	if len(height) == 0 {
		return 0
	}
	leftMax := -1
	rightMax := -1
	water := 0
	d := make([]int, len(height))
	length := len(height)
	for i := 0; i < length; i++ {
		leftMax = utils.Max(height[i], leftMax)
		d[i] = leftMax
	}

	for i := length - 1; i >= 0; i-- {
		rightMax = utils.Max(rightMax, height[i])
		d[i] = utils.Min(d[i], rightMax)
		water += (d[i] - height[i])
	}
	return water
}

func trapSpaceO1(height []int) int {
	if len(height) == 0 {
		return 0
	}
	leftMax := -1
	rightMax := -1
	water := 0
	length := len(height)
	i := 0
	j := length - 1
	for i < j {
		leftMax = utils.Max(leftMax, height[i])
		rightMax = utils.Max(rightMax, height[j])
		if leftMax < rightMax {
			water += (leftMax - height[i])
			i++
		} else {
			water += (rightMax - height[j])
			j--
		}
	}
	return water
}
