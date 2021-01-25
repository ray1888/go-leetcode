package DynamicProgramming

func maxAreaBurtalForce(height []int) int {
	if len(height) == 0 {
		return 0
	}
	maxValue := 0

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			cur := min(height[i], height[j]) * (j - i)
			maxValue = max(maxValue, cur)
		}
	}
	return maxValue
}

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	maxValue := 0
	i := 0
	j := len(height) - 1
	for i < j {
		cur := min(height[i], height[j]) * (j - i)
		maxValue = max(maxValue, cur)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxValue
}
