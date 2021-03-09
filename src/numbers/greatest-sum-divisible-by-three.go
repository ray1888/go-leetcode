package numbers

import (
	"go-leetcode/src/utils"
	"math"
)

func maxSumDivThree(nums []int) int {
	state := make([]int, 3)

	state[1] = math.MinInt32
	state[2] = math.MinInt32

	for _, item := range nums {
		if item%3 == 0 {
			state = []int{state[0] + item, state[1] + item, state[2] + item}
		} else if item%3 == 1 {
			a := utils.Max(state[2]+item, state[0])
			b := utils.Max(state[0]+item, state[1])
			c := utils.Max(state[1]+item, state[2])
			state = []int{a, b, c}
		} else if item%3 == 2 {
			a := utils.Max(state[1]+item, state[0])
			b := utils.Max(state[2]+item, state[1])
			c := utils.Max(state[0]+item, state[2])
			state = []int{a, b, c}
		}
	}
	return state[0]
}
