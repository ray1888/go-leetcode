package tree

import "go-leetcode/src/datastructure"

func inorder(root *datastructure.TreeNode, result *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, result)
	*result = append(*result, root.Val)
	inorder(root.Right, result)
}

func inorderTraversal(root *datastructure.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	inorder(root, &result)
	return result
}
