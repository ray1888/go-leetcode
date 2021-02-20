package tree

import "go-leetcode/src/datastructure"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepthNAnyTree(root *datastructure.NTreeNode) int {
	if root == nil {
		return 0
	}
	level := 1
	originStart := 1
	maxDepthFunc(root, &originStart, &level)
	return level
}

func maxDepthFunc(root *datastructure.NTreeNode, depth *int, maxDepth *int) {
	if root == nil {
		*maxDepth = max(*maxDepth, *depth)
		return
	}
	if len(root.Children) == 0 {
		*maxDepth = max(*maxDepth, *depth)
		return
	}
	for i := 0; i < len(root.Children); i++ {
		*depth++
		maxDepthFunc(root.Children[i], depth, maxDepth)
		*depth--
	}
}
