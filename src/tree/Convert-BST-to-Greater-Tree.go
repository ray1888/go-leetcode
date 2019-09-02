package tree

import "go-leetcode/src/datastructure"

func _converBST(root *datastructure.TreeNode, sum int) int {
	if root == nil {
		return sum
	}
	cur := _converBST(root.Right, sum)
	root.Val += cur
	return _converBST(root.Left, root.Val)
}

func convertBST(root *datastructure.TreeNode) *datastructure.TreeNode {
	if root == nil {
		return nil
	}
	_converBST(root, 0)
	return root
}
