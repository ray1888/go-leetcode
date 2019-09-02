package tree

import "go-leetcode/src/datastructure"

func deleteNodeFail(root *datastructure.TreeNode, key int) *datastructure.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		deleteNode(root.Left, key)
	} else if root.Val < key {
		deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		leftmax := root.Left
		for leftmax.Right != nil {
			leftmax = leftmax.Right
		}
		leftmax.Right = root.Right
		deleteNode(leftmax.Right, leftmax.Val)
		root = root.Left
	}
	return root
}

func deleteNode(root *datastructure.TreeNode, key int) *datastructure.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		deleteNode(root.Left, key)
	} else if root.Val < key {
		deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		root.Val = findMax(root.Left).Val
		root.Left = deleteNode(root.Left, root.Val)
	}
	return root
}

func findMax(root *datastructure.TreeNode) *datastructure.TreeNode {
	if root == nil {
		return nil
	} else if root.Right == nil {
		return root
	}
	return findMax(root.Right)
}
