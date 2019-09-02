package tree

import "go-leetcode/src/datastructure"

func _searchBST(root *datastructure.TreeNode, val int) *datastructure.TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > val {
		return _searchBST(root.Left, val)
	} else if root.Val == val {
		return root
	} else {
		return _searchBST(root.Right, val)
	}
}

func searchBST(root *datastructure.TreeNode, val int) *datastructure.TreeNode {
	if root == nil {
		return nil
	}
	return _searchBST(root, val)
}
