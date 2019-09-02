package tree

import "go-leetcode/src/datastructure"

func _sortArrayToBST(nums []int, start, end int) *datastructure.TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	root := &datastructure.TreeNode{nums[mid], nil, nil}
	root.Left = _sortArrayToBST(nums, start, mid-1)
	root.Right = _sortArrayToBST(nums, mid+1, end)
	return root
}

func sortedArrayToBST(nums []int) *datastructure.TreeNode {
	if len(nums) == 0 || nums == nil {
		return nil
	}
	return _sortArrayToBST(nums, 0, len(nums)-1)
}
