package tree

import "go-leetcode/src/datastructure"

func _invertTree(root *datastructure.TreeNode) *datastructure.TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	_invertTree(root.Left)
	_invertTree(root.Right)
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	return root
}

func invertTree(root *datastructure.TreeNode) *datastructure.TreeNode {
	if root == nil {
		return root
	}
	return _invertTree(root)
}
