package tree

import "go-leetcode/src/datastructure"

func _pathSum3(root *datastructure.TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	result := 0
	if root.Val == sum {
		result += 1
	}
	result += (_pathSum3(root.Left, sum-root.Val) + _pathSum3(root.Right, sum-root.Val))
	return result
}

func pathSum3(root *datastructure.TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	result := _pathSum3(root, sum) + pathSum3(root.Left, sum) + pathSum3(root.Right, sum)
	return result
}
