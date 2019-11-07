package tree

import "go-leetcode/src/datastructure"

func postOrder(root *datastructure.TreeNode, result *[]int) {
	if root == nil {
		return
	}
	postOrder(root.Left, result)
	postOrder(root.Right, result)
	*result = append(*result, root.Val)
}

func postorderTraversal(root *datastructure.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0, 10)
	postOrder(root, &result)
	return result
}
