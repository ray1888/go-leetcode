package tree

import "go-leetcode/src/datastructure"

func Reverse(input []int) []int {
	i := 0
	j := len(input) - 1
	for i <= j {
		input[i], input[j] = input[j], input[i]
		i += 1
		j -= 1
	}
	r := make([]int, len(input))
	copy(r, input)
	return r
}

func zigzagLevelOrder(root *datastructure.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	stack := make([]*datastructure.TreeNode, 0, 100)
	stack = append(stack, root)
	reverse := false
	cur := stack[0]
	for len(stack) > 0 {
		length := len(stack)
		s_result := make([]int, 0, 10)
		s_result = append(s_result, cur.Val)
		for i := 0; i < length; i++ {
			cur = stack[0]
			stack = stack[1:]
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
		}
		if reverse {
			s_result = Reverse(s_result)
			reverse = false
		} else {
			reverse = true
		}
		result = append(result, s_result)
	}
	return result
}
