package tree

import "go-leetcode/src/datastructure"

func largestBSTSubtree(root *datastructure.TreeNode) int {
	if isBST(root) {
		n := nodes(root)
		return n
	}
	l := largestBSTSubtree(root.Left)
	r := largestBSTSubtree(root.Right)
	if l > r {
		return l
	}
	return r
}

func isBST(node *datastructure.TreeNode) bool {
	if node == nil {
		return true
	}
	if node.Left != nil && maxOf(node.Left) >= node.Val {
		return false
	}
	if node.Right != nil && minOf(node.Right) <= node.Val {
		return false
	}
	return isBST(node.Left) && isBST(node.Right)
}

func maxOf(node *datastructure.TreeNode) int {
	if node.Right != nil {
		return maxOf(node.Right)
	}
	return node.Val
}

func minOf(node *datastructure.TreeNode) int {
	if node.Left != nil {
		return minOf(node.Left)
	}
	return node.Val
}

func nodes(node *datastructure.TreeNode) int {
	if node == nil {
		return 0
	}
	l := nodes(node.Left)
	r := nodes(node.Right)
	return l + r + 1
}
