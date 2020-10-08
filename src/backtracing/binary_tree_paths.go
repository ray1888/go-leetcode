package backtracing

import (
	"go-leetcode/src/datastructure"
	"strconv"
	"strings"
)

func dfsTreePath(node *datastructure.TreeNode, elem []int, result *[]string) {
	if node.Left == nil && node.Right == nil {
		elem = append(elem, node.Val)
		s := make([]string, 0)
		for _, item := range elem {
			s = append(s, strconv.Itoa(item))
		}
		ss := strings.Join(s, "->")
		*result = append(*result, ss)
		elem = elem[:len(elem)-1]
		return
	}
	elem = append(elem, node.Val)
	if node.Left != nil {
		dfsTreePath(node.Left, elem, result)
	}
	if node.Right != nil {
		dfsTreePath(node.Right, elem, result)
	}
	elem = elem[:len(elem)-1]
}

func binaryTreePaths(root *datastructure.TreeNode) []string {
	if root == nil {
		return []string{}
	}
	result := make([]string, 0)
	elem := make([]int, 0)
	dfsTreePath(root, elem, &result)
	return result
}
