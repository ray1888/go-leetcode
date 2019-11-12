package nsum

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := 0
	min := int(math.Pow(2, 32))
	for j := len(nums) - 1; j >= 2; j-- {
		i := 0
		k := j - 1
		for i < k {
			sum := nums[i] + nums[k] + nums[j]
			if target == sum {
				return sum
			} else if target < sum {
				k -= 1
			} else {
				i += 1
			}
			diff := target - sum
			if diff < 0 {
				diff = -1 * diff
			}
			if diff < min {
				result = sum
				min = diff
			}
		}
	}
	return result
}
