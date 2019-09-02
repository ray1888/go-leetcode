package tree

import "go-leetcode/src/datastructure"

func IsSameTree(s *datastructure.TreeNode, t *datastructure.TreeNode) bool {
	if s == nil && t == nil {
		return true
	} else if s == nil || t == nil {
		return false
	}
	return s.Val == t.Val && IsSameTree(s.Left, t.Left) && IsSameTree(s.Right, t.Right)
}

func IsSubtree(s *datastructure.TreeNode, t *datastructure.TreeNode) bool {
	if s == nil && t != nil {
		return false
	} else if s != nil && t == nil {
		return true
	}
	return IsSameTree(s, t) || IsSubtree(s.Left, t) || IsSubtree(s.Right, t)
}
