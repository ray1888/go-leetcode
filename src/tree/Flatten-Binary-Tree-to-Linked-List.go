package tree

import "go-leetcode/src/datastructure"

var prev *datastructure.TreeNode = nil

func _flattenReversePreOrder(root *datastructure.TreeNode) *datastructure.TreeNode {
	// work in seq right left root
	if root == nil {
		return nil
	}
	_flattenReversePreOrder(root.Right)
	_flattenReversePreOrder(root.Left)
	root.Left = nil
	root.Right = prev
	prev = root
	return root
}

func flatten(root *datastructure.TreeNode) *datastructure.TreeNode {
	// work in seq right left root
	if root == nil {
		return nil
	}
	prev = nil
	return _flattenReversePreOrder(root)
}
