package tree

import "go-leetcode/src/datastructure"

func _sortArrayToBST(nums []int, start, end int) *datastructure.TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	root := &datastructure.TreeNode{Val: nums[mid], Left: nil, Right: nil}
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

type Item struct {
	parent *datastructure.TreeNode
	start  int
	end    int
	isLeft bool
}

func sortedArrayToBSTIterative(nums []int) *datastructure.TreeNode {
	if len(nums) == 0 {
		return &datastructure.TreeNode{}
	}
	stack := make([]Item, 0)

	dummy := &datastructure.TreeNode{Val: 0}
	newItem := Item{parent: dummy, start: 0, end: len(nums) - 1, isLeft: true}
	stack = append(stack, newItem)
	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if item.start <= item.end {
			mid := item.start + (item.end-item.start)/2
			child := &datastructure.TreeNode{Val: nums[mid]}
			if item.isLeft {
				item.parent.Left = child
			} else {
				item.parent.Right = child
			}
			stack = append(stack, Item{parent: child, start: item.start, end: mid - 1, isLeft: true})
			stack = append(stack, Item{parent: child, start: mid + 1, end: item.end, isLeft: false})
		}
	}
	return dummy.Left
}
