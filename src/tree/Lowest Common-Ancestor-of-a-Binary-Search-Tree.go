package tree

import "go-leetcode/src/datastructure"

func lowestCommonAncestor(root, p, q *datastructure.TreeNode) *datastructure.TreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
