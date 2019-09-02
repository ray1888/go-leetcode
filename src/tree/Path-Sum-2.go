package tree

import "go-leetcode/src/datastructure"

func _pathsum(root *datastructure.TreeNode, sum int, stack *[]int, tr *[][]int) {
	if root == nil {
		return
	}
	*stack = append(*stack, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == sum {
		sr := make([]int, len(*stack))
		copy(sr, *stack)
		*tr = append(*tr, sr)
	}
	_pathsum(root.Left, sum-root.Val, stack, tr)
	_pathsum(root.Right, sum-root.Val, stack, tr)
	*stack = (*stack)[:len(*stack)-1]
}

func pathSum(root *datastructure.TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)
	sr := make([]int, 0)
	_pathsum(root, sum, &sr, &result)
	return result
}
