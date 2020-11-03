package tree

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	// 二叉搜索树的性质
	if root.Val < L {
		return rangeSumBST(root.Right, L, R)
	}
	if root.Val > R {
		return rangeSumBST(root.Left, L, R)
	}
	return root.Val + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)
}
