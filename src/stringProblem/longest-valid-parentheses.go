package stringProblem

import "go-leetcode/src/utils"

func longestValidParentheses(s string) int {
	maxAns := 0
	stack := []int{}
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxAns = utils.Max(maxAns, i-stack[len(stack)-1])
			}
		}
	}
	return maxAns

}
