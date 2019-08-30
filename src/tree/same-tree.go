package tree

import "go-leetcode/src/datastructure"

func isSameTree(p *datastructure.TreeNode, q *datastructure.TreeNode) bool {
	if p != nil && q != nil {
		return p.Val == q.Val &&
			isSameTree(p.Left, q.Left) &&
			isSameTree(p.Right, q.Right)
	}
	return p == nil && q == nil
}

func isSameTreeRecursive(p *datastructure.TreeNode, q *datastructure.TreeNode) bool {
	return isSameTree(p, q)
}
