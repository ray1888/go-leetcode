package tree

import "go-leetcode/src/datastructure"

func levelOrderNAryTree(root *datastructure.NTreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)

	queue := make([]*datastructure.NTreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		tmpResult := make([]int, 0)
		for i := 0; i < length; i++ {
			node := queue[i]
			tmpResult = append(tmpResult, node.Val)
			if len(node.Children) > 0 {
				queue = append(queue, node.Children...)
			}
		}
		queue = queue[length:]
		result = append(result, tmpResult)
	}
	return result
}
