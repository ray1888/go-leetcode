package tree

import "go-leetcode/src/datastructure"

func _buildTreeByInPost(InOrderMap map[int]int, poststart int, postend int, instart int, postorder []int) *datastructure.TreeNode {
	if poststart > postend {
		return nil
	}
	rootId := InOrderMap[postorder[postend]]
	leftLength := rootId - instart
	root := &datastructure.TreeNode{postorder[postend], nil, nil}
	root.Left = _buildTreeByInPost(InOrderMap, poststart, poststart+leftLength-1, instart, postorder)
	root.Right = _buildTreeByInPost(InOrderMap, poststart+leftLength, postend-1, rootId+1, postorder)
	return root
}

func buildTreeByInPost(inorder []int, postorder []int) *datastructure.TreeNode {
	if postorder == nil && inorder == nil {
		return nil
	}
	InOrder := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		InOrder[inorder[i]] = i
	}
	return _buildTreeByInPost(InOrder, 0, len(postorder)-1, 0, postorder)
}
