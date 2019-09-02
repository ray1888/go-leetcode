package tree

import "go-leetcode/src/datastructure"

func _buildtree(InOrderMap map[int]int, prestartpos int, preendpos int, instart int, preorder []int) *datastructure.TreeNode {
	if prestartpos > preendpos {
		return nil
	}
	rootId := InOrderMap[preorder[prestartpos]]
	leftLength := rootId - instart
	root := &datastructure.TreeNode{preorder[prestartpos], nil, nil}
	root.Left = _buildtree(InOrderMap, prestartpos+1, prestartpos+leftLength, instart, preorder)
	root.Right = _buildtree(InOrderMap, prestartpos+leftLength+1, preendpos, rootId+1, preorder)
	return root
}

func buildTree(preorder []int, inorder []int) *datastructure.TreeNode {
	if preorder == nil && inorder == nil {
		return nil
	}
	InOrder := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		InOrder[inorder[i]] = i
	}
	return _buildtree(InOrder, 0, len(inorder)-1, 0, preorder)
}
