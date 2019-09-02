package tree

import "go-leetcode/src/datastructure"

func _mergeTrees(t1 *datastructure.TreeNode, t2 *datastructure.TreeNode) *datastructure.TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil && t2 != nil {
		t1 = t2
		t1.Left = _mergeTrees(nil, t2.Left)
		t1.Right = _mergeTrees(nil, t2.Right)
	} else if t1 != nil {
		if t2 != nil {
			t1.Val += t2.Val
			t1.Left = _mergeTrees(t1.Left, t2.Left)
			t1.Right = _mergeTrees(t1.Right, t2.Right)
		} else {
			t1.Left = _mergeTrees(t1.Left, nil)
			t1.Right = _mergeTrees(t1.Right, nil)
		}
	}
	return t1
}

func _mergeTreesClear(t1 *datastructure.TreeNode, t2 *datastructure.TreeNode) *datastructure.TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	} else if t1 == nil && t2 != nil {
		return t2
	} else if t1 != nil && t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = _mergeTreesClear(t1.Left, t2.Left)
	t1.Right = _mergeTreesClear(t1.Right, t2.Right)
	return t1
}

func mergeTrees(t1 *datastructure.TreeNode, t2 *datastructure.TreeNode) *datastructure.TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	return _mergeTrees(t1, t2)
}
