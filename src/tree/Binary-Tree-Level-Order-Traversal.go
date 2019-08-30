package tree

import "go-leetcode/src/datastructure"

func levelOrder(root *datastructure.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := make([]*datastructure.TreeNode, 0)
	if root.Left != nil {
		queue = append(queue, root.Left)
	}
	if root.Right != nil {
		queue = append(queue, root.Right)
	}
	var result [][]int = [][]int{}
	result = append(result, []int{root.Val})
	var cur *datastructure.TreeNode
	for len(queue) != 0 {
		length := len(queue)
		tmp_result := make([]int, 0)
		for i := 0; i < length; i++ {
			//if len(queue)
			cur, queue = queue[0], queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			tmp_result = append(tmp_result, cur.Val)
		}
		result = append(result, tmp_result)
	}
	return result
}
