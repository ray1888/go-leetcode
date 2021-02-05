package houserobber

import (
	"go-leetcode/src/datastructure"
	"go-leetcode/src/utils"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var valueMap = map[*datastructure.TreeNode]int{}

func rob3(root *datastructure.TreeNode) int {
	if root == nil {
		return 0
	}
	if val, ok := valueMap[root]; ok {
		return val
	}

	notDo := rob3(root.Left) + rob3(root.Right)
	do := root.Val
	if root.Left != nil {
		do += (rob3(root.Left.Left) + rob3(root.Left.Right))
	}
	if root.Right != nil {
		do += (rob3(root.Right.Right) + rob3(root.Right.Left))
	}

	maxValue := utils.Max(do, notDo)
	valueMap[root] = maxValue
	return maxValue
}
